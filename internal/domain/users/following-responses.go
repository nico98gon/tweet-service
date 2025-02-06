package users

type FollowingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"` 
}

type User struct {
	ID				string	`bson:"_id,omitempty" json:"id"`
	Name     	string 	`bson:"name" json:"name,omitempty"`
	LastName 	string 	`bson:"last_name" json:"last_name,omitempty"`
	Email    	string 	`bson:"email" json:"email"`
	Avatar 		string 	`bson:"avatar" json:"avatar,omitempty"`
}