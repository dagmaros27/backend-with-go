package models

type User struct {
	ID       string `json:"_id " bson:"_id,omitempty"`
	Username string `json:"username" binding:"required" bson:"username"`
	Password string `json:"password" binding:"required" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

