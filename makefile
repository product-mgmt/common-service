db-migration-up:
	@go run migration/databases/main.go up

db-migration-down:
	@go run migration/databases/main.go down

db-clean: db-migration-down

db-setup: db-migration-up
	@go run migration/procedures/main.go