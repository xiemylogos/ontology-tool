OntCversion = '2.0.0'
"""
An Example of OEP-9
"""
from ontology.interop.Ontology.Contract import InitMetaData
from ontology.interop.Ontology.Runtime import Base58ToAddress
from ontology.interop.Ontology.Shard import GetShardId
from ontology.interop.System.Runtime import Deserialize, Notify, Serialize
from ontology.interop.System.Storage import GetContext, Put, Get
from ontology.libont import str

ctx = GetContext()

OWNER = Base58ToAddress("AKuMYaCm7LeBHqNeKzvj7qQb3USakDr5yg")
X_SHARD_NOTIFY_KEY = "xshardNotify"
X_SHARD_INVOKE_KEY = "xshardInvoke"


def Main(operation, args):
    """
    :param operation:
    :param args:
    :return:
    """
    if operation == 'init':
        return init()
    if operation == 'notifyCallee':
        if len(args) < 1:
            return False
        return notifyCallee(args[0])
    if operation == 'invokeCallee':
        if len(args) < 1:
            return False
        return invokeCallee(args[0])
    if operation == 'getXShardCount':
        return getXShardCount()
    return False


def init():
    """
    initialize the contract, init contract meta data
    :return:
    """
    allShard = True
    frozen = False
    shardId = GetShardId()
    res = InitMetaData(OWNER, allShard, frozen, shardId, [])
    assert (res)


def notifyCallee(argsByteArray):
    """

    :param argsByteArray: byte array
    :return:
    """
    notifyCount = Get(ctx, X_SHARD_NOTIFY_KEY)
    notifyCount += 1
    Put(ctx, X_SHARD_NOTIFY_KEY, notifyCount)
    args = Deserialize(argsByteArray)
    a = args[0]
    b = args[1]
    Notify([a, b, notifyCount])
    return True


def invokeCallee(argsByteArray):
    """

    :param argsByteArray: byte array
    :return:
    """
    invokeCount = Get(ctx, X_SHARD_INVOKE_KEY)
    invokeCount += 1
    Put(ctx, X_SHARD_INVOKE_KEY, invokeCount)
    args = Deserialize(argsByteArray)
    a = args[0]
    b = args[1]
    Notify([a, b, invokeCount])
    res = Serialize(True)
    return res


def getXShardCount():
    """
    query x-shard call count, pre-execute
    :return:
    """
    invokeCount = Get(ctx, X_SHARD_INVOKE_KEY)
    notifyCount = Get(ctx, X_SHARD_NOTIFY_KEY)
    return [str(invokeCount), str(notifyCount)]
