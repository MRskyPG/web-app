package web

type Client struct {
	ID      int    `json:"id"`
	FIO     string `json:"fio"`
	Phone   string `json:"phone"`
	Adress  string `json:"adress"`
	OrderID int    `json:"order_id"`
}

type WorkingPosition struct {
	PositionID int    `json:"position_id"`
	Name       string `json:"name"`
}

type Staff struct {
	StaffID     int    `json:"staff_id"`
	FIO         string `json:"fio"`
	PositionID  int    `json:"position_id"`
	DateOfBirth string `json:"date_of_birth"`
	Salary      int    `json:"salary"`
	Phone       string `json:"phone"`
	Adress      string `json:"adress"`
}

type Order struct {
	OrderID       int     `json:"order_id"`
	ClientID      int     `json:"client_id"`
	Name          string  `json:"name"`
	Cost          float64 `json:"cost"`
	PaymentMethod string  `json:"payment_method"`
	Date          string  `json:"date"`
	FinishDate    string  `json:"finish_date"`
	Description   string  `json:"description"`
}
