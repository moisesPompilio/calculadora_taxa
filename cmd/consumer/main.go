package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/moisesPompilio/calculadora_taxa/internal/order/infra/database"
	"github.com/moisesPompilio/calculadora_taxa/internal/order/usecase"
	"github.com/moisesPompilio/calculadora_taxa/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)

	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)

	qtdWorkers := 100
	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i)
	}

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUC := usecase.GetTotalUseCase{OrderRepository: repository}
		total, err := getTotalUC.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(total)
	})

	http.ListenAndServe(":8080", nil)

}
func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int) {
	for msg := range deliveryMessage {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Worker %d has processed order %s \n", workerID, outputDTO.Id)
		time.Sleep(1 * time.Second)
	}
}
