package position

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

func init() {
	service.RegisterPosition(New())
}

type sPosition struct{}

func (a *sPosition) Create(ctx context.Context, in model.PositionInfoCreateInput) (*model.PositionInfoCreateOutput, error) {

	id, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	return &model.PositionInfoCreateOutput{
		Id: int(id),
	}, nil
}

func (s *sPosition) Update(ctx context.Context, in model.PositionInfoUpdateInput) (err error) {

	err = dao.PositionInfo.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PositionInfo.Ctx(ctx).Data(in).FieldsEx(dao.PositionInfo.Columns().Id).Where(dao.PositionInfo.Columns().Id, in.Id).Update()
		return err
	})
	return
}

func New() *sPosition {
	return new(sPosition)
}
