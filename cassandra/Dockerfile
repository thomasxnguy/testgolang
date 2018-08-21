FROM cassandra:3.10

COPY cassandra.yaml /etc/cassandra/cassandra.yaml
COPY records.cql /records.cql
COPY script.cql /script.cql

VOLUME /var/lib/cassandra

# 7000: intra-node communication
# 7001: TLS intra-node communication
# 7199: JMX
# 9042: CQL
# 9160: thrift service
EXPOSE 7000 7001 7199 9042 9160
CMD ["cassandra", "-f"]
