cd ../ontology/
./ontology asset transfer --from 1 --to 1 --amount 1 --gasprice 0
./ontology contract deploy --needstore --code ../ontology-tool/params/shardasset/xshardasstdemo.avm.str --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
./ontology contract deploy --needstore --code ../ontology-tool/params/shardcontract/xshardcaller.avm.str --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
./ontology contract deploy --needstore --code ../ontology-tool/params/shardcontract/xshardcallee.avm.str --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
sleep 10
# xshard asset init
./ontology contract invoke --address a4b7710d5286352c4a23a51eef5e87d02c590617 --params string:init,[int:0] --gasprice 0 --gaslimit 3000000
# xshard callee init
./ontology contract invoke --address 88493a7ebae5e0431854f3f0b7e8f791f5e2d089 --params string:init,[int:0] --gasprice 0 --gaslimit 3000000
sleep 10
# caller dependent callee
./ontology contract invoke --address 0d9b18e994330002d823cd9809543cada9a5a2c1 --params string:init,[int:0] --gasprice 0 --gaslimit 3000000
./ontology asset withdrawong 1 --gasprice 0
./ontology asset transfer --from 1 --to AZqk4i7Zhfhc1CRUtZYKrLw4YTSq4Y9khN --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to ARpjnrnHEjXhg4aw7vY6xsY6CfQ1XEWzWC --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AQs2BmzzFVk7pQPfTQQi9CTEz43ejSyBnt --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AKBSRLbFNvUrWEGtKxNTpe2ZdkepQjYKfM --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AduX7odaWGipkdvzBwyaTgsumRbRzhhiwe --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to ANFfWhk3A5iFXQrVBHKrerjDDapYmLo5Bi --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AK3YRcRvKrASQ6nTfW48Z4iMZ2sDTDRiMC --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AN9PD1zC4moFWjDzY4xG9bAr7R7UvHwmLL --amount 100000 --gasprice 0
sleep 10

./ontology asset transfer --asset ong --from 1 --to AZqk4i7Zhfhc1CRUtZYKrLw4YTSq4Y9khN --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to ARpjnrnHEjXhg4aw7vY6xsY6CfQ1XEWzWC --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to AQs2BmzzFVk7pQPfTQQi9CTEz43ejSyBnt --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to AKBSRLbFNvUrWEGtKxNTpe2ZdkepQjYKfM --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to AduX7odaWGipkdvzBwyaTgsumRbRzhhiwe --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to ANFfWhk3A5iFXQrVBHKrerjDDapYmLo5Bi --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to AK3YRcRvKrASQ6nTfW48Z4iMZ2sDTDRiMC --amount 1000 --gasprice 0
./ontology asset transfer --asset ong --from 1 --to AN9PD1zC4moFWjDzY4xG9bAr7R7UvHwmLL --amount 1000 --gasprice 0
sleep 10
cd ../ontology-tool
./main -t ShardInit
./main -t ShardCreate
./main -t ShardConfig
./main -t ShardPeerApply
./main -t ShardPeerApprove
./main -t ShardPeerJoin
./main -t ShardActivate
./main -t ShardAssetInit