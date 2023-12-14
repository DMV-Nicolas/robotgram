CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "avatar" varchar NOT NULL,
  "description" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "images" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "url" varchar NOT NULL
);

CREATE INDEX ON "posts" ("owner");

CREATE INDEX ON "images" ("post_id");

ALTER TABLE "posts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "images" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
