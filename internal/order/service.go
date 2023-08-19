package order

import (
	"github.com/hardstylez72/cry-pay/internal/account"
)

type Service struct {
	repo         *Repository
	accountsRepo *account.Repository
}

type InitConfig struct {
	OrdersRepo   *Repository
	AccountsRepo *account.Repository
}

func NewService(initConfig *InitConfig) (*Service, error) {
	return &Service{
		repo:         initConfig.OrdersRepo,
		accountsRepo: initConfig.AccountsRepo,
	}, nil
}
