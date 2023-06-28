.PHONY: migrate
migrate:
	./scripts/migrate.sh

.PHONY: run
run: migrate
	air

.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest
