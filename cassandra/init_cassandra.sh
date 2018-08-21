
#!/bin/bash

docker exec challenger_cassandra cqlsh -u cassandra -p cassandra -f challenger.cql
docker exec challenger_cassandra cqlsh -u cassandra -p cassandra -f devauthrecords.cql


