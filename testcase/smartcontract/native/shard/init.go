package shard

import "github.com/ontio/ontology-tool/testframework"

func TestShardMgmtContract() {
	testframework.TFramework.RegTestCase("ShardInit", TestShardInit)
	testframework.TFramework.RegTestCase("ShardCreate", TestShardCreate)
	testframework.TFramework.RegTestCase("ShardConfig", TestShardConfig)
	testframework.TFramework.RegTestCase("ShardPeerJoin", TestShardPeerJoin)
	testframework.TFramework.RegTestCase("ShardActivate", TestShardActivate)
	testframework.TFramework.RegTestCase("ShardInfoQuery", TestShardInfoQuery)

	testframework.TFramework.RegTestCase("ShardGasInit", TestShardGasInit)
	testframework.TFramework.RegTestCase("ShardDepositGas", TestShardDespoitGas)
	testframework.TFramework.RegTestCase("ShardQueryGas", TestShardQueryGas)
}
