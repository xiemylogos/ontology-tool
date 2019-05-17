cd ../ontology/
./ontology asset transfer --from 1 --to 1 --amount 1 --gasprice 0
sleep 10
./ontology asset withdrawong 1 --gasprice 0
./ontology asset transfer --from 1 --to AZqk4i7Zhfhc1CRUtZYKrLw4YTSq4Y9khN --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to ARpjnrnHEjXhg4aw7vY6xsY6CfQ1XEWzWC --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AQs2BmzzFVk7pQPfTQQi9CTEz43ejSyBnt --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AKBSRLbFNvUrWEGtKxNTpe2ZdkepQjYKfM --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AduX7odaWGipkdvzBwyaTgsumRbRzhhiwe --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to ANFfWhk3A5iFXQrVBHKrerjDDapYmLo5Bi --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AK3YRcRvKrASQ6nTfW48Z4iMZ2sDTDRiMC --amount 100000 --gasprice 0
./ontology asset transfer --from 1 --to AN9PD1zC4moFWjDzY4xG9bAr7R7UvHwmLL --amount 100000 --gasprice 0
cd ../ontology-tool
./main -t ShardInit
./main -t ShardCreate
./main -t ShardConfig
./main -t ShardPeerApply
./main -t ShardPeerApprove
./main -t ShardPeerJoin
./main -t ShardActivate
./main -t ShardAssetInit