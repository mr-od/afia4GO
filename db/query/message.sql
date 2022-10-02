-- name: SaveMessage :one
INSERT INTO chat_messages (
  chat_room_id,
  username,
  public_id,
  body
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetMessage :one
SELECT * FROM chat_messages
WHERE id = $1 LIMIT 1;

-- name: GetMessageForUpdate :one
SELECT * FROM chat_messages
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetChatHistory :many
SELECT * FROM chat_messages
WHERE chat_room_id = $1;


-- name: ListMessages :many
SELECT * FROM chat_messages
WHERE chat_room_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateMessage :one
UPDATE chat_messages
SET chat_room_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM chat_messages
WHERE id = $1;
