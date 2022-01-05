package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/rey/micro-demo/dao/mysql"
	"github.com/rey/micro-demo/model"
	"github.com/rey/micro-demo/pkg/e"
	"github.com/rey/micro-demo/pkg/password"
	"github.com/rey/micro-demo/proto/account"
)

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
// type UnimplementedAccountServiceServer struct {
// }
type AccountServer struct {
	// 见注释
	account.UnimplementedAccountServiceServer
}

func (as *AccountServer) GetAccountList(ctx context.Context, req *account.PagingReq) (res *account.AccountListRes, err error) {
	res = new(account.AccountListRes)
	var accountList []*model.Account
	result := mysql.DB.Scopes(mysql.Paginate(int(req.PageNo), int(req.PageSize))).Find(&accountList)
	if result.Error != nil {
		return nil, result.Error
	}

	res.Total = int32(result.RowsAffected)

	for _, account := range accountList {
		res.AccountList = append(res.AccountList, account.ToPBAccountInfo())
	}
	return res, nil
}

func (as *AccountServer) GetAccountByMobile(ctx context.Context, req *account.MobileReq) (res *account.MobileRes, err error) {
	res = new(account.MobileRes)
	var account model.Account
	result := mysql.DB.Where(&model.Account{Mobile: req.Mobile}).First(&account)
	fmt.Println(result)
	if result.RowsAffected == 0 {
		return nil, errors.New(e.AccountNotFund)
	}

	res.Account = account.ToPBAccountInfo()
	return res, nil
}

func (as *AccountServer) GetAccountByID(ctx context.Context, req *account.IDReq) (res *account.IDRes, err error) {
	res = new(account.IDRes)
	var account model.Account
	result := mysql.DB.First(&account, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(e.AccountNotFund)
	}
	res.Account = account.ToPBAccountInfo()
	return res, nil
}

func (as *AccountServer) AddAccount(ctx context.Context, req *account.AddAccountReq) (res *account.AddAccountRes, err error) {
	res = new(account.AddAccountRes)
	var account model.Account
	result := mysql.DB.Where(&model.Account{Mobile: req.Account.Mobile}).First(&account)
	if result.RowsAffected != 0 {
		return nil, errors.New(e.AccountAlreadyExitst)
	}

	account.Mobile = req.Account.Mobile
	account.UserName = req.Account.Username
	account.Gender = req.Account.Gender
	account.Role = int(req.GetAccount().Role)

	salt, encodedPassword := password.MD5(req.Account.Password)
	account.Salt = salt
	account.Password = encodedPassword

	if err := mysql.DB.Create(&account); err != nil {
		return nil, errors.New(e.InternalBusy)
	}
	res.Ok = true

	return res, nil
}
func (as *AccountServer) UpdateAccount(ctx context.Context, req *account.UpdateAccountReq) (res *account.UpdateAccountRes, err error) {
	res = new(account.UpdateAccountRes)
	var account model.Account
	result := mysql.DB.Where(&model.Account{Mobile: req.Account.Mobile}).First(&account)
	if result.RowsAffected == 0 {
		return nil, errors.New(e.AccountNotFund)
	}

	account.Password = req.Account.Password
	account.UserName = req.Account.Username
	account.Gender = req.Account.Gender
	account.Role = int(req.GetAccount().Role)

	if err := mysql.DB.Save(&account).Error; err != nil {
		return nil, errors.New(e.InternalBusy)
	}
	res.Ok = true

	return res, nil
}

func (as *AccountServer) CheckAccountByID(ctx context.Context, req *account.CheckAccountByIDReq) (res *account.CheckAccountByIDRes, err error) {
	res = new(account.CheckAccountByIDRes)
	var act model.Account
	result := mysql.DB.First(&act, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(e.AccountNotFund)
	}

	if act.Salt == "" {
		return nil, errors.New(e.InternalBusy)
	}

	res.Ok = password.MD5Verify(req.Password, act.Salt, act.Password)

	return res, nil
}

func (as *AccountServer) DeleceAccount(ctx context.Context, req *account.DeleteAccountReq) (res *account.DeleteAccountRes, err error) {
	res = new(account.DeleteAccountRes)

	// 查人并验证密码
	var act model.Account
	result := mysql.DB.First(&act, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(e.AccountNotFund)
	}

	if act.Salt == "" {
		return nil, errors.New(e.InternalBusy)
	}

	ok := password.MD5Verify(req.Password, act.Salt, act.Password)
	if !ok {
		// res.ok = false
		return res, errors.New(e.AccountPasswordNotMatch)
	}

	result = mysql.DB.Delete(&act)
	if result.RowsAffected == 0 {
		// res.ok = false
		return res, errors.New(e.InternalBusy)
	}

	res.Ok = ok
	return res, nil
}
