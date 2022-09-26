-- name: CreateProduct :one
INSERT INTO products (
  owner,
  name,
  price,
  description,
  imgs_url,
  imgs
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListProducts :many
SELECT * FROM products
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateProduct :one
UPDATE products
SET price = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
