package mailapi

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/ludusrusso/kannon/generated/pb"
	sqlc "github.com/ludusrusso/kannon/internal/db"
	"github.com/ludusrusso/kannon/internal/domains"
	"github.com/ludusrusso/kannon/internal/pool"
	"github.com/ludusrusso/kannon/internal/templates"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mailAPIService struct {
	domains     domains.DomainManager
	templates   templates.Manager
	sendingPoll pool.SendingPoolManager
}

func (s mailAPIService) SendHTML(ctx context.Context, in *pb.SendHTMLRequest) (*pb.SendResponse, error) {
	domain, err := s.getCallDomainFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid or wrong auth")
	}

	template, err := s.templates.CreateTemplate(ctx, in.Html, domain.Domain)
	if err != nil {
		logrus.Errorf("cannot create template %v\n", err)
		return nil, status.Errorf(codes.Internal, "cannot create template %v", err)
	}

	sender := pool.Sender{
		Email: in.Sender.Email,
		Alias: in.Sender.Alias,
	}
	pool, err := s.sendingPoll.AddPool(ctx, template, in.To, sender, in.Subject, domain.Domain)

	if err != nil {
		logrus.Errorf("cannot create pool %v\n", err)
		return nil, err
	}

	response := pb.SendResponse{
		MessageId:     pool.MessageID,
		TemplateId:    template.TemplateID,
		ScheduledTime: timestamppb.New(time.Now()),
	}

	return &response, nil
}

func (s mailAPIService) SendTemplate(ctx context.Context, in *pb.SendTemplateRequest) (*pb.SendResponse, error) {
	domain, err := s.getCallDomainFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid or wrong auth")
	}
	template, err := s.templates.FindTemplate(ctx, domain.Domain, in.TemplateId)
	if err != nil {
		logrus.Errorf("cannot create template %v\n", err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot find template with id: %v", in.TemplateId)
	}

	sender := pool.Sender{
		Email: in.Sender.Email,
		Alias: in.Sender.Alias,
	}
	pool, err := s.sendingPoll.AddPool(ctx, template, in.To, sender, in.Subject, domain.Domain)

	if err != nil {
		logrus.Errorf("cannot create pool %v\n", err)
		return nil, err
	}

	response := pb.SendResponse{
		MessageId:     pool.MessageID,
		TemplateId:    template.TemplateID,
		ScheduledTime: timestamppb.New(time.Now()),
	}

	return &response, nil
}

func (s mailAPIService) Close() error {
	return s.domains.Close()
}

func (s mailAPIService) getCallDomainFromContext(ctx context.Context) (sqlc.Domain, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return sqlc.Domain{}, fmt.Errorf("cannot find metatada")
	}

	auths := m.Get("authorization")
	if len(auths) != 1 {
		return sqlc.Domain{}, fmt.Errorf("cannot find authorization header")
	}

	auth := auths[0]
	if !strings.HasPrefix(auth, "Basic ") {
		return sqlc.Domain{}, fmt.Errorf("no prefix Basic in auth: %v", auth)
	}

	token := strings.Replace(auth, "Basic ", "", 1)
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return sqlc.Domain{}, fmt.Errorf("decode token error: %v", token)
	}

	authData := string(data)

	var d, k string
	_, err = fmt.Sscanf(authData, "%v:%v", &d, &k)
	if err != nil {
		return sqlc.Domain{}, fmt.Errorf("invalid token: %w", err)
	}

	domain, err := s.domains.FindDomainWithKey(ctx, d, k)
	if err != nil {
		return sqlc.Domain{}, fmt.Errorf("cannot find domain: %w", err)
	}

	return domain, nil
}

func NewMailAPIService(dbi *sql.DB, q *sqlc.Queries) (pb.MailerServer, error) {
	domainsCli := domains.NewDomainManager(q)

	sendingPoolCli, err := pool.NewSendingPoolManager(dbi)
	if err != nil {
		return nil, err
	}

	templates, err := templates.NewTemplateManager(dbi)
	if err != nil {
		return nil, err
	}

	return &mailAPIService{
		domains:     domainsCli,
		sendingPoll: sendingPoolCli,
		templates:   templates,
	}, nil
}
