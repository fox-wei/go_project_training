package handler

import (
	"errors"
	"net/http"

	mylog "github.com/fox-wei/go_project_training/account_management/MyLog"
	"github.com/fox-wei/go_project_training/account_management/handler/param"
	"github.com/fox-wei/go_project_training/account_management/middleware"
	"github.com/fox-wei/go_project_training/account_management/model"
	mytoken "github.com/fox-wei/go_project_training/account_management/model/my_token"
	"github.com/fox-wei/go_project_training/account_management/myerr"
	"github.com/fox-wei/go_project_training/account_management/res"
	"github.com/fox-wei/go_project_training/account_management/service"
	"github.com/fox-wei/go_project_training/account_management/token"
	"github.com/fox-wei/go_project_training/account_management/utils"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Srv service.AccountService
}

//?新建account
func (ah *AccountHandler) AccountCreate(ctx *gin.Context) {
	var r param.AccountRequest

	id := ctx.Param("id")

	//?数据绑定
	if err := ctx.Bind(&r); err != nil {
		res.SendResponse(ctx, myerr.ErrBind, err)
		return
	}

	//?数据验证
	if err := utils.CheckParam(r.AccountName, r.Password); err.Err != nil {
		res.SendResponse(ctx, &err, r)
		return
	}

	accountName := r.AccountName
	mylog.Log.Infof("用户名: %s", accountName) //?打印日志

	desc := ctx.Query("desc") //?url参数获取，参数不存在则为空
	mylog.Log.Infof("desc:%s", desc)

	contentType := ctx.GetHeader("Content-Type")
	mylog.Log.Infof("Header Content-Type:%s", contentType)

	//!数据加密
	md5wd, err := utils.Encrypt(r.Password)
	if err != nil {
		res.SendResponse(ctx, myerr.ErrEncrypt, err.Error())
		return
	}

	if err != nil {
		res.SendResponse(ctx, myerr.InternalServerError, nil)
		return
	}

	//*添加到数据库
	a := model.Account{
		AccountId:   id,
		AccountName: r.AccountName,
		Password:    md5wd,
	}

	if err := ah.Srv.CreatedAccount(a); err != nil {
		res.SendResponse(ctx, myerr.ErrDatabase, nil)
		return
	}

	rsp := res.AccountResp{
		AccountId:   id,
		AccountName: r.AccountName,
	}

	res.SendResponse(ctx, nil, rsp)

}

//?软删除
func (ah *AccountHandler) AccountDelete1(ctx *gin.Context) {
	id := ctx.Param("id") //?获取动态参数/:id
	mylog.Log.Infof("id:%s", id)

	//?查询id是否存在
	_, err := ah.Srv.Repo.GetAccountById(id)

	if err != nil {
		res.SendResponse(ctx, errors.New("account isn't exist"), nil)
		return
	}

	if err := ah.Srv.Repo.DeleteAccount(id); err != nil {
		res.SendResponse(ctx, myerr.ErrDatabase, err)
		return
	}

	res.SendResponse(ctx, nil, nil)
}

//*账户更新
func (ah *AccountHandler) AccountUpdate(ctx *gin.Context) {
	var m model.Account //?数据绑定

	if err := ctx.Bind(&m); err != nil {
		res.SendResponse(ctx, myerr.ErrBind, err.Error())
		return
	}

	//?密码加密
	if m.Password != "" {
		md5pd, err := utils.Encrypt(m.Password)
		if err != nil {
			res.SendResponse(ctx, myerr.ErrEncrypt, err.Error())
			return
		}
		m.Password = md5pd
	}

	//*保存更新
	if err := ah.Srv.UpdateAccount(m); err != nil {
		res.SendResponse(ctx, myerr.ErrDatabase, err.Error())
		return
	}

	res.SendResponse(ctx, nil, nil)
}

//?根据名字返回账户信息
func (ah *AccountHandler) GetAccount(ctx *gin.Context) {
	name := ctx.Param("name")

	account, err := ah.Srv.GetAccount(name)
	if err != nil {
		res.SendResponse(ctx, myerr.ErrAccountNotFound, err.Error())
		return
	}

	res.SendResponse(ctx, nil, gin.H{"name": account.AccountName,
		"accoundId": account.AccountId})
}

//*账户列表
func (ah *AccountHandler) AccountList(ctx *gin.Context) {
	var r param.AccountList

	//*数据绑定获取offset和limit
	if err := ctx.Bind(&r); err != nil {
		res.SendResponse(ctx, myerr.ErrBind, err.Error())
		return
	}

	mylog.Log.Infof("offset:%d limt:%d", r.Offset, r.Limit)
	if r.Offset < 0 {
		r.Offset = 0
	}

	if r.Limit < 1 {
		r.Limit = utils.Limit
	}

	accountlist, count, err := ah.Srv.ListAccount(r.Offset, r.Limit)
	if err != nil {
		res.SendResponse(ctx, err, nil)
		return
	}

	var resp []*res.AccountResp

	for _, item := range accountlist {
		r := res.AccountResp{AccountId: item.AccountId, AccountName: item.AccountName}
		resp = append(resp, &r)
	}

	//*获取查询人姓名
	auth := ctx.Request.Header.Get("authorization")
	claims, _ := middleware.ParseToken(auth)
	name := claims["username"]

	ctx.JSON(http.StatusOK, res.ListReponse{
		Username:    name,
		TotalCount:  count,
		AccountList: resp,
	})
}

//?登录
func (ah *AccountHandler) Login(ctx *gin.Context) {
	var m model.Account
	//*数据绑定，获取账户和密码
	if err := ctx.Bind(&m); err != nil {
		res.SendResponse(ctx, myerr.ErrBind, err.Error())
		return
	}

	//?根据id获取账户信息
	account, err := ah.Srv.Repo.GetAccountById(m.AccountId)
	if err != nil {
		res.SendResponse(ctx, err, nil)
		return
	}

	//?密码检验
	if err := utils.Compare(account.Password, m.Password); err != nil {
		res.SendResponse(ctx, myerr.ErrPassword, err)
		return
	}

	//*生成jwt
	sign, err := token.Sign(ctx, token.Context{ID: account.AccountId, Username: account.AccountName}, "")
	if err != nil {
		res.SendResponse(ctx, myerr.ErrToken, err.Error())
		return
	}

	res.SendResponse(ctx, nil, mytoken.Token{Token: sign})
}
