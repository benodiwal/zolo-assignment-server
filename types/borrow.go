package types

import "time"

type Borrow struct {
	ID              uint      `gorm:"primaryKey"`
	BookID          uint      `json:"book_id"`
	BorrowStartTime time.Time `json:"borrow_start_time"`
	BorrowEndTime   time.Time `json:"borrow_end_time"`
	Returned        bool      `json:"returned"`
}

type BorrowBookRequestBody struct {
	BorrowStartTime time.Time `json:"borrow_start_time"`
	BorrowEndTime   time.Time `json:"borrow_end_time"`
}
