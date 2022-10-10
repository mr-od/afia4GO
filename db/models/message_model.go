package model

import (
	"time"

	"github.com/google/uuid"
	db "github.com/oddinnovate/a4go/db/sqlc"
)

type ChatMessageDTO struct {
	ID         string        `json:"id"`
	ChatRoomID string        `json:"chatRoomId"`
	CreatedAt  time.Time     `json:"createdAt"`
	Body       string        `json:"body"`
	User       PublicUserDTO `json:"user"`
}

func NewChatMessageDTO(id, roomID, userID, username, body string, createdAt time.Time) ChatMessageDTO {
	return ChatMessageDTO{
		ID:         id,
		ChatRoomID: roomID,
		CreatedAt:  createdAt,
		Body:       body,
		User:       NewPublicUserDTO(userID, username),
	}
}

type ChatMessageRequest struct {
	Body        string `json:"body"`
	UserID      uint   `json:"userId"`
	ChatRoomID  uint   `json:"chatRoomId"`
	MessageType string `json:"messageType"`
}

func NewChatMessage(body string, userID, chatRoomID int64) ChatMessage {
	return ChatMessage{
		PublicID:   uuid.NewString(),
		Body:       body,
		UserID:     userID,
		ChatRoomID: chatRoomID,
	}
}

type ChatMessage struct {
	PublicID   string
	Body       string
	User       db.User
	UserID     int64
	ChatRoom   ChatRoom
	ChatRoomID int64
}
