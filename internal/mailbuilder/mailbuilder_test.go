package mailbuilder_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"net/mail"
	"os"
	"testing"

	schema "github.com/ludusrusso/kannon/db"
	"github.com/ludusrusso/kannon/generated/pb"
	sqlc "github.com/ludusrusso/kannon/internal/db"
	"github.com/ludusrusso/kannon/internal/mailbuilder"
	"github.com/ludusrusso/kannon/internal/pool"
	"github.com/ludusrusso/kannon/internal/statssec"
	"github.com/ludusrusso/kannon/internal/tests"
	"github.com/ludusrusso/kannon/pkg/api/adminapi"
	"github.com/ludusrusso/kannon/pkg/api/mailapi"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var db *sql.DB
var q *sqlc.Queries
var mb mailbuilder.MailBulder
var ma pb.MailerServer
var adminAPI pb.ApiServer
var pm pool.SendingPoolManager

func TestMain(m *testing.M) {
	var purge tests.PurgeFunc
	var err error

	db, purge, err = tests.TestPostgresInit(schema.Schema)
	if err != nil {
		logrus.Fatalf("Could not start resource: %s", err)
	}

	q := sqlc.New(db)

	mb = mailbuilder.NewMailBuilder(q, statssec.NewStatsService(q))
	ma = mailapi.NewMailAPIService(q)
	adminAPI = adminapi.CreateAdminAPIService(q)
	pm = pool.NewSendingPoolManager(q)

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := purge(); err != nil {
		logrus.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestPrepareMail(t *testing.T) {
	d, err := adminAPI.CreateDomain(context.Background(), &pb.CreateDomainRequest{
		Domain: "test.com",
	})
	assert.Nil(t, err)

	token := base64.StdEncoding.EncodeToString([]byte(d.Domain + ":" + d.Key))
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Basic "+token))

	sendRes, err := ma.SendHTML(ctx, &pb.SendHTMLReq{
		Sender: &pb.Sender{
			Email: "test@test.com",
			Alias: "Test",
		},
		To:            []string{"test@emailtest.com"},
		Subject:       "Test",
		Html:          "test",
		ScheduledTime: timestamppb.Now(),
	})
	assert.Nil(t, err)

	emails, err := pm.PrepareForSend(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(emails))

	m, err := mb.PerpareForSend(context.Background(), emails[0])
	parsed, err := mail.ReadMessage(bytes.NewReader(m.Body))
	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, "bump_dGVzdEBlbWFpbHRlc3QuY29t+"+sendRes.MessageId, parsed.Header.Get("Return-Path"))
	assert.Equal(t, "test@emailtest.com", parsed.Header.Get("To"))
	assert.Equal(t, "Test <test@test.com>", parsed.Header.Get("From"))
}