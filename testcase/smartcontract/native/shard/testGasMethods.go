package shard

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ontio/ontology-tool/testframework"
)

func TestShardGasInit(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardGasInit.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardInitParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard gas init param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardGasInit(ctx, user); err != nil {
		ctx.LogError("shard init failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardDepositGasParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
	Amount  uint64 `json:"amount"`
}

func TestShardDespoitGas(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardDepositGas.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardDepositGasParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardDepositGas(ctx, user, param.ShardID, param.Amount); err != nil {
		ctx.LogError("shard deposit gas failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardQueryGasParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
}

func TestShardQueryGas(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardQueryGas.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardQueryGasParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardQueryGas(ctx, user, param.ShardID); err != nil {
		ctx.LogError("shard deposit gas failed: %s", err)
		return false
	}

	return true
}
