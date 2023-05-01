module CRUD

go 1.20

require (
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.5.1
	statemachine v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace statemachine => ../statemachine_in_go
