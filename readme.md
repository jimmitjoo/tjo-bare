# Tjo Skeleton

![Tjo Logo](https://raw.githubusercontent.com/jimmitjoo/tjo-bare/main/public/images/tjo-logo.webp)

Minimal starting point for Tjo applications.

## Quick Start

```bash
# Create project
tjo new myapp
cd myapp

# Configure - tjo new already wrote .env with a generated KEY.
# Edit it with your database settings. Do NOT overwrite it with
# .env.example, which ships an empty KEY.
$EDITOR .env

# Run
tjo run
```

Cloning this repository directly instead of using `tjo new`? Then you do need
`cp .env.example .env`, and you must generate `KEY` yourself — see the comment
above it in `.env.example`.


Open [http://localhost:4000](http://localhost:4000)

## Adding Features

```bash
tjo make auth                    # Authentication system
tjo make model Post              # Data model
tjo make handler Blog            # HTTP handler
tjo make migration create_posts  # Database migration
tjo make middleware ratelimit    # Middleware
tjo make mail welcome            # Email template
```

## Project Structure

```
myapp/
├── main.go              # Entry point
├── bootstrap.go         # Framework init
├── routes.go            # Routes
├── handlers/            # HTTP handlers
├── data/                # Models
├── middleware/          # Middleware
├── views/               # Jet templates
├── public/              # Static assets
├── migrations/          # DB migrations (created by tjo make migration)
├── .env.example         # Config template
└── docker-compose.yml   # Dev services
```

## Configuration

Copy `.env.example` to `.env`:

| Variable | Description | Example |
|----------|-------------|---------|
| `APP_NAME` | Application name | `myapp` |
| `PORT` | Server port | `4000` |
| `DATABASE_TYPE` | Database driver | `postgres`, `mysql` |
| `DATABASE_HOST` | Database host | `localhost` |
| `DATABASE_PORT` | Database port | `5432` |
| `DATABASE_NAME` | Database name | `myapp` |
| `DATABASE_USER` | Database user | `postgres` |
| `DATABASE_PASS` | Database password | `secret` |

## Development

### Hot Reload

```bash
tjo run -w
# or
air
```

### Docker Services

```bash
docker-compose up -d    # PostgreSQL, Redis, MinIO
```

### Tailwind CSS

```bash
npm run watch
```

### Working against a local framework checkout

Use a `go.work` (gitignored) — never a `replace` in `go.mod`, that breaks the
build for everyone who does not have the framework as a sibling directory:

```bash
go work init . ../tjo ../tjo/email ../tjo/otel ../tjo/sms ../tjo/websocket
```

## Documentation

- [Tjo Framework](https://github.com/jimmitjoo/tjo)
- [Jet Templates](https://github.com/CloudyKit/jet)
- [Chi Router](https://github.com/go-chi/chi)

## License

MIT
