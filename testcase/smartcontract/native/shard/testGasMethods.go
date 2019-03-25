package shard

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
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
	if err := ShardGasInit(ctx, pubKeys, users); err != nil {
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

type QueryShardGasParam struct {
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

	param := &QueryShardGasParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal query shard gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardQueryGas(ctx, user, param.ShardID); err != nil {
		ctx.LogError("shard query gas failed: %s", err)
		return false
	}

	return true
}

type ShardUserWithdrawGasParam struct {
	Path     string `json:"path"`
	ShardID  uint64 `json:"shard_id"`
	Amount   uint64 `json:"amount"`
	ShardUrl string `json:"shard_url"`
}

func TestShardUserWithdrawGas(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardUserWithdrawGas.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardUserWithdrawGasParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal withdraw shard gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardUserWithdrawGas(ctx, user, param.ShardID, param.Amount, param.ShardUrl); err != nil {
		ctx.LogError("user withdraw shard gas failed: %s", err)
		return false
	}

	return true
}

type QueryShardUserUnFinishWithdrawParam struct {
	Path     string `json:"path"`
	ShardID  uint64 `json:"shard_id"`
	ShardUrl string `json:"shard_url"`
}

func TestQueryShardUserUnFinishWithdraw(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardUserQueryUnFinishWithdraw.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &QueryShardUserUnFinishWithdrawParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := QueryShardUserUnFinishWithdraw(ctx, user, param.ShardID, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	return true
}

type ShardRetryWithdrawParam struct {
	Path       string `json:"path"`
	ShardID    uint64 `json:"shard_id"`
	WithdrawId uint64 `json:"withdraw_id"`
	ShardUrl   string `json:"shard_url"`
}

func TestShardUserRetryWithdraw(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardUserRetryWithdrawGas.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardRetryWithdrawParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardUserRetryWithdraw(ctx, user, param.ShardID, param.WithdrawId, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	return true
}

type ShardCommitDposParam struct {
	Path     string `json:"path"`
	ShardID  uint64 `json:"shard_id"`
	ShardUrl string `json:"shard_url"`
}

func TestShardCommitDpos(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardCommitDpos.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardCommitDposParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardCommitDpos(ctx, user, param.ShardID, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	return true
}

type ShardSendPingParam struct {
	Path        string `json:"path"`
	FromShardID uint64 `json:"from_shard_id"`
	ToShardID   uint64 `json:"to_shard_id"`
	Param       string `json:"param"`
}

func TestShardSendPing(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/ShardSendPing.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardSendPingParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard deposit gas param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	ctx.LogInfo("send shard ping from %d to %d, param %s", param.FromShardID, param.ToShardID, param.Param)
	if err := ShardSendPing(ctx, user, param.FromShardID, param.ToShardID, param.Param); err != nil {
		ctx.LogError("shard ping failed: %s", err)
		return false
	}

	return true
}
