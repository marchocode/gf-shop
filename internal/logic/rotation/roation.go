package rotation

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type sRotationInfo struct{}

func init() {
	service.RegisterRotationInfo(New())
}

func (s *sRotationInfo) Create(ctx context.Context, in model.RotationInfoCreateInput) error {

	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.RotationInfo.Ctx(ctx).Data(in).Insert()

		if err != nil {
			return err
		}

		return nil
	})

}

func (s *sRotationInfo) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.RotationInfo.Ctx(ctx).Where(dao.RotationInfo.Columns().Id, id).Delete()
	return
}

func (s *sRotationInfo) Update(ctx context.Context, in model.RotationInfoUpdateInput) (err error) {

	err = dao.RotationInfo.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.RotationInfo.Ctx(ctx).Data(in).FieldsEx(dao.RotationInfo.Columns().Id).Where(dao.RotationInfo.Columns().Id, in.Id).Update()
		return err
	})
	return
}

func New() *sRotationInfo {
	return &sRotationInfo{}
}
