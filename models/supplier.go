package models

type Supplier struct {
	ID       string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	BranchId string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
	Address  string `gorm:"type:text;" json:"address"`
	Phone    string `gorm:"type:varchar(100);" json:"phone"`
	Email    string `gorm:"type:varchar(100);" json:"email"`
}
