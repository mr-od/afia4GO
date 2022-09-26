CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "status" varchar NOT NULL,
  "delivery_fee" bigint NOT NULL,
  "subtotal" bigint NOT NULL,
  "total" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
ALTER TABLE "orders" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX ON "orders" ("owner");
