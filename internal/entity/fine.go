package entity

type Fine struct {
	Id        int     `json:"id" db:"id"`
	LoanId    int     `json:"loan_id" db:"loan_id"`
	UserId    int     `json:"user_id" db:"user_id"`
	Amount    float64 `json:"amount" db:"amount"`
	Reason    *string `json:"reason,omitempty" db:"reason"`
	Paid      bool    `json:"paid" db:"paid"`
	PaidAt    *string `json:"paid_at,omitempty" db:"paid_at"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
}
