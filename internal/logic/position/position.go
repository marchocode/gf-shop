package position

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
)

func init() {
	service.RegisterPosition(New())
}

type sPosition struct{}

func (a *sPosition) Create(ctx context.Context, in model.PositionInfoCreateInput) (*model.PositionInfoCreateOutput, error) {

	id, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	return &model.PositionInfoCreateOutput{
		Id: int(id),
	}, nil
}

func New() *sPosition {
	return new(sPosition)
}
