package postgres

import (
	"time"

	"github.com/bcmmbaga/go-graphql/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/theckman/yacspin"
)

func createSchema(cn *pg.Conn) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Running Migration \n",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	for _, model := range []interface{}{
		(*models.User)(nil),
		(*models.Meetup)(nil)} {
		err := cn.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		})

		if err != nil {
			spinner.Stop()
			return err
		}
	}

	spinner.Stop()

	return nil
}

func createUsers(cn *pg.Conn) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Creating Users \n",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	users := []models.User{
		{
			Username: "golangdev",
			Email:    "golangdev@gmail.com",
		},
		{
			Username: "sharkstewart",
			Email:    "shark.stewart.gmail.com",
		},
	}

	for _, user := range users {
		if _, err := cn.Model(&user).WherePK().Insert(); err != nil {
			return err
		}
	}

	spinner.Stop()

	return nil
}

func createMeeutps(cn *pg.Conn) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Creating Meetups \n",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	meetups := []models.Meetup{
		{
			ID:          "1",
			Name:        "AI , Data Science& Machine Learning CodeSeeker's",
			Description: "Better late than never - Ride the AI ,Data Science , Machine Learning revolution wave",
			UserID:      "shark.stewart.gmail.com",
		},
		{
			ID:          "2",
			Name:        "GoJakarta",
			Description: "A meetup group to discuss the Go Programming Language in Jakarta.",
			UserID:      "golangdev@gmail.com",
		},
		{
			ID:          "3",
			Name:        "Golang Bangalore",
			Description: "A meetup group to discuss the Go Programming Language (https://golang.org/).",
			UserID:      "golangdev@gmail.com",
		},
	}

	for _, meetup := range meetups {
		if _, err := cn.Model(&meetup).WherePK().Insert(); err != nil {
			return err
		}
	}

	spinner.Stop()

	return nil
}

func seed(cn *pg.Conn) error {

	if err := createSchema(cn); err != nil {
		return err
	}

	if err := createUsers(cn); err != nil {
		return err
	}

	if err := createMeeutps(cn); err != nil {
		return err
	}

	return nil
}
