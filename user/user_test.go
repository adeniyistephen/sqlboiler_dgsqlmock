package user

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestInsertUser(t *testing.T) {
	// Mock DB instance by sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

  	// Inject mock instance into boil.
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)

	query := regexp.QuoteMeta("INSERT INTO `user` (`firstname`,`lastname`) VALUES (?,?)")

	mock.ExpectExec(query).WithArgs("smith", "rowe").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(query).WithArgs("smith", "rowe").WillReturnError(errors.New("internal server error"))
}
