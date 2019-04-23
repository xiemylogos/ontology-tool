package shard

import (
	"time"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
)

func getAccountByPassword(ctx *testframework.TestFrameworkContext, path string) (*sdk.Account, bool) {
	wallet, err := ctx.Ont.OpenWallet(path)
	if err != nil {
		ctx.LogError("open wallet error:%s", err)
		return nil, false
	}
	//pwd, err := password.GetPassword()
	//if err != nil {
	//	ctx.LogError("getPassword error:%s", err)
	//	return nil, false
	//}
	user, err := wallet.GetDefaultAccount([]byte("123"))
	if err != nil {
		ctx.LogError("getDefaultAccount error:%s", err)
		return nil, false
	}
	return user, true
}

func waitForBlock(ctx *testframework.TestFrameworkContext) bool {
	_, err := ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("WaitForGenerateBlock error:%s", err)
		return false
	}
	return true
}

func ConcatKey(args ...[]byte) []byte {
	temp := []byte{}
	for _, arg := range args {
		temp = append(temp, arg...)
	}
	return temp
}
