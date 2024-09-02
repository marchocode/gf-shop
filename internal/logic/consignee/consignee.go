package consignee

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sConsignee struct{}

func init() {
	service.RegisterConsignee(New())
}

func New() *sConsignee {
	return new(sConsignee)
}

func (s *sConsignee) Craete(ctx context.Context, input *model.ConsigneeCreateInput) (out *model.ConsigneeCreateOutput, err error) {

	tx, err := g.DB().Begin(ctx)

	if err != nil {
		return
	}

	defer func() {

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	// if default.
	if input.IsDefault == 1 {
		// update other to 0
		dao.ConsigneeInfo.Ctx(ctx).TX(tx).
			Data(g.Map{
				dao.ConsigneeInfo.Columns().IsDefault: 0}).
			WherePri(input.UserId).Update()
	}

	// insert
	id, err := dao.ConsigneeInfo.Ctx(ctx).TX(tx).InsertAndGetId(input)

	if err != nil {
		return
	}

	out = &model.ConsigneeCreateOutput{
		Id: int(id),
	}

	return
}
