package service

import "errors"

// AuthRequest request 结构体 用于接口入参效验
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

// CheckAuth 使用客户端传入的认证信息作为筛选条件获取数据行 根据是否取到认证信息 ID 判定认证信息 ID 是否存在
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist.")
}
