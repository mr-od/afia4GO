CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner" varchar NOT NULL,
  "price" bigint NOT NULL,
  "description" varchar NOT NULL,
  "imgs_url" text[],
  "imgs" text[],
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "products" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX ON "products" ("owner");