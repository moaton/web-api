package revenue_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/repository/revenue"
	"github.com/moaton/web-api/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRevenues(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := revenue.NewRevenueStorage(db)
	logger.SetLogger(false)

	var limit int64 = 10
	var offset int64 = 0

	//	Case: count error
	var countError error = errors.New("sql: no result")
	mock.ExpectQuery("^SELECT (.+) FROM revenues*").WithArgs(limit, offset).WillReturnError(countError)
	_, _, err = storage.GetRevenues(context.Background(), limit, offset)
	require.Error(t, err)

	//	Case: rows error
	rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("^SELECT (.+) FROM revenues*").WithArgs(limit, offset).WillReturnRows(rowsCount)

	var rowsError error = errors.New("sql: no result")
	mock.ExpectQuery("^SELECT id, title, description, amount, type, createdat, updatedat FROM revenues*").WithArgs(limit, offset).WillReturnError(rowsError)
	_, _, err = storage.GetRevenues(context.Background(), limit, offset)
	require.Error(t, err)

	//	Case: 1 element
	now := time.Now()
	revenuesExpected := []models.Revenue{
		{
			ID:          1,
			Title:       "test",
			Description: "test",
			Amount:      1,
			Type:        "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	rowsCount = sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("^SELECT (.+) FROM revenues*").WithArgs(limit, offset).WillReturnRows(rowsCount)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "amount", "type", "createdat", "updatedat"}).AddRow(1, "test", "test", 1, "test", now, now)
	mock.ExpectQuery("^SELECT id, title, description, amount, type, createdat, updatedat FROM revenues*").WithArgs(limit, offset).WillReturnRows(rows)

	revenues, count, err := storage.GetRevenues(context.Background(), limit, offset)
	require.NoError(t, err)

	assert.Equal(t, revenuesExpected, revenues, "they should be equal")
	assert.Equal(t, int64(1), count, "they should be equal")

	//	Case: rows.Next error
	revenuesExpected = []models.Revenue{
		{
			ID:          1,
			Title:       "test",
			Description: "test",
			Amount:      1,
			Type:        "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	rowsCount = sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("^SELECT (.+) FROM revenues*").WithArgs(limit, offset).WillReturnRows(rowsCount)

	rows = sqlmock.NewRows([]string{"id", "title", "description", "amount", "type", "createdat", "updatedat"}).AddRow(1, "test", "test", 1, "test", now, now).AddRow(1, nil, nil, 1, nil, now, now)
	mock.ExpectQuery("^SELECT id, title, description, amount, type, createdat, updatedat FROM revenues*").WithArgs(limit, offset).WillReturnRows(rows)
	revenues, count, err = storage.GetRevenues(context.Background(), limit, offset)
	require.NoError(t, err)

	assert.Equal(t, revenuesExpected, revenues, "they should be equal")
	assert.Equal(t, int64(2), count, "they should be equal")
}

func TestGetRevenueById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := revenue.NewRevenueStorage(db)
	logger.SetLogger(false)

	var id int64 = 10
	now := time.Now()
	//	Case: rows error
	var rowsError error = errors.New("sql: no result")
	mock.ExpectQuery("^SELECT id, title, description, amount, type, createdat, updatedat FROM revenues*").WithArgs(id).WillReturnError(rowsError)
	_, err = storage.GetRevenueById(context.Background(), id)
	require.Error(t, err)

	//	Case: 1 element
	revenuesExpected := models.Revenue{
		ID:          1,
		Title:       "test",
		Description: "test",
		Amount:      1,
		Type:        "test",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "amount", "type", "createdat", "updatedat"}).AddRow(1, "test", "test", 1, "test", now, now)
	mock.ExpectQuery("^SELECT id, title, description, amount, type, createdat, updatedat FROM revenues*").WithArgs(id).WillReturnRows(rows)

	revenues, err := storage.GetRevenueById(context.Background(), id)
	require.NoError(t, err)

	assert.Equal(t, revenuesExpected, revenues, "they should be equal")
}

func TestInsertRevenue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := revenue.NewRevenueStorage(db)

	//	Case: insert
	revenue := models.Revenue{
		Title:       "test",
		Description: "test",
		Amount:      1,
		Type:        "test",
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^INSERT INTO revenues*").WillReturnRows(rows)

	id, err := storage.InsertRevenue(context.Background(), revenue)
	require.NoError(t, err)

	assert.Equal(t, int64(1), id, "they should be equal")

	//	Case: error
	var rowsError error = errors.New("sql: deadlock")
	mock.ExpectQuery("^INSERT INTO revenues*").WillReturnError(rowsError)

	_, err = storage.InsertRevenue(context.Background(), revenue)
	require.Error(t, err)
}

func TestUpdateRevenue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := revenue.NewRevenueStorage(db)

	//	Case: update
	revenue := models.Revenue{
		Title:       "test",
		Description: "test",
		Amount:      1,
		Type:        "test",
	}

	mock.ExpectExec("UPDATE revenues*").WithArgs(revenue.Title, revenue.Description, revenue.Amount, revenue.Type, time.Now().UTC(), revenue.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.UpdateRevenue(context.Background(), revenue)
	require.NoError(t, err)

	//	Case: error
	var updateError error = errors.New("sql: problem")
	mock.ExpectExec("UPDATE revenues*").WithArgs(revenue.Title, revenue.Description, revenue.Amount, revenue.Type, time.Now().UTC(), revenue.ID).WillReturnError(updateError)
	err = storage.UpdateRevenue(context.Background(), revenue)
	require.Error(t, err)
}

func TestDeleteRevenue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := revenue.NewRevenueStorage(db)

	//	Case: delete
	var id int64 = 1

	mock.ExpectExec("DELETE FROM revenues*").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.DeleteRevenue(context.Background(), id)
	require.NoError(t, err)

	//	Case: error
	var updateError error = errors.New("sql: problem")
	mock.ExpectExec("DELETE FROM revenues*").WithArgs(id).WillReturnError(updateError)
	err = storage.DeleteRevenue(context.Background(), id)
	require.Error(t, err)
}

func TestClose(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	storage := revenue.NewRevenueStorage(db)

	var wg sync.WaitGroup
	wg.Add(1)
	storage.Close(&wg)

	var limit int64 = 10
	var offset int64 = 0

	var countError error = errors.New("sql: no result")
	mock.ExpectQuery("^SELECT (.+) FROM revenues*").WithArgs(limit, offset).WillReturnError(countError)
	_, _, err = storage.GetRevenues(context.Background(), limit, offset)
	require.Error(t, err)
}
