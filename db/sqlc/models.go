// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID          int64     `json:"id"`
	Owner       string    `json:"owner"`
	Status      string    `json:"status"`
	DeliveryFee int64     `json:"delivery_fee"`
	Subtotal    int64     `json:"subtotal"`
	Total       int64     `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderItem struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	OrderID   int64     `json:"order_id"`
	ProductID int64     `json:"product_id"`
	Status    string    `json:"status"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	ImgsUrl     []string  `json:"imgs_url"`
	Imgs        []string  `json:"imgs"`
	CreatedAt   time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	IsBlocked    bool      `json:"is_blocked"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
