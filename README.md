# Bearly Secure

Bearly Secure is an intentionally vulnerable starter web application built for the **Learn Web Security in Go** course.

It is a small plushie shop application built using **Go**, Go's standard `net/http` package, and **SQLite**. The project provides a realistic environment for learning how web applications can be attacked, secured, monitored, and improved through practical security engineering.

> [!IMPORTANT]
>
> This repository is an intentionally vulnerable course project. Course assignments may change its behavior, so this README should be treated as the current setup and structure reference rather than a general-purpose security guide.

---

## Motivation

The motivation behind Bearly Secure is to gain practical experience with **secure web application development** using Go.

Security is an important part of modern software development and should not be treated as an afterthought. A web application can contain vulnerabilities at many different layers, including authentication, authorization, session management, input validation, file handling, APIs, browser security, cryptography, infrastructure, and operational monitoring.

Bearly Secure provides a hands-on environment where these concepts can be explored in the context of a realistic application.

Throughout the project, security concepts and controls include:

- Authentication and authorization
- Session management
- Password security
- Password reset protection
- Multi-factor authentication
- Access control
- Input validation
- File upload security
- Secure file handling
- Encryption
- Key management
- Security headers
- Content Security Policy
- Cross-Origin Resource Sharing
- Rate limiting
- Load shedding
- Bot detection
- Logging and security monitoring
- Request tracing and request IDs
- Security alerts
- Incident response
- Session revocation
- DDoS mitigation concepts
- TLS and certificate trust
- Secure integrations
- OAuth concepts
- Security incident reporting
- Postmortem analysis

The project also focuses on understanding **why** security controls are necessary rather than simply implementing them.

By working through the application, the goal is to understand how multiple security layers work together to protect:

- Users
- Authentication credentials
- Sessions
- Customer information
- Application data
- APIs
- Uploaded files
- Internal services
- Infrastructure

Bearly Secure is therefore both a software development project and a practical security-learning environment.

---

## Quick Start

### Requirements

Before running Bearly Secure, make sure the following software is installed:

- Go 1.27.0 or newer
- Git
- SQLite-compatible environment
- A terminal

You can verify your Go installation with:

```bash
go version

---

## Usage

Bearly Secure is designed as a hands-on environment for learning web application security with Go.

### Start the Application

After completing the Quick Start setup, start the main application with:

```bash
go run ./cmd/server


---

## Contributing

Bearly Secure is primarily a learning project for the Learn Web Security in Go course.

When contributing changes:

1. Keep changes focused and clearly documented.
2. Follow standard Go formatting.
3. Run the test suite before committing changes.
4. Run static analysis before committing changes.
5. Do not commit `.env` files or secrets.
6. Do not commit API keys, passwords, private keys, or other sensitive information.
7. Keep changes consistent with the educational purpose of the project.

Before committing, check your changes:

```bash
git status