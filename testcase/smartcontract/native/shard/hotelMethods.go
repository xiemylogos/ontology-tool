package shard

import (
	"bytes"
	"fmt"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/shardhotel"
	"github.com/ontio/ontology/smartcontract/service/native/shardmgmt/utils"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func ShardHotelInit(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomCount int) error {
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

func ShardReserveRoom(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomNo int) error {
	param := shardhotel.ShardHotelReserveParam{
		User:   user.Address,
		RoomNo: roomNo,
	}
	buf := new(bytes.Buffer)
	if err := param.Serialize(buf); err != nil {
		return fmt.Errorf("failed to ser shard deposit gas param: %s", err)
	}

	method := shardhotel.SHARD_RESERVE_ROOM_NAME
	contractAddress := utils.ShardHotelAddress
	txHash, err := ctx.Ont.Native.InvokeShardNativeContract(shardID, ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{buf.Bytes()})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("shard reserve room txHash is :%s", txHash.ToHexString())
	return nil
}

func ShardCheckoutRoom(ctx *testframework.TestFrameworkContext, user *sdk.Account, shardID uint64, roomNo int) error {
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
	roomBytes := shardutil.GetUint32Bytes(uint32(room))
	key := ConcatKey([]byte(shardhotel.KEY_ROOM), roomBytes)
	value, err := ctx.Ont.GetStorage(utils.ShardHotelAddress.ToHexString(), key)
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("shardQuery, get storage: %s", err)
	}
	var userAddr common.Address
	copy(userAddr[:], value)
	return userAddr, nil
}
