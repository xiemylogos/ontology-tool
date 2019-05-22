OntCversion = '2.0.0'
"""
An Example of OEP-4
"""
from ontology.interop.Ontology.Shard import GetShardId
from ontology.interop.Ontology.Contract import Migrate, InitMetaData
from ontology.interop.Ontology.Native import Invoke
from ontology.interop.Ontology.Runtime import Base58ToAddress
from ontology.interop.System.Runtime import Notify, CheckWitness
from ontology.interop.System.Storage import GetContext

ctx = GetContext()

NAME = 'MyToken'
SYMBOL = 'MYT'
DECIMALS = 8
FACTOR = 100000000
OWNER = Base58ToAddress("AKuMYaCm7LeBHqNeKzvj7qQb3USakDr5yg")
TOTAL_AMOUNT = 1000000000
BALANCE_PREFIX = bytearray(b'\x01')
APPROVE_PREFIX = b'\x02'
SUPPLY_KEY = 'TotalSupply'

# xshard asset contract Big endian Script Hash: 0x0900000000000000000000000000000000000000
XSHARD_ASSET_ADDR = Base58ToAddress("AFmseVrdL9f9oyCzZefL9tG6UbviRj6Fv6")

SHARD_VERSION = 1


def Main(operation, args):
    """
    :param operation:
    :param args:
    :return:
    """
    # 'init' has to be invokded first after deploying the contract to store the necessary and important info in the blockchain
    if operation == 'init':
        return init()
    if operation == 'name':
        return name()
    if operation == 'symbol':
        return symbol()
    if operation == 'decimals':
        return decimals()
    if operation == 'totalSupply':
        return totalSupply()
    if operation == 'shardSupply':
        return shardSupply()
    if operation == 'wholeSupply':
        return wholeSupply()
    if operation == 'supplyInfo':
        return supplyInfo()
    if operation == 'balanceOf':
        if len(args) != 1:
            return False
        acct = args[0]
        return balanceOf(acct)
    if operation == 'assetId':
        return assetId()
    if operation == 'transfer':
        if len(args) != 3:
            return False
        else:
            from_acct = args[0]
            to_acct = args[1]
            amount = args[2]
            return transfer(from_acct, to_acct, amount)
    if operation == 'transferMulti':
        return transferMulti(args)
    if operation == 'transferFrom':
        if len(args) != 4:
            return False
        spender = args[0]
        from_acct = args[1]
        to_acct = args[2]
        amount = args[3]
        return transferFrom(spender, from_acct, to_acct, amount)
    if operation == 'approve':
        if len(args) != 3:
            return False
        owner = args[0]
        spender = args[1]
        amount = args[2]
        return approve(owner, spender, amount)
    if operation == 'allowance':
        if len(args) != 2:
            return False
        owner = args[0]
        spender = args[1]
        return allowance(owner, spender)
    if operation == 'mint':
        if len(args) != 2:
            return False
        return mint(args[0], args[1])
    if operation == 'burn':
        if len(args) != 2:
            return False
        return burn(args[0], args[1])
    if operation == 'xshardTransfer':
        if len(args) != 4:
            return False
        return xshardTransfer(args[0], args[1], args[2], args[3])
    if operation == 'xshardTransferRetry':
        if len(args) != 2:
            return False
        return xshardTransferRetry(args[0], args[1])
    if operation == 'getXShardTransferDetail':
        if len(args) != 2:
            return False
        return getXShardTransferDetail(args[0], args[1])
    if operation == 'getPendignXShardTransfer':
        if len(args) != 1:
            return False
        return getPendignXShardTransfer(args[0])
    if operation == 'xshardTransferOng':
        if len(args) != 4:
            return False
        return xshardTransferOng(args[0], args[1], args[2], args[3])
    if operation == 'xshardTransferOngRetry':
        if len(args) != 2:
            return False
        return xshardTransferOngRetry(args[0], args[1])
    if operation == "migrate":
        if len(args) != 7:
            return False
        return migrate(args[0], args[1], args[2], args[3], args[4], args[5], args[6])
    return False


def init():
    """
    initialize the contract, call xshard asset contract register function
    :return:
    """

    if len(OWNER) != 20:
        Notify(["Owner illegal!"])
        return False
    allShard = True
    frozen = False
    shardId = 0
    res = InitMetaData(OWNER, allShard, frozen, shardId)
    assert (res)
    param = state(NAME, SYMBOL, DECIMALS, TOTAL_AMOUNT * FACTOR, OWNER)
    registerRes = Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Register', param)
    assert (registerRes)
    shardId = GetShardId()
    Notify(["Register success, current shard", shardId])


def name():
    """
    :return: name of the token
    """
    return NAME


def symbol():
    """
    :return: symbol of the token
    """
    return SYMBOL


def decimals():
    """
    :return: the decimals of the token
    """
    return DECIMALS


def totalSupply():
    """
    :return: query total supply, if invoked at shard, there are no value
    """
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4TotalSupply', [])


def shardSupply():
    """
    :return: query shard supply at root
    """
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4ShardSupply', [])


