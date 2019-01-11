package shard

import (
	"github.com/ontio/ontology-tool/testframework"
	sdk "github.com/ontio/ontology-go-sdk"
	"fmt"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt"
	"bytes"
)

func ShardInit(ctx *testframework.TestFrameworkContext, user *sdk.Account) error {
	method := shardmgmt.INIT_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard init txHash is :", txHash.ToHexString())
	waitForBlock(ctx)

	return nil
}

func ShardCreate(ctx *testframework.TestFrameworkContext, user *sdk.Account, parentID uint64) error {
	param := &shardmgmt.CreateShardParam{
		ParentShardID: parentID,
		Creator: user.Address,
	}

	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser param: %s", err)
	}
	method := shardmgmt.CREATE_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard create txHash is :", txHash.ToHexString())
	waitForBlock(ctx)

	return nil
}
