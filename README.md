# Go server starter

This repo provides a starting point to develop a backend server with Go. It relies on the Iris Framework for http and socket servers. The tests actually run the server and test against it endpoints (both of them pings at the moment).

## Dev Environment

Run container (from this repo at /)
```sh
./src/scripts/docker.sh
```

This will run the golang image and mount the source code and lib dir (will be created the first time you install dependencies) into the container - the lib dir is gitignored as it is the place where all the dependencies will be stored locally - think of it as a node_modules dir; it doesn't get committed and you do not have to install all dependencies everytime you run the container.

When the container is running and the app is running it exposes the http server to http://localhost:8877

Install dependencies (inside running container at /)
```sh
# -t installs testing dependencies
go get -v -t ./...
```

Build (inside running container anywhere)
```sh
go installgo-server-starter 
```

Execute (inside running container at /)
```sh
./scripts/run.sh
```

Enter running container if needed
```sh
docker exec -it go-server-starter /bin/bash
```

Tests (inside running container at /)
```sh
./scripts/test.sh
```

