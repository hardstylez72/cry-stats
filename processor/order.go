package processor

import (
	"context"
	"log"
	"math"

	"github.com/hardstylez72/cry-pay/internal/order"
	"github.com/pkg/errors"
)

type orderForProcess struct {
	OrderId   string
	AccountId string
}

func (p *Processor) findOrder(ctx context.Context, rec *recForProcess) (res *orderForProcess, err error) {
	o, isFound, err := p.ordersRepo.FindOrder(ctx, order.FindOrderReq{
		Addr:           rec.Addr,
		Net:            rec.Net,
		IncomeExpected: rec.IncomeReceived,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ordersRepo find")
	}

	if !isFound {
		return nil, nil
	}

	if o.Status != order.StatusCreated {
		return nil, nil
	}

	return &orderForProcess{
		OrderId:   o.Id,
		AccountId: o.AccountId,
	}, nil
}

func (p *Processor) processOrder(ctx context.Context, txId string, rec *recForProcess, o *orderForProcess) (err error) {
	err = p.ordersRepo.SetStatus(ctx, o.OrderId, order.StatusInProgress)
	if err != nil {
		return errors.Wrap(err, "ordersRepo SetStatus")
	}

	defer func() {
		var resultStatus string
		if err != nil {
			resultStatus = order.StatusError
		} else {
			resultStatus = order.StatusProcessed
		}

		err = p.ordersRepo.SetStatus(ctx, o.OrderId, resultStatus)
		if err != nil {
			log.Printf("Err @ defer: %+v \n", err)
		}
	}()

	_, err = p.accountsRepo.AddFundsById(ctx, o.AccountId, rec.IncomeReceived)
	if err != nil {
		return errors.Wrap(err, "accountsRepo add funds by id")
	}

	err = p.ordersRepo.Fill(ctx, o.OrderId, &order.FillReq{
		ConfirmedTxId:  txId,
		Meta:           "",
		ExchangeRate:   exchangeRate,
		IncomeReceived: rec.IncomeReceived,
	})
	if err != nil {
		return errors.Wrap(err, "ordersRepo Fill")
	}

	return nil
}

func incomeToFunds(income int, exchangeRate float32) int {
	res := float32(income) / exchangeRate

	return int(math.Round(float64(res)))
}
