ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";
ALTER TABLE IF EXISTS "chat_messages" DROP CONSTRAINT IF EXISTS "chat_messages_username_fkey";

ALTER TABLE IF EXISTS "chat_subscriptions" DROP CONSTRAINT IF EXISTS "chat_subscriptions_username_fkey";




DROP TABLE IF EXISTS "users"
