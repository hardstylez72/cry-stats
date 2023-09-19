package order

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/rand"
)

type CreateOrderReq struct {
	UserId        string
	Net           string
	DesiredAmount int64
}

type CreateOrderRes struct {
	Id             string
	Addr           string
	ExpectedAmount float64
}

func (s *Service) Create(ctx context.Context, req *CreateOrderReq) (*CreateOrderRes, error) {
	if req.DesiredAmount < 0 {
		return nil, errors.New("Cannot create order with negative value")
	}

	_, err := s.accountsRepo.GetAccount(ctx, req.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "@ get account")
	}

	expectedAddr, addrFound, _ := getAddrByNet(req.Net)
	if !addrFound {
		return nil, errors.New("Payment receiver address not set")
	}

	var am float64
	for {
		am = GenerateAm(req.DesiredAmount)
		uniq, err := s.repo.AmountUniq(ctx, am)
		if err != nil {
			return nil, err
		}
		if uniq {
			break
		}
	}

	insertRes, err := s.repo.InsertOrder(ctx, &InsertOrderReq{
		AccountId:      req.UserId,
		Net:            req.Net,
		IncomeExpected: am,
		Addr:           expectedAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "@ InsertOrder")
	}

	return &CreateOrderRes{
		Id:             insertRes.Id,
		Addr:           expectedAddr,
		ExpectedAmount: am,
	}, nil
}

func GenerateAm(base int64) float64 {
	seed := rand.NewSource(uint64(time.Now().UnixMicro()))

	first := strconv.Itoa(int(base))

	first = first + ".0"
	precision := 1
	i := 0
	for i < precision {
		rInt := int(rand.New(seed).Int63n(9))
		if rInt == 0 {
			continue
		}
		first = first + strconv.Itoa(rInt)
		i++
	}

	out, _ := StringToFloat(first)
	return out
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 10)
}
