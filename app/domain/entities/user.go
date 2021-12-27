package entities

type User struct {
	UserId 		int		`json:"id_user"`
	Name 		string 	`json:"name"`
	LastName 	string 	`json:"last_name"`
	Nickname 	string	`json:"nickname"`
	Mail 		string 	`json:"mail"`
	Password 	string 	`json:"password"`
	Active 		bool 	`json:"active"`
}
