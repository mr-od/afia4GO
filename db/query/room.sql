-- name: CreateRoom :one
INSERT INTO chat_rooms (
  owner,
  name,
  public_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetRoom :one
SELECT * FROM chat_rooms
WHERE name = $1 LIMIT 1;

-- name: GetRoomForUpdate :one
SELECT * FROM chat_rooms
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- -- name: ListRooms :many
-- SELECT * FROM chat_rooms
-- WHERE owner = $1
-- ORDER BY id
-- LIMIT $2
-- OFFSET $3;

-- name: ListRooms :many
SELECT * FROM chat_rooms
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRoom :one
UPDATE chat_rooms
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM chat_rooms
WHERE id = $1;
