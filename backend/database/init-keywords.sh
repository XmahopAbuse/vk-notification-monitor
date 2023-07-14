#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE "public"."keywords" (
        "id" serial  PRIMARY KEY NOT NULL,
        "keyword" character varying(200) UNIQUE NOT NULL);
EOSQL




