# GoCoreFoundation

A production-ready Go web application foundation providing secure authentication, user management, and multi-channel communication capabilities. Built with clean architecture principles and designed for scalability.

## Overview

**GoCoreFoundation** is a comprehensive boilerplate/starter template for building enterprise-grade RESTful APIs in Go. It provides complete infrastructure for authentication (JWT, session management), user CRUD operations, device tracking, OTP verification, and integrated communication channels (SMS via Twilio, Email).

This foundation is designed to accelerate development of secure, scalable web applications by providing battle-tested patterns and implementations of common requirements.

## Features

- **🔐 Authentication & Authorization**
  - JWT-based authentication (HMAC and ECDSA support)
  - Session management with multiple providers (JWT, XWT)
  - Device tracking and blocking
  - Login audit trail

- **👥 User Management**
  - Complete CRUD operations
  - Role-based access control (Admin, User, Guest)
  - User aliases support
  - Soft and hard delete capabilities

- **🔒 Security**
  - OTP (One-Time Password) verification system
  - Device fingerprinting and blocking
  - Secure password hashing with bcrypt
  - CORS middleware
  - Request validation and sanitization

- **📱 Multi-Channel Communication**
  - SMS notifications via Twilio
  - Email sending capabilities
  - Template-based messaging

- **🗄️ Database**
  - MySQL 8.0 with connection pooling
  - Migration management (up/down)
  - Comprehensive schema for users, devices, sessions, OTPs

- **🚀 Deployment Ready**
  - Docker and Docker Compose support
  - AWS EC2 deployment scripts
  - Nginx load balancing configuration
  - Health check endpoints

- **🛠️ Developer Experience**
  - Graceful shutdown handling
  - Background job management
  - Structured logging
  - Internationalization support (i18n)
  - Comprehensive utility libraries

## Technology Stack

### Core
- **Go 1.24.4** - Primary language
- **MySQL 8.0** - Database
- **Docker & Docker Compose** - Containerization
- **Nginx** - Reverse proxy and load balancing

### Key Dependencies
- `golang-jwt/jwt` - JWT authentication
- `go-sql-driver/mysql` - MySQL driver
- `rs/cors` - CORS middleware
- `twilio/twilio-go` - SMS integration
- `gopkg.in/gomail.v2` - Email service
- `golang.org/x/crypto` - Cryptography
- `google/uuid` - UUID generation
- `joho/godotenv` - Environment configuration

## Project Structure

```
GoCoreFoundation/
├── cmd/server/              # Application entry point
│   └── main.go             # Main executable
│
├── internal/               # Private application code
│   ├── app/               # Application initialization
│   │   ├── routes/        # HTTP route definitions
│   │   ├── services/      # Service container
│   │   └── resource/      # Application resources
│   ├── configs/           # Configuration management
│   ├── constants/         # Application constants
│   ├── db/               # Database connections
│   ├── handlers/         # HTTP request handlers
│   │   ├── login/        # Authentication
│   │   ├── users/        # User management
│   │   ├── otp/          # OTP verification
│   │   ├── device/       # Device tracking
│   │   ├── block/        # Blocking system
│   │   └── health/       # Health checks
│   ├── middlewares/      # HTTP middlewares
│   ├── services/         # Business logic
│   │   ├── sms/          # SMS service
│   │   └── mail/         # Email service
│   ├── sessions/         # Session management
│   ├── jobs/            # Background jobs
│   └── utils/           # Utility functions
│
├── pkg/                  # Public reusable packages
│   ├── twilio/          # Twilio client wrapper
│   └── mailer/          # Email client wrapper
│
├── root/                # Core foundation packages
│   ├── jwt/            # JWT utilities
│   ├── session/        # Session storage
│   └── sessionprovider/ # Session providers
│
├── migrations/          # Database migrations
│   ├── up/             # Apply migrations
│   └── down/           # Rollback migrations
│
├── bin/                # Build and deployment scripts
├── wrk/nginx/          # Nginx configurations
├── Makefile            # Build automation
├── Dockerfile          # Container definition
└── docker-compose.yml  # Multi-container setup
```

## Database Schema

The application manages these core entities:

- **users** - User accounts with role management
- **logins** - Authentication credentials
- **aliases** - Alternative user identifiers
- **devices** - Device tracking for security
- **login_logs** - Audit trail for authentication
- **otp_codes** - One-time password storage
- **blocks** - User/device blocking system

## API Endpoints

### Authentication
- `POST /login` - User authentication

### Users
- `GET /users/list` - List all users
- `GET /users/{id}` - Get specific user
- `GET /users/profile` - Get current user profile
- `POST /users/create` - Create new user
- `POST /users/update` - Update user information
- `POST /users/delete` - Soft delete user
- `POST /users/force-delete` - Permanently delete user

### OTP Verification
- `POST /otp/send` - Send OTP code
- `POST /otp/verify` - Verify OTP code

### Blocks
- `GET /blocks/list` - List blocked entities

### Health
- `GET /healths/ping` - Health check endpoint

