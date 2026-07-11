# Docker Compose Commands Reference for Celebut Application

## Basic Commands

### Start all services
```bash
docker-compose up -d
```

### Start with development overrides (live reload, etc.)
```bash
docker-compose -f docker-compose.yml -f docker-compose.override.yml up -d
```

### Stop and remove containers
```bash
docker-compose down
```

### Stop and remove containers with volumes
```bash
docker-compose down -v
```

### View logs
```bash
# Follow logs for all services
docker-compose logs -f

# Follow logs for specific service
docker-compose logs -f backend
docker-compose logs -f ai-service
```

### View running services
```bash
docker-compose ps
```

### Execute commands in running containers
```bash
# Execute bash in backend container
docker-compose exec backend sh

# Execute bash in ai-service container
docker-compose exec ai-service sh

# Run a one-time command
docker-compose run --rm backend go test ./...
```

### View resource usage
```bash
docker-compose stats
```

## Development Workflow

### With Live Reload (using docker-compose.override.yml)
1. Start development environment:
   ```bash
   docker-compose -f docker-compose.yml -f docker-compose.override.yml up -d
   ```

2. The backend will automatically reload when Go files change (using Air)
3. The AI service will automatically reload when Python files change (using Uvicorn reload)

### Without Overrides (Production-like)
1. Start environment:
   ```bash
   docker-compose up -d
   ```

2. To rebuild after code changes:
   ```bash
   docker-compose up -d --build
   ```

## Testing the Setup

### Check if services are healthy
```bash
docker-compose ps
# Look for "healthy" status in the State column
```

### Access API endpoints
- Backend health: http://localhost:7070/healthz
- API docs: http://localhost:7070/swagger/index.html
- AI service health: http://localhost:8000/ai/health
- AI docs: http://localhost:8000/docs

## Common Issues and Solutions

### Port conflicts
If you see errors like "port is already allocated":
- Stop any existing services using those ports
- Or modify the port mappings in docker-compose.yml

### Database connection issues
If the backend can't connect to PostgreSQL:
- Wait a moment for the database to fully initialize
- Check the postgres container logs: `docker-compose logs postgres`
- Ensure the POSTGRES_* environment variables match

### Slow startup
First startup may take several minutes as:
- Docker images are downloaded
- Go modules are fetched
- Python dependencies are installed
- AI models are downloaded (on first AI service start)

### Memory issues
If containers fail due to insufficient memory:
- Increase Docker's memory allocation in Docker Desktop preferences
- Or reduce resource usage by disabling optional services (tempo, prometheus, grafana)

## Production Considerations

For production deployments:
1. Use specific image tags instead of latest
2. Enable resource limits (memory, CPU) in docker-compose.yml
3. Use a proper secrets management system instead of .env files
4. Configure proper logging drivers
5. Set up backup strategies for persistent volumes
6. Consider using Docker Swarm or Kubernetes for orchestration

## Maintenance

### Updating images
```bash
docker-compose pull
docker-compose up -d
```

### Cleaning up unused resources
```bash
docker system prune -f
docker volume prune -f  # Be careful with this - it will delete all volumes!
```

### Backing up data
```bash
# Backup PostgreSQL
docker-compose exec postgres pg_dump -U $POSTGRES_USER $POSTGRES_DB > backup.sql

# Backup Redis (requires redis-cli)
docker-compose exec redis redis-cli save
# Then copy the dump.rdb file from the container
```