def wholeSupply():
    """
    :return: sum supply at all shard, only can be invoked at root
    """
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4WholeSupply', [])


def supplyInfo():
    """
    :return: query every shard supply at root
    """
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4SupplyInfo', [])


def balanceOf(account):
    """
    :param account:
    :return: the token balance of account
    """
    if len(account) != 20:
        raise Exception("address length error")
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4BalanceOf', account)


def assetId():
    """
    each xshard asset own a unique asst id
    :return: xshard asset id, big integer
    """
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4AssetId', [])


def transfer(from_acct, to_acct, amount):
    """
    Transfer amount of tokens from from_acct to to_acct
    :param from_acct: the account from which the amount of tokens will be transferred
    :param to_acct: the account to which the amount of tokens will be transferred
    :param amount: the amount of the tokens to be transferred, >= 0
    :return: True means success, False or raising exception means failure.
    """
    if len(to_acct) != 20 or len(from_acct) != 20:
        raise Exception("address length error")
    param = state(from_acct, to_acct, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Transfer', param)


def transferMulti(args):
    """
    :param args: the parameter is an array, containing element like [from, to, amount]
    :return: True means success, False or raising exception means failure.
    """
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4TransferMulti', args)


def approve(owner, spender, amount):
    """
    owner allow spender to spend amount of token from owner account
    Note here, the amount should be less than the balance of owner right now.
    :param owner:
    :param spender:
    :param amount: amount>=0
    :return: True means success, False or raising exception means failure.
    """
    if len(spender) != 20 or len(owner) != 20:
        raise Exception("address length error")
    param = state(owner, spender, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Approve', param)


def transferFrom(spender, from_acct, to_acct, amount):
    """
    spender spends amount of tokens on the behalf of from_acct, spender makes a transaction of amount of tokens
    from from_acct to to_acct
    :param spender:
    :param from_acct:
    :param to_acct:
    :param amount:
    :return:
    """
    if len(spender) != 20 or len(from_acct) != 20 or len(to_acct) != 20:
        raise Exception("address length error")
    param = state(spender, from_acct, to_acct, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4TransferFrom', param)


def allowance(owner, spender):
    """
    check how many token the spender is allowed to spend from owner account
    :param owner: token owner
    :param spender:  token spender
    :return: the allowed amount of tokens
    """
    param = state(owner, spender)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Allowance', param)


def mint(user, amount):
    """
    need check witness by myself
    :param user: 
    :param amount: 
    :return: 
    """
    if CheckWitness(OWNER) == False:
        return False
    param = state(user, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Mint', param)


def burn(user, amount):
    param = state(user, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Burn', param)


def xshardTransfer(from_acc, to_acc, to_shard, amount):
    """
    cross shard transfer asset
    :param from_acc: 
    :param to_acc: 
    :param to_shard: 
    :param amount: 
    :return: xshard transfer id
    """
    param = state(from_acc, to_acc, to_shard, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4XShardTransfer', param)


def xshardTransferRetry(from_acc, transferId):
    """
    if cross shard transfer failed, invoke this method to retry
    :param from_acc: 
    :param transferId: xshard transfer id
    :return: 
    """
    param = state(from_acc, transferId)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4XShardTransferRetry', param)


def getXShardTransferDetail(user, transferId):
    """
    query user xshard transfer detail
    :param user: 
    :param transferId: 
    :return: xshard transfer info
    """
    param = state(user, assetId(), transferId)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'getOep4Transfer', param)


def getPendignXShardTransfer(user):
    """
    get all pending transfer from user
    :param user: 
    :return: 
    """
    param = state(user, assetId())
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'getOep4PendingTransfer', param)


"""
the follow is how to xshard transfer ong
"""


def xshardTransferOng(from_acc, to_acc, to_shard, amount):
    """
    cross shard transfe ong
    :param from_acc: 
    :param to_acc: 
    :param to_shard: 
    :param amount: 
    :return: xshard transfer id
    """
    param = state(from_acc, to_acc, to_shard, amount)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'ongXShardTransfer', param)


def xshardTransferOngRetry(from_acc, transferId):
    """
    if cross shard transfer failed, invoke this method to retry
    :param from_acc: 
    :param transferId: xshard transfer id
    :return: 
    """
    param = state(from_acc, transferId)
    return Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'ongXShardTransferRetry', param)


def migrate(code, needStorage, name, version, author, email, description):
    newAddr = AddressFromVmCode(code)
    res = Invoke(SHARD_VERSION, XSHARD_ASSET_ADDR, 'oep4Migrate', newAddr)
    assert (res)
    res = Migrate(code, needStorage, name, version, author, email, description)
    assert (res)
    Notify(["Migrate successfully"])
    return True


def AddressFromVmCode(code):
    from ontology.builtins import hash160
    Address = None
    assert (len(code) > 0)
    addr = hash160(code)

    for i in reversed(range(0, 21)):
        if i < 1:
            break
        Address = concat(Address, addr[i - 1:i])

    return Address
