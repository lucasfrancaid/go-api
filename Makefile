ENV_FILE := .env.dev

run:
	@echo "--> Starting Go Api"
	@for VAR in `cat ${ENV_FILE}`; do export $$VAR; done && \
		export DB_HOST=localhost && \
		go run .

run_docker:
	@echo "--> Starting Go Api with Docker"
	@docker-compose up

run_postgres:
	@echo "--> Starting Postgres Database"
	@docker-compose up -d postgres

stop_postgres:
	@echo "--> Stopping Postgres Database"
	@docker-compose stop postgres
