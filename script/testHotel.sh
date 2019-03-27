#!/bin/bash

PASSWD="1"

# test on shard-1
sed -i -e 's/localhost:[0-9]*/localhost:20346/g' config_test.json
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardHotelInit.json
echo $PASSWD | ./main -t ShardHotelInit
sleep 5

sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardHotelReserve.json
echo $PASSWD | ./main -t ShardHotelReserve
sleep 5

# test on shard-2
sed -i -e 's/localhost:[0-9]*/localhost:20356/g' config_test.json
sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardHotelInit.json
echo $PASSWD | ./main -t ShardHotelInit
sleep 5

sed -i -e 's/"shard_id".*/"shard_id": 2,/g' params/ShardHotelReserve.json
echo $PASSWD | ./main -t ShardHotelReserve
sleep 5

sed -i -e 's/localhost:[0-9]*/localhost:20336/g' config_test.json

