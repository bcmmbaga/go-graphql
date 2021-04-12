package postgres

import (
	"time"

	"github.com/bcmmbaga/go-graphql/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/theckman/yacspin"
)

func createSchema(db *pg.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        "Migrating Schema",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	for _, model := range []interface{}{
		(*models.User)(nil),
		(*models.Meetup)(nil)} {
		err := db.Model(&model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		})

		if err != nil {
			return err
		}
	}

	spinner.Stop()

	return nil
}
