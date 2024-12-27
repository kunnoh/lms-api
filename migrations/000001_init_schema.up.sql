CREATE TYPE "user_roles" AS ENUM (
  'admin',
  'moderator'
  'user'
);

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid ()),
    "email" text,
    "user_role" user_roles NOT NULL DEFAULT 'user',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "users" ("id");

CREATE TABLE "books" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid ()),
    "descriptions" text,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "books" ("id");