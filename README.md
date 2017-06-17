# Go server starter [![Build Status](https://travis-ci.org/AquiGorka/go-server-starter.svg?branch=master)](https://travis-ci.org/AquiGorka/go-server-starter) [![Coverage Status](https://coveralls.io/repos/github/AquiGorka/go-server-starter/badge.svg?branch=master)](https://coveralls.io/github/AquiGorka/go-server-starter?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/AquiGorka/go-server-starter)](https://goreportcard.com/report/github.com/AquiGorka/go-server-starter)

This repo provides a starting point to develop a backend server with Go. It relies on the Iris Framework for http and socket servers. The tests actually run the server and test against both ping endpoints.

## Go getable
```sh
go get github.com/AquiGorka/go-server-starter
```

## Dev Environment

Run container (from this repo at /)
```sh
./scripts/docker.sh
```

This will run the golang image and mount the source code and lib dir (will be created the first time you install dependencies) into the container - the lib dir is gitignored as it is the place where all the dependencies will be stored locally - think of it as a node_modules dir; it doesn't get committed and you do not have to install all dependencies everytime you run the container.

When the container is running and the app is running it exposes the server to http://localhost:8877

Install dependencies (inside running container at ...go-server-starter/)
```sh
# -t installs testing dependencies
go get -v -t ./...
```

Build (inside running container at ...go-server-starter/)
```sh
go install
# run
go-server-starter
```

Execute (inside running container at ...go-server-starter/)
```sh
./scripts/run.sh
```

Enter running container if needed
```sh
docker exec -it go-server-starter /bin/bash
```

Tests (inside running container at ...go-server-starter/)
```sh
./scripts/test.sh
```
