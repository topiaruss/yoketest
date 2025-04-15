# yoketest

From: https://yokecd.github.io/docs/examples/basics/

install go and yoke
```bash
brew install go
brew install yoke
```

try example1
```bash
cd example1
GOOS=wasip1 GOARCH=wasm go build -o example.wasm ./example.go
yoke takeoff example ./example.wasm
cd ..
```

Then check your cluster

```bash
% kubectl get deployments
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
example-app   2/2     2            2           2m55s
```

remove the deployment then...
```bash
cd example2
GOOS=wasip1 GOARCH=wasm go build -o example.wasm ./example.go
yoke takeoff example ./example.wasm
cd ..
```

And check the cluster again
