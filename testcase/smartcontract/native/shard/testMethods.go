package shard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common/config"
)

type ShardInitParam struct {
	Path []string
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
	var users []*sdk.Account
	var pubKeys []keypair.PublicKey
	time.Sleep(1 * time.Second)
	for _, path := range param.Path {
		user, ok := getAccountByPassword(ctx, path)
		if !ok {
			return false
		}
		users = append(users, user)
		pubKeys = append(pubKeys, user.PublicKey)
	}
	if err := ShardInit(ctx, pubKeys, users); err != nil {
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
	Path        string             `json:"path"`
	ShardID     uint64             `json:"shard_id"`
	NetworkSize uint32             `json:"network_size"`
	VbftConfig  *config.VBFTConfig `json:"vbft"`
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

	if err := ShardConfig(ctx, user, param.ShardID, param.NetworkSize, param.VbftConfig); err != nil {
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

func TestShardInfoQuery(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardInfoQuery.json"
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

	s, err := ShardStateQuery(ctx, param.ShardID)
	if err != nil {
		ctx.LogError("shard query: %s", err)
		return false
	}

	buf := new(bytes.Buffer)
	s.Serialize(buf)
	fmt.Printf("shard: %v", string(buf.Bytes()))
	return true
}
