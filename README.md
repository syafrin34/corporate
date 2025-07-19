// cara buat migrate di go
migrate create -ext sql  -dir database/migrations -seq create_portofolio_testimonials_table

// migrate database
migrate -database "postgres://postgres:password@localhost:5432/corporate?sslmode=disable" -path database/migrations up
// migrate 
//migrate -database "postgres://postgres:password@172.17.0.2:5432/corporate?sslmode=disable" -path database/migrations up

// go get github.com/golang-jwt/jwt/v5
// go get github.com/go-playground/validator/v10
// go get github.com/go-playground/universal-translator
// go get github.com/supabase-community/storage-go