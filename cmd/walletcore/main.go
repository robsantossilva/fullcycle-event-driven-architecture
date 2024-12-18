package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/database"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/event"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/usecase/create_account"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/usecase/create_client"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/usecase/create_transaction"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/web"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/web/webserver"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/pkg/events"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher := events.NewEventDispatcher()
	//eventDispatcher.Register(transactionCreatedEvent.GetName(), )

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	//transactionDb := database.NewTransactionDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/client", clientHandler.CreateClient)
	webserver.AddHandler("/account", accountHandler.CreateAccount)
	webserver.AddHandler("/transaction", transactionHandler.CreateTransaction)

	fmt.Println("Server running on port 3000")
	webserver.Start()
}
