KNEX=./node_modules/.bin/knex

# ge
gql-generate:
	go generate ./...

# Define targets
.PHONY: migrate-up migrate-down create-migration

migrate-up:
	$(KNEX) migrate:latest --knexfile=./knexfile.js

migrate-down:
	$(KNEX) migrate:rollback --knexfile=./knexfile.js

create-migration:
	$(KNEX) migrate:make $(name) --knexfile=./knexfile.js