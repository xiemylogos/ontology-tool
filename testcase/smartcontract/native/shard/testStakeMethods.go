package shard

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ontio/ontology-tool/testframework"
)

type UserStakeParam struct {
	Path       string   `json:"path"`
	ShardId    uint64   `json:"shard_id"`
	PeerPubKey []string `json:"peer_pub_key"`
	Amount     []uint64 `json:"amount"`
}

func TestShardUserStake(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserStake.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserStakeParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		return false
	}

	if err := ShardUserStake(ctx, user, param.ShardId, param.PeerPubKey, param.Amount); err != nil {
		ctx.LogError("shard init failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}
