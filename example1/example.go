package main

import "fmt"

var deployment = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-app
  labels:
    app: example-app
spec:
  replicas: 2
  selector:
    matchLabels:
      app: example-app
  template:
    metadata:
      labels:
        app: example-app
    spec:
      containers:
      - name: example-app-container
        image: nginx:latest  # Replace with your actual container image
        ports:
        - containerPort: 80
`

func main() {
  fmt.Println(deployment)
}