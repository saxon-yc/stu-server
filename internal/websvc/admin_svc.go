package websvc

import (
	"fmt"
	"student-server/internal/model"
	errorcode "student-server/pkg/errors"
	"student-server/pkg/util"

	"github.com/gin-gonic/gin"
)

type AdminManager interface {
	LoginWeb(c *gin.Context, params model.LoginRequest) (output model.UserDb, errString error)
	QueryUser(params model.DupuserRequest) (result model.UserDb, err error)
}

func (svc *WebService) QueryUser(params model.DupuserRequest) (result model.UserDb, err error) {
	result, err = svc.dbService.FindUser(params)
	if err != nil {
		return result, errorcode.New(errorcode.INVALID_ACCOUNT_NOT_FOUND_CODE, "checkQueryAccount", errorcode.INVALID_ACCOUNT_NOT_FOUND_MSG)
	}
	return result, nil
}

func (svc *WebService) LoginWeb(c *gin.Context, params model.LoginRequest) (result model.UserDb, errString error) {
	user1, err1 := svc.dbService.FindUser(model.DupuserRequest{UserName: params.UserName})
	// 登录并注册
	if params.IsCreateAccount {
		if err1 == nil {
			return result, errorcode.New(errorcode.INVALID_ACCOUNT_EXISTED_CODE, "checkAccount", errorcode.INVALID_ACCOUNT_EXISTED_MSG)
		}

		err2 := svc.dbService.CreateUser(params)
		if err2 != nil {
			// 账户创建失败
			return result, errorcode.New(errorcode.INVALID_ACCOUNT_CREATED_CODE, "checkAccountCreate", errorcode.INVALID_ACCOUNT_CREATED_MSG)
		}

		user2, err2 := svc.dbService.FindUser(model.DupuserRequest{UserName: params.UserName})
		if err2 != nil {
			return user2, errorcode.New(errorcode.DB_QUERY_FAILED_CODE, "checkQueryAccount", err2.Error())
		}
	} else { // 只登录

		if err1 != nil {
			return result, errorcode.New(errorcode.INVALID_ACCOUNT_NOT_FOUND_CODE, "checkQueryAccount", errorcode.INVALID_ACCOUNT_NOT_FOUND_MSG)
		}

		if params.Password != user1.Password {
			return result, errorcode.New(errorcode.INVALID_PASSWORD_FAILED_CODE, "checkPassword", errorcode.INVALID_PASSWORD_FAILED_MSG)
		}
	}

	token, terr := util.GenToken(params.UserName)
	if terr == nil {
		svc.dbService.RefreshToken(params, fmt.Sprintf("Bearer %s", token))
	}
	ud, err := svc.dbService.FindUser(model.DupuserRequest{UserName: params.UserName})
	return ud, err
}
