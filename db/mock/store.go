// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oddinnovate/a4go/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/oddinnovate/a4go/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(arg0 context.Context, arg1 db.CreateEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 db.CreateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreateOrderItem mocks base method.
func (m *MockStore) CreateOrderItem(arg0 context.Context, arg1 db.CreateOrderItemParams) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItem indicates an expected call of CreateOrderItem.
func (mr *MockStoreMockRecorder) CreateOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockStore)(nil).CreateOrderItem), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateRoom mocks base method.
func (m *MockStore) CreateRoom(arg0 context.Context, arg1 db.CreateRoomParams) (db.ChatRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", arg0, arg1)
	ret0, _ := ret[0].(db.ChatRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockStoreMockRecorder) CreateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockStore)(nil).CreateRoom), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateSubscription mocks base method.
func (m *MockStore) CreateSubscription(arg0 context.Context, arg1 db.CreateSubscriptionParams) (db.ChatSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", arg0, arg1)
	ret0, _ := ret[0].(db.ChatSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockStoreMockRecorder) CreateSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockStore)(nil).CreateSubscription), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteMessage mocks base method.
func (m *MockStore) DeleteMessage(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockStoreMockRecorder) DeleteMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockStore)(nil).DeleteMessage), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0, arg1)
}

// DeleteOrderItem mocks base method.
func (m *MockStore) DeleteOrderItem(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderItem indicates an expected call of DeleteOrderItem.
func (mr *MockStoreMockRecorder) DeleteOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderItem", reflect.TypeOf((*MockStore)(nil).DeleteOrderItem), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// DeleteRoom mocks base method.
func (m *MockStore) DeleteRoom(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockStoreMockRecorder) DeleteRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockStore)(nil).DeleteRoom), arg0, arg1)
}

// DeleteSubscription mocks base method.
func (m *MockStore) DeleteSubscription(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription.
func (mr *MockStoreMockRecorder) DeleteSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockStore)(nil).DeleteSubscription), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetChatHistory mocks base method.
func (m *MockStore) GetChatHistory(arg0 context.Context, arg1 int64) ([]db.ChatMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatHistory", arg0, arg1)
	ret0, _ := ret[0].([]db.ChatMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatHistory indicates an expected call of GetChatHistory.
func (mr *MockStoreMockRecorder) GetChatHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatHistory", reflect.TypeOf((*MockStore)(nil).GetChatHistory), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), arg0, arg1)
}

// GetMessage mocks base method.
func (m *MockStore) GetMessage(arg0 context.Context, arg1 int64) (db.ChatMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", arg0, arg1)
	ret0, _ := ret[0].(db.ChatMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockStoreMockRecorder) GetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockStore)(nil).GetMessage), arg0, arg1)
}

// GetMessageForUpdate mocks base method.
func (m *MockStore) GetMessageForUpdate(arg0 context.Context, arg1 int64) (db.ChatMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.ChatMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageForUpdate indicates an expected call of GetMessageForUpdate.
func (mr *MockStoreMockRecorder) GetMessageForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageForUpdate", reflect.TypeOf((*MockStore)(nil).GetMessageForUpdate), arg0, arg1)
}

// GetOrder mocks base method.
func (m *MockStore) GetOrder(arg0 context.Context, arg1 int64) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockStoreMockRecorder) GetOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockStore)(nil).GetOrder), arg0, arg1)
}

// GetOrderForUpdate mocks base method.
func (m *MockStore) GetOrderForUpdate(arg0 context.Context, arg1 int64) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderForUpdate indicates an expected call of GetOrderForUpdate.
func (mr *MockStoreMockRecorder) GetOrderForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderForUpdate", reflect.TypeOf((*MockStore)(nil).GetOrderForUpdate), arg0, arg1)
}

// GetOrderItem mocks base method.
func (m *MockStore) GetOrderItem(arg0 context.Context, arg1 int64) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItem", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItem indicates an expected call of GetOrderItem.
func (mr *MockStoreMockRecorder) GetOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItem", reflect.TypeOf((*MockStore)(nil).GetOrderItem), arg0, arg1)
}

// GetOrderItemForUpdate mocks base method.
func (m *MockStore) GetOrderItemForUpdate(arg0 context.Context, arg1 int64) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemForUpdate indicates an expected call of GetOrderItemForUpdate.
func (mr *MockStoreMockRecorder) GetOrderItemForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemForUpdate", reflect.TypeOf((*MockStore)(nil).GetOrderItemForUpdate), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetProductForUpdate mocks base method.
func (m *MockStore) GetProductForUpdate(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductForUpdate indicates an expected call of GetProductForUpdate.
func (mr *MockStoreMockRecorder) GetProductForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductForUpdate", reflect.TypeOf((*MockStore)(nil).GetProductForUpdate), arg0, arg1)
}

// GetRoom mocks base method.
func (m *MockStore) GetRoom(arg0 context.Context, arg1 string) (db.ChatRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", arg0, arg1)
	ret0, _ := ret[0].(db.ChatRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockStoreMockRecorder) GetRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockStore)(nil).GetRoom), arg0, arg1)
}

// GetRoomForUpdate mocks base method.
func (m *MockStore) GetRoomForUpdate(arg0 context.Context, arg1 int64) (db.ChatRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.ChatRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomForUpdate indicates an expected call of GetRoomForUpdate.
func (mr *MockStoreMockRecorder) GetRoomForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomForUpdate", reflect.TypeOf((*MockStore)(nil).GetRoomForUpdate), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetSubscription mocks base method.
func (m *MockStore) GetSubscription(arg0 context.Context, arg1 int64) (db.ChatSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscription", arg0, arg1)
	ret0, _ := ret[0].(db.ChatSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription.
func (mr *MockStoreMockRecorder) GetSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscription", reflect.TypeOf((*MockStore)(nil).GetSubscription), arg0, arg1)
}

// GetSubscriptionForUpdate mocks base method.
func (m *MockStore) GetSubscriptionForUpdate(arg0 context.Context, arg1 int64) (db.ChatSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.ChatSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionForUpdate indicates an expected call of GetSubscriptionForUpdate.
func (mr *MockStoreMockRecorder) GetSubscriptionForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionForUpdate", reflect.TypeOf((*MockStore)(nil).GetSubscriptionForUpdate), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), arg0, arg1)
}

// ListMessages mocks base method.
func (m *MockStore) ListMessages(arg0 context.Context, arg1 db.ListMessagesParams) ([]db.ChatMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessages", arg0, arg1)
	ret0, _ := ret[0].([]db.ChatMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMessages indicates an expected call of ListMessages.
func (mr *MockStoreMockRecorder) ListMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessages", reflect.TypeOf((*MockStore)(nil).ListMessages), arg0, arg1)
}

// ListOrderItems mocks base method.
func (m *MockStore) ListOrderItems(arg0 context.Context, arg1 db.ListOrderItemsParams) ([]db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrderItems", arg0, arg1)
	ret0, _ := ret[0].([]db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrderItems indicates an expected call of ListOrderItems.
func (mr *MockStoreMockRecorder) ListOrderItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrderItems", reflect.TypeOf((*MockStore)(nil).ListOrderItems), arg0, arg1)
}

// ListOrders mocks base method.
func (m *MockStore) ListOrders(arg0 context.Context, arg1 db.ListOrdersParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockStoreMockRecorder) ListOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockStore)(nil).ListOrders), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0, arg1)
}

// ListRooms mocks base method.
func (m *MockStore) ListRooms(arg0 context.Context, arg1 db.ListRoomsParams) ([]db.ChatRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRooms", arg0, arg1)
	ret0, _ := ret[0].([]db.ChatRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRooms indicates an expected call of ListRooms.
func (mr *MockStoreMockRecorder) ListRooms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRooms", reflect.TypeOf((*MockStore)(nil).ListRooms), arg0, arg1)
}

// ListSubscriptions mocks base method.
func (m *MockStore) ListSubscriptions(arg0 context.Context, arg1 db.ListSubscriptionsParams) ([]db.ChatSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubscriptions", arg0, arg1)
	ret0, _ := ret[0].([]db.ChatSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscriptions indicates an expected call of ListSubscriptions.
func (mr *MockStoreMockRecorder) ListSubscriptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscriptions", reflect.TypeOf((*MockStore)(nil).ListSubscriptions), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), arg0, arg1)
}

// SaveMessage mocks base method.
func (m *MockStore) SaveMessage(arg0 context.Context, arg1 db.SaveMessageParams) (db.ChatMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveMessage", arg0, arg1)
	ret0, _ := ret[0].(db.ChatMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveMessage indicates an expected call of SaveMessage.
func (mr *MockStoreMockRecorder) SaveMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveMessage", reflect.TypeOf((*MockStore)(nil).SaveMessage), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateMessage mocks base method.
func (m *MockStore) UpdateMessage(arg0 context.Context, arg1 db.UpdateMessageParams) (db.ChatMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessage", arg0, arg1)
	ret0, _ := ret[0].(db.ChatMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMessage indicates an expected call of UpdateMessage.
func (mr *MockStoreMockRecorder) UpdateMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockStore)(nil).UpdateMessage), arg0, arg1)
}

// UpdateOrder mocks base method.
func (m *MockStore) UpdateOrder(arg0 context.Context, arg1 db.UpdateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStoreMockRecorder) UpdateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStore)(nil).UpdateOrder), arg0, arg1)
}

// UpdateOrderItem mocks base method.
func (m *MockStore) UpdateOrderItem(arg0 context.Context, arg1 db.UpdateOrderItemParams) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderItem", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderItem indicates an expected call of UpdateOrderItem.
func (mr *MockStoreMockRecorder) UpdateOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderItem", reflect.TypeOf((*MockStore)(nil).UpdateOrderItem), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockStore) UpdateProduct(arg0 context.Context, arg1 db.UpdateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), arg0, arg1)
}

// UpdateRoom mocks base method.
func (m *MockStore) UpdateRoom(arg0 context.Context, arg1 db.UpdateRoomParams) (db.ChatRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", arg0, arg1)
	ret0, _ := ret[0].(db.ChatRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockStoreMockRecorder) UpdateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockStore)(nil).UpdateRoom), arg0, arg1)
}

// UpdateSubscription mocks base method.
func (m *MockStore) UpdateSubscription(arg0 context.Context, arg1 db.UpdateSubscriptionParams) (db.ChatSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscription", arg0, arg1)
	ret0, _ := ret[0].(db.ChatSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubscription indicates an expected call of UpdateSubscription.
func (mr *MockStoreMockRecorder) UpdateSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscription", reflect.TypeOf((*MockStore)(nil).UpdateSubscription), arg0, arg1)
}
