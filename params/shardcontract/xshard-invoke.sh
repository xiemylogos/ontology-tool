./ontology contract invoke --address 839b3d9e93b20dd4106090035a14b260a890cca1 \
--params string:xshardNotify,[int:0,int:1] --gasprice 0 --gaslimit 3000000

./ontology contract invoke --address 839b3d9e93b20dd4106090035a14b260a890cca1 \
--params string:xshardInvoke,[int:0,int:1] --gasprice 0 --gaslimit 3000000 --rpcport 30336 --ShardID 1

sleep 20

./ontology contract invoke --address a95ee9e4c5ed9f927dda46c29af9f203195e40a1 \
--params string:getXShardCount,[int:0] --gasprice 0 --rpcport 40336 --prepare