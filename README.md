# Bearly Secure

Bearly Secure is a Go-based web application built to explore practical **web application security, secure software development, and defensive engineering**.

The project is a small plushie shop application built with **Go**, Go's standard `net/http` package, and **SQLite**. It provides a realistic environment for identifying security weaknesses, implementing defensive controls, monitoring application behavior, and improving the overall security posture of a web application.

> [!IMPORTANT]
>
> This project is designed for security experimentation and defensive development. Some components may intentionally demonstrate insecure behavior before security controls are applied.

---

## Motivation

Modern web applications face security risks at many different layers. Authentication, authorization, sessions, file handling, APIs, browser security, cryptography, infrastructure, and monitoring all need to work together to protect an application.

The motivation behind Bearly Secure is to gain practical experience designing and implementing these security controls in a real Go application rather than studying them only in theory.

The project focuses on understanding **why a security control is necessary, how it should be implemented, and what can happen when it is missing or incorrectly configured**.

The project covers a wide range of security concepts, including:

* Authentication and authorization
* Session management
* Password hashing and password security
* Password reset protection
* Multi-factor authentication
* Access control
* Input validation
* File upload security
* Secure file handling
* Encryption and key management
* Security headers
* Content Security Policy (CSP)
* Cross-Origin Resource Sharing (CORS)
* Rate limiting
* Load shedding
* Bot detection
* Structured logging
* Request IDs and request tracing
* Security monitoring and alerting
* Session revocation
* DDoS mitigation concepts
* TLS and certificate trust
* Secure third-party integrations
* OAuth concepts
* Security incident handling
* Post-incident analysis

Rather than treating security as a single feature, Bearly Secure approaches security as a **defense-in-depth problem**, where multiple layers work together to reduce risk.

---

## Features

### Application

* Go HTTP server using `net/http`
* SQLite-based application storage
* Plushie product catalog
* User authentication
* User sessions
* Account management
* Password reset functionality
* File upload and download functionality
* Product APIs
* Support functionality
* Image preview functionality

### Authentication & Account Security

* Secure password hashing
* Password reset protection
* Session management
* Session revocation
* Authentication controls
* Multi-factor authentication
* TOTP-based authentication
* Login attempt monitoring
* Protection against authentication abuse
* Secure session cookies
* Authentication security alerts

### Authorization & Access Control

* User authentication checks
* Authorization checks
* Protected application routes
* Access control around sensitive operations
* Session-based authorization

### Data Protection

* AES-256-GCM encryption
* Encryption key management
* Encryption key rotation
* Versioned encryption keys
* Secure signing keys
* Sensitive-data protection
* Sensitive-field redaction in logs

### Browser & HTTP Security

* Security response headers
* Content Security Policy
* CSP nonce support
* `X-Content-Type-Options`
* `X-Frame-Options`
* `Referrer-Policy`
* HSTS support
* Secure cookie attributes
* CORS configuration
* Request ID handling

### Abuse Prevention

* Rate limiting
* Search throttling
* Global load shedding
* Bot detection
* Signup protection
* Authentication throttling
* Request-level protection

### File Security

* Secure file uploads
* File validation
* Protected file handling
* Secure download mechanisms
* Download signing
* Image preview handling

### Logging & Monitoring

* Structured logging
* Request IDs
* Authentication event logging
* Security event monitoring
* Security alerts
* Sensitive information redaction
* Operational metrics
* Error tracking

### Infrastructure

* Environment-based configuration
* `.env` configuration support
* Docker support
* Multi-stage Docker builds
* Runtime security considerations
* Health and metrics endpoints
* Trusted proxy configuration

---

## Tech Stack

| Technology   | Purpose                 |
| ------------ | ----------------------- |
| **Go**       | Backend application     |
| **net/http** | HTTP server and routing |
| **SQLite**   | Application database    |
| **Docker**   | Containerization        |
| **HTML/CSS** | Web interface           |
| **Git**      | Version control         |

---

## Project Structure

