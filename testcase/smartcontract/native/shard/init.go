
package shard

import "github.com/ontio/ontology-tool/testframework"

func TestShardMgmtContract() {
	testframework.TFramework.RegTestCase("ShardInit", TestShardInit)
	testframework.TFramework.RegTestCase("ShardCreate", TestShardCreate)
}


