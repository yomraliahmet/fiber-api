package resources

import "github.com/yomraliahmet/fiber-api/models"

type User struct {
	ID        uint   `json:"id" extensions:"x-order=1" example:"1" gorm:"primaryKey"`
	FirstName string `json:"first_name" extensions:"x-order=2" example:"Musa"`
	LastName  string `json:"last_name" extensions:"x-order=3" example:"UZUN"`
	City      string `json:"city" extensions:"x-order=4" example:"Kayseri"`
}

func UserResource(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		City:      userModel.City,
	}
}
