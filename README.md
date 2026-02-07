# GoWatch

GoWatch is a lightweight observability tool tailored for Go developers building microservices in Docker stacks, with planned AWS Serverless support. It enables fast, customizable, and extensible monitoring of local services during development, providing real-time insights into performance metrics, logs, distributed tracing, resource consumption, and API status.

## Features

- **Local Service Monitoring**: Track and visualize metrics for services running locally
- **Performance Metrics**: Monitor key performance indicators such as response times, throughput, and error rates
- **Log Aggregation**: Collect and display logs from multiple sources for easier debugging
- **Distributed Tracing**: Trace requests across microservices to identify bottlenecks and dependencies
- **Resource Usage Tracking**: Monitor CPU, memory, and other resource consumption
- **API Status Monitoring**: Check the health and status of APIs in real-time
- **Docker Integration**: Seamlessly monitor services running in Docker containers
- **Future AWS Integration**: Planned support for AWS services like CloudWatch, XRay, Lambda, and CloudFormation
- **Terminal UI**: Interactive dashboard for real-time monitoring and visualization

## Development Environment

GoWatch includes a fully configured Docker development environment:

### Docker Development Container

The project includes a complete Docker setup for development:

- **Dockerfile**: Multi-stage build using Go tip-alpine and Alpine runtime
- **Docker Compose**: Pre-configured service with volume mounts and Docker socket access
- **Development Tools**: Includes security scanning and field alignment tools

### Quick Start with Docker

```bash
# Build the Docker development environment
make docker-build

# Run the observability tool in Docker
docker compose up
```

### Manual Installation

To install GoWatch locally, ensure you have Go 1.25.3 or later installed:

```bash
git clone https://github.com/b92c/gowatch.git
cd gowatch
make install
make run
```

## Build Commands

```bash
# Install dependencies and hooks
make install

# Build the binary
make build

# Run the application
make run

# Build Docker image
make docker-build

# Run tests
make test

# Security scan
make go-sec

# Fix struct field alignment
make field-fix
```

## Project Structure

```
gowatch/
├── cmd/gowatch/          # Application entry point
├── internal/
│   ├── aws/              # Future AWS integrations (XRay, CloudWatch, Lambda, CloudFormation)
│   ├── docker/           # Docker monitoring and collection
│   ├── trace/            # Distributed tracing functionality
│   ├── ui/               # Terminal UI components and dashboard
│   └── config/           # Configuration management
├── pkg/metrics/          # Metrics types and definitions
├── docker-compose.yaml   # Docker development environment
├── Dockerfile           # Multi-stage build configuration
└── makefile            # Build and development commands
```

## Configuration

GoWatch is designed to work out-of-the-box with Docker containers and AWS services. The tool automatically detects:

- Running Docker containers
- Docker daemon socket for container metrics
- Future AWS credentials for serverless monitoring
- Local services and endpoints

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and security checks
5. Submit a pull request

## License

MIT