```text
bearly-secure/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   │
│   ├── auth/
│   │   ├── passwords/
│   │   └── ...
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── httpserver/
│   │   ├── app.go
│   │   ├── auth.go
│   │   ├── auth_advanced.go
│   │   ├── middleware.go
│   │   └── ...
│   │
│   ├── imagepreview/
│   │   └── service.go
│   │
│   ├── logging/
│   │   └── logger.go
│   │
│   ├── storage/
│   │   ├── encryption.go
│   │   ├── keyring.go
│   │   └── ...
│   │
│   ├── support/
│   │   └── handler.go
│   │
│   └── uploads/
│       └── handler.go
│
├── templates/
│
├── Dockerfile
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

The project uses a modular internal structure so authentication, HTTP handling, configuration, logging, storage, encryption, uploads, and other security-sensitive components remain separated.

---

# Quick Start

## Requirements

Before running Bearly Secure, make sure the following software is installed:

* Go 1.27.0 or newer
* Git
* SQLite-compatible environment
* Docker (optional)
* A terminal

Verify your Go installation:

```bash
go version
```

Example:

```text
go version go1.27.1 linux/amd64
```

---

## Clone the Repository

Clone the project:

```bash
git clone <your-repository-url>
```

Move into the project directory:

```bash
cd bearly-secure
```

---

## Install Dependencies

Download the Go dependencies:

```bash
go mod download
```

You can also verify the project dependencies with:

```bash
go mod tidy
```

> Run `go mod tidy` only when you intend to update the module dependency files.

---

# Configuration

Bearly Secure uses environment variables for configuration.

A sample configuration is provided in:

```text
.env.example
```

Create your local environment file:

```bash
cp .env.example .env
```

Then update the values according to your local environment.

Configuration may include:

* Application configuration
* Database settings
* Authentication settings
* Encryption configuration
* Encryption key versions
* Download signing keys
* Trusted proxy configuration
* Third-party service credentials
* Security-related configuration

### Important

Never commit `.env` to Git.

Sensitive values should remain outside version control, including:

```text
.env
API keys
Passwords
Private keys
Encryption keys
Signing keys
OAuth credentials
Service credentials
Authentication tokens
```

---

# Running the Application

## Start the Server

Run the application with:

```bash
go run ./cmd/server
```

The server will start using the configured environment variables.

Open the application using the address configured by the server.

---

# Development

## Format the Code

Format Go source files using:

```bash
gofmt -w .
```

To check which files are not formatted:

```bash
gofmt -l .
```

---

## Run Tests

Run the complete test suite:

```bash
go test ./...
```

For verbose output:

```bash
go test -v ./...
```

Run tests with the race detector:

```bash
go test -race ./...
```

---

## Static Analysis

Run Go's built-in static analyzer:

```bash
go vet ./...
```

---

# Docker

Bearly Secure can also be built and run as a Docker container.

## Build the Image

```bash
docker build -t bearly-secure .
```

## Run the Container

```bash
docker run --env-file .env -p 8080:8080 bearly-secure
```

The exact port depends on the application configuration.

You can check running containers with:

```bash
docker ps
```

Stop a running container with:

```bash
docker stop <container-id>
```

---

# Security Architecture

Bearly Secure follows a **defense-in-depth** approach.

Instead of depending on a single security mechanism, the application uses multiple layers of protection.

```text
                         ┌──────────────────────┐
                         │       Client         │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │  HTTP Security       │
                         │  Headers / CSP /     │
                         │  CORS / Cookies      │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │ Abuse Protection     │
                         │ Rate Limiting        │
                         │ Load Shedding        │
                         │ Bot Detection        │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │ Authentication       │
                         │ Authorization        │
                         │ Session Security     │
                         │ MFA / TOTP           │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │ Application Logic    │
                         │ Input Validation     │
                         │ Access Control       │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │ Data Protection      │
                         │ Encryption            │
                         │ Key Management        │
                         │ Secure Storage        │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │ Monitoring & Logging │
                         │ Alerts / Metrics     │
                         └──────────────────────┘
