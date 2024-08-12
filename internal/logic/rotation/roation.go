package rotation

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func (s *sRotation) Create(ctx context.Context, in model.RotationInput) error {

	return dao.Rotation.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.Rotation.Ctx(ctx).Data(in).Insert()
		if err != nil {
			return err
		}

		return nil
	})

}

func New() *sRotation {
	return &sRotation{}
}
