package entity

import "fmt"

const (
	ReservationStatusWaiting   = "waiting"
	ReservationStatusAvailable = "available"
	ReservationStatusCancelled = "cancelled"
	ReservationStatusCompleted = "completed"
)

type Reservation struct {
	Id        int     `json:"id" db:"id"`
	UserId    int     `json:"user_id" db:"user_id"`
	BookId    int     `json:"book_id" db:"book_id"`
	Position  *int    `json:"position,omitempty" db:"position"`
	Status    string  `json:"status" db:"status"`
	ExpiresAt *string `json:"expires_at,omitempty" db:"expires_at"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
}

func (r *Reservation) SetDefault() {
	if r.Status == "" {
		r.Status = ReservationStatusWaiting
	}
}

func (r *Reservation) Validate() error {
	if r.UserId <= 0 {
		return fmt.Errorf("Invalid field: UserId is required")
	}
	if r.BookId <= 0 {
		return fmt.Errorf("Invalid field: BookId is required")
	}
	switch r.Status {
	case ReservationStatusWaiting, ReservationStatusAvailable, ReservationStatusCancelled, ReservationStatusCompleted:
	default:
		return fmt.Errorf("Invalid field: Status must be waiting, available, cancelled or completed")
	}
	if r.Position != nil && *r.Position < 0 {
		return fmt.Errorf("Invalid field: Position cannot be negative")
	}
	return nil
}
