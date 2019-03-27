#!/bin/bash -x

WALLET="../ontology/wallet.dat"
PASSWD="1"
ONTOLOGY="../ontology/ontology"

PK=`$ONTOLOGY account list -v -w $WALLET | grep 'Public key:' | awk '{print $3}'`
echo $PK

# withdraw ong
echo $PASSWD | $ONTOLOGY asset transfer  --from 1 --to 1 --amount 1 -w $WALLET
sleep 10
echo $PASSWD | $ONTOLOGY asset withdrawong 1 -w $WALLET
sleep 10

sed -i -e 's/localhost:[0-9]*/localhost:20336/g' config_test.json
echo $PASSWD | ./main -t ShardInit
sleep 5

# create shard-1
echo $PASSWD | ./main -t ShardCreate
sleep 5

# create shard-2
echo $PASSWD | ./main -t ShardCreate
sleep 5

# config shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardConfig.json
echo $PASSWD | ./main -t ShardConfig
sleep 5

# config shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardConfig.json
echo $PASSWD | ./main -t ShardConfig
sleep 5

# peer apply for shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardPeerApplyJoin.json
sed -i -e "s/\"peer_pub_key\".[*]/\"peer_pub_key\": [\"$PK\"]/g" params/ShardPeerApplyJoin.json
echo $PASSWD | ./main -t ShardPeerApply
sleep 5

# peer apply for shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardPeerApplyJoin.json
sed -i -e "s/\"peer_pub_key\".[*]/\"peer_pub_key\": [\"$PK\"]/g" params/ShardPeerApplyJoin.json
echo $PASSWD | ./main -t ShardPeerApply
sleep 5

# approve peer for shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardPeerApproveJoin.json
sed -i -e "s/\"peer_pub_key\".[*]/\"peer_pub_key\": [\"$PK\"]/g" params/ShardPeerApproveJoin.json
echo $PASSWD | ./main -t ShardPeerApprove
sleep 5

# approve peer for shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardPeerApproveJoin.json
sed -i -e "s/\"peer_pub_key\".[*]/\"peer_pub_key\": [\"$PK\"]/g" params/ShardPeerApproveJoin.json
echo $PASSWD | ./main -t ShardPeerApprove
sleep 5

# join shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardPeerJoin.json
sed -i -e "s/\"peer_pub_key\".[*]/\"peer_pub_key\": \"$PK\",/g" params/ShardPeerJoin.json
echo $PASSWD | ./main -t ShardPeerJoin
sleep 5

# join shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardPeerJoin.json
sed -i -e "s/\"peer_pub_key\".[*]/\"peer_pub_key\": \"$PK\",/g" params/ShardPeerJoin.json
echo $PASSWD | ./main -t ShardPeerJoin
sleep 5

# activate shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1/g' params/ShardActivate.json
echo $PASSWD | ./main -t ShardActivate
sleep 5

# activate shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2/g' params/ShardActivate.json
echo $PASSWD | ./main -t ShardActivate
sleep 5

# gas ini
echo $PASSWD | ./main -t ShardGasInit
sleep 5

# deposit gas to shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardDepositGas.json
echo $PASSWD | ./main -t ShardDepositGas
sleep 5


# deposit gas to shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardDepositGas.json
echo $PASSWD | ./main -t ShardDepositGas
sleep 5


