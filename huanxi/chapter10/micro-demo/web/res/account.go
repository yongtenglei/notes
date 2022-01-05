package res

import "github.com/rey/micro-demo/proto/account"

type AccountRes struct {
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
	Gender   string `json:"gender"`
}

func AccountInfo2AccountRes(ai *account.AccountInfo) *AccountRes {
	return &AccountRes{
		Mobile:   ai.Mobile,
		UserName: ai.Username,
		Gender:   ai.Gender,
	}

}
