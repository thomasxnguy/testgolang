DROP USER IF EXISTS test_user;
CREATE USER IF NOT EXISTS test_user WITH PASSWORD 'password' NOSUPERUSER;

CREATE KEYSPACE IF NOT EXISTS test_db WITH
REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 } AND 
DURABLE_WRITES = true;

GRANT AUTHORIZE ON KEYSPACE test_db TO test_user;
GRANT CREATE ON KEYSPACE test_db TO test_user;
GRANT SELECT ON KEYSPACE test_db TO test_user;
GRANT MODIFY ON KEYSPACE test_db TO test_user;
GRANT DROP ON KEYSPACE test_db TO test_user;
GRANT ALTER ON KEYSPACE test_db TO test_user;

USE test_db;

DROP TABLE IF EXISTS challenger.selector;
DROP TABLE IF EXISTS challenger.auth;
DROP TABLE IF EXISTS challenger.ipbl;
DROP TABLE IF EXISTS challenger.ipbl_up;
DROP TABLE IF EXISTS challenger.idbl;
DROP TABLE IF EXISTS challenger.idbl_up;
DROP TABLE IF EXISTS challenger.easyidbl;
DROP TABLE IF EXISTS challenger.easyidbl_up;
DROP TABLE IF EXISTS challenger.rtokenbl;
DROP TABLE IF EXISTS challenger.rtokenbl_up;
DROP TABLE IF EXISTS challenger.useridbl;
DROP TABLE IF EXISTS challenger.useridbl_up;
DROP TABLE IF EXISTS challenger.combinedbl;
DROP TABLE IF EXISTS challenger.combinedbl_up;

CREATE TABLE IF NOT EXISTS test_db.TABLEA
(
	pid UUID,
	a INT,
	b INT,
	ts TIMESTAMP,
	name TEXT,
	PRIMARY KEY (pid)
);

CREATE TABLE IF NOT EXISTS test_db.TABLEB
(
	sid UUID,
	ts TIMESTAMP,
	permissions INT,
	name TEXT,
	PRIMARY KEY (sid)
);