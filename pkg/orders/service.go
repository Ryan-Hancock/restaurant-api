package orders

import (
	"errors"
	"fmt"
)

// ErrNotFound for when an Order can not found.
var ErrNotFound = errors.New("order not found")

// ErrUnderPaid for when an amount is lower than the bill.
var ErrUnderPaid = errors.New("amount was under paid")

// Repository provides access to the Order storage
type Repository interface {
	GetOrder(int) (Order, error)
	InsertOrder(Order) (int, error)
	UpdateOrder(Order) error
	InsertLine(Line) (int, error)
	GetLinesByOrderID(int) ([]Line, error)
	GetLinesPrice(ID int) (float32, error)
}

// Service provides Order opertions.
type Service interface {
	NewOrder() (int, error)
	AppendLine(line Line) (int, error)
	Pay(orderID int, paid float32) error
	GetOrderWithLines(orderID int) (Order, error)
}

type service struct {
	r Repository
}

// NewService creates a new Item service
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) NewOrder() (int, error) {
	return s.r.InsertOrder(Order{IsPaid: false})
}

func (s *service) AppendLine(l Line) (int, error) {
	return s.r.InsertLine(l)
}

func (s *service) Pay(orderID int, paid float32) error {
	amount, err := s.r.GetLinesPrice(orderID)
	if err != nil {
		return err
	}
	fmt.Println("here 2 ", amount)

	if paid < amount {
		return ErrUnderPaid
	}
	fmt.Println("we paid")

	order, err := s.r.GetOrder(orderID)
	fmt.Println("any order ", order)
	if err != nil {
		return err
	}
	order.IsPaid = true

	return s.r.UpdateOrder(order)
}

func (s *service) GetOrderWithLines(orderID int) (Order, error) {
	order, err := s.r.GetOrder(orderID)
	if err != nil {
		return order, err
	}

	lines, err := s.r.GetLinesByOrderID(orderID)
	if err != nil {
		return order, err
	}

	order.Lines = &lines
	return order, nil
}
