package shard

import "github.com/ontio/ontology-tool/testframework"

func TestShardMgmtContract() {
	testframework.TFramework.RegTestCase("ShardInit", TestShardInit)
	testframework.TFramework.RegTestCase("ShardCreate", TestShardCreate)
	testframework.TFramework.RegTestCase("ShardConfig", TestShardConfig)
	testframework.TFramework.RegTestCase("ShardPeerJoin", TestShardPeerJoin)
	testframework.TFramework.RegTestCase("ShardPeerApply", TestShardPeerApplyJoin)
	testframework.TFramework.RegTestCase("ShardPeerApprove", TestShardPeerApproveJoin)
	testframework.TFramework.RegTestCase("ShardActivate", TestShardActivate)
	testframework.TFramework.RegTestCase("ShardInfoQuery", TestShardInfoQuery)

	testframework.TFramework.RegTestCase("ShardGasInit", TestShardGasInit)
	testframework.TFramework.RegTestCase("ShardDepositGas", TestShardDespoitGas)
	testframework.TFramework.RegTestCase("ShardQueryGas", TestShardQueryGas)
	testframework.TFramework.RegTestCase("ShardUserWithdrawGas", TestShardUserWithdrawGas)
	testframework.TFramework.RegTestCase("ShardQueryUnFinishWithdraw", TestQueryShardUserUnFinishWithdraw)
	testframework.TFramework.RegTestCase("ShardRetryWithdraw", TestShardUserRetryWithdraw)
	testframework.TFramework.RegTestCase("ShardCommitDpos", TestShardCommitDpos)

	testframework.TFramework.RegTestCase("ShardUserStake", TestShardUserStake)

	testframework.TFramework.RegTestCase("ShardSendPing", TestShardSendPing)
	testframework.TFramework.RegTestCase("ShardHotelInit", TestShardHotelInit)
	testframework.TFramework.RegTestCase("ShardHotelQuery", TestShardHotelQuery)
	testframework.TFramework.RegTestCase("ShardHotelReserve", TestShardHotelReserve)
	testframework.TFramework.RegTestCase("ShardHotelCheckout", TestShardHotelCheckout)
	testframework.TFramework.RegTestCase("ShardHotelReserve2", TestShardHotelReserve2)
	testframework.TFramework.RegTestCase("ShardHotelCheckout2", TestShardHotelCheckout2)
}
