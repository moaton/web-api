package user_test

import (
	"context"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/repository/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := user.NewUserStorage(db)

	//	Case: Когда юзер есть

	rows := sqlmock.NewRows([]string{"id", "email", "name", "password"}).AddRow(5, "test", "test", "123")

	mock.ExpectQuery("^SELECT id, email, name, password FROM users*").WithArgs(5).WillReturnRows(rows)

	user, err := storage.GetUserById(context.Background(), 5)
	require.NoError(t, err)

	userEx := models.User{
		ID:       5,
		Email:    "test",
		Name:     "test",
		Password: "123",
	}
	assert.Equal(t, userEx, user, "they should be equal")

	//	Case: Когда юзера нет

	rows = sqlmock.NewRows([]string{"id", "email", "name", "password"}).AddRow(0, "", "", "")

	mock.ExpectQuery("^SELECT id, email, name, password FROM users*").WithArgs(5).WillReturnRows(rows)

	user, err = storage.GetUserById(context.Background(), 5)
	require.NoError(t, err)

	userEx = models.User{}
	assert.Equal(t, userEx, user, "they should be equal")
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := user.NewUserStorage(db)

	//	Case: Когда юзер есть

	rows := sqlmock.NewRows([]string{"id", "email", "name", "password"}).AddRow(5, "test", "test", "123")

	mock.ExpectQuery("^SELECT id, email, name, password FROM users*").WithArgs("test").WillReturnRows(rows)

	user, err := storage.GetUserByEmail(context.Background(), "test")
	require.NoError(t, err)

	userEx := models.User{
		ID:       5,
		Email:    "test",
		Name:     "test",
		Password: "123",
	}
	assert.Equal(t, userEx, user, "they should be equal")

	//	Case: Когда юзера нет

	rows = sqlmock.NewRows([]string{"id", "email", "name", "password"}).AddRow(0, "", "", "")

	mock.ExpectQuery("^SELECT id, email, name, password FROM users*").WithArgs("test2").WillReturnRows(rows)

	user, err = storage.GetUserByEmail(context.Background(), "test2")
	require.NoError(t, err)

	userEx = models.User{}
	assert.Equal(t, userEx, user, "they should be equal")
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := user.NewUserStorage(db)

	user := models.User{
		Email:    "test",
		Name:     "test",
		Password: "123",
	}
	row := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("INSERT INTO users*").WillReturnRows(row)

	id, err := storage.InsertUser(context.Background(), user)
	require.NoError(t, err)

	assert.Equal(t, int64(1), id, "they should be equal")
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := user.NewUserStorage(db)

	user := models.User{
		ID:       1,
		Email:    "test",
		Name:     "test",
		Password: "123",
	}

	mock.ExpectExec("UPDATE users*").WithArgs(user.Email, user.Name, user.Password, user.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.UpdateUser(context.Background(), user)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := user.NewUserStorage(db)

	mock.ExpectExec("DELETE FROM users*").WithArgs("test").WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.DeleteUser(context.Background(), "test")
	require.NoError(t, err)
}

func TestClose(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := user.NewUserStorage(db)

	var wg sync.WaitGroup
	wg.Add(1)
	storage.Close(&wg)

	rows := sqlmock.NewRows([]string{"id", "email", "name", "password"}).AddRow(5, "test", "test", "123")
	mock.ExpectQuery("^SELECT id, email, name, password FROM users*").WithArgs(5).WillReturnRows(rows)
	_, err = storage.GetUserById(context.Background(), 5)
	require.Error(t, err)
}
