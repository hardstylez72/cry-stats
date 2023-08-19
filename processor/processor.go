package processor

import (
	"time"

	"github.com/hardstylez72/cry-pay/internal/account"
	"github.com/hardstylez72/cry-pay/internal/order"
	"github.com/hardstylez72/cry-pay/processor/history"
)

type Processor struct {
	historyRepo  *history.Repository
	ordersRepo   *order.Repository
	accountsRepo *account.Repository
}

type TxRec struct {
	Addr           string
	Net            string
	Hash           string
	IncomeReceived int
	IncomeTime     time.Time
}

type InitConfig struct {
	HistoryRepo  *history.Repository
	OrdersRepo   *order.Repository
	AccountsRepo *account.Repository
}

func New(initConfig *InitConfig) (*Processor, error) {
	return &Processor{
		historyRepo:  initConfig.HistoryRepo,
		ordersRepo:   initConfig.OrdersRepo,
		accountsRepo: initConfig.AccountsRepo,
	}, nil
}
