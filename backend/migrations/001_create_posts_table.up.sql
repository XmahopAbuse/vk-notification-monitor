-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "posts" (
                         "id" serial PRIMARY KEY NOT NULL,
                         "vk_id" integer,
                         "author" character varying(200) NOT NULL,
                         "text" character varying(4000),
                         "author_id" character varying(200),
                         "group_id" character varying(200),
                         "hash" character varying(200) UNIQUE NOT NULL,
                         "status" boolean DEFAULT false,
                         "from_id" integer,
                         "post_url" character varying(200),
                         "date" TIMESTAMP WITH TIME ZONE NOT NULL
);