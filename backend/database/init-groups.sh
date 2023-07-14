#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE "public"."groups" (
        "id" serial  PRIMARY KEY NOT NULL,
        "group_id" int UNIQUE NOT NULL,
        "group_address" character varying(200) NOT NULL,
        "photo_url" character varying(200) NOT NULL,
        "name" character varying(200) NOT NULL,
        "full_address" character varying(200) NOT NULL);
EOSQL




