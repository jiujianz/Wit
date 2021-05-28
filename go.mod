module github.com/jiujianz/Wit

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jiujianz/Wit/Config v1.0.0
	github.com/jiujianz/Wit/Controllers v1.0.0 // indirect
	github.com/jiujianz/Wit/Models v1.0.0
	github.com/jiujianz/Wit/Routes v1.0.0
	github.com/jiujianz/Wit/Services v1.0.0 // indirect
	github.com/jiujianz/Wit/Utility v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)

replace (
	github.com/jiujianz/Wit/Config v1.0.0 => ./api/Config
	github.com/jiujianz/Wit/Controllers v1.0.0 => ./api/Controllers
	github.com/jiujianz/Wit/Models v1.0.0 => ./api/Models
	github.com/jiujianz/Wit/Routes v1.0.0 => ./api/Routes
	github.com/jiujianz/Wit/Services v1.0.0 => ./api/Services
	github.com/jiujianz/Wit/Utility v1.0.0 => ./api/Utility
)
