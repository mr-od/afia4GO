-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-10-02T03:09:55.501Z

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "public_id" varchar UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner" varchar NOT NULL,
  "price" bigint NOT NULL,
  "description" varchar NOT NULL,
  "imgs_url" text[],
  "imgs_name" text[],
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "status" varchar NOT NULL,
  "delivery_fee" bigint NOT NULL,
  "subtotal" bigint NOT NULL,
  "total" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "status" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "chat_rooms" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "public_id" varchar UNIQUE NOT NULL,
  "owner" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "chat_subscriptions" (
  "id" bigserial PRIMARY KEY,
  "chat_room_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "chat_messages" (
  "id" bigserial PRIMARY KEY,
  "chat_room_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "public_id" varchar NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

CREATE INDEX ON "products" ("owner");

CREATE INDEX ON "orders" ("owner");

CREATE INDEX ON "order_items" ("owner");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "orders" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "order_items" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "chat_rooms" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "chat_subscriptions" ADD FOREIGN KEY ("chat_room_id") REFERENCES "chat_rooms" ("id");

ALTER TABLE "chat_subscriptions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "chat_messages" ADD FOREIGN KEY ("chat_room_id") REFERENCES "chat_rooms" ("id");

ALTER TABLE "chat_messages" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
