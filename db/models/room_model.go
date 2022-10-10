package model

import db "github.com/oddinnovate/a4go/db/sqlc"

type ChatRoom struct {
	Name          string
	PublicID      string
	Owner         *db.User
	OwnerID       uint
	Subscriptions []*db.ChatSubscription
	Messages      []*db.ChatMessage
}
