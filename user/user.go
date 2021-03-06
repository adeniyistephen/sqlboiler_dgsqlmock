package user

import (
	"context"
	"database/sql"
	"fmt"
	"mysql/example/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type user struct {
	ID        int
	FirstName string
	LastName  string
}

type Api struct {
	db *sql.DB
}

func (a *Api) InsertUser() error {

	u := user{
		FirstName: "John",
		LastName:  "Doe",
	}

	u1 := models.UserTable1{Firstname: u.FirstName, Lastname: u.LastName}
	u2 := models.UserTable2{Firstname: u.FirstName, Lastname: u.LastName}
	u3 := models.UserTable3{Firstname: u.FirstName, Lastname: u.LastName}

	_, err := a.db.BeginTx(context.Background(), nil)
	dieIf(err)

	err = u1.Insert(context.Background(), a.db, boil.Infer())
	dieIf(err)
	fmt.Println("Inserted user", u.ID)
	err = u2.Insert(context.Background(), a.db, boil.Infer())
	dieIf(err)
	fmt.Println("Inserted user", u2.ID)
	err = u3.Insert(context.Background(), a.db, boil.Infer())
	dieIf(err)
	fmt.Println("Inserted user", u3.ID)

	// Delete users from tables
	user1, err := models.UserTable1s().DeleteAll(context.Background(), a.db)
	dieIf(err)
	fmt.Println("Deleted user", user1)
	user2, err := models.UserTable2s().DeleteAll(context.Background(), a.db)
	dieIf(err)
	fmt.Println("Deleted user", user2)
	user3, err := models.UserTable3s().DeleteAll(context.Background(), a.db)
	dieIf(err)
	fmt.Println("Deleted user", user3)

	return nil
}


func dieIf(err error) error {
	if err != nil {
		return err
	}
	return nil
}
