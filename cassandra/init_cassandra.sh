
#!/bin/bash

docker exec challenger_cassandra cqlsh -u cassandra -p cassandra -f script.cql
docker exec challenger_cassandra cqlsh -u cassandra -p cassandra -f records.cql


