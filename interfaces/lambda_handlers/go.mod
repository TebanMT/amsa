module github.com/TebanMT/amsa/interfaces/lambda_handlers

go 1.21.5

require (
	github.com/TebanMT/amsa v0.0.0
	github.com/TebanMT/amsa/infrastructure/authentication v0.0.0
	github.com/TebanMT/amsa/infrastructure/db v0.0.0
	github.com/aws/aws-lambda-go v1.43.0
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/gorm v1.25.5 // indirect
)

replace github.com/TebanMT/amsa => ../../

replace github.com/TebanMT/amsa/infrastructure/db => ../../infrastructure/db

replace github.com/TebanMT/amsa/infrastructure/authentication => ../../infrastructure/authentication
