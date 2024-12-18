package create_account

import (
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/entity"
	"github.com/robsantossilva/fullcycle-event-driven-architecture/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (us *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := us.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)

	err = us.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	output := &CreateAccountOutputDTO{
		ID: account.ID,
	}
	return output, nil
}
