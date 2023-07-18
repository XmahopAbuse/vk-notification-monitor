-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "public"."keywords" (
                                     "id" serial  PRIMARY KEY NOT NULL,
                                     "keyword" character varying(200) UNIQUE NOT NULL
);