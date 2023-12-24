```shell
PGPASSWORD=habi123 psql -U habi -h 127.0.0.1 rest

migrate create -ext sql -dir migrations -seq create_todo_table

export POSTGRESQL_URL='postgres://habi:habi123@localhost:5432/rest?sslmode=disable'

go run cmd/app/*.go
``````