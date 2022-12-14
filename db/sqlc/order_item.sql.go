// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: order_item.sql

package db

import (
	"context"
)

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (
  owner,
  order_id,
  product_id,
  status,
  quantity
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, owner, order_id, product_id, status, quantity, created_at
`

type CreateOrderItemParams struct {
	Owner     string `json:"owner"`
	OrderID   int64  `json:"order_id"`
	ProductID int64  `json:"product_id"`
	Status    string `json:"status"`
	Quantity  int64  `json:"quantity"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.Owner,
		arg.OrderID,
		arg.ProductID,
		arg.Status,
		arg.Quantity,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.OrderID,
		&i.ProductID,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOrderItem = `-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1
`

func (q *Queries) DeleteOrderItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrderItem, id)
	return err
}

const getOrderItem = `-- name: GetOrderItem :one
SELECT id, owner, order_id, product_id, status, quantity, created_at FROM order_items
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrderItem(ctx context.Context, id int64) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderItem, id)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.OrderID,
		&i.ProductID,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const getOrderItemForUpdate = `-- name: GetOrderItemForUpdate :one
SELECT id, owner, order_id, product_id, status, quantity, created_at FROM order_items
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetOrderItemForUpdate(ctx context.Context, id int64) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderItemForUpdate, id)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.OrderID,
		&i.ProductID,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const listOrderItems = `-- name: ListOrderItems :many
SELECT id, owner, order_id, product_id, status, quantity, created_at FROM order_items
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListOrderItemsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListOrderItems(ctx context.Context, arg ListOrderItemsParams) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrderItems, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.OrderID,
			&i.ProductID,
			&i.Status,
			&i.Quantity,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderItem = `-- name: UpdateOrderItem :one
UPDATE order_items
SET status = $2
WHERE id = $1
RETURNING id, owner, order_id, product_id, status, quantity, created_at
`

type UpdateOrderItemParams struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, updateOrderItem, arg.ID, arg.Status)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.OrderID,
		&i.ProductID,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}
