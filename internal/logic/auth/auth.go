package auth

import (
	"context"
	"gf-shop/internal/dao"
	"gf-shop/internal/model"
	"gf-shop/internal/model/entity"
	"gf-shop/internal/service"
	"gf-shop/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return new(sAuth)
}

func (a *sAuth) Login(ctx context.Context, in model.AdminLoginInput) error {

	adminEntity := new(entity.AdminInfo)
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(adminEntity)

	if err != nil {
		return err
	}

	pas := utility.EncryptPassword(in.Password, adminEntity.UserSalt)

	if pas != adminEntity.Password {
		return gerror.New("用户名或密码不正确")
	}

	// 登陆成功
	if err := service.Session().SetUser(ctx, adminEntity); err != nil {
		return err
	}

	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:   adminEntity.Id,
		Name: adminEntity.Name,
	})

	return nil
}
