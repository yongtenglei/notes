# chapter8 微服务 demo

## 生成表

```go
package model

import (
	"github.com/rey/micro-demo/proto/account"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;varchar(11;not null)"`
	Password string `gorm:"type:varchar(64);not null"`
	UserName string `gorm:"type:varchar(32)"`
	Gender   string `gorm:"type:varchar(6);default:male"`
	Salt     string `gorm:"type:varchar(64)"`
	Role     int    `gorm:"type:tinyint;defaut:1;comment:'1:user 0:admin'"`
}

// 更改gorm自动迁移时的表名
func (Account) TableName() string {
	return "micro_account"
}

// Account -> pb.AccountInfo
func (a Account) ToPBAccountInfo() (pa *account.AccountInfo) {
	pa = &account.AccountInfo{}
	pa.Id = int64(a.ID)
	pa.Mobile = a.Mobile
	pa.Password = a.Password
	pa.Username = a.UserName
	pa.Gender = a.Gender

	return
}

```

## MYSQL 的配置

```go
package mysql

import (
	"log"
	"os"
	"time"

	"github.com/rey/micro-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/micro_demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&model.Account{})
	if err != nil {
		panic("AutoMigrate Failed")
	}
}

func Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo == 0 {
			pageNo = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 1
		}

		offset := pageNo - 1
		return db.Offset(offset).Limit(pageSize)
	}

}

```

## protobuffer

```go
syntax="proto3";

package account;

option go_package="/proto/account;account"; // 完整路径 ; 别名

service AccountService{
  rpc GetAccountList(PagingReq) returns (AccountListRes);
  rpc GetAccountByMobile(MobileReq) returns (MobileRes);
  rpc GetAccountByID(IDReq) returns (IDRes);
  rpc AddAccount(AddAccountReq) returns (AddAccountRes);
  rpc UpdateAccount(UpdateAccountReq) returns (UpdateAccountRes);
  rpc CheckAccountByID(CheckAccountByIDReq) returns (CheckAccountByIDRes);
  rpc DeleceAccount(DeleteAccountReq) returns (DeleteAccountRes);
}

message PagingReq {
  uint32 pageNo=1;
  uint32 pageSize=2;
}

message AccountInfo {
  int64 id=1;
  string mobile=2;
  string password=3;
  string username=4;
  string gender=5;
  uint32 role=6;
}

message AccounttRes {
  AccountInfo account=1;
}

message AccountListRes {
  int32 total=1;
  repeated AccountInfo accountList=2;
}

message MobileReq {
  string mobile=1;
}

message MobileRes {
  AccountInfo account=1;
}

message IDReq {
  int64 id=1;
}

message IDRes {
  AccountInfo account=1;
}

message AddAccountReq {
  AccountInfo account=1;
}

message AddAccountRes {
  bool ok=1;
}

message UpdateAccountReq {
  AccountInfo account=1;
}

message UpdateAccountRes {
  bool ok=1;
}

message CheckAccountByIDReq {
  int64 id=1;
  string password=2;
}

message CheckAccountByIDRes {
  bool ok=1;
}


message DeleteAccountReq {
  int64 id=1;
  string password=2;
}

message DeleteAccountRes {
  bool ok=1;
}

```

生成代码: `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./account/account.proto`

## 定义错误代码

```go
package e

const (
	AccountNotFund          = "账户不存在"
	AccountAlreadyExitst    = "账户已存在"
	AccountPasswordNotMatch = "账户密码错误"

	InternalBusy = "服务器繁忙"
)

```

## 编写业务代码

```go
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

```

## 编写测试

```go
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

```

## 使 md5 更加安全

```go
package password

import (
	"crypto/md5"

	password "github.com/anaskhan96/go-password-encoder"
)

//var options
var options = password.Options{
	SaltLen:      16,
	Iterations:   100,
	KeyLen:       32,
	HashFunction: md5.New,
}

func MD5(p string) (salt, encoded string) {
	salt, encoded = password.Encode(p, &options)
	return
}

func MD5Verify(reqPwd, salt, pwd string) bool {
	return password.Verify(reqPwd, salt, pwd, &options)
}

```

推荐包: `github.com/anaskhan96/go-password-encoder`

本实现对包进行了第二次封装

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gxlc0e6d9cj30cc0dlq5b.jpg">

</div>
