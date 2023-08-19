package order

import (
	"context"

	"github.com/hardstylez72/cry-pay/internal/config"
	v1 "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
)

type ControllerAdapter struct {
	s *Service
	v1.UnimplementedOrdersServiceServer
}

func NewControllerAdapter(s *Service) *ControllerAdapter {
	return &ControllerAdapter{
		s: s,
	}
}

func (c *ControllerAdapter) CheckOrder(ctx context.Context, req *v1.CheckOrderReq) (*v1.CheckOrderResp, error) {
	res, err := c.s.repo.GetOrder(ctx, &GetOrderReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CheckOrderResp{
		Status: res.Status,
	}, nil
}

func (c *ControllerAdapter) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (*v1.CreateOrderResp, error) {
	res, err := c.s.Create(ctx, &CreateOrderReq{
		UserId:        req.UserId,
		Net:           req.Net,
		DesiredAmount: req.Am,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateOrderResp{
		Id:          res.Id,
		CoinAddrUrl: res.Addr,
		Am:          res.ExpectedAmount,
		ToWallet:    config.CFG.ARBIUSDTWalletAddr,
	}, nil
}

func (c *ControllerAdapter) GetOrderHistory(ctx context.Context, req *v1.GetOrderHistoryReq) (*v1.GetOrderHistoryRes, error) {

	records, err := c.s.repo.GetOrderHistory(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.Order, len(records))

	for i, r := range records {
		tmp[i] = &v1.Order{
			Id:          r.Id,
			Net:         r.Net,
			CoinAddrUrl: r.Addr,
			Status:      r.Status,
			CreatedAt:   r.CreatedAt.Unix(),
			ConfirmedAt: r.ConfirmedAt.Unix(),
			Am:          r.Amount,
			ToWallet:    config.CFG.ARBIUSDTWalletAddr,
		}
	}

	return &v1.GetOrderHistoryRes{
		Orders: tmp,
	}, nil
}
