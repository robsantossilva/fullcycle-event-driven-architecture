package gateway

import "github.com/robsantossilva/fullcycle-event-driven-architecture/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
