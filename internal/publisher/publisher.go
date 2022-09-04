package publisher

import (
	"fmt"

	sq "github.com/ludusrusso/kannon/internal/stats_db"
	mailertypes "github.com/ludusrusso/kannon/proto/kannon/mailer/types"
	"github.com/ludusrusso/kannon/proto/kannon/stats/types"
	"google.golang.org/protobuf/proto"
)

type Publisher interface {
	Publish(subj string, data []byte) error
}

func SendEmail(p Publisher, email *mailertypes.EmailToSend) error {
	msg, err := proto.Marshal(email)
	if err != nil {
		return err
	}
	err = p.Publish("kannon.sending", msg)
	if err != nil {
		return err
	}
	return nil
}

func PublishStat(p Publisher, stats *types.Stats) error {
	stype := sq.GetStatsType(stats)
	subj := fmt.Sprintf("kannon.stats.%s", stype)

	data, err := proto.Marshal(stats)
	if err != nil {
		return fmt.Errorf("cannot marshal protoc: %v", err)
	}
	return p.Publish(subj, data)
}
