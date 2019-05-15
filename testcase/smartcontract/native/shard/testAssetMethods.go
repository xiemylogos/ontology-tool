package shard

import (
	"encoding/json"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/shardasset/oep4"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"io/ioutil"
)

type AssetInitParam struct {
	Path string
}

func TestAssetInit(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/AssetInit.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &AssetInitParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}
	if err := AssetInit(ctx, user); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type XShardTransferParam struct {
	Path     []string
	To       []common.Address
	Amount   []uint64
	ToShard  common.ShardID
	ShardUrl string
	Contract string
}

func TestXShardTransferOep4(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/XShardTransferOep4.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &XShardTransferParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
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
	contract, err := common.AddressFromHexString(param.Contract)
	if err != nil {
		ctx.LogError("decode contract addr failed, err: %s", err)
		return false
	}
	if err := XShardTransfer(ctx, users, contract, param.To, param.Amount, param.ToShard, param.ShardUrl,
		oep4.XSHARD_TRANSFER); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

func TestXShardTransferOng(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/XShardTransferOng.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &XShardTransferParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
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
	if err := XShardTransfer(ctx, users, utils.ShardAssetAddress, param.To, param.Amount, param.ToShard, param.ShardUrl,
		oep4.ONG_XSHARD_TRANSFER); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type XShardTransferRetryParam struct {
	Path       []string
	TransferId []uint64
	ShardUrl   string
	Contract   string
}

func TestXShardTransferOngRetry(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/XShardTransferOngRetry.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &XShardTransferRetryParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
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
	if err := XShardTransferRetry(ctx, users, utils.ShardAssetAddress, param.TransferId, param.ShardUrl,
		oep4.ONG_XSHARD_TRANSFER_RETRY); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

func TestXShardTransferOep4Retry(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/XShardTransferOepRetry.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &XShardTransferRetryParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
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
	contract, err := common.AddressFromHexString(param.Contract)
	if err != nil {
		ctx.LogError("decode contract addr failed, err: %s", err)
		return false
	}
	if err := XShardTransferRetry(ctx, users, contract, param.TransferId, param.ShardUrl, oep4.XSHARD_TRANFSER_RETRY);
		err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type GetPendingTransferParam struct {
	Path     string
	AssetId  uint64
	ShardUrl string
}

func TestGetPendingTransfer(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/GetPendingTransfer.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetPendingTransferParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		ctx.LogError("get account failed")
		return false
	}
	if err := GetPendingTransfer(ctx, user, param.AssetId, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}

type GetTransferDetailParam struct {
	User       common.Address
	AssetId    uint64
	ShardUrl   string
	TransferId uint64
}

func TestGetTransferDetail(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/GetTransferDetail.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetTransferDetailParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	if err := GetTransferDetail(ctx, param.User, param.AssetId, param.TransferId, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}

type GetSupplyInfoParam struct {
	AssetId  uint64
	ShardUrl string
}

func TestGetSupplyInfo(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shardasset/GetSupplyInfo.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetSupplyInfoParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	if err := GetSupplyInfo(ctx, param.AssetId, param.ShardUrl); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}
