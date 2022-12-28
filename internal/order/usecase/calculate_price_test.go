package usecase

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"

	"github.com/moisesPompilio/calculadora_taxa/internal/order/entity"
	"github.com/moisesPompilio/calculadora_taxa/internal/order/infra/database"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec("CREATE TABLE orders (id varchar(255) PRIMARY KEY, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL)")
	suite.NoError(err)
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseTestSuite) TearDowTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("dsS", 10.0, 2.23)
	suite.NoError(err)
	order.CalculateFinalPrice()
	calculateFinalPriceInput := OrderInputDTO{
		Id:    order.Id,
		Price: order.Price,
		Tax:   order.Tax,
	}
	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)
	suite.Equal(order.FinalPrice, output.FinalPrice)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
