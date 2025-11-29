# go-emailVerify

A small Go service to send and verify OTPs (SMS) using Twilio. It provides a minimal HTTP API built with the Gin web framework and validates request payloads using `go-playground/validator`.

This repository is organized for a simple microservice that exposes two endpoints:
- `POST /otp` — request an OTP to be sent to a phone number
- `POST /verifyOTP` — verify an OTP code for a phone number

## Features
- Gin-based HTTP server
- Request validation with `validator/v10`
- Environment configuration via `github.com/joho/godotenv`
- Twilio integration for sending/verifying OTPs (you must provide Twilio credentials)

## Prerequisites
- Go 1.25+ installed
- A Twilio account with a Messaging or Verify service (and credentials)

## Quickstart

1. Clone the repository (or use your local copy):

   git clone https://github.com/imshubham024/go-emailVerify.git
   cd go-emailVerify

2. Create a `.env` file in the project root with the required environment variables (see below).

3. Download and tidy dependencies:

```bash
go mod tidy
```

4. Run the service:

```bash
go run ./cmd
```

By default the server listens on port `:8082` (as configured in `cmd/main.go`).

## Environment Variables
Create a `.env` file in the project root with the following keys:

```
TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTHTOKEN=your_auth_token
TWILIO_SERVICES_SID=your_service_sid
PORT=8082
```

Notes:
- The code uses `godotenv` to read `.env` in development. In production you can provide real environment variables instead.

## API

1) Request OTP

- Endpoint: `POST /otp`
- Payload:

```json
{ "phoneNumber": "+15551234567" }
```

- Success response: HTTP 202 Accepted with JSON payload indicating success.

Example curl:

```bash
curl -X POST http://localhost:8082/otp \
  -H "Content-Type: application/json" \
  -d '{"phoneNumber":"+15551234567"}'
```

2) Verify OTP

- Endpoint: `POST /verifyOTP`
- Payload:

```json
{
  "user": {"phoneNumber": "+15551234567"},
  "code": "123456"
}
```

- Success response: HTTP 202 Accepted with a message confirming OTP verification.

Example curl:

```bash
curl -X POST http://localhost:8082/verifyOTP \
  -H "Content-Type: application/json" \
  -d '{"user":{"phoneNumber":"+15551234567"},"code":"123456"}'
```

## Project Structure

- `cmd/` — application entrypoint (`cmd/main.go`)
- `api/` — request handlers, routes and helper functions
- `data/` — request/response data models
- `go.mod` / `go.sum` — module definitions

## Development & Testing

- Format code:

```bash
gofmt -w .
```

- Run static checks and build:

```bash
go vet ./...
go build ./...
```

There are currently no unit tests included in this repo. Consider adding tests for handlers and Twilio integration logic.

## Contributing

Contributions are welcome. Please open issues for bugs or feature requests, and submit pull requests for fixes or improvements.

Guidelines:
- Keep changes small and focused.
- Run `gofmt` and `go vet` before submitting.

## License

This project is provided without explicit license in the repository. If you want to add a license, create a `LICENSE` file (e.g., MIT) to make the terms clear.


