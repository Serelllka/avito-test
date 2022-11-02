package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestTransaction_Deposit(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	s := NewTransactionPostgres(db)

	for _, testCase := range []struct {
		name         string
		id           int
		item         dto.Deposit
		mockBehavior func(int, dto.Deposit)
		isErr        bool
	}{
		{
			name: "OK",
			item: dto.Deposit{
				ConsumerId:  1,
				Description: "just Deposit",
				Amount:      700,
			},
			id: 1,
			mockBehavior: func(id int, item dto.Deposit) {
				rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)

				mock.ExpectQuery(fmt.Sprintf("INSERT INTO %s", usersTransactionTable)).
					WithArgs(model.Deposit, item.ConsumerId, item.Amount, item.Description).WillReturnRows(rows)
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.id, testCase.item)

			ans, err := s.CreateDeposit(testCase.item)
			if testCase.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, ans)
			}
		})
	}
}
