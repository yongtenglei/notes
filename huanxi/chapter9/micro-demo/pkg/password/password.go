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
