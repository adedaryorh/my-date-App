# Celebut Application - Docker Compose Setup

This directory contains Docker Compose files for running the Celebut application with all its dependencies.

## Services

The following services are defined in the docker-compose.yml:

- **postgres**: PostgreSQL 15 database for persistent storage
- **redis**: Redis 7 for caching and session storage
- **rabbitmq**: RabbitMQ 3 with management plugin for message queuing
- **backend**: Go/Gin application server ( Celebut backend API )
- **ai-service**: Python/FastAPI microservice for AI features (recommendations, moderation, search)
- **tempo**: Grafana Tempo for distributed tracing (optional)
- **prometheus**: Prometheus for metrics collection (optional)
- **grafana**: Grafana for visualization (optional)

## Environment Variables

The application uses environment variables for configuration. Default values are provided in the docker-compose.yml file, but you can override them by:

1. Creating a `.env` file in the root directory
2. Setting environment variables in your shell before running docker-compose
3. Using docker-compose.override.yml for development-specific overrides

See the individual service sections in docker-compose.yml for specific environment variables used.

## Getting Started

1. Make sure you have Docker and Docker Compose installed

2. Copy the example environment file (if provided) or create your own:
   ```bash
   cp .env.example .env  # If an example exists
   # Or create your own .env file with required variables
   ```

3. Start the application:
   ```bash
   docker-compose up -d
   ```

4. To view logs:
   ```bash
   docker-compose logs -f
   ```

5. To stop and remove containers:
   ```bash
   docker-compose down
   ```

6. To start with development overrides (if using docker-compose.override.yml):
   ```bash
   docker-compose -f docker-compose.yml -f docker-compose.override.yml up -d
   ```

## Service Endpoints

Once the containers are running, you can access:

- **Backend API**: http://localhost:7070
- **API Documentation**: http://localhost:7070/swagger/index.html
- **AI Service**: http://localhost:8000
- **AI Service Documentation**: http://localhost:8000/docs
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (default admin/admin)
- **RabbitMQ Management**: http://localhost:15672 (default guest/guest)
- **Adminer (DB GUI)**: http://localhost:8080 (if enabled)

```

## Data Persistence

The following volumes are used for persistent data:
- `postgres_data`: PostgreSQL database files
- `redis_data`: Redis data (append-only file)
- `rabbitmq_data`: RabbitMQ message store
- `tempo_data`: Tempo trace data
- `prometheus_data`: Prometheus metrics data
- `grafana_data`: Grafana configuration and dashboards

## Development Workflow

For development, you might want to:
1. Use the docker-compose.override.yml for live reload capabilities
2. Mount your local code directories into the containers
3. Use development-specific configurations

The docker-compose.override.yml includes:
- Live reload for Go backend using Air
- Live reload for Python AI service using Uvicorn's reload feature
- Additional tools like Adminer for database inspection

## Customizing

To customize the deployment:
1. Modify the docker-compose.yml for production settings
2. Use docker-compose.override.yml for development overrides
3. Adjust resource limits (memory, CPU) as needed for your environment
4. Add additional services as required by your specific setup

## Troubleshooting

If you encounter issues:
1. Check container logs: `docker-compose logs [service-name]`
2. Verify ports are not already in use on your host
3. Ensure sufficient resources (RAM/CPU) are available
4. Check that dependencies are properly configured (databases initialized, etc.)

For database migrations, the next level all services will be running




cd /Users/adedaryorh/Documents/myDate-App/backend/app
go build -o celebut-api .
./celebut-api

cd /Users/adedaryorh/Documents/myDate-App/ai-service
rm -rf .venv

/opt/homebrew/bin/python3.11 -m venv .venv
source .venv/bin/activate


python -m pip install --upgrade pip setuptools wheel
python -m pip install -r requirements.txt

Then create the local storage directories:
mkdir -p models data
Set the AI environment:
export DATABASE_URL="postgresql://celebut:celebut_local_password@127.0.0.1:5432/celebut_db?sslmode=disable"
export MODEL_CACHE_DIR="./models"
export FAISS_INDEX_PATH="./data/faiss_index.bin"
export EMBEDDING_MODEL_NAME="all-MiniLM-L6-v2"
export TOXICITY_MODEL_NAME="unitary/toxic-bert"
export TOXICITY_THRESHOLD="0.8"

python -m uvicorn app.main:app --host 0.0.0.0 --port 8000 --reload


cd /Users/adedaryorh/Documents/myDate-App/ai-service
source .venv/bin/activate
uvicorn app.main:app --reload --host 0.0.0.0 --port 8000