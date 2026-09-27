# Bearly Secure

Bearly Secure is the intentionally vulnerable starter app for Learn Web Security in Go. It's a tiny plushie shop built with Go, `net/http`, and SQLite.

> [!IMPORTANT]
>
> This repository is an intentionally vulnerable course project. Course assignments may change its behavior, so treat this README as the current setup and structure reference rather than a security guide.

## Motivation

Bearly Secure was developed as a hands-on project to understand how security should be integrated into a real-world web application rather than treated as an afterthought.

The main motivation behind this project was to learn how to identify, prevent, and respond to common web application security risks while working with Go. Throughout the project, security controls were implemented across authentication, authorization, sessions, password management, input handling, file uploads, encryption, rate limiting, security headers, logging, monitoring, and incident response.

The project also focuses on understanding the reasoning behind security decisions. Instead of simply implementing individual protections, Bearly Secure demonstrates how multiple layers of security work together to protect users, application data, and infrastructure.

This project was built as a practical learning experience in secure backend development, with the goal of developing stronger skills in Go, web security, defensive programming, and secure system design.


## Requirements

- Go 1.27.0 or newer

## Run the Starter

Create a local environment file from the example. Keep `.env` private; it is ignored by Git.

```sh
cp .env.example .env
```

Set `PAWPAL_API_KEY`, `DOWNLOAD_SIGNING_KEY`, and `DATA_ENCRYPTION_KEY_V1` in `.env` before starting the application. The signing and encryption values must be 64 hexadecimal characters. Use disposable local values only.

Seed the local database:

```sh
go run ./cmd/seed
```

Start the app at <http://localhost:3030>:

```sh
go run ./cmd/server
```

## Attacker Lab

In another terminal, start the browser-based attacker lab at <http://localhost:4040>:

```sh
go run ./cmd/attackerlab
```

It runs as a separate process and stays on a separate origin so you can explore cross-origin browser security behavior.

## Project Checks

Run the test suite:

```sh
go test ./...
```

Check static analysis:

```sh
go vet ./...
```

You can restore the deterministic starter data at any time with `go run ./cmd/seed`.

## Baseline Features

- Public storefront with product listing, search, detail pages, and reviews
- Account creation, login, logout, password reset, and session cookies
- Account profiles, order history, review management, and tax-document uploads
- Authenticated shopping cart and checkout with simulated PawPal and Acorn integrations
- Support and admin areas for order, tax-document, and product workflows
- JSON product and order APIs
- Browser attacker lab and embedded shipping widget
- Deterministic local order-assistant simulation
- SQLite seed data, local file storage, and JSON-lines application logs
- Multi-stage container build that compiles the Bearly Secure and attacker-lab binaries

## Security Warning

Bearly Secure is deliberately unsafe. It contains exploitable authentication, authorization, injection, browser-security, data-exposure, infrastructure, and operational weaknesses for course exercises.

Do not deploy it or use its security patterns in a real application. Its credentials, integrations, payments, and third-party services are local simulations that use fake data only.

## Baseline Structure

- `cmd/server`: starts Bearly Secure
- `cmd/attackerlab`: starts the attacker lab
- `cmd/seed`: resets the deterministic SQLite data
- `internal/`: contains application behavior
- `internal/database/`: contains migrations, sqlc queries, and seed data
- `internal/httpserver/`: composes the HTTP server and middleware
- `internal/auth/`: contains authentication, session, TOTP, passkey, and access-control helpers
- `internal/integrations/`: contains simulated external-service integrations
- `internal/uploads/`: contains upload metadata, middleware, and archive extraction
- `web/`: contains server-rendered templates and static assets
- `attacker-lab/`: contains the browser attacker lab assets
- `data/fixtures/`: contains the public educational sample used by lesson checks and the container build
- `data/uploads/`: contains runtime uploads and is excluded from Git
- `data/bulk-tax-documents/`: receives documents extracted from support ZIP imports
- `Dockerfile`: defines the multi-stage container image for Bearly Secure and the attacker lab
