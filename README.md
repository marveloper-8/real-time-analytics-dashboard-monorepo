# Real-time Analytics Dashboard

This project is a real-time analytics dashboard built with Go, GraphQL, InfluxDB, and Next.js. It demonstrates a robust, scalable architecture for handling time-series data and real-time updates.

## Features

- Go backend with GraphQL API
- InfluxDB for time-series data storage
- Next.js frontend with real-time updates using subscriptions
- Docker containerization
- Kubernetes deployment
- CI/CD with GitHub Actions
- Infrastructure as Code with Terraform

## Prerequisites

- Docker and Docker Compose
- Go 1.18+
- Node.js 16+
- kubectl
- Terraform
- DigitalOcean account (for deployment)

## Getting Started

1. Clone the repository:
git clone https://github.com/marveloper-8/real-time-analytics-dashboard-monorepo
cd analytics-dashboard

2. Set up environment variables:
Create a `.env` file in the root directory and add the following:
INFLUXDB_TOKEN=your_influxdb_token
INFLUXDB_ORG=your_org
INFLUXDB_BUCKET=your_bucket

3. Build and run the project:
make build
make run

4. Access the dashboard at `http://localhost:3000`

## Development

- Generate GraphQL code: `make generate-backend`
- Run tests: `make test`
- Lint code: `make lint`
- Format code: `make format`

## Deployment

1. Set up DigitalOcean infrastructure:
cd infra/terraform
terraform init
terraform apply

2. Deploy to Kubernetes:
make deploy

## Project Structure

- `backend/`: Go backend with GraphQL API
- `frontend/`: Next.js frontend
- `infra/`: Kubernetes and Terraform configurations
- `.github/workflows/`: CI/CD pipelines
- `docker-compose.yml`: Local development setup

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.