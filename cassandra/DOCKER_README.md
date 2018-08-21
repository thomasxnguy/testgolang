- Run the docker compose file with `docker-compose up` to start up zookeeper,
kafka and cassandra (only docker needs to be installed).
Use `docker-compose -f docker-compose-windows.yaml up -d` when you are on Windows 7 (with Docker toolbox).
Then you might want to use `challenger-test-windows.conf` for configuration.

- If the cassandra database needs to be initialised (if it's being started for the first time)
then wait until it has completely started up and then run the bash script init_cassandra.sh
with `sh init_cassandra.sh`.

- To stop them all run `docker-compose down`.
