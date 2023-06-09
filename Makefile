setup:
	@echo "Setting up the frontend and backend folders"
	@cd frontend && yarn install
	@cd backend && go install

run-backend:
	@echo "Running the backend"
	@cd backend && go run main.go

run-frontend:
	@echo "Running the frontend"
	@cd frontend && yarn dev --host
