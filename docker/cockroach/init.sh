#!/bin/sh

[ "$(/cockroach/cockroach sql --host things-db --insecure --execute='show users;' | grep -w ${POSTGRES_USER:=mainflux})" ] && exit 0;

# Initialise the cockroach users and databases
# Meant to be executed via docker exec in a systemd PostStart section
cat > /tmp/create.sql <<EOL
CREATE DATABASE IF NOT EXISTS ${POSTGRES_DB:=things};
CREATE USER ${POSTGRES_USER:=mainflux};
GRANT ALL ON DATABASE ${POSTGRES_DB:=mainflux} TO ${POSTGRES_USER:=mainflux};
SET DATABASE = ${POSTGRES_DB:=things};
CREATE TABLE IF NOT EXISTS things (id BIGSERIAL, owner VARCHAR(254), type VARCHAR(10) NOT NULL, key CHAR(36) UNIQUE NOT NULL, name TEXT, metadata TEXT, PRIMARY KEY (id, owner));
CREATE TABLE IF NOT EXISTS channels (id BIGSERIAL, owner VARCHAR(254), name TEXT, PRIMARY KEY (id, owner));
CREATE TABLE IF NOT EXISTS connections (channel_id BIGINT, channel_owner VARCHAR(254), thing_id BIGINT, thing_owner VARCHAR(254),FOREIGN KEY (channel_id, channel_owner) REFERENCES channels (id, owner) ON DELETE CASCADE ON UPDATE CASCADE, FOREIGN KEY (thing_id, thing_owner) REFERENCES things (id, owner) ON DELETE CASCADE ON UPDATE CASCADE, PRIMARY KEY (channel_id, channel_owner, thing_id, thing_owner));
EOL

/cockroach/cockroach sql --host things-db --insecure < /tmp/create.sql && rm /tmp/create.sql