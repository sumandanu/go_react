.PHONY: dev

dev:
	@echo "Starting backend..."
	@cd apps/backend && air &
	
	@echo "Starting frontend..."
	@cd apps/frontend && npm start