```

This approach helps reduce the impact of individual security failures by placing multiple controls between an attacker and sensitive application resources.

---

# Authentication & Session Security

Authentication is one of the core security areas of the application.

The project includes mechanisms for:

* Password hashing
* Secure authentication
* Session management
* Session cookie protection
* Session revocation
* Password reset protection
* Multi-factor authentication
* TOTP authentication
* Login attempt tracking
* Authentication alerts
* Protection against repeated authentication attempts

Security-sensitive authentication events are also integrated with the application's logging and monitoring systems.

---

# Password Security

Passwords should never be stored as plaintext.

Bearly Secure uses modern password hashing techniques and security-focused password handling.

The authentication layer is designed around:

* Strong password hashing
* Salted password storage
* Secure password verification
* Protection against repeated login attempts
* Secure password reset flows
* Protection of authentication-related information

---

# Encryption & Key Management

Sensitive application data requires protection both in storage and during application processing.

The project includes:

* AES-256-GCM encryption
* Versioned encryption keys
* Active encryption key configuration
* Encryption key rotation
* Key management
* Signing key support

Encryption configuration is kept outside source control through environment variables.

Example configuration variables may include:

```text
DATA_ENCRYPTION_ACTIVE_VERSION
DATA_ENCRYPTION_KEY_V1
DOWNLOAD_SIGNING_KEY
```

Actual secret values should never be committed to the repository.

---

# HTTP Security

The application uses multiple HTTP security controls.

These include:

* Content Security Policy
* CSP nonces
* `X-Content-Type-Options`
* `X-Frame-Options`
* `Referrer-Policy`
* HSTS support
* Secure cookies
* HTTP-only cookies
* SameSite cookie protection
* CORS restrictions
* Request IDs

These controls help reduce common browser-based attack risks such as:

* Cross-site scripting
* Clickjacking
* MIME confusion
* Cross-origin abuse
* Session theft

---

# Content Security Policy

The application uses Content Security Policy to restrict which resources browsers are allowed to load and execute.

Where required, CSP nonces can be used to allow specific trusted inline resources without broadly allowing arbitrary inline scripts.

This provides an additional layer of protection against script injection vulnerabilities.

---

# CORS

Cross-Origin Resource Sharing is configured according to the intended API behavior.

The goal is to ensure that browser-based requests from unauthorized origins cannot freely interact with protected resources.

CORS configuration should always be kept as restrictive as the application's legitimate requirements allow.

---

# Rate Limiting

Rate limiting is used to reduce abuse of application endpoints.

It can help protect against:

* Brute-force authentication attempts
* Excessive API requests
* Automated abuse
* Resource exhaustion
* Search endpoint abuse

The project also includes specialized throttling for sensitive operations.

---

# Load Shedding

Rate limiting alone is not always sufficient during periods of high traffic.

Bearly Secure includes load-shedding mechanisms designed to prevent the application from becoming overwhelmed when incoming traffic exceeds available resources.

The goal is to preserve availability for legitimate users while reducing unnecessary system load.

---

# Bot Detection

Automated requests can be particularly harmful to authentication and signup endpoints.

The application includes bot-detection mechanisms that can be used to identify and restrict suspicious automated activity.

This is especially useful for:

* Signup protection
* Authentication endpoints
* Repeated automated requests
* Abuse prevention

---

# File Upload Security

File uploads are treated as untrusted input.

Security considerations include:

* Input validation
* File validation
* Controlled file handling
* Safe storage
* Secure file access
* Protection against malicious uploads

Uploaded files should never automatically be considered trustworthy simply because they originate from an authenticated user.

---

# Secure Downloads

Sensitive or protected files should not be exposed through predictable public URLs.

The application includes download-signing mechanisms that can be used to ensure that download requests are properly authorized.

Signing keys are stored through environment-based configuration rather than being hard-coded into the application.

---

# Logging & Observability

Security requires visibility into what the application is doing.

Bearly Secure includes structured logging and observability features for tracking application and security events.

These include:

* Structured logs
* Request IDs
* Authentication events
* Security events
* Error information
* Metrics
* Security alerts
* Sensitive-field redaction

Sensitive information should never be unnecessarily written to logs.

---

# Security Alerts

Authentication and security events can be monitored for suspicious activity.

Examples include:

* Repeated failed login attempts
* Password reset abuse
* Authentication threshold violations
* Suspicious authentication patterns
* Other security-sensitive events

Security alerts are intended to make important events visible to operators so they can be investigated.

---

# Request IDs

Each request can be associated with a unique request identifier.

Request IDs make it easier to correlate:

```text
Client Request
      │
      ▼
HTTP Server
      │
      ├── Authentication
      │
      ├── Application Logic
      │
      ├── Database
      │
      └── Logging
