-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "public"."groups" (
                                   "id" serial  PRIMARY KEY NOT NULL,
                                   "group_id" int UNIQUE NOT NULL,
                                   "group_address" character varying(200) NOT NULL,
                                   "photo_url" character varying(200) NOT NULL,
                                   "name" character varying(200) NOT NULL,
                                   "full_address" character varying(200) NOT NULL
);