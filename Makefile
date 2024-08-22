# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=15
## Show help.
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

assets/schemas/global.schema.json: packages/configuration
	@echo "Generating global schema"
	go run ./cmd/schemas/main.go global > assets/schemas/global.schema.json

assets/schemas/hooks.schema.json: packages/configuration
	@echo "Generating hooks schema"
	go run ./cmd/schemas/main.go hooks > assets/schemas/hooks.schema.json

schemas: assets/schemas/global.schema.json assets/schemas/hooks.schema.json

.git/hooks/pre-commit: scripts/pre-commit.sh
	@echo "Installing pre-commit hook"
	cp scripts/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

.git/hooks/commit-msg: scripts/commit-msg.sh
	@echo "Installing commit-msg hook"
	cp scripts/commit-msg.sh .git/hooks/commit-msg
	chmod +x .git/hooks/commit-msg

hooks: .git/hooks/pre-commit .git/hooks/commit-msg
