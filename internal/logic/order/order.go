package order

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sOrder struct{}

func New() *sOrder {
	return new(sOrder)
}

func init() {
	service.RegisterOrder(New())
}

func (s *sOrder) TodaysOrders(ctx context.Context) (*model.OrderTodaysNumOutput, error) {

	now := gtime.New("2023-02-07 15:50:15")
	num, err := dao.OrderInfo.Ctx(ctx).WhereBetween(dao.OrderInfo.Columns().CreatedAt, now.StartOfDay(), now.EndOfDay()).Count()

	if err != nil {
		return nil, err
	}

	return &model.OrderTodaysNumOutput{
		Nums: num,
		Date: now.Format("Y-m-d"),
	}, nil
}
