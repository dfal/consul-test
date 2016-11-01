curl -XPUT 172.20.20.10:9200/_template/logstash --output  /dev/null --silent -d '{ "template" : "logstash-*", "mappings" : { "_default_" : { "properties": { "@timestamp": { "type": "date"}, "CheckID": {"type":"string", "index":"not_analyzed"}, "Name" : {"type":"string","index":"not_analyzed"},"Status":{"type": "string", "index":"not_analyzed"}, "ServiceName":{"type":"string", "index":"not_analyzed"}  } } } }'
curl -sL https://deb.nodesource.com/setup_5.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo apt-get install -y npm
sudo npm install elasticdump -g
elasticdump --input=/n0/kibana-configuration.json --output=http://localhost:9200/.kibana --type=data