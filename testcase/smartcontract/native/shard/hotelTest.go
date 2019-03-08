package shard

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ontio/ontology-tool/testframework"
)

type ShardHotelInitParam struct {
	Path      string `json:"path"`
	ShardID   uint64 `json:"shard_id"`
	RoomCount int    `json:"room_count"`
}

func TestShardHotelInit(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardHotelInit.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardHotelInitParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	ctx.LogInfo("send shard hotel init to %d, room count: %d", param.ShardID, param.RoomCount)
	if err := ShardHotelInit(ctx, user, param.ShardID, param.RoomCount); err != nil {
		ctx.LogError("shard ping failed: %s", err)
		return false
	}

	return true
}

type ShardQueryRoomParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
	RoomNo  int    `json:"room_no"`
}

func TestShardHotelQuery(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardHotelQuery.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardQueryRoomParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	ctx.LogInfo("send shard hotel reserve to %d, room %d", param.ShardID, param.RoomNo)
	if user, err := ShardQueryRoom(ctx, user, param.ShardID, param.RoomNo); err != nil {
		ctx.LogError("shard query failed: %s", err)
		return false
	} else {
		ctx.LogInfo("user address: %s", user.ToBase58())
	}

	return true
}

type ShardReserveHotelParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
	RoomNo  int    `json:"room_no"`
}

func TestShardHotelReserve(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardHotelReserve.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardReserveHotelParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	ctx.LogInfo("send shard hotel reserve to %d, room %d", param.ShardID, param.RoomNo)
	if err := ShardReserveRoom(ctx, user, param.ShardID, param.RoomNo); err != nil {
		ctx.LogError("shard ping failed: %s", err)
		return false
	}

	return true
}

type ShardCheckoutHotelParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
	RoomNo  int    `json:"room_no"`
}

func TestShardHotelCheckout(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardHotelCheckout.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardCheckoutHotelParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	ctx.LogInfo("send shard hotel checkout to %d, room %d", param.ShardID, param.RoomNo)
	if err := ShardCheckoutRoom(ctx, user, param.ShardID, param.RoomNo); err != nil {
		ctx.LogError("shard ping failed: %s", err)
		return false
	}

	return true
}
