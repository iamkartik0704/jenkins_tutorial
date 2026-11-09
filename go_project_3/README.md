# Go Jenkins CI/CD Demo

A small Go application integrated with Jenkins and DockerHub. The pipeline formats, analyzes, tests, and builds the application before packaging it into a Docker image.

## Project Structure

```text
maven-jenkins-cicd-demo/
├── main.go
├── main_test.go
├── go.mod
├── Dockerfile
└── Jenkinsfile
```

## Requirements

- Go 1.22 or later
- Docker
- Jenkins with Docker access and a Go-enabled agent

## Run Locally

```bash
go test ./...
go run .
```

Expected output:

```text
Hello Jenkins!
```

Build a standalone binary with:

```bash
go build -o dist/app .
```

## Docker

Build and run the container:

```bash
docker build -t theshubhamgour/maven-jenkins-demo:latest .
docker run --rm theshubhamgour/maven-jenkins-demo:latest
```

The Dockerfile compiles the program in a Go build stage and copies only the executable into a minimal runtime image.

## Jenkins

The Jenkins pipeline checks Go formatting, runs `go vet` and unit tests, and builds a binary. It then builds and pushes a Docker image, runs it as a container, and cleans up the temporary container and image.

Configure the Jenkins username/password credential with ID `DockerHub`. The pipeline expects Docker access on the Jenkins agent and publishes images under `theshubhamgour/maven-jenkins-demo`.

## License

This project is licensed under the MIT License.
