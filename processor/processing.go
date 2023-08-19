package processor

import (
	"context"
	"fmt"
	"time"

	"github.com/hardstylez72/cry-pay/processor/history"
	"github.com/pkg/errors"
)

const exchangeRate = 1e4 // TODO fix
const orderExpirationDuration = (24 * time.Hour)

type ProcessingTask struct {
	processor *Processor
}

func NewProcessingTask(processor *Processor) *ProcessingTask {
	return &ProcessingTask{
		processor: processor,
	}
}

func (s *ProcessingTask) RunTask(ctx context.Context, interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {

		err := s.processor.Process(ctx)
		if err != nil {
			fmt.Printf("Err sched RunTask loop run %+v \n", err)
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return
		}
	}
}

func (p *Processor) ProcessRec(ctx context.Context, rec *TxRec) error {
	// TODO override ctx
	txId, err := p.historyRepo.Append(ctx, history.AppendTxReq{
		Addr:           rec.Addr,
		Net:            rec.Net,
		Hash:           rec.Hash,
		IncomeReceived: rec.IncomeReceived,
		IncomeTime:     rec.IncomeTime,
	})
	if err != nil {
		return errors.Wrap(err, "historyRepo append")
	}

	err = p.findAndProcessOrder(ctx, txId, &recForProcess{
		Addr:           rec.Addr,
		Net:            rec.Net,
		IncomeReceived: float64(rec.IncomeReceived) / 10e5,
	})
	if err != nil {
		return errors.Wrap(err, "findAndProcessOrder")
	}

	err = p.historyRepo.MarkAsProcessed(ctx, txId)
	if err != nil {
		return errors.Wrap(err, "historyRepo MarkAsProcessed")
	}

	return nil
}

func (p *Processor) Process(ctx context.Context) error {
	txList, err := p.historyRepo.GetUnprocessed(ctx)
	if err != nil {
		return errors.Wrap(err, "historyRepo GetUnprocessed")
	}

	for i := range txList {
		tx := txList[i]
		err = p.findAndProcessOrder(ctx, tx.Id, &recForProcess{
			Addr:           tx.Addr,
			Net:            tx.Net,
			IncomeReceived: float64(tx.IncomeReceived) / (10e5),
		})
		if err != nil {
			return errors.Wrap(err, "findAndProcessOrder")
		}

		err = p.historyRepo.MarkAsProcessed(ctx, tx.Id)
		if err != nil {
			return errors.Wrap(err, "historyRepo MarkAsProcessed")
		}
	}

	return nil
}

type recForProcess struct {
	Addr           string
	Net            string
	IncomeReceived float64
	CreatedSince   *time.Time
}

func (p *Processor) findAndProcessOrder(ctx context.Context, txId string, rec *recForProcess) (err error) {
	order, err := p.findOrder(ctx, rec)
	if err != nil {
		return errors.Wrap(err, "findOrder")
	}

	if order == nil {
		return nil
	}

	err = p.processOrder(ctx, txId, rec, order)
	if err != nil {
		return errors.Wrap(err, "processOrder")
	}

	return nil
}
