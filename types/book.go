package types

type Book struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `json:"name"`
}
