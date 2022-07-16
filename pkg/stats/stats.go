package stats

import (
	"context"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/ludusrusso/kannon/generated/pb"
	sq "github.com/ludusrusso/kannon/internal/stats_db"
	"github.com/ludusrusso/kannon/internal/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	"github.com/nats-io/nats.go"
)

func Run(ctx context.Context, vc *viper.Viper) {
	vc.SetEnvPrefix("STATS")
	vc.AutomaticEnv()

	dbURL := vc.GetString("database_url")
	natsURL := vc.GetString("nats_url")

	logrus.Info("🚀 Starting stats")

	db, q, err := sq.Conn(ctx, dbURL)
	if err != nil {
		logrus.Fatalf("cannot connect to database: %v", err)
	}
	defer db.Close()

	_, js, closeNats := utils.MustGetNats(natsURL)
	defer closeNats()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		handleErrors(ctx, js, q)
		wg.Done()
	}()
	go func() {
		handleDelivereds(ctx, js, q)
		wg.Done()
	}()
	go func() {
		handleSends(ctx, js, q)
		wg.Done()
	}()
	wg.Wait()
}

func handleSends(ctx context.Context, js nats.JetStreamContext, q *sq.Queries) {
	con := utils.MustGetPullSubscriber(js, "kannon.sending", "kannon-stats-sending-logs")
	for {
		msgs, err := con.Fetch(10, nats.MaxWait(10*time.Second))
		if err != nil {
			if err != nats.ErrTimeout {
				logrus.Errorf("error fetching messages: %v", err)
			}
			continue
		}
		for _, msg := range msgs {
			sendMsg := pb.EmailToSend{}
			err = proto.Unmarshal(msg.Data, &sendMsg)
			if err != nil {
				logrus.Errorf("cannot marshal message %v", err.Error())
			} else {
				logrus.Printf("[✅ bump] %v %v - %v", sendMsg.To, sendMsg.MessageId)
				err := q.InsertDelivered(ctx, sq.InsertDeliveredParams{
					Email:     sendMsg.To,
					MessageID: sendMsg.MessageId,
				})
				if err != nil {
					logrus.Errorf("Cannot insert sent: %v", err)
				}
			}
			if err := msg.Ack(); err != nil {
				logrus.Errorf("Cannot hack msg to nats: %v\n", err)
			}
		}
	}
}

func handleErrors(ctx context.Context, js nats.JetStreamContext, q *sq.Queries) {
	con := utils.MustGetPullSubscriber(js, "kannon.stats.error", "kannon-stats-error-logs")
	for {
		msgs, err := con.Fetch(10, nats.MaxWait(10*time.Second))
		if err != nil {
			if err != nats.ErrTimeout {
				logrus.Errorf("error fetching messages: %v", err)
			}
			continue
		}
		for _, msg := range msgs {
			errMsg := pb.Error{}
			err = proto.Unmarshal(msg.Data, &errMsg)
			if err != nil {
				logrus.Errorf("cannot marshal message %v", err.Error())
			} else {
				logrus.Printf("[🛑 bump] %v %v - %v", errMsg.Email, errMsg.MessageId, errMsg.Msg)
				err := q.InsertBounced(ctx, sq.InsertBouncedParams{
					Email:     errMsg.Email,
					MessageID: errMsg.MessageId,
					ErrCode:   0, // TODO
					ErrMsg:    errMsg.Msg,
					Timestamp: errMsg.Timestamp.AsTime(),
				})
				if err != nil {
					logrus.Errorf("Cannot insert bounced: %v", err)
				}
			}
			if err := msg.Ack(); err != nil {
				logrus.Errorf("Cannot hack msg to nats: %v\n", err)
			}
		}
	}
}

func handleDelivereds(ctx context.Context, js nats.JetStreamContext, q *sq.Queries) {
	con := utils.MustGetPullSubscriber(js, "kannon.stats.delivered", "kannon-stats-delivered-logs")
	for {
		msgs, err := con.Fetch(10, nats.MaxWait(10*time.Second))
		if err != nil {
			if err != nats.ErrTimeout {
				logrus.Errorf("error fetching messages: %v", err)
			}
			continue
		}
		for _, msg := range msgs {
			deliveredMsg := pb.Delivered{}
			err = proto.Unmarshal(msg.Data, &deliveredMsg)
			if err != nil {
				logrus.Errorf("cannot marshal message %v", err.Error())
			} else {
				logrus.Printf("[🚀 delivered] %v %v", deliveredMsg.Email, deliveredMsg.MessageId)
				err := q.InsertDelivered(ctx, sq.InsertDeliveredParams{
					Email:     deliveredMsg.Email,
					MessageID: deliveredMsg.MessageId,
					Timestamp: deliveredMsg.Timestamp.AsTime(),
				})
				if err != nil {
					logrus.Errorf("Cannot insert delivered: %v", err)
				}
			}
			if err := msg.Ack(); err != nil {
				logrus.Errorf("Cannot hack msg to nats: %v", err)
			}
		}
	}
}