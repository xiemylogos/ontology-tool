OntCversion = '2.0.0'
"""
An Example of OEP-11
"""
from ontology.interop.Ontology.Contract import InitMetaData
from ontology.interop.Ontology.Runtime import Base58ToAddress
from ontology.interop.Ontology.Shard import GetShardId, NotifyRemoteShard, InvokeRemoteShard
from ontology.interop.System.Runtime import Serialize, Deserialize
from ontology.interop.System.Storage import GetContext, Put, Get
from ontology.libont import hexstring2address

ctx = GetContext()

OWNER = Base58ToAddress("AZ3BTJt7jNGwJjVLsYJAyfLtCJ38Cd8Uri")
X_SHARD_INVOKED_CONTRACT = hexstring2address("88493a7ebae5e0431854f3f0b7e8f791f5e2d089")
SHARD_VERSION = 1

X_SHARD_INVOKE_KEY = "xshard_invoke"


def Main(operation, args):
    """
    :param operation:
    :param args:
    :return:
    """
    if operation == 'init':
        return init()
    if operation == 'xshardNotify':
        if len(args) < 2:
            return False
        xshardNotify(args[0], args[1])
        return False
    if operation == 'xshardInvoke':
        if len(args) < 2:
            return False
        xshardInvoke(args[0], args[1])
        return False
    if operation == 'getXShardInvokeRes':
        return getXShardInvokeRes()
    return False


def init():
    """
    initialize the contract, init contract meta data
    :return:
    """
    allShard = True
    frozen = False
    shardId = GetShardId()
    res = InitMetaData(OWNER, allShard, frozen, shardId, [X_SHARD_INVOKED_CONTRACT])
    assert (res)


def xshardNotify(a, b):
    list = [a, b]
    argsByteArray = Serialize(list)
    targetShardId = 2
    res = NotifyRemoteShard(targetShardId, X_SHARD_INVOKED_CONTRACT, 30000, "notifyCallee", argsByteArray)
    assert (res)
    return True


def xshardInvoke(a, b):
    list = [a, b]
    argsByteArray = Serialize(list)
    targetShardId = 2
    res = InvokeRemoteShard(targetShardId, X_SHARD_INVOKED_CONTRACT, "invokeCallee", argsByteArray)
    Put(ctx, X_SHARD_INVOKE_KEY, Deserialize(res))
    return True


def getXShardInvokeRes():
    return Get(ctx, X_SHARD_INVOKE_KEY)
