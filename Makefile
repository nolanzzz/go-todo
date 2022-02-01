migrate_up:
	migrate -path database/migration -database "mysql://root:@tcp(localhost:3306)/demo" -verbose up

migrate_down:
	migrate -path database/migration -database "mysql://root:@tcp(localhost:3306)/demo" -verbose down

.PHONY: migrate_up migrate_down