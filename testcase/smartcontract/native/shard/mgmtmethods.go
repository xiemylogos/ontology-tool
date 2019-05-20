package shard

import (
	"bytes"
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	com "github.com/ontio/ontology-tool/testcase/smartcontract/native/common"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/config"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt/states"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func ShardInit(ctx *testframework.TestFrameworkContext, pubKeys []keypair.PublicKey, users []*sdk.Account) error {
	method := shardmgmt.INIT_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash := common.Uint256{}
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
	ctx.LogInfo("shard init txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardCreate(ctx *testframework.TestFrameworkContext, user *sdk.Account, parentID common.ShardID) error {
	param := &shardmgmt.CreateShardParam{
		ParentShardID: parentID,
		Creator:       user.Address,
	}

	method := shardmgmt.CREATE_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard create txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardConfig(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID common.ShardID, networkSize uint32,
	vbft *config.VBFTConfig) error {
	cfgBuff := new(bytes.Buffer)
	if err := vbft.Serialize(cfgBuff); err != nil {
		return fmt.Errorf("serialize vbft config failed, err: %s", err)
	}
	param := &shardmgmt.ConfigShardParam{
		ShardID:           shardID,
		NetworkMin:        networkSize,
		GasPrice:          0,
		GasLimit:          20000,
		StakeAssetAddress: utils.OntContractAddress,
		GasAssetAddress:   utils.OngContractAddress,
		VbftConfigData:    cfgBuff.Bytes(),
	}
	method := shardmgmt.CONFIG_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard config txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardApplyJoin(ctx *testframework.TestFrameworkContext, shardID common.ShardID, user []*sdk.Account, peerPubKey []string) error {
	for index, acc := range user {
		applyJoinParam := &shardmgmt.ApplyJoinShardParam{
			ShardId:    shardID,
			PeerOwner:  acc.Address,
			PeerPubKey: peerPubKey[index],
		}
		method := shardmgmt.APPLY_JOIN_SHARD_NAME
		contractAddress := utils.ShardMgmtContractAddress
		txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), acc, 0,
			contractAddress, method, []interface{}{applyJoinParam})
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("apply join shard txHash is: %s", txHash.ToHexString())
	}
	return nil
}

func ApproveJoin(ctx *testframework.TestFrameworkContext, user []*sdk.Account, shardID common.ShardID, peerPubKey []string) error {
	applyJoinParam := &shardmgmt.ApproveJoinShardParam{
		ShardId:    shardID,
		PeerPubKey: peerPubKey,
	}
	method := shardmgmt.APPROVE_JOIN_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	var txHash common.Uint256
	var err error = nil
	if len(user) == 1 {
		txHash, err = ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user[0], 0,
			contractAddress, method, []interface{}{applyJoinParam})
	} else {
		pubKeys := make([]keypair.PublicKey, 0)
		for _, u := range user {
			pubKeys = append(pubKeys, u.PublicKey)
		}
		txHash, err = com.InvokeNativeContractWithMultiSign(ctx, ctx.GetGasPrice(), ctx.GetGasLimit(), pubKeys, user, 0,
			contractAddress, method, []interface{}{applyJoinParam})
	}
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("approve join shard txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardPeerJoin(ctx *testframework.TestFrameworkContext, shardID common.ShardID, user []*sdk.Account, peers []*JoinShardPeer) error {
	for index, u := range user {
		peer := peers[index]
		param := &shardmgmt.JoinShardParam{
			ShardID:     shardID,
			IpAddress:   peer.IpAddress,
			PeerOwner:   u.Address,
			PeerPubKey:  peer.PubKey,
			StakeAmount: peer.StakeAmount,
		}

		method := shardmgmt.JOIN_SHARD_NAME
		contractAddress := utils.ShardMgmtContractAddress
		txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), u, 0,
			contractAddress, method, []interface{}{param})
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("join shard txHash is: %s", txHash.ToHexString())
	}
	return nil
}

func ShardActivate(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID common.ShardID) error {
	param := &shardmgmt.ActivateShardParam{
		ShardID: shardID,
	}

	method := shardmgmt.ACTIVATE_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("activate shard txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardPeerExit(ctx *testframework.TestFrameworkContext, owner *sdk.Account, shardID common.ShardID, peer string) error {
	param := shardmgmt.ExitShardParam{
		ShardId:    shardID,
		PeerOwner:  owner.Address,
		PeerPubKey: peer,
	}
	method := shardmgmt.EXIT_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), owner, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("peer exit shard txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardStateQuery(ctx *testframework.TestFrameworkContext, shardID common.ShardID) (*shardstates.ShardState, error) {
	globalStateValue, err := ctx.Ont.GetStorage(utils.ShardMgmtContractAddress.ToHexString(), []byte(shardmgmt.KEY_GLOBAL_STATE))
	if err != nil {
		return nil, fmt.Errorf("shardQeury, get global storage: %s", err)
	}
	gs := &shardstates.ShardMgmtGlobalState{}
	if err := gs.Deserialization(common.NewZeroCopySource(globalStateValue)); err != nil {
		return nil, fmt.Errorf("failed to parse global state: %s", err)
	}
	fmt.Printf("global state: %v \n", gs)

	shardIDBytes := utils.GetUint64Bytes(shardID.ToUint64())
	key := ConcatKey([]byte(shardmgmt.KEY_SHARD_STATE), shardIDBytes)
	value, err := ctx.Ont.GetStorage(utils.ShardMgmtContractAddress.ToHexString(), key)
	if err != nil {
		return nil, fmt.Errorf("shardQuery, get storage: %s", err)
	}

	s := &shardstates.ShardState{}
	if err := s.Deserialization(common.NewZeroCopySource(value)); err != nil {
		return nil, fmt.Errorf("shardQuery, deserialize shard state: %s", err)
	}

	return s, nil
}

func NotifyParentCommitDpos(ctx *testframework.TestFrameworkContext, shardId common.ShardID, user *sdk.Account,
	shardUrl string) error {
	contractAddress := utils.ShardMgmtContractAddress
	method := shardmgmt.NOTIFY_PARENT_COMMIT_DPOS
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardId.ToUint64(), ctx.GetGasPrice(), ctx.GetGasLimit(),
		user, 0, contractAddress, method, []interface{}{})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func NotifyShardCommitDpos(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardId common.ShardID) error {
	contractAddress := utils.ShardMgmtContractAddress
	method := shardmgmt.NOTIFY_SHARD_COMMIT_DPOS
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{shardId})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func ShardRetryCommitDpos(ctx *testframework.TestFrameworkContext, shardId common.ShardID, user *sdk.Account,
	shardUrl string) error {
	contractAddress := utils.ShardMgmtContractAddress
	method := shardmgmt.SHARD_RETRY_COMMIT_DPOS
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardId.ToUint64(), ctx.GetGasPrice(), ctx.GetGasLimit(),
		user, 0, contractAddress, method, []interface{}{})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func GetShardCommitDposInfo(ctx *testframework.TestFrameworkContext, shardUrl string, shardId common.ShardID) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	method := shardmgmt.GET_SHARD_COMMIT_DPOS_INFO
	contractAddress := utils.ShardMgmtContractAddress
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, shardId.ToUint64(),
		[]interface{}{})
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

func UpdateShardConfig(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID common.ShardID,
	cfg *utils.Configuration) error {
	contractAddress := utils.ShardMgmtContractAddress
	method := shardmgmt.SHARD_RETRY_COMMIT_DPOS
	param := &shardmgmt.UpdateConfigParam{
		ShardId: shardID,
		NewCfg:  cfg,
	}
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{param})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}