```

This makes debugging and security investigation easier because events belonging to the same request can be correlated.

---

# Security Principles

The project follows several important security principles.

## Defense in Depth

Security should not depend on a single control.

Multiple independent protections help reduce the impact of individual failures.

## Least Privilege

Users and components should receive only the permissions necessary for their intended operations.

## Secure by Default

Security-sensitive functionality should use safer defaults whenever possible.

## Fail Securely

When an operation fails, the application should avoid accidentally granting access or exposing sensitive information.

## Input Validation

All external input should be treated as untrusted until it has been appropriately validated.

## Protect Sensitive Data

Credentials, tokens, encryption keys, and other sensitive information should not be unnecessarily exposed through:

* Logs
* HTTP responses
* Error messages
* Source code
* Configuration files

---

# Monitoring & Incident Response

Security does not end with prevention.

A secure application should also be able to detect, investigate, and respond to suspicious activity.

The project therefore includes concepts around:

* Security event monitoring
* Authentication alerts
* Structured logging
* Request tracing
* Incident investigation
* Security reporting
* Post-incident analysis

A typical investigation can follow a flow such as:

```text
Security Event
      │
      ▼
Log / Alert
      │
      ▼
Request ID
      │
      ▼
Application Logs
      │
      ▼
Identify Cause
      │
      ▼
Apply Mitigation
      │
      ▼
Verify Fix
      │
      ▼
Document Findings
```

---

# Development Workflow

A typical development workflow looks like:

```bash
# Check repository state
git status

# Format code
gofmt -w .

# Run tests
go test ./...

# Run static analysis
go vet ./...

# Review modifications
git diff

# Stage changes
git add .

# Commit changes
git commit -m "Describe the change"
```

Keeping changes focused makes security-sensitive modifications easier to review and test.

---

# Contributing

Contributions and improvements are welcome.

When contributing to Bearly Secure:

1. Keep changes focused and clearly documented.
2. Follow standard Go formatting conventions.
3. Add or update tests when changing application behavior.
4. Run the complete test suite before committing.
5. Run static analysis before committing.
6. Do not commit `.env` files or secrets.
7. Never commit API keys, passwords, private keys, encryption keys, signing keys, or authentication tokens.
8. Keep security-sensitive changes easy to review.
9. Document significant security or architectural decisions.
10. Avoid introducing unnecessary dependencies.
11. Review the security implications of changes that affect authentication, authorization, storage, HTTP handling, or user input.

Before submitting changes, run:

```bash
git status
gofmt -w .
go test ./...
go vet ./...
git diff
```

---

# Security Reporting

If you discover a security issue, avoid publicly disclosing sensitive details before the issue has been investigated.

A useful security report should include:

* Clear description of the issue
* Steps to reproduce
* Expected behavior
* Actual behavior
* Potential impact
* Relevant logs or screenshots where appropriate
* Suggested mitigation, if known

Do not include sensitive information such as:

```text
Passwords
API keys
Private keys
Encryption keys
Authentication tokens
Session cookies
```

in security reports.

---

# Learning Goals

Bearly Secure provides practical experience with:

* Go web development
* Secure application architecture
* Authentication
* Authorization
* Session security
* Password security
* Multi-factor authentication
* Cryptography
* Encryption
* Key management
* API security
* Browser security
* File security
* Rate limiting
* Load shedding
* Bot detection
* Security monitoring
* Logging
* Observability
* Docker
* Infrastructure security
* Incident response

The overall goal is to develop the ability to **identify security risks, understand their impact, implement appropriate defensive controls, and verify that those controls work as intended**.

---

# Project Status

Bearly Secure is an actively developed security-focused Go application.

Security controls and application behavior may evolve as the project is improved and additional defensive mechanisms are introduced.

---

## Useful Commands

### Run the application

```bash
go run ./cmd/server
```

### Run tests

```bash
go test ./...
```

### Run tests with race detection

```bash
go test -race ./...
```

### Run static analysis

```bash
go vet ./...
```

### Format code

```bash
gofmt -w .
```

### Check formatting

```bash
gofmt -l .
```

### Build the application

```bash
go build ./...
```

### Build Docker image

```bash
docker build -t bearly-secure .
```

### Check Git status

```bash
git status
```

### Review changes

```bash
git diff
```

---

# Final Notes

Bearly Secure is built around the idea that application security is not a single feature that can simply be added at the end of development.

A secure application requires multiple layers working together:

```text
Secure Code
     +
Authentication
     +
Authorization
     +
Input Validation
     +
Encryption
     +
Secure Sessions
     +
HTTP Security
     +
Abuse Prevention
     +
Logging
     +
Monitoring
     +
Incident Response
     =
Defense in Depth
```

The project brings these concepts together in a practical Go application so that security decisions can be implemented, tested, observed, and improved as part of the development process.
