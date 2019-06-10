# Ontology Shard Tool

Usage: ./main -t {method_name}

There are many method to use.

## Shard Management Contract

以下命令用来管理分片的生命周期以及共识切换。

分片创建及初始化参考[shard.sh](./shard.sh)。

### ShardInit

初始化ShardMgmt合约。交易只能发往root。

### ShardCreate

创建分片，分片的创建者就是该分片的管理员。创建完成后分片的状态为```SHARD_STATE_CREATED```

### ShardConfig

配置创建的分片。配置完成后分片状态为```SHARD_STATE_CONFIGURED```

### ShardPeerApply

节点申请加入分片分片共识，要加入的节点必须是在root上的共识或者候选节点。

ontology-tool可以一次做多个节点的申请。

### ShardPeerApprove

分片管理员同意已经做过申请的节点加入分片，可以一次同意多个节点的申请。

### ShardPeerJoin

节点加入分片，当有一个节点加入后，分片的状态变为```SHARD_PEER_JOIND```。

ontology-tool可以一次做多个节点的加入。

每个节点加入后会将参与质押的token转入到shard stake合约，如果余额不足，加入分片会失败。

### ShardActivate

由分片管理员调用，激活分片，分片状态必须是```SHARD_PEER_JOIND```，激活后状态变为```SHARD_STATE_ACTIVE```。

激活时，会做一点检查。

### ShardInfoQuery

查询分片详情。honglei使用的，wangcheng未使用过。

### ShardPeerExit

节点退出分片，共识方面的退出机制与主链一样。

节点质押的token需要在节点完全退出后才能提取，节点的手续费收益正常提取。

### NotifyParentCommitDpos

分片通知parent自己要切换共识。分片周期性切共识的入口也是这个函数，当分片周期性切共识的请求没有上发到parent，可以调用这个函数当做重试机制。

也可以主动使用这个命令，让分片提前切共识，但是这并不是强制切共识。

交易只能发往分片上。

### NotifyShardRootCommitDpos

parent上切完了分片的共识，通知到分片我已经准备好了，你可以切了。当parent下发这个通知失败时，可以使用这个命令当做重试机制。

当没有处于分片共识切换的过程中时，这个函数将无法调用。

交易只能发往parent上。

### ShardRetryCommitDpos

分片收到主链切换完自己共识的notify后，开始自己切共识，将自己的手续费和切换的block的height以及hash发送到parent，并调用parent的shard stake合约的方法记录下来。

分片在最后notify主链之前，会将notify去主链的信息记录下来。最后notify主链如果失败了，可以使用此命令重试。

交易只能发往分片上。

### GetShardCommitDposInfo

发往分片上，查询分片最新的共识切换信息，就是[ShardRetryCommitDpos](#ShardRetryCommitDpos)提到的保存下来的信息。结构如下：
```
type ShardCommitDposInfo struct {
	TransferId          *big.Int                   `json:"transfer_id"`
	Height              uint32                     `json:"height"`
	Hash                common.Uint256             `json:"hash"`
	FeeAmount           uint64                     `json:"fee_amount"`
	XShardHandleFeeInfo *shard_stake.XShardFeeInfo `json:"xshard_handle_fee_info"`
}
```

### UpdateShardConfig

强制切分片的共识。使用的配置可以跟之前的不一样也可以一样，不一样的时候顺便连配置也更新了。

交易只能发往parent上。

## Shard Asset Contract

用来管理跨分片资产。ontology tool存了一份编译好的跨分片oep4资产的[avm code](./params/shardasset/xshardasstdemo.avm)以及对应的[源码](./params/shardasset/xshardasstdemo.py)。可以部署到主链上去，然后调用init方法初始化，之后就可以正常使用了。

使用ontology-cli可以很方便的实现合约部署与init的步骤，示例合约的地址是```18fb9366f1a9fa0a7cdd4d71661b4ac8c78ea762```，参考命令如下：

```
./ontology contract deploy --needstore --code ../ontology-tool/params/shardasset/xshardasstdemo.avm --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
sleep 10 # wait block generate
./ontology contract invoke --address 18fb9366f1a9fa0a7cdd4d71661b4ac8c78ea762 --params string:init,[int:0] --gasprice 0
```

### ShardAssetInit

初始化ShardAsset合约。

交易只能发往parent上。

### XShardTransferOep4

跨分片转移oep4资产。可以在任意分片上调用。但是资产只能在parent和child之间转移。

### XShardTransferOng

跨分片转移ong。可以在任意分片上调用。但是资产只能在parent和child之间转移。

### XShardTransferOngRetry

跨分片转移ong失败时，调用该函数重试。该函数使用的transferId可以使用[ShardGetPendingTransfer](#ShardGetPendingTransfer)查询出某个用户所有未完成transfer得到。

### XShardTransferOep4Retry

同[XShardTransferOngRetry](#XShardTransferOngRetry)。

### ShardGetPendingTransfer

查询某用户在某个asset的未完成的跨分片转移。ong的资产ID是0，其他的资产的ID按照注册的先后从1开始递增。

### ShardGetTransferDetail

查询某笔跨分片交易的详情。

### ShardGetSupplyInfo

查询资产在各个分片上的发行量，只能在parent上查。

### ShardGetOep4Balance

查询oep4资产的余额。

ong的余额可以通过ontology-cli查询，加上--rpcport可以指定发送到哪个端口（对应哪个分片）。

## Shard Stake Contract

用来管理节点质押，用户质押，节点和用户提取收益。

用户质押节点的方法参考[stake.sh](./stake.sh)

### ShardChangePeerMaxAuth

节点调用改变自己接受的票数，即调大了可以接受用户质押。参数是多少就代表接受多少量的token的质押。

ontology-tool可以一次让多个节点接受投票。

### ShardChangePeerProportion

节点改变分给用户手续费的比例，参数是多少，代表分给百分之多少给用户。到下一轮共识周期生效。

ontology-tool可以一次让多个节点改变这个比例。

### ShardUserWithdrawOng

提取质押的ONT自然解绑的ONG。节点和用户都可以调用。到下一轮共识周期生效。

### ShardUserStake

用户质押节点，可以一次质押多个节点，但是节点必须要先[ShardChangePeerMaxAuth](#ShardChangePeerMaxAuth)接受投票。到下一轮共识周期生效。

### ShardUserUnfreezeStake

用户取消质押，可以一次取消多个节点的质押。到下一轮共识周期才能把质押的钱提出来，即质押的token在下一轮才解冻。

### ShardUserWithdrawStake

用户提取已经解冻的质押token。包括之前轮解冻的和当前轮质押的。

### ShardUserWithdrawFee

用户提取积累的所有手续费收益。

### ShardAddInitPos

节点增加init pos。

### ShardReduceInitPos

节点减少init pos。

### ShardQueryView

查询分片当前的共识周期，到了第几轮共识。

### ShardQueryPeerInfo

根据参数查询分片的节点质押信息，参数是第几轮共识。

### ShardQueryUserInfo

根据参数查询用户对节点的质押信息，参数是第几轮共识。