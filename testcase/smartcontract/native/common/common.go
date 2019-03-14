package common

import (
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	scommon "github.com/ontio/ontology/common"
)

func InvokeNativeContractWithMultiSign(
	ctx *testframework.TestFrameworkContext,
	gasPrice,
	gasLimit uint64,
	pubKeys []keypair.PublicKey,
	singers []*sdk.Account,
	cversion byte,
	contractAddress scommon.Address,
	method string,
	params []interface{},
) (scommon.Uint256, error) {
	tx, err := ctx.Ont.Native.NewNativeInvokeTransaction(gasPrice, gasLimit, cversion, contractAddress, method, params)
	if err != nil {
		return scommon.UINT256_EMPTY, err
	}
	for _, singer := range singers {
		err = ctx.Ont.MultiSignToTransaction(tx, uint16((5*len(pubKeys)+6)/7), pubKeys, singer)
		if err != nil {
			return scommon.UINT256_EMPTY, err
		}
	}
	return ctx.Ont.SendTransaction(tx)
}
