#!/bin/bash -x

sed -i -e 's/localhost:[0-9]*/localhost:20336/g' config_test.json

./main -t ShardInit
sleep 5

# create shard-1
./main -t ShardCreate
sleep 5

# create shard-2
./main -t ShardCreate
sleep 5

# config shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardConfig.json
./main -t ShardConfig
sleep 5

# config shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardConfig.json
./main -t ShardConfig
sleep 5

# join shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardPeerJoin.json
./main -t ShardPeerJoin
sleep 5

# join shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardPeerJoin.json
./main -t ShardPeerJoin
sleep 5

# activate shard-1
sed -i -e 's/"shard_id".*/"shard_id": 1/g' params/ShardActivate.json
./main -t ShardActivate
sleep 5

# activate shard-2
sed -i -e 's/"shard_id".*/"shard_id": 2/g' params/ShardActivate.json
./main -t ShardActivate
sleep 5


