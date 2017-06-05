docker run -it --rm --volume "$PWD/src:/go/src/go-server-starter" --volume "$PWD/lib:/go/src" --name go-server-starter --workdir "/go/src/go-server-starter" --entrypoint /bin/bash -p 8877:8080 golang
