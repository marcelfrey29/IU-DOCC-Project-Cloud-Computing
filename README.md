# IU-DOCC-Project-Cloud-Computing

[![Node.js CI](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-node.yml/badge.svg)](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-node.yml)
[![Go CI](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-go.yml/badge.svg)](https://github.com/marcelfrey29/IU-DOCC-Project-Cloud-Computing/actions/workflows/ci-go.yml)
[![SAST](https://github.com/marcelfrey29/marcelfrey29.github.io/actions/workflows/sast.yml/badge.svg)](https://github.com/marcelfrey29/marcelfrey29.github.io/actions/workflows/sast.yml)

## Documentation

See [Docs](docs/docs.md)

## Requirements

- Current Node.js LTS (Node.js 22)
- Golang >= 1.23.0
- [Bearer CLI](https://github.com/Bearer/bearer) must be installed for Security Scanning (SAST)
    - Run SAST with `bearer scan .` in project root

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
