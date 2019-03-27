package shard

import (
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/smartcontract/service/native/shard_stake"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func ShardUserStake(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardId uint64, pubKeys []string,
	amount []uint64) error {
	param := &shard_stake.UserStakeParam{
		ShardId:    shardId,
		User:       user.Address,
		PeerPubKey: pubKeys,
		Amount:     amount,
	}
	method := shard_stake.USER_STAKE
	contractAddress := utils.ShardStakeAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("ShardUserStake txHash is: %s", txHash.ToHexString())
	return nil
}
