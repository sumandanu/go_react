.PHONY: build start

build:
	@echo "Building backend..."
	@cd apps/backend && go build -o ../bin/backend ./cmd/server
	
	@echo "Building frontend..."
	@cd apps/frontend && npm run build

test:
	@echo "Testing backend..."
	@cd apps/backend && go test ./...
	
	@echo "Testing frontend..."
	@cd apps/frontend && npm run test

start:
	@echo "Starting backend..."
	@cd apps/backend && air &
	
	@echo "Starting frontend..."
	@cd apps/frontend && npm run start