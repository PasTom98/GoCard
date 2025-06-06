# GoCard

A full-stack flashcard application designed for Magic: The Gathering cards, built with Go backend API and Electron/React frontend.

## Overview

GoCard is a comprehensive flashcard application that allows users to manage Magic: The Gathering card collections and create study decks. The application features a RESTful API backend built with Go and a cross-platform desktop frontend built with Electron and React.

## Architecture

### Backend (Go)
- **API Framework**: Chi router
- **Database**: PostgreSQL 16.3
- **Caching**: Redis 6.2
- **Migration**: golang-migrate
- **Documentation**: Swagger/OpenAPI

### Frontend (Electron/React)
- **Framework**: React 19.1.0
- **Desktop**: Electron 36.4.0
- **Build Tool**: Vite 6.3.5
- **Language**: TypeScript 5.8.3

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Database Management**: PostgreSQL with migrations
- **Cache Management**: Redis with Redis Commander UI

## Project Structure

```
GoCard/
├── app/                    # Go backend application
│   ├── api/               # API handlers and models
│   ├── config/            # Configuration management
│   ├── main.go            # Application entry point
│   ├── go.mod             # Go modules
│   └── go.sum
├── frontend/              # Electron/React frontend
│   ├── src/
│   │   ├── electron/      # Electron main process
│   │   └── ui/            # React UI components
│   ├── package.json
│   └── vite.config.ts
├── cmd/                   # Command line tools
│   ├── api/               # API server
│   └── migrate/           # Database migrations
├── data/                  # Database files
│   └── postgres/
│       ├── migrations/    # SQL migration files
│       └── db_init.sql    # Database initialization
├── config/                # Configuration files
├── docker/                # Docker configuration
├── scripts/               # Utility scripts
├── docs/                  # Documentation
├── public/                # Static assets
├── docker-compose.yml     # Multi-container setup
├── Dockerfile             # Container image definition
├── Makefile              # Build automation
└── env.dist              # Environment variables template
```

## Quick Start

### Prerequisites

- Go 1.24+
- Node.js 18+
- Docker & Docker Compose
- PostgreSQL 16.3 (if running locally)
- Redis 6.2+ (if running locally)

### 1. Clone the Repository

```bash
git clone <repository-url>
cd GoCard
```

### 2. Environment Setup

```bash
# Copy environment template
cp env.dist .env
cp env.dist config.env

# Edit .env and config.env with your database credentials
```

### 3. Start with Docker Compose (Recommended)

```bash
# Start all services (database, redis, migrations)
docker-compose up -d

# View logs
docker-compose logs -f
```

### 4. Manual Setup (Alternative)

#### Backend Setup
```bash
cd app

# Install Go dependencies
go mod download

# Run database migrations
make migrate-up

# Start the API server
go run main.go
```

#### Frontend Setup
```bash
cd frontend

# Install Node.js dependencies
npm install

# Start React development server
npm run dev:react

# Or start Electron app
npm run dev:electron
```

## Available Commands

### Make Commands
```bash
make test           # Run Go tests
make migration      # Create new database migration
make migrate-up     # Apply database migrations
make migrate-down   # Rollback database migrations
make seed          # Seed database with initial data
make gen-docs      # Generate API documentation
```

### Frontend Commands
```bash
cd frontend
npm run dev:react     # Start React development server
npm run dev:electron  # Start Electron desktop app
npm run build        # Build for production
npm run lint         # Run ESLint
npm run preview      # Preview production build
```

### Docker Commands
```bash
docker-compose up -d              # Start all services
docker-compose down               # Stop all services
docker-compose logs -f            # View logs
docker-compose exec db psql       # Connect to database
```

## Services

The application runs the following services:

| Service | Port | Description |
|---------|------|-------------|
| API Server | 8080 | Go backend API |
| PostgreSQL | 5432 | Main database |
| Redis | 6379 | Caching layer |
| Redis Commander | 8081 | Redis management UI |
| Frontend Dev | 5173 | Vite development server |

## Database Schema

The application uses PostgreSQL with the following main entities:

- **Cards**: Magic: The Gathering card information
  - Scryfall ID integration
  - Card metadata (name, mana cost, oracle text, types, set)
- **Users**: User management system
- **Decks**: Flashcard deck management
- **GoCard**: Main application table

## API Endpoints

The API provides RESTful endpoints for:

- `/api/health` - Health check
- `/api/decks` - Deck management (CRUD operations)
- `/api/cards` - Card management
- `/api/docs` - Swagger documentation

Visit `http://localhost:8080/api/docs` for interactive API documentation.

## Testing

```bash
# Run Go backend tests
cd app
make test

# Run frontend tests (if configured)
cd frontend
npm test
```

## Development

### Adding New Migrations

```bash
make migration <migration_name>
# Edit the generated SQL files in cmd/migrate/migrations/
make migrate-up
```

### API Development

1. Add new models in `app/api/models/`
2. Implement handlers in `app/api/`
3. Update routes in the main application
4. Generate documentation with `make gen-docs`

### Frontend Development

1. Add React components in `frontend/src/ui/`
2. Update Electron main process in `frontend/src/electron/`
3. Build and test with `npm run build`

## Configuration

Key configuration options in `.env`:

```bash
# Database
PG_ADRESS=:8080
PG_MAX_OPEN_CONNS=30
PG_MAX_IDLE_CONNS=30
PG_IDLE_TIME=15min

# PostgreSQL
POSTGRES_DB=your_db_name
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
DATABASE_URL=postgres://user:pass@localhost/dbname?sslmode=disable
```

## Deployment

### Docker Production Build
```bash
docker build -t gocard .
docker run -p 8080:8080 gocard
```

### Manual Deployment
1. Build the Go application: `go build -o bin/gocard app/main.go`
2. Build the frontend: `cd frontend && npm run build`
3. Deploy with your preferred method (systemd, pm2, etc.)

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Run tests: `make test`
5. Commit your changes: `git commit -am 'Add feature'`
6. Push to the branch: `git push origin feature-name`
7. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Additional Resources

- [Go Chi Router Documentation](https://github.com/go-chi/chi)
- [React Documentation](https://react.dev/)
- [Electron Documentation](https://www.electronjs.org/docs)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)

## Support

For questions and support, please open an issue in the GitHub repository. 