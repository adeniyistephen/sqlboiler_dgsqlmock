package user

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertUser(t *testing.T) {
	// Mock DB instance by sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	type fileds struct {
		conn *sql.DB
	}

	type args struct {
		user user
	}

	tests := []struct {
		name    string
		fields  fileds
		args    args
		wantID  int
		wantErr bool
	}{
		{
			name: "Test query error",
			fields: fileds{
				conn: db,
			},
			args: args{user: user{
				FirstName: "John",
				LastName:  "Doe",
			}},
			wantID:  0,
			wantErr: true,
		},
		{
			name: "Test unique",
			fields: fileds{
				conn: db,
			},
			args: args{user: user{
				FirstName: "John",
				LastName:  "Doe",
			}},
			wantID:  0,
			wantErr: true,
		},
		{
			name: "Test failed entry",
			fields: fileds{
				conn: db,
			},
			args: args{user: user{
				FirstName: "John",
				LastName:  "Doe",
			}},
			wantID:  0,
			wantErr: true,
		},
		{
			name: "Test success",
			fields: fileds{
				conn: db,
			},
			args: args{user: user{
				FirstName: "John",
				LastName:  "Doe",
			}},
			wantID:  1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		app := Api{db: tt.fields.conn}

		if tt.name == "Test query error" {
			mock.ExpectQuery("SELECT (.+) FROM `user` WHERE .*").WillReturnError(errors.New("error trying to validate the article being added"))
			err = app.InsertUser()
			fmt.Printf("%+v\n", err)
		}
		if tt.name == "Test unique" {
			row := sqlmock.NewRows([]string{"id", "firstname", "lastname"}).
				AddRow(1, "John", "Doe")
			mock.ExpectQuery("SELECT (.+) FROM `user_table1` WHERE .*").WillReturnRows(row)
			err = app.InsertUser()
			fmt.Printf("%+v\n", err)
		}
		if tt.name == "Test failed entry" {

			row := sqlmock.NewRows([]string{"id", "firstname", "lastname"})
			mock.ExpectQuery("SELECT (.+) FROM `user_table1` WHERE .*").WillReturnRows(row)

			mock.ExpectExec("^INSERT INTO `user_table1`. *$").
				WithArgs(tt.args.user.FirstName, tt.args.user.LastName).
				WillReturnError(errors.New("error trying to insert a new entry"))

			err = app.InsertUser()
			fmt.Printf("%+v\n", err)
		}
		if tt.name == "Test success" {
			// Pass empty records
			row := sqlmock.NewRows([]string{"id", "firstname", "lastname"})
			mock.ExpectQuery("SELECT (.+) FROM `user_table1` WHERE .*").WillReturnRows(row)

			// Pass insert
			mock.ExpectExec("^INSERT INTO `article`.*$").
				WithArgs(tt.args.user.FirstName, tt.args.user.LastName).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err = app.InsertUser()
			if err != nil {
				t.Errorf("InsertUser() error = %v", err)
			}
			t.Log(t, tt.wantID, "failed at: \"%v\"", tt.name)
		}

	}

}
