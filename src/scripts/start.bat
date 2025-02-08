@echo off
REM Load environment variables
if exist .env (
  for /F "tokens=* usebackq" %%F in (`type .env`) do set %%F
)

REM Run the application
echo Starting the Go application...
go run cmd/app/main.go
