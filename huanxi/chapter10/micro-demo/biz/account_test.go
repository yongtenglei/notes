package biz

import (
	"context"
	"fmt"
	"testing"

	"github.com/rey/micro-demo/dao/mysql"
	_ "github.com/rey/micro-demo/dao/mysql"
	"github.com/rey/micro-demo/model"
	"github.com/rey/micro-demo/proto/account"
)

func TestAddAcount(t *testing.T) {
	accountServer := AccountServer{}

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%s%d", "1300000000", i)

		accountInfo := &account.AccountInfo{
			Mobile:   s,
			Password: s,
			Username: s,
			Gender:   "male",
			Role:     1,
		}

		if _, err := accountServer.AddAccount(context.Background(), &account.AddAccountReq{Account: accountInfo}); err != nil {
			continue

		}

		var account model.Account
		if err := mysql.DB.Where(&model.Account{Mobile: s}).First(&account); err != nil {
			panic(err)
		}

		t.Logf("%d\n", account.ID)
	}

}

func TestGetAccountByMobile(t *testing.T) {
	accountServer := &AccountServer{}
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%s%d", "1300000000", i)

		res, err := accountServer.GetAccountByMobile(context.Background(), &account.MobileReq{Mobile: s})
		if err != nil {
			fmt.Println("bong")
			fmt.Println(err)
			panic(err)
		}

		fmt.Println(res.Account.Id)
		fmt.Println(res.Account.Password)

	}
}

func TestGetAccountByID(t *testing.T) {
	accountServer := &AccountServer{}

	res, err := accountServer.GetAccountByID(context.Background(), &account.IDReq{Id: 1})
	if err != nil {
		fmt.Println("bong")
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(res.Account.Id)
	fmt.Println(res.Account.Password)

}

func TestUPdateAccount(t *testing.T) {
	accountServer := &AccountServer{}

	s := "13000000000"

	var act model.Account
	result := mysql.DB.Where("mobile = ?", s).First(&act)
	if result.RowsAffected == 0 {
		panic("bong")
	}

	act.UserName = "rey"

	res, err := accountServer.UpdateAccount(context.Background(), &account.UpdateAccountReq{Account: act.ToPBAccountInfo()})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Ok)

}

func TestCheckAccountByID(t *testing.T) {
	accountServer := &AccountServer{}

	s := "13000000000"

	res, err := accountServer.CheckAccountByID(context.Background(), &account.CheckAccountByIDReq{Id: 1, Password: s})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Ok)
}

func TestDeleteAccountByID(t *testing.T) {
	accountServer := &AccountServer{}

	s := "13000000009"

	res, err := accountServer.DeleceAccount(context.Background(), &account.DeleteAccountReq{Id: 10, Password: s})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Ok)
}
