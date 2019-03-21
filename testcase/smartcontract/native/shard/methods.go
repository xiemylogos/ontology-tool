package shard

import (
	"bytes"
	"fmt"

	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"

	sdk "github.com/ontio/ontology-go-sdk"
	com "github.com/ontio/ontology-tool/testcase/smartcontract/native/common"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common/config"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt/states"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt/utils"
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
	ctx.LogInfo("shard init txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardCreate(ctx *testframework.TestFrameworkContext, user *sdk.Account, parentID uint64) error {
	tShardId, _ := types.NewShardID(parentID)
	param := &shardmgmt.CreateShardParam{
		ParentShardID: tShardId,
		Creator:       user.Address,
	}

	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser createshard param: %s", err)
	}
	method := shardmgmt.CREATE_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard create txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardConfig(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, networkSize uint32, vbft *config.VBFTConfig) error {
	tShardId, _ := types.NewShardID(shardID)
	param := &shardmgmt.ConfigShardParam{
		ShardID:           tShardId,
		NetworkMin:        networkSize,
		StakeAssetAddress: utils.OntContractAddress,
		GasAssetAddress:   utils.OngContractAddress,
		VbftConfigData:    vbft,
	}

	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser config shard param: %s", err)
	}

	method := shardmgmt.CONFIG_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard config txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardApplyJoin(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, peerPubKey string) error {
	tShardId, _ := types.NewShardID(shardID)
	applyJoinParam := &shardmgmt.ApplyJoinShardParam{
		ShardId:    tShardId,
		PeerOwner:  user.Address,
		PeerPubKey: peerPubKey,
	}
	buf := new(bytes.Buffer)
	if err := applyJoinParam.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser param: %s", err)
	}
	method := shardmgmt.APPLY_JOIN_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("apply join shard txHash is :%s", txHash.ToHexString())
	return nil
}

func ApproveJoin(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, peerPubKey []string) error {
	tShardId, _ := types.NewShardID(shardID)
	applyJoinParam := &shardmgmt.ApproveJoinShardParam{
		ShardId:    tShardId,
		PeerPubKey: peerPubKey,
	}
	buf := new(bytes.Buffer)
	if err := applyJoinParam.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser param: %s", err)
	}
	method := shardmgmt.APPROVE_JOIN_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("apply join shard txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardPeerJoin(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, peerPubKey string,
	stakeAmount uint64) error {
	tShardId, _ := types.NewShardID(shardID)
	param := &shardmgmt.JoinShardParam{
		ShardID:     tShardId,
		PeerOwner:   user.Address,
		PeerPubKey:  peerPubKey,
		StakeAmount: stakeAmount,
	}

	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser join shard param: %s", err)
	}

	method := shardmgmt.JOIN_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("join shard txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardActivate(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64) error {
	tShardId, _ := types.NewShardID(shardID)
	param := &shardmgmt.ActivateShardParam{
		ShardID: tShardId,
	}

	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser activate shard param: %s", err)
	}

	method := shardmgmt.ACTIVATE_SHARD_NAME
	contractAddress := utils.ShardMgmtContractAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("activate shard txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardStateQuery(ctx *testframework.TestFrameworkContext, shardID uint64) (*shardstates.ShardState, error) {
	globalStateValue, err := ctx.Ont.GetStorage(utils.ShardMgmtContractAddress.ToHexString(), []byte(shardmgmt.KEY_GLOBAL_STATE))
	if err != nil {
		return nil, fmt.Errorf("shardQeury, get global storage: %s", err)
	}
	gs := &shardstates.ShardMgmtGlobalState{}
	if err := gs.Deserialize(bytes.NewBuffer(globalStateValue)); err != nil {
		return nil, fmt.Errorf("failed to parse global state: %s", err)
	}
	fmt.Printf("global state: %v \n", gs)

	shardIDBytes, err := shardutil.GetUint64Bytes(shardID)
	if err != nil {
		return nil, fmt.Errorf("get shard ID bytes: %s", err)
	}
	key := ConcatKey([]byte(shardmgmt.KEY_SHARD_STATE), shardIDBytes)
	value, err := ctx.Ont.GetStorage(utils.ShardMgmtContractAddress.ToHexString(), key)
	if err != nil {
		return nil, fmt.Errorf("shardQuery, get storage: %s", err)
	}

	s := &shardstates.ShardState{}
	if err := s.Deserialize(bytes.NewBuffer(value)); err != nil {
		return nil, fmt.Errorf("shardQuery, deserialize shard state: %s", err)
	}

	return s, nil
}
