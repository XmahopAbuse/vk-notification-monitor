-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "public"."notification" (
                                         "id" SERIAL NOT NULL PRIMARY KEY,
                                         "type" character varying(200) NOT NULL,
                                         "value" character varying(200) UNIQUE NOT NULL
);