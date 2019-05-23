./main -t TransferOntMultiSign
./main -t TransferFromOngMultiSign
cd ../ontology
./ontology contract deploy --needstore --code ../ontology-tool/params/shardasset/xshardasstdemo.avm --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
sleep 30
./ontology contract invoke --address c53729d41eb1953b4d7ed66d421dd58462a8cf9d --params string:init,[int:0] --gasprice 0
cd ../ontology-tool
./main -t ShardInit
./main -t ShardCreate
./main -t ShardConfig
./main -t ShardPeerApply
./main -t ShardPeerApprove
./main -t ShardPeerJoin
./main -t ShardActivate
./main -t ShardAssetInit