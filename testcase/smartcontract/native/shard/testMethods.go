
package shard

import (
	"github.com/ontio/ontology-tool/testframework"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"bytes"
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

type ShardQueryParam struct {
	ShardID uint64 `json:"shard_id"`
}

func TestShardQuery(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardQuery.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardQueryParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard create param: %s", err)
		return false
	}

	s, err := ShardQuery(ctx, param.ShardID)
	if err != nil {
		ctx.LogError("shard query: %s", err)
		return false
	}

	buf := new(bytes.Buffer)
	s.Serialize(buf)
	fmt.Printf("shard: %s", string(buf.Bytes()))
	return true
}