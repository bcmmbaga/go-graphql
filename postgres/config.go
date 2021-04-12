package postgres

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type queryHook struct{}

func (*queryHook) BeforeQuery(ctx context.Context, query *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (*queryHook) AfterQuery(ctx context.Context, query *pg.QueryEvent) error {
	formattedQuery, err := query.FormattedQuery()
	if err != nil {
		return err
	}

	logrus.WithField("query", string(formattedQuery)).Infoln()
	return nil
}

// Connect connects to a database.
func Connect(addr, user, password, dbName string) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr: addr,
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			cn.AddQueryHook(&queryHook{})

			if err := seed(cn); err != nil {
				return nil
			}

			return nil
		},
		User:            user,
		Password:        password,
		Database:        dbName,
		ApplicationName: "go-graphql",
	})

	db.Ping(context.Background())

	return db
}
