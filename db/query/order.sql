-- name: CreateOrder :one
INSERT INTO orders (
  owner,
  status,
  delivery_fee,
  subtotal,
  total
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrderForUpdate :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListOrders :many
SELECT * FROM orders
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateOrder :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;
