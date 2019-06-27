./ontology contract invoke --address 0d9b18e994330002d823cd9809543cada9a5a2c1 \
--params string:xshardNotify,[int:0,int:1] --gasprice 0 --gaslimit 3000000

./ontology contract invoke --address 0d9b18e994330002d823cd9809543cada9a5a2c1 \
--params string:xshardInvoke,[int:0,int:1] --gasprice 0 --gaslimit 3000000 --rpcport 30336 --ShardID 1

sleep 20

./ontology contract invoke --address 88493a7ebae5e0431854f3f0b7e8f791f5e2d089 \
--params string:getXShardCount,[int:0] --gasprice 0 --rpcport 40336 --prepare