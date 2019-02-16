package shard

import (
	"bytes"
	"fmt"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/smartcontract/service/native/shardgas"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/ontio/ontology/errors"
	"github.com/ontio/ontology/common/serialization"
	"github.com/ontio/ontology/smartcontract/service/native/shardping"
)

func ShardGasInit(ctx *testframework.TestFrameworkContext, user *sdk.Account) error {
	method := shardgas.INIT_NAME
	contractAddress := utils.ShardGasMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard gas init txHash is :", txHash.ToHexString())
	return nil
}

func ShardDepositGas(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, amount uint64) error {
	param := shardgas.DepositGasParam{
		UserAddress: user.Address,
		ShardID:     shardID,
		Amount:      amount,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard deposit gas param: %s", err)
	}

	method := shardgas.DEPOSIT_GAS_NAME
	contractAddress := utils.ShardGasMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard deposit gas txHash is :", txHash.ToHexString())
	return nil
}

func ShardQueryGas(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64) error {
	contractAddr := utils.OngContractAddress.ToHexString()
	value, err := ctx.Ont.GetShardStorage(shardID, contractAddr, user.Address[:])
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "get shard storage error")
	}
	amount, err := serialization.ReadUint64(bytes.NewBuffer(value))
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "parse ong amount")
	}
	ctx.LogInfo("shard %d, address: %s, amount: %d", shardID, user.Address.ToHexString(), amount)
	return nil
}

func ShardSendPing(ctx *testframework.TestFrameworkContext, user *sdk.Account, fromShardID, toShardID uint64, txt string) error {
	param := shardping.ShardPingParam{
		FromShard: fromShardID,
		ToShard: toShardID,
		Param: txt,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard deposit gas param: %s", err)
	}

	method := shardping.SEND_SHARD_PING_NAME
	contractAddress := utils.ShardPingAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(fromShardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard send ping txHash is :", txHash.ToHexString())
	return nil
}