### Miscellaneous
- `GET /misc/sessions/dump` - Debug session information

## Getting Started

### Prerequisites

- Go 1.24.4 or higher
- MySQL 8.0 or higher
- Docker & Docker Compose (for containerized deployment)
- Make utility

### Local Development Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd GoCoreFoundation
   ```

2. **Configure environment**
   ```bash
   cp .example.env .env
   ```
   Edit `.env` and fill in your configuration:
   - Database credentials
   - JWT secrets
   - Twilio credentials (for SMS)
   - Email service credentials
   - Application port and settings

3. **Create MySQL database**
   ```bash
   mysql -u root -p
   CREATE DATABASE your_database_name;
   ```

4. **Run database migrations**
   ```bash
   make migrate-up
   ```
   To rollback migrations:
   ```bash
   make migrate-down
   ```

5. **Run the application**
   ```bash
   make run
   ```

6. **Test the API**
   ```
   Base URL: http://localhost:8080
   ```

### Docker Deployment

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd GoCoreFoundation
   ```

2. **Configure Docker environment**
   ```bash
   cp .example.docker.env .env.docker
   ```
   Edit `.env.docker` with your Docker-specific configuration

3. **Start containers**
   ```bash
   make docker-compose-up
   ```
   This will start:
   - 3 server replicas (load balanced)
   - MySQL database
   - Nginx reverse proxy

4. **Run migrations**
   ```bash
   make migrate-up-docker
   ```
   To rollback:
   ```bash
   make migrate-down-docker
   ```

5. **Test the API**
   ```
   Base URL: http://localhost:70 (Nginx port)
   Database: localhost:33066
   ```

### AWS EC2 Deployment

1. **Build for AWS ARM64**
   ```bash
   make build-ec2
   ```

2. **Initialize deployment**
   ```bash
   make init-deploy
   ```

3. **Deploy to EC2**
   ```bash
   make deploy-ec2-remote
   ```

4. **Run migrations on AWS**
   ```bash
   make migrate-up-aws
   ```

## Development

### Creating Database Migrations

```bash
./bin/create_migration.sh migration_name
```

This creates two files:
- `migrations/up/YYYYMMDDHHMMSS_migration_name.sql`
- `migrations/down/YYYYMMDDHHMMSS_migration_name.sql`

### Running Specific Migrations

**Local:**
```bash
./bin/run_migration_file.sh path/to/migration.sql
```

**Docker:**
```bash
./bin/run_migration_file_docker.sh path/to/migration.sql
```

**AWS:**
```bash
./bin/run_migration_file_aws.sh path/to/migration.sql
```

### Available Make Commands

```bash
make run                  # Run application locally
make build                # Build binary
make build-ec2           # Build for AWS ARM64
make test                # Run tests
make docker-compose-up   # Start Docker containers
make docker-compose-down # Stop Docker containers
make migrate-up          # Apply all migrations (local)
make migrate-down        # Rollback all migrations (local)
make migrate-up-docker   # Apply migrations (Docker)
make migrate-down-docker # Rollback migrations (Docker)
make migrate-up-aws      # Apply migrations (AWS)
make deploy-ec2-remote   # Deploy to EC2
make init-deploy         # Initialize deployment
```

## Middleware Pipeline

The application processes requests through the following middleware (in order):

1. **RootSessionMiddleware** - Session initialization and management
2. **LocaleMiddleware** - Internationalization (default: "en")
3. **DeviceBlockMiddleware** - Block requests from banned devices
4. **DeviceMiddleware** - Device tracking and fingerprinting
5. **LogRequestMiddleware** - Request logging and monitoring

## Background Jobs

The application includes a job manager for scheduled tasks:
- OTP cleanup (expire old codes)
- User maintenance tasks
- Session cleanup

## Configuration

Key environment variables:

```env
# Server
APP_PORT=8080
APP_ENV=development

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database

# JWT
JWT_SECRET=your_jwt_secret
JWT_EXPIRATION=24h

# Twilio (SMS)
TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_PHONE_NUMBER=your_phone_number

# Email
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASSWORD=your_app_password
```

## Architecture Highlights

### Clean Architecture
- Clear separation between handlers, services, and data layers
- Dependency injection through service containers
- Interface-based design for testability

### Session Management
- Multiple session provider support (JWT, XWT)
- Session serialization for persistence
- Graceful session cleanup on shutdown

### Security Features
- JWT with HMAC and ECDSA signing
- Device fingerprinting and blocking
- OTP verification for sensitive operations
- Secure password hashing with bcrypt
- CORS configuration
- Request validation and sanitization

## Contributing

Contributions are welcome! Please ensure:
- Code follows Go best practices and conventions
- Tests are included for new features
- Documentation is updated accordingly
- Commits are clear and descriptive

## License

[Add your license information here]

## Author Contact

**Tran Phuoc Anh Quoc**

- Email: anquoctpdev@gmail.com
- Facebook: https://www.facebook.com/tranphuocanhquoc2003

---

<p align="center">Thank you for using GoCoreFoundation!</p>
