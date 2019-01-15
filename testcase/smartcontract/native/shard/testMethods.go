package shard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ontio/ontology-tool/testframework"
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
	Path          string `json:"path"`
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

type ShardConfigParam struct {
	Path        string `json:"path"`
	ShardID     uint64 `json:"shard_id"`
	NetworkSize uint32 `json:"network_size"`
}

func TestShardConfig(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardConfig.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardConfigParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard create param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardConfig(ctx, user, param.ShardID, param.NetworkSize); err != nil {
		ctx.LogError("shard init failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardPeerJoinParam struct {
	Path        string `json:"path"`
	ShardID     uint64 `json:"shard_id"`
	PeerAddress string `json:"peer_address"`
	PeerPubKey  string `json:"peer_pub_key"`
	StakeAmount uint64 `json:"stake_amount"`
}

func TestShardPeerJoin(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardPeerJoin.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardPeerJoinParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard peer join param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardPeerJoin(ctx, user, param.ShardID, param.PeerAddress, param.PeerPubKey, param.StakeAmount); err != nil {
		ctx.LogError("shard peer join failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardActivateParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
}

func TestShardActivate(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardActivate.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardActivateParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard activate param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardActivate(ctx, user, param.ShardID); err != nil {
		ctx.LogError("shard activate failed: %s", err)
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

type ShardDepositGasParam struct {
	Path    string `json:"path"`
	ShardID uint64 `json:"shard_id"`
	Amount uint64 `json:"amount"`
}

func TestShardDespoitGas(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardDespositGas.json"
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