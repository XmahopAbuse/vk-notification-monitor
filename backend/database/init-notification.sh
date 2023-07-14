package database


set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE TABLE notification (
"id" SERIAL NOT NULL PRIMARY KEY,
"type" character varying(200) NOT NULL,
"value" character varying(200) UNIQUE NOT NULL
);
EOSQL

