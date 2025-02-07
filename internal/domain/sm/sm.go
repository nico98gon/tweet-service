package domain

type Secret struct {
	Host 			string `json:"host"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	JWTSign 	string `json:"jwtSign"`
	IsSrv			bool	 `json:"isSrv"`
	Database 	string `json:"database"`
}