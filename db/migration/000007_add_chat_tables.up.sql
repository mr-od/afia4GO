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


ALTER TABLE "chat_subscriptions" ADD FOREIGN KEY ("chat_room_id") REFERENCES "chat_rooms" ("id");

ALTER TABLE "chat_subscriptions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "chat_messages" ADD FOREIGN KEY ("chat_room_id") REFERENCES "chat_rooms" ("id");

ALTER TABLE "chat_messages" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "chat_rooms" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
