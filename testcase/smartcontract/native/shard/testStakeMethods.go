package shard

import (
	"encoding/json"
	"io/ioutil"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
)

type ChangePeerAttrParam struct {
	ShardId     uint64   `json:"shard_id"`
	PeerWallets []string `json:"peer_wallets"`
	Amount      []uint64 `json:"amount"`
}

func TestShardChangePeerMaxAuthorization(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardPeerChangeMaxAuth.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ChangePeerAttrParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	users := make([]*sdk.Account, 0)
	for _, peerWallet := range param.PeerWallets {
		user, ok := getAccountByPassword(ctx, peerWallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}
	if err := ShardPeerChangeMaxAuth(ctx, param.ShardId, users, param.Amount); err != nil {
		ctx.LogError("TestShardChangePeerMaxAuthorization failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

func TestShardChangePeerProportion(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardPeerChangeProportion.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ChangePeerAttrParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	users := make([]*sdk.Account, 0)
	for _, peerWallet := range param.PeerWallets {
		user, ok := getAccountByPassword(ctx, peerWallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}
	if err := ShardPeerChangeProportion(ctx, param.ShardId, users, param.Amount); err != nil {
		ctx.LogError("TestShardChangePeerProportion failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

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
		ctx.LogError("TestShardUserStake failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type UserWithdrawOngParam struct {
	Wallets []string `json:"wallets"`
}

func TestShardUserWithdrawOng(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserWithdrawOng.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserWithdrawOngParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	users := make([]*sdk.Account, 0)
	for _, peerWallet := range param.Wallets {
		user, ok := getAccountByPassword(ctx, peerWallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}
	if err := ShardUserWithdrawOng(ctx, users); err != nil {
		ctx.LogError("TestShardUserWithdrawOng failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}
