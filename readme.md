# Go Prometheus Grafana

A sample Go application demonstrating the integration of Prometheus metrics monitoring and Grafana visualization using the Gin web framework.

## Features

- REST API endpoints with Gin framework
- Prometheus metrics integration
- Grafana dashboards
- Load testing with k6
- Docker compose setup

## Prerequisites

- Go 1.25+
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Project Structure

```
.
├── docker-compose.yml    # Main Docker compose file for Prometheus and Grafana
├── go.mod               # Go module file
├── main.go              # Main application file
├── Makefile            # Build and run commands
├── k6/                 # Load testing
│   ├── api_test.js     # k6 test scripts
│   └── docker-compose.yml
├── middleware/         # Custom middleware
│   └── prometheus.go   # Prometheus metrics middleware
└── prometheus/        # Prometheus configuration
    └── prometheus.yml
```

## Metrics

The application tracks the following metrics:

- `my_app_request_total`: Total number of processed requests
- `my_request_error_total`: Total number of error requests

## API Endpoints

- `GET /`: Hello endpoint
- `GET /metrics`: Prometheus metrics endpoint
- `GET /get-user`: User endpoint
- `GET /get-role`: Role endpoint
- `GET /get-level`: Level endpoint

## Getting Started

1. Start the application:
```bash
make run
```

2. Start Prometheus and Grafana:
```bash
make install-prometheus
```

3. Run load tests:
```bash
make run-k6
```

4. Access services:
- Application: http://localhost:8080
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (default credentials: admin/admin)

## Stopping Services

To stop the services:

```bash
make down-prometheus
make down-k6
```

## Load Testing

The project includes k6 load tests that simulate:
- Ramp-up to 10 users over 1 minute
- Maintain 10 users for 2 minutes
- Ramp-down to 0 users over 1 minute

The tests cover both successful and error scenarios for all endpoints.

## Contributing

Feel free to submit issues and pull requests.

## License

This project is open source and available under the [MIT License](LICENSE).