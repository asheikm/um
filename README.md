# User Management

This project is a user management system built in Go, with a REST API for managing users.

## Getting Started

### Prerequisites

- Go 1.17 or higher
- Docker
- Kubernetes (minikube for local development)

### Building the Application

To build the application, run:

```sh
go build
```

### Running the Application

To run the application locally, use the following command:

```sh
./um
```

### Docker
To build a Docker image, use the following command:

```sh
docker build -t um:latest .
```

To run the Docker container, use the following command:

```sh
docker run -p 8080:8080 um:latest
```

### Kubernetes
To deploy the application to Kubernetes, make sure you have a Kubernetes cluster running (e.g., minikube or private setup or public cloud services like EKS ).

Apply the Kubernetes deployment file:

```sh
kubectl apply -f um-deployment.yaml
```

### Helm
To deploy the application using Helm, make sure you have Helm installed and initialized.

```sh
helm install um um-chart
```
