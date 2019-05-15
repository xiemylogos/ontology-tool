package shard

import (
	"encoding/json"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
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
	configFile := "./params/shardmgmt/ShardInit.json"
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
	Path          string         `json:"path"`
	ParentShardID common.ShardID `json:"parent_shard_id"`
}

func TestShardCreate(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardCreate.json"
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
	ShardID     common.ShardID     `json:"shard_id"`
	NetworkSize uint32             `json:"network_size"`
	VbftConfig  *config.VBFTConfig `json:"vbft"`
}

func TestShardConfig(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardConfig.json"
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

type ShardPeerApplyJoinParam struct {
	Path       []string       `json:"path"`
	ShardId    common.ShardID `json:"shard_id"`
	PeerPubKey []string       `json:"peer_pub_key"`
}

func TestShardPeerApplyJoin(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardPeerApplyJoin.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardPeerApplyJoinParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	users := make([]*sdk.Account, 0)
	for _, path := range param.Path {
		user, ok := getAccountByPassword(ctx, path)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}

	if err := ShardApplyJoin(ctx, param.ShardId, users, param.PeerPubKey); err != nil {
		ctx.LogError("shard peer apply join failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardPeerApproveJoinParam struct {
	Path       []string       `json:"path"`
	ShardId    common.ShardID `json:"shard_id"`
	PeerPubKey []string       `json:"peer_pub_key"`
}

func TestShardPeerApproveJoin(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardPeerApproveJoin.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardPeerApproveJoinParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	users := make([]*sdk.Account, 0)
	for _, path := range param.Path {
		user, ok := getAccountByPassword(ctx, path)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}

	if err := ApproveJoin(ctx, users, param.ShardId, param.PeerPubKey); err != nil {
		ctx.LogError("shard peer approve join failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type JoinShardPeer struct {
	Wallet      string `json:"wallet"`
	IpAddress   string `json:"ip_address"`
	PubKey      string `json:"pub_key"`
	StakeAmount uint64 `json:"stake_amount"`
}

type ShardPeerJoinParam struct {
	ShardId common.ShardID   `json:"shard_id"`
	Peers   []*JoinShardPeer `json:"peers"`
}

func TestShardPeerJoin(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardPeerJoin.json"
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

	users := make([]*sdk.Account, 0)
	for _, peer := range param.Peers {
		user, ok := getAccountByPassword(ctx, peer.Wallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}

	if err := ShardPeerJoin(ctx, param.ShardId, users, param.Peers); err != nil {
		ctx.LogError("shard peer join failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardActivateParam struct {
	Path    string         `json:"path"`
	ShardID common.ShardID `json:"shard_id"`
}

func TestShardActivate(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardActivate.json"
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

type ShardPeerExitParam struct {
	ShardId common.ShardID `json:"shard_id"`
	Wallet  string         `json:"wallet"`
	Peer    string         `json:"peer"`
}

func TestShardPeerExit(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardPeerExit.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardPeerExitParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard activate param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Wallet)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardPeerExit(ctx, user, param.ShardId, param.Peer); err != nil {
		ctx.LogError("shard activate failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardQueryParam struct {
	ShardID common.ShardID `json:"shard_id"`
}

func TestShardInfoQuery(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardInfoQuery.json"
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

	sink := common.NewZeroCopySink(0)
	s.Serialization(sink)
	fmt.Printf("shard: %s", string(sink.Bytes()))
	return true
}

type NotifyRootCommitDposParam struct {
	Path     string
	ShardId  common.ShardID
	ShardUrl string
}

func TestNotifyRootCommitDpos(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/NotifyRootCommitDpos.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &NotifyRootCommitDposParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := NotifyRootCommitDpos(ctx, param.ShardId, user, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type NotifyShardCommitDposParam struct {
	Path    string
	ShardId common.ShardID
}

func TestNotifyShardCommitDpos(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/NotifyShardCommitDpos.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &NotifyShardCommitDposParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := NotifyShardCommitDpos(ctx, user, param.ShardId); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type ShardRetryCommitDposParam struct {
	Path     string
	ShardId  common.ShardID
	ShardUrl string
}

func TestShardRetryCommitDpos(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/ShardRetryCommitDpos.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ShardRetryCommitDposParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := ShardRetryCommitDpos(ctx, param.ShardId, user, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type GetShardCommitDposInfoParam struct {
	ShardUrl string
	ShardId common.ShardID
}

func TestGetShardCommitDposInfo(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/GetShardCommitDposInfo.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetShardCommitDposInfoParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}
	if err := GetShardCommitDposInfo(ctx, param.ShardUrl, param.ShardId); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}

	waitForBlock(ctx)
	return true
}

type UpdateShardConfigParam struct {
	Path    string
	ShardId common.ShardID
	*utils.Configuration
}

func TestUpdateShardConfig(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardmgmt/UpdateShardConfig.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UpdateShardConfigParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}

	if err := UpdateShardConfig(ctx, user, param.ShardId, param.Configuration); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}
