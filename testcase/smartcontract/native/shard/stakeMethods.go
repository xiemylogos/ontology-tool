package shard

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/service/native/shard_stake"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func ShardPeerChangeMaxAuth(ctx *testframework.TestFrameworkContext, shardId uint64, peers []*sdk.Account, amount []uint64) error {
	for index, peer := range peers {
		param := &shard_stake.ChangeMaxAuthorizationParam{
			ShardId: types.NewShardIDUnchecked(shardId),
			User:    peer.Address,
			Value: &shard_stake.PeerAmount{
				PeerPubKey: hex.EncodeToString(keypair.SerializePublicKey(peer.PublicKey)),
				Amount:     amount[index],
			},
		}
		method := shard_stake.CHANGE_MAX_AUTHORIZATION
		contractAddress := utils.ShardStakeAddress
		txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), peer, 0,
			contractAddress, method, []interface{}{param})
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("ShardPeerChangeMaxAuth txHash is: %s", txHash.ToHexString())
	}
	return nil
}

func ShardPeerChangeProportion(ctx *testframework.TestFrameworkContext, shardId uint64, peers []*sdk.Account, amount []uint64) error {
	for index, peer := range peers {
		param := &shard_stake.ChangeProportionParam{
			ShardId: types.NewShardIDUnchecked(shardId),
			User:    peer.Address,
			Value: &shard_stake.PeerAmount{
				PeerPubKey: hex.EncodeToString(keypair.SerializePublicKey(peer.PublicKey)),
				Amount:     amount[index],
			},
		}
		method := shard_stake.CHANGE_PROPORTION
		contractAddress := utils.ShardStakeAddress
		txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), peer, 0,
			contractAddress, method, []interface{}{param})
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("ShardPeerChangeProportion txHash is: %s", txHash.ToHexString())
	}
	return nil
}

func ShardUserStake(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardId uint64, pubKeys []string,
	amount []uint64) error {
	param := &shard_stake.UserStakeParam{
		ShardId: types.NewShardIDUnchecked(shardId),
		User:    user.Address,
	}
	param.Value = make([]*shard_stake.PeerAmount, 0)
	for index, key := range pubKeys {
		param.Value = append(param.Value, &shard_stake.PeerAmount{
			PeerPubKey: key,
			Amount:     amount[index],
		})
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

func ShardUserWithdrawOng(ctx *testframework.TestFrameworkContext, users []*sdk.Account) error {
	for _, user := range users {
		method := shard_stake.WITHDRAW_ONG
		contractAddress := utils.ShardStakeAddress
		txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
			contractAddress, method, []interface{}{user.Address})
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("ShardUserWithdrawOng txHash is: %s", txHash.ToHexString())
	}
	return nil
}
