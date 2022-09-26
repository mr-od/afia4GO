CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "status" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "order_items" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE INDEX ON "order_items" ("owner");
