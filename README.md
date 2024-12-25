# IU-DOCC-Project-Cloud-Computing

[![Node.js CI](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-node.yml/badge.svg)](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-node.yml)
[![Go CI](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-go.yml/badge.svg)](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-go.yml)
[![SAST](https://github.com/marcelfrey29/marcelfrey29.github.io/actions/workflows/sast.yml/badge.svg)](https://github.com/marcelfrey29/marcelfrey29.github.io/actions/workflows/sast.yml)

## Documentation

See [Docs](docs/docs.md)

## Requirements

- Current Node.js LTS (Node.js 22)
- Golang >= 1.23.0
- Docker (incl. Docker Compose)
- AWS CLI must be installed and configured
    - For local execution (e.g. DynamoDB Local) Dummy-Access-Keys can be used
    ```conf
    # ~/.aws/credentials
    [default]
    aws_access_key_id = XXXXX
    aws_secret_access_key = XXXXX
    ```
- [Bearer CLI](https://github.com/Bearer/bearer) must be installed for Security Scanning (SAST)
    - Run SAST with `bearer scan .` in project root
- [Bruno](https://www.usebruno.com/) (Optional) as HTTP Client
    - Collection with Assertions is located in `tools/bruno`

## Running the Project

All services can be managed at once via Docker Compose.

```bash
# Run latest version of all services
docker compose up --build
docker compose up --build -d

# Run all services
docker compose up
docker compose up -d

# Remove all services
docker compose down

# Stop & Start
docker compose stop
docker compose start

# Get logs when running in detached mode
# docker logs <container_name>
docker logs iu-project-cc-backend
```

### Initial Setup of DynamoDB Local

When the DynamoDB Local Container started for the first time, we need to create a Table manually (just like we have to when using real AWS infrastructure).
To perform this task, the `tools/create-local-table.sh` script can be used. 

> [!IMPORTANT]
> DynamoDB Local must be running (Run `docker compose up`).

```bash
# Create Table
./tools/create-local-table.sh

# Verify Table Creation
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

Expected Output: 
```json
{
    "TableNames": [
        "TravelGuides"
    ]
}
```
