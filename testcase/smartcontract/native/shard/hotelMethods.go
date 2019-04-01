package shard

import (
	"bytes"
	"fmt"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/service/native/shardhotel"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func ShardHotelInit(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomCount uint64) error {
	param := shardhotel.ShardHotelInitParam{
		Count: roomCount,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard deposit gas param: %s", err)
	}

	method := shardhotel.SHARD_HOTEL_INIT_NAME
	contractAddress := utils.ShardHotelAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard room init txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardReserveRoom(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomNo uint64) error {
	param := shardhotel.ShardHotelReserveParam{
		User:   user.Address,
		RoomNo: roomNo,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard deposit gas param: %s", err)
	}

	method := shardhotel.SHARD_RESERVE_NAME
	contractAddress := utils.ShardHotelAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard reserve room txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardCheckoutRoom(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomNo uint64) error {
	param := shardhotel.ShardHotelCheckoutParam{
		User:   user.Address,
		RoomNo: roomNo,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard deposit gas param: %s", err)
	}

	method := shardhotel.SHARD_CHECKOUT_NAME
	contractAddress := utils.ShardHotelAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard checkout room txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardQueryRoom(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, room int) (common.Address, error) {
	roomBytes := utils.GetUint32Bytes(uint32(room))
	key := ConcatKey([]byte(shardhotel.KEY_ROOM), roomBytes)
	value, err := ctx.Ont.GetStorage(utils.ShardHotelAddress.ToHexString(), key)
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("shardQuery, get storage: %s", err)
	}
	var userAddr common.Address
	copy(userAddr[:], value)
	return userAddr, nil
}

func ShardHotelReserve2(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomNo uint64,
	shardID2 uint64, roomNo2 uint64, transactional bool) error {
	shardId2, err := types.NewShardID(shardID2)
	if err != nil {
		return fmt.Errorf("invalid shard id2: %d", shardID2)
	}
	param := shardhotel.ShardHotelReserve2Param{
		User:             user.Address,
		RoomNo1:          roomNo,
		Shard2:           shardId2,
		ContractAddress2: utils.ShardHotelAddress,
		RoomNo2:          roomNo2,
		Transactional:    transactional,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard hotel reserve2 param: %s", err)
	}

	method := shardhotel.SHARD_DOUBLE_RESERVE_NAME
	contractAddress := utils.ShardHotelAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard hotel reserve 2 txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardHotelCheckout2(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomNo uint64,
	shardID2 uint64, roomNo2 uint64, transactional bool) error {
	shardId2, err := types.NewShardID(shardID2)
	if err != nil {
		return fmt.Errorf("invalid shard id2: %d", shardID2)
	}
	param := shardhotel.ShardHotelReserve2Param{
		User:             user.Address,
		RoomNo1:          roomNo,
		Shard2:           shardId2,
		ContractAddress2: utils.ShardHotelAddress,
		RoomNo2:          roomNo2,
		Transactional:    transactional,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard hotel checkout2 param: %s", err)
	}

	method := shardhotel.SHARD_DOUBLE_CHECKOUT_NAME
	contractAddress := utils.ShardHotelAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard hotel checkout2 txHash is :%s", txHash.ToHexString())
	return nil
}
