ALTER TABLE IF EXISTS "chat_messages" DROP CONSTRAINT IF EXISTS "chat_messages_chat_room_id_fkey";

ALTER TABLE IF EXISTS "chat_subscriptions" DROP CONSTRAINT IF EXISTS "chat_subscriptions_chat_room_id_fkey";

DROP TABLE IF EXISTS "chat_rooms";
DROP TABLE IF EXISTS "chat_subscriptions";
DROP TABLE IF EXISTS "chat_messages";