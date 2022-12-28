package database

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"

	"github.com/moisesPompilio/calculadora_taxa/internal/order/entity"
)

type OrderRepositoryTestSuit struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuit) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec("CREATE TABLE orders (id varchar(255) PRIMARY KEY, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL)")
	suite.NoError(err)
	suite.Db = db
}

func (suite *OrderRepositoryTestSuit) TearDowTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {

	suite.Run(t, new(OrderRepositoryTestSuit))

}

func (suite *OrderRepositoryTestSuit) TestGivenAnOrder_WhenSave_TheShouldSaverOrder() {
	order, err := entity.NewOrder("1jk23", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	log.Println(suite.Db)
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price from orders where id = ?", order.Id).
		Scan(&orderResult.Id, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.Id, orderResult.Id)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
