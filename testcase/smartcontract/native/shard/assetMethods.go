package shard

import (
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/shardasset/oep4"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"math/big"
)

func AssetInit(ctx *testframework.TestFrameworkContext, user *sdk.Account) error {
	method := oep4.INIT
	contractAddress := utils.ShardAssetAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func XShardTransfer(ctx *testframework.TestFrameworkContext, user *sdk.Account, contractAddress, to common.Address,
	amount uint64, toShard common.ShardID, shardUrl, method string) error {
	param := &oep4.XShardTransferParam{
		From:    user.Address,
		To:      to,
		ToShard: toShard,
		Amount:  new(big.Int).SetUint64(amount),
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func XShardTransferRetry(ctx *testframework.TestFrameworkContext, user *sdk.Account, contractAddress common.Address,
	transferId uint64, shardUrl, method string) error {
	param := &oep4.XShardTransferRetryParam{
		From:       user.Address,
		TransferId: new(big.Int).SetUint64(transferId),
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func GetPendingTransfer(ctx *testframework.TestFrameworkContext, user *sdk.Account, assetId uint64, shardUrl string) error {
	method := oep4.GET_PENDING_TRANSFER
	contractAddress := utils.ShardAssetAddress
	param := &oep4.GetPendingXShardTransferParam{
		Account: user.Address,
		Asset:   assetId,
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, 0,
		[]interface{}{param})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("pending transfer is: %s", info)
	return nil
}

func GetTransferDetail(ctx *testframework.TestFrameworkContext, user common.Address, assetId, transferId uint64,
	shardUrl string) error {
	method := oep4.GET_TRANSFER
	contractAddress := utils.ShardAssetAddress
	param := &oep4.GetXShardTransferInfoParam{
		Account:    user,
		Asset:      assetId,
		TransferId: new(big.Int).SetUint64(transferId),
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, 0,
		[]interface{}{param})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("transfer is: %s", info)
	return nil
}

func GetSupplyInfo(ctx *testframework.TestFrameworkContext, assetId uint64, shardUrl string) error {
	method := oep4.GET_TRANSFER
	contractAddress := utils.ShardAssetAddress
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, 0,
		[]interface{}{assetId})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("supply info is: %s", info)
	return nil
}
