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

func ShardQuery(ctx *testframework.TestFrameworkContext, shardID uint64) (*shardmgmt.ShardState, error) {
	globalStateValue, err := ctx.Ont.GetStorage(utils.ShardMgmtContractAddress.ToHexString(), []byte(shardmgmt.KEY_GLOBAL_STATE))
	if err != nil {
		return nil, fmt.Errorf("shardQeury, get global storage: %s", err)
	}
	gs := &shardmgmt.ShardMgmtGlobalState{}
	if err := gs.Deserialize(bytes.NewBuffer(globalStateValue)); err != nil {
		return nil, fmt.Errorf("failed to parse global state: %s", err)
	}
	fmt.Printf("global state: %v \n", gs)

	shardIDBytes, err := shardmgmt.GetUint64Bytes(shardID)
	if err != nil {
		return nil, fmt.Errorf("get shard ID bytes: %s", err)
	}
	key := ConcatKey([]byte(shardmgmt.KEY_SHARD_STATE), shardIDBytes)
	value, err := ctx.Ont.GetStorage(utils.ShardMgmtContractAddress.ToHexString(), key)
	if err != nil {
		return nil, fmt.Errorf("shardQuery, get storage: %s", err)
	}

	s := &shardmgmt.ShardState{}
	if err := s.Deserialize(bytes.NewBuffer(value)); err != nil {
		return nil, fmt.Errorf("shardQuery, deserialize shard state: %s", err)
	}

	return s, nil
}