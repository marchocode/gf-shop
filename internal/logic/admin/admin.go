package admin

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
	"gf-shop/utility"

	"github.com/gogf/gf/v2/util/grand"
)

func init() {
	service.RegisterAdmin(New())
}

type sAdmin struct{}

func (a *sAdmin) Create(ctx context.Context, in model.AdminInfoCreateInput) (*model.AdminInfoCreateOutput, error) {

	// salt
	salt := grand.S(10)

	in.UserSalt = salt
	in.Password = utility.EncryptPassword(in.Password, salt)

	id, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	return &model.AdminInfoCreateOutput{
		Id: int(id),
	}, nil
}

func New() *sAdmin {
	return new(sAdmin)
}
