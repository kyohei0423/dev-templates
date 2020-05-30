gmig:
	migrate create -dir db/migrations -ext sql -seq $(subject)
migrate:
	migrate -database $(POSTGRESQL_URL) -path db/migrations up $(version)
force:
	migrate -database $(POSTGRESQL_URL) -path db/migrations force $(version)
down:
	migrate -database $(POSTGRESQL_URL) -path db/migrations down $(version)
