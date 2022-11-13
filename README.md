# golang-united-homework
This repository contains code of gRPC service for homework.

## Requirements
- Docker installed
- PostgreSQL database installed locally or in the docker container

## How to run the project
1. Set environment variables with correct values:

```bash
export HOMEWORK_DB_HOST=
export HOMEWORK_DB_PORT=
export HOMEWORK_DB_USER=
export HOMEWORK_DB_PASSWORD=
export HOMEWORK_DB_DATABASE=
```

2. To start gRPC server:

```bash
make start
```

3. To stop gRPC server:

```bash
make stop
```
