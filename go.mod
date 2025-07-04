// File: go.mod
// Purpose: Go module definition and dependency tracking for the shorty project

module github.com/casjaysdev/shorty

go 1.21

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/redis/go-redis/v9 v9.2.1
	github.com/golang-jwt/jwt/v5 v5.1.0
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v5 v5.5.4
	github.com/jinzhu/inflection v1.0.0
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.19
	github.com/rs/cors v1.9.0
	github.com/swaggo/files v1.0.1
	github.com/swaggo/http-swagger v1.3.4
	github.com/swaggo/swag v1.16.3
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.22.0
	golang.org/x/net v0.22.0
	golang.org/x/sys v0.19.0
	gorm.io/driver/postgres v1.3.3
	gorm.io/driver/sqlite v1.3.3
	gorm.io/gorm v1.30.0
)

require (
	github.com/stretchr/testify v1.8.4 // indirect
)
