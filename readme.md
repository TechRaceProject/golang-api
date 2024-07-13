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

## Testing

Run api tests
```bash
  docker-compose exec api go test ./src/tests/... -v
```

Run every api tests from a single test repository
```bash
  docker-compose exec api go test ./src/tests/auth/login -v
```

Run a single test file 
```bash
  docker-compose exec api go test ./src/tests/auth/login -v -run Test_can_login_if_valid_email_and_password_are_provided
```
