-- name: CreateOrderItem :one
INSERT INTO order_items (
  owner,
  order_id,
  product_id,
  status,
  quantity
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE id = $1 LIMIT 1;

-- name: GetOrderItemForUpdate :one
SELECT * FROM order_items
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListOrderItems :many
SELECT * FROM order_items
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateOrderItem :one
UPDATE order_items
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;
