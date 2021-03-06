docker run -it --rm \
  --volume "$PWD/main.go:/go/src/github.com/AquiGorka/go-server-starter/main.go" \
  --volume "$PWD/server:/go/src/github.com/AquiGorka/go-server-starter/server" \
  --volume "$PWD/scripts:/go/src/github.com/AquiGorka/go-server-starter/scripts" \
  --volume "$PWD/pkg:/go/src" \
  --name go-server-starter \
  --workdir "/go/src/github.com/AquiGorka/go-server-starter" \
  --entrypoint /bin/bash \
  -p 8877:8080 golang
