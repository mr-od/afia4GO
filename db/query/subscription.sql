-- name: CreateSubscription :one
INSERT INTO chat_subscriptions (
  chat_room_id,
  username
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetSubscription :one
SELECT * FROM chat_subscriptions
WHERE id = $1 LIMIT 1;

-- name: GetSubscriptionForUpdate :one
SELECT * FROM chat_subscriptions
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListSubscriptions :many
SELECT * FROM chat_subscriptions
WHERE chat_room_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateSubscription :one
UPDATE chat_subscriptions
SET chat_room_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSubscription :exec
DELETE FROM chat_subscriptions
WHERE id = $1;
