# Golang with Docker

**ONLY for DEV, not for production**

A very simple Docker-compose to use Go and Docker
## Run Locally

Clone the project

```bash
 git@github.com:TechRaceProject/golang-api.git
```

Run the docker-compose

```bash
  docker compose build --no-cache --pull
  docker-compose up -d
```

Log into the GO container

```bash
  docker exec -it  nom-du-container bash
```