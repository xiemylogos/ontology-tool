#!/bin/bash

# test on shard-1
sed -i -e 's/localhost:[0-9]*/localhost:20346/g' config_test.json
sed -i -e 's/"shard_id".*/"shard_id": 1,/g' params/ShardHotelReserve2.json
sed -i -e 's/"shard_id_2".*/"shard_id_2": 2,/g' params/ShardHotelReserve2.json
./main -t ShardHotelReserve2
sleep 5

sed -i -e 's/localhost:[0-9]*/localhost:20336/g' config_test.json

