package native

import (
	"github.com/ontio/ontology-tool/testcase/smartcontract/native/governance_feeSplit"
	"github.com/ontio/ontology-tool/testcase/smartcontract/native/shard"
)

func TestNative() {
	//governance_feeSplit.TestGovernanceContract()
	//governance_feeSplit.TestGovernanceContractError()
	governance_feeSplit.TestGovernanceMethods()
	//governance_feeSplit.TestGovernanceBatch()
	shard.TestShardMgmtContract()
}
