# GoCard Project Index

## Directory Structure Overview

### Root Level
- `README.md` - Project documentation and setup guide
- `docker-compose.yml` - Multi-service container orchestration
- `Dockerfile` - Container image build instructions
- `Makefile` - Build automation and common tasks
- `env.dist` - Environment variables template
- `.gitignore` - Git ignore patterns

## Backend (Go Application)

### `/app` - Main Go Application
- `main.go` - Application entry point and server configuration
- `go.mod` / `go.sum` - Go module dependencies
- `/api` - API layer components
  - `/models` - Data models and structs
    - `card.go` - Magic: The Gathering card model
    - `user.go` - User management model
  - `/fetch` - Database access layer
    - `fetch.go` - Storage interface definitions
    - `cards.go` - Card data operations
    - `users.go` - User data operations
  - `/swagger` - API documentation
    - `swagger.json` - OpenAPI specification
    - `swagger-ui.html` - Interactive API documentation
  - `decks.go` - Deck management handlers
  - Route handlers and middleware
- `/config` - Configuration management
  - Environment variable handling
- `/test` - Backend unit tests

### `/cmd` - Command Line Tools
- `/api` - API server entry point
- `/migrate` - Database migration tools
  - `/migrations` - SQL migration files
  - `/seed` - Database seeding utilities

## Frontend (Electron/React Application)

### `/frontend` - Desktop Application
- `package.json` - Node.js dependencies and scripts
- `package-lock.json` - Dependency lock file
- `tsconfig.json` / `tsconfig.*.json` - TypeScript configuration
- `vite.config.ts` - Vite build tool configuration
- `eslint.config.js` - ESLint code quality rules
- `index.html` - Main HTML entry point
- `/src` - Source code
  - `/electron` - Electron main process
    - Application window management
    - System integration
  - `/ui` - React user interface
    - Component library
    - Application views
    - State management
- `/node_modules` - Node.js dependencies
- `/dist-react` - Built React application

## Database & Data

### `/data` - Database Files
- `/postgres` - PostgreSQL configuration
  - `db_init.sql` - Database initialization script
  - `/migrations` - SQL schema migrations
    - Sequential migration files
    - Schema evolution tracking

## Infrastructure

### `/docker` - Docker Configuration
- Container-specific configurations
- Development vs production setups

### `/config` - Application Configuration
- Environment-specific settings
- Configuration templates
- Service configurations

## Scripts & Automation

### `/scripts` - Utility Scripts
- Development helpers
- Deployment scripts
- Data management tools

### `/bin` - Binary Files
- Compiled executables
- Build artifacts

## Documentation

### `/docs` - Documentation Files
- Technical specifications
- API documentation
- User guides
- Development notes

### `/public` - Static Assets
- Public web assets
- Documentation resources
- Static files served by the application

## CI/CD

### `/ci-cd` - Continuous Integration/Deployment
- Pipeline configurations
- Deployment scripts
- Testing automation

## Key Technologies & Dependencies

### Backend Stack
```
Go 1.24
├── github.com/go-chi/chi/v5 (HTTP router)
├── github.com/lib/pq (PostgreSQL driver)
├── golang-migrate (Database migrations)
└── Swagger/OpenAPI (API documentation)
```

### Frontend Stack
```
Node.js
├── React 19.1.0 (UI framework)
├── Electron 36.4.0 (Desktop app framework)
├── TypeScript 5.8.3 (Type safety)
├── Vite 6.3.5 (Build tool)
└── ESLint (Code quality)
```

### Infrastructure Stack
```
Docker & Docker Compose
├── PostgreSQL 16.3 (Primary database)
├── Redis 6.2 (Caching layer)
├── Redis Commander (Redis management UI)
└── Migrate (Database migration tool)
```

## File Types & Extensions

### Go Files (`.go`)
- Backend application logic
- API handlers and middleware
- Database operations
- Business logic

### TypeScript/JavaScript (`.ts`, `.tsx`, `.js`)
- Frontend React components
- Electron main process
- Build configurations
- Test files

### SQL Files (`.sql`)
- Database schemas
- Migration scripts
- Initial data seeds
- Query definitions

### Configuration Files
- `.json` - Package manifests, configs
- `.yml`/`.yaml` - Docker Compose, CI/CD
- `.env` - Environment variables
- `Makefile` - Build automation
- `Dockerfile` - Container definitions

### Documentation (`.md`)
- README files
- Technical documentation
- API guides
- Development notes

## Integration Points

### API Endpoints
- Health checks (`/api/health`)
- Deck management (`/api/decks/*`)
- Card operations (`/api/cards/*`)
- User management (`/api/users/*`)
- Documentation (`/api/docs`)

### Database Connections
- PostgreSQL primary database
- Redis caching layer
- Migration management
- Data seeding

### External Services
- Scryfall API (Magic: The Gathering card data)
- Docker registry (container images)
- Package managers (Go modules, npm)

##  Development Workflow

### Local Development
1. Environment setup (`.env` configuration)
2. Database initialization (`make migrate-up`)
3. Backend development (`go run main.go`)
4. Frontend development (`npm run dev:react` or `npm run dev:electron`)
5. Testing (`make test`, `npm test`)

### Build Process
1. Go compilation (`go build`)
2. Frontend building (`npm run build`)
3. Docker image creation (`docker build`)
4. Documentation generation (`make gen-docs`)

### Deployment Pipeline
1. Code quality checks (ESLint, go fmt)
2. Unit testing (Go tests, React tests)
3. Integration testing
4. Container building
5. Service deployment

##  Quality Assurance

### Code Quality Tools
- **Go**: `go fmt`, `go vet`, `go test`
- **Frontend**: ESLint, TypeScript compiler
- **SQL**: Migration validation
- **Docker**: Dockerfile linting

### Testing Strategy
- Unit tests for backend logic
- Component tests for React UI
- Integration tests for API endpoints
- End-to-end testing for user workflows

This index provides a comprehensive overview of the GoCard project structure, making it easier to navigate and understand the codebase architecture. 