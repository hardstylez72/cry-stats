package account

import (
	"context"

	"github.com/hardstylez72/cry-pay/internal/config"
	v1 "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
)

type ControllerAdapter struct {
	repo *Repository
	v1.UnimplementedFundsServiceServer
}

func NewControllerAdapter(repo *Repository) *ControllerAdapter {
	return &ControllerAdapter{
		repo: repo,
	}
}

const funds = 10.0
const defaultTaskPrice = 0.2

func (c *ControllerAdapter) GetAccount(ctx context.Context, req *v1.GetAccountReq) (*v1.GetAccountRes, error) {
	acc, err := c.repo.GetAccount(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	r := &v1.GetAccountRes{
		Account: &v1.Account{
			Id:        acc.Id,
			Login:     acc.Login,
			Status:    acc.Status,
			Funds:     acc.Funds,
			TaskPrice: acc.TaskPrice,
		},
	}
	if acc.Promo.Valid {
		r.Account.Promo = &acc.Promo.String
	}

	return r, nil
}

func (c *ControllerAdapter) CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (*v1.CreateAccountResp, error) {

	var price = defaultTaskPrice
	var funds = funds
	if config.CFG.Standalone {
		price = 0
		funds = 9999
	}

	err := c.repo.CreateAccount(ctx, req.Id, req.Login, funds, price)
	if err != nil {
		return nil, err
	}

	return &v1.CreateAccountResp{}, nil
}

func (c *ControllerAdapter) AccountExist(ctx context.Context, req *v1.AccountExistReq) (*v1.AccountExistResp, error) {

	exist, err := c.repo.AccountExist(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.AccountExistResp{
		Exist: *exist,
	}, nil
}

func (c *ControllerAdapter) GetFunds(ctx context.Context, req *v1.GetFundsReq) (*v1.GetFundsResp, error) {
	funds, isFound, err := c.repo.GetFundsByLogin(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	return &v1.GetFundsResp{
		FundsLeft: funds,
		IsFound:   isFound,
	}, nil
}

func (c *ControllerAdapter) TaskCompleted(ctx context.Context, req *v1.TaskCompletedReq) (*v1.TaskCompletedRes, error) {
	acc, err := c.repo.GetAccount(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	price := acc.TaskPrice

	exist, err := c.repo.TaskExist(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	if exist {
		return &v1.TaskCompletedRes{}, nil
	}

	if err := c.repo.AddTaskRecord(ctx, &TaskHistoryRecord{
		TaskId:    req.TaskId,
		TaskType:  req.TaskType,
		UserId:    req.UserId,
		ProcessId: req.ProcessId,
		Price:     price,
	}); err != nil {
		return nil, err
	}

	err = c.repo.DecrementFundsByUserId(ctx, acc.Id, price)
	if err != nil {
		return nil, err
	}

	return &v1.TaskCompletedRes{}, nil
}

func (c *ControllerAdapter) UserTaskHistory(ctx context.Context, req *v1.UserTaskHistoryReq) (*v1.UserTaskHistoryRes, error) {
	records, err := c.repo.GetHistoryRecords(ctx, req.UserId, int(req.Offset))
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.TaskHistoryRecord, len(records))

	for i, r := range records {
		tmp[i] = &v1.TaskHistoryRecord{
			ProcessId: r.ProcessId,
			TaskId:    r.TaskId,
			TaskType:  r.TaskType,
			TaskPrice: r.Price,
		}
	}

	return &v1.UserTaskHistoryRes{
		Total:   0,
		Records: tmp,
	}, nil
}

func (c *ControllerAdapter) AddPromo(ctx context.Context, req *v1.AddPromoReq) (*v1.AddPromoRes, error) {

	a, err := c.repo.GetAccount(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	if a.Promo.Valid {
		return &v1.AddPromoRes{
			Valid: false,
			Bonus: 0,
		}, nil
	}

	bonus := PromoBonus(req.Promo)
	if bonus == 0 {
		return &v1.AddPromoRes{
			Valid: false,
			Bonus: 0,
		}, nil
	}

	if err := c.repo.AddPromo(ctx, req.UserId, req.GetPromo()); err != nil {
		return nil, err
	}

	if _, err := c.repo.AddFundsById(ctx, req.UserId, bonus); err != nil {
		return nil, err
	}
	return &v1.AddPromoRes{
		Valid: true,
		Bonus: bonus,
	}, nil
}
