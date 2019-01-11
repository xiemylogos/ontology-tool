
package shard

import (
	"github.com/ontio/ontology-tool/testframework"
	"io/ioutil"
	"encoding/json"
)

type ShardInitParam struct {
	Path string `json:"path"`
}

func TestShardInit(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardInit.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardInitParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardInit(ctx, user); err != nil {
		ctx.LogError("shard init failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardCreateParam struct {
	Path string `json:"path"`
	ParentShardID uint64 `json:"parent_shard_id"`
}

func TestShardCreate(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardCreate.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardCreateParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard create param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardCreate(ctx, user, param.ParentShardID); err != nil {
		ctx.LogError("shard init failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}