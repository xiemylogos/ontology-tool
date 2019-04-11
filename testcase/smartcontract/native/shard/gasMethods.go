package shard

import (
	"bytes"
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	com "github.com/ontio/ontology-tool/testcase/smartcontract/native/common"
	"github.com/ontio/ontology-tool/testframework"
	comm "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/errors"
	"github.com/ontio/ontology/http/base/common"
	"github.com/ontio/ontology/smartcontract/service/native/shardgas"
	"github.com/ontio/ontology/smartcontract/service/native/shardping"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func ShardGasInit(ctx *testframework.TestFrameworkContext, pubKeys []keypair.PublicKey, users []*sdk.Account) error {
	method := shardgas.INIT_NAME
	contractAddress := utils.ShardGasMgmtContractAddress

	txHash := comm.Uint256{}
	var err error
	if len(users) == 1 {
		txHash, err = ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), users[0], 0,
			contractAddress, method, []interface{}{})
	} else {
		txHash, err = com.InvokeNativeContractWithMultiSign(ctx, ctx.GetGasPrice(), ctx.GetGasLimit(), pubKeys, users, 0,
			contractAddress, method, []interface{}{})
	}
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard gas init txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardDepositGas(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, amount uint64) error {
	param := shardgas.DepositGasParam{
		User:    user.Address,
		ShardId: types.NewShardIDUnchecked(shardID),
		Amount:  amount,
	}
	method := shardgas.DEPOSIT_GAS_NAME
	contractAddress := utils.ShardGasMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard deposit gas txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardQueryGas(ctx *testframework.TestFrameworkContext, shardID uint64) error {
	contractAddr := utils.ShardGasMgmtContractAddress
	preTx, err := common.NewNativeInvokeTransaction(0, 0, contractAddr, byte(0),
		shardgas.GET_SHARD_GAS_BALANCE_NAME, []interface{}{shardID})
	if err != nil {
		return fmt.Errorf("ShardQueryGas: build tx failed, err: %s", err)
	}
	preTx.ShardID = shardID
	value, err := ctx.Ont.PreExecTransaction(preTx)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "get shard storage error")
	}
	amount, err := value.Result.ToInteger()
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "parse ong amount")
	}
	ctx.LogInfo("shard %d, amount: %d", shardID, amount)
	return nil
}

func ShardUserWithdrawGas(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, amount uint64, shardUrl string) error {
	contractAddr := utils.ShardGasMgmtContractAddress
	param := &shardgas.UserWithdrawGasParam{
		User:   user.Address,
		Amount: amount,
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	tx, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddr, shardgas.USER_WITHDRAW_GAS_NAME, []interface{}{param})
	if err != nil {
		return fmt.Errorf("send tx failed, err: %s", err)
	}
	ctx.LogInfo("success, withdraw shard gas tx hash is %s", tx.ToHexString())
	return nil
}

func QueryShardUserUnFinishWithdraw(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, shardUrl string) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	contractAddr := utils.ShardGasMgmtContractAddress
	preTx, err := common.NewNativeInvokeTransaction(0, 0, contractAddr, byte(0),
		shardgas.GET_UN_FINISH_WITHDRAW, []interface{}{user.Address})
	if err != nil {
		return fmt.Errorf("QueryShardUserUnFinishWithdraw: build tx failed, err: %s", err)
	}
	preTx.Version = 1
	preTx.ShardID = shardID
	value, err := ctx.Ont.PreExecTransaction(preTx)
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	amount, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("shard %d, amount: %s", shardID, amount)
	return nil
}

func ShardUserRetryWithdraw(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID, withdrawId uint64, shardUrl string) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	contractAddr := utils.ShardGasMgmtContractAddress
	param := &shardgas.UserRetryWithdrawParam{
		User:       user.Address,
		WithdrawId: withdrawId,
	}
	tx, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddr, shardgas.WITHDRAW_RETRY_NAME, []interface{}{param})
	if err != nil {
		return fmt.Errorf("send tx failed, err: %s", err)
	}
	ctx.LogInfo("success, withdraw shard gas tx hash is %s", tx.ToHexString())
	return nil
}

func ShardCommitDpos(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, shardUrl string) error {
	contractAddr := utils.ShardGasMgmtContractAddress
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	tx, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddr, shardgas.SHARD_COMMIT_DPOS, []interface{}{})
	if err != nil {
		return fmt.Errorf("send tx failed, err: %s", err)
	}
	ctx.LogInfo("success, shard commit dpos tx hash is %s", tx.ToHexString())
	return nil
}

func ShardSendPing(ctx *testframework.TestFrameworkContext, user *sdk.Account, fromShardID, toShardID uint64, txt string) error {
	tFromShardId, _ := types.NewShardID(fromShardID)
	tToShardId, _ := types.NewShardID(toShardID)
	param := shardping.ShardPingParam{
		FromShard: tFromShardId,
		ToShard:   tToShardId,
		Param:     txt,
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
	ctx.LogInfo("shard send ping txHash is :%s", txHash.ToHexString())
	return nil
}
