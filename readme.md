# Tjo Skeleton

![Tjo Logo](https://raw.githubusercontent.com/jimmitjoo/tjo-bare/main/public/images/tjo-logo.png)

Minimal starting point for Tjo applications.

## Quick Start

### 1. Create a new project

```bash
tjonew myapp
cd myapp
```

### 2. Configure environment

```bash
cp .env.example .env
# Edit .env with your settings (database, etc.)
```

### 3. Start development server

**macOS/Linux:**
```bash
make run
```

**Windows:**
```bash
make -f Makefile.windows run
```

### 4. Open your browser

Navigate to [http://localhost:4000](http://localhost:4000)

## Adding Features

Use the Tjo CLI to generate code:

```bash
# Add complete authentication (login, register, password reset)
tjomake auth

# Create a new data model
tjomake model Post

# Create a new HTTP handler
tjomake handler Blog

# Create a RESTful API controller
tjomake api-controller Products

# Create database migrations
tjomake migration create_posts_table

# Add middleware
tjomake middleware ratelimit
```

## Project Structure

```
myapp/
├── main.go              # Application entry point
├── bootstrap.go         # Framework initialization
├── routes.go            # Route definitions
├── handlers/            # HTTP request handlers
│   ├── handlers.go      # Handler struct and methods
│   └── convenience.go   # Helper methods
├── data/                # Data models and database access
│   └── models.go        # Model definitions
├── middleware/          # Request middleware
│   └── middleware.go    # Middleware definitions
├── views/               # Jet templates
│   ├── home.jet         # Home page template
│   └── layouts/         # Layout templates
│       └── base.jet     # Base HTML layout
├── public/              # Static assets (CSS, JS, images)
│   ├── css/             # Compiled Tailwind CSS
│   └── images/          # Image files
├── migrations/          # Database migration files
├── .env.example         # Environment variable template
├── docker-compose.yml   # Local development services
└── Makefile.mac         # Build commands (macOS)
```

## Configuration

Copy `.env.example` to `.env` and configure:

| Variable | Description | Example |
|----------|-------------|---------|
| `APP_NAME` | Application name | `myapp` |
| `PORT` | Server port | `4000` |
| `DATABASE_TYPE` | Database driver | `postgres`, `mysql`, `mariadb` |
| `DATABASE_HOST` | Database host | `localhost` |
| `DATABASE_PORT` | Database port | `5432` |
| `DATABASE_NAME` | Database name | `myapp` |
| `DATABASE_USER` | Database user | `postgres` |
| `DATABASE_PASS` | Database password | `secret` |

See `.env.example` for all available options.

## Development

### Hot Reload

The project includes [Air](https://github.com/air-verse/air) configuration for hot reloading:

```bash
# Install Air (one time)
go install github.com/air-verse/air@latest

# Run with hot reload
air
```

### Docker Services

Start local PostgreSQL, Redis, and MinIO:

```bash
docker-compose up -d
```

### Tailwind CSS

CSS is built with Tailwind. Watch for changes:

```bash
npm run watch
```

## Documentation

- [Tjo Framework](https://github.com/jimmitjoo/tjo) - Full documentation
- [Jet Templates](https://github.com/CloudyKit/jet) - Template engine docs
- [Chi Router](https://github.com/go-chi/chi) - HTTP router docs

## License

MIT
