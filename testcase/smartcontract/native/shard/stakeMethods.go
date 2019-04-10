package shard

import (
	"encoding/hex"
	"fmt"

	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/errors"
	bComm "github.com/ontio/ontology/http/base/common"
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

func ShardQueryView(ctx *testframework.TestFrameworkContext, shardID uint64) error {
	contractAddr := utils.ShardStakeAddress
	preTx, err := bComm.NewNativeInvokeTransaction(0, 0, contractAddr, byte(0),
		shard_stake.GET_CURRENT_VIEW, []interface{}{shardID})
	if err != nil {
		return fmt.Errorf("ShardQueryView: build tx failed, err: %s", err)
	}
	preTx.ShardID = shardID
	value, err := ctx.Ont.PreExecTransaction(preTx)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "get shard storage error")
	}
	view, err := value.Result.ToInteger()
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "parse view")
	}
	ctx.LogInfo("shard %d, view: %d", shardID, view)
	return nil
}

func ShardQueryPeerInfo(ctx *testframework.TestFrameworkContext, shardID, view uint64) error {
	contractAddr := utils.ShardStakeAddress
	param := &shard_stake.GetPeerInfoParam{
		ShardId: types.NewShardIDUnchecked(shardID),
		View:    view,
	}
	preTx, err := bComm.NewNativeInvokeTransaction(0, 0, contractAddr, byte(0),
		shard_stake.GET_PEER_INFO, []interface{}{param})
	if err != nil {
		return fmt.Errorf("ShardQueryPeerInfo: build tx failed, err: %s", err)
	}
	preTx.ShardID = shardID
	value, err := ctx.Ont.PreExecTransaction(preTx)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "get shard storage error")
	}
	info, err := value.Result.ToString()
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "parse view")
	}
	ctx.LogInfo("shard %d, info: %s", shardID, info)
	return nil
}

func ShardQueryUserInfo(ctx *testframework.TestFrameworkContext, shardID, view uint64, user common.Address) error {
	contractAddr := utils.ShardStakeAddress
	param := &shard_stake.GetUserStakeInfoParam{
		ShardId: types.NewShardIDUnchecked(shardID),
		View:    view,
		User:    user,
	}
	preTx, err := bComm.NewNativeInvokeTransaction(0, 0, contractAddr, byte(0),
		shard_stake.GET_USER_INFO, []interface{}{param})
	if err != nil {
		return fmt.Errorf("ShardQueryUserInfo: build tx failed, err: %s", err)
	}
	preTx.ShardID = shardID
	value, err := ctx.Ont.PreExecTransaction(preTx)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "get shard storage error")
	}
	info, err := value.Result.ToString()
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "parse view")
	}
	ctx.LogInfo("shard %d, info: %s", shardID, info)
	return nil
}
