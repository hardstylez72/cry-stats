package task

import (
	"context"
	"fmt"
	"time"

	"github.com/hardstylez72/cry-pay/internal/arbitrum"
	"github.com/hardstylez72/cry-pay/internal/config"
	"github.com/hardstylez72/cry-pay/internal/order"
	"github.com/hardstylez72/cry-pay/processor"
	"github.com/pkg/errors"
)

type ArbiscanImporter struct {
	arbiscanClient *arbitrum.Client
	processor      *processor.Processor
	offsetsRepo    *Repository
}

func NewArbiscanImporter(arbiscanClient *arbitrum.Client, processor *processor.Processor, offsetsRepo *Repository) *ArbiscanImporter {
	return &ArbiscanImporter{
		arbiscanClient: arbiscanClient,
		processor:      processor,
		offsetsRepo:    offsetsRepo,
	}
}

func (s *ArbiscanImporter) RunImportTask(ctx context.Context, interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		_, err := s.importFromArbiscan(ctx, &ArbiscanImportReq{Addr: config.CFG.ARBIUSDTWalletAddr, Net: order.TheOnlyNet})
		if err != nil {
			fmt.Printf("Err sched RunImportTask loop run %+v \n", err)
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return
		}
	}
}

type ArbiscanImportReq struct {
	Net  string
	Addr string
}

func (s *ArbiscanImporter) importFromArbiscan(ctx context.Context, req *ArbiscanImportReq) (*bool, error) {
	offsets, err := s.offsetsRepo.GetOffsets(ctx, &GetOffsetsReq{
		Net:  req.Net,
		Addr: req.Addr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "offsetsRepo GetOffsets")
	}

	currentOffset := offsets.TxOffset
	for {
		// use req.Net
		list, err := s.arbiscanClient.GetListTx(ctx, req.Net, req.Addr, currentOffset)
		if err != nil {
			return nil, errors.Wrap(err, "api GetListTx")
		}
		if list.TxOffsetIncrBy == 0 {
			// up-to-date, nothing to process
			return nil, nil
		}
		currentOffset += list.TxOffsetIncrBy

		err = s.processList(ctx, req.Net, req.Addr, list)
		if err != nil {
			return nil, errors.Wrap(err, "processList")
		}

		err = s.offsetsRepo.UpdateOffsets(ctx, &UpdateOffsetsReq{
			Net:            req.Net,
			Addr:           req.Addr,
			TxTotal:        list.TxTotal,
			TxOffsetIncrBy: list.TxOffsetIncrBy,
		})
		if err != nil {
			return nil, errors.Wrap(err, "offsetsRepo UpdateOffsets")
		}
	}
}

func (s *ArbiscanImporter) processList(ctx context.Context, net, addr string, list *arbitrum.TxList) error {
	for i := range list.Incomes {
		err := s.processor.ProcessRec(ctx, &processor.TxRec{
			Addr:           addr,
			Net:            net,
			Hash:           list.Incomes[i].Hash,
			IncomeReceived: list.Incomes[i].Received,
			IncomeTime:     list.Incomes[i].Time,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
