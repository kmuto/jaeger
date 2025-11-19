# Send traces to Mackerel

```
$ git clone https://github.com/kmuto/jaeger.git
$ cd jaeger/examples/hotrod
$ git checkout vaxila-demo

# Make hotrod-linux binary
$ GOARCH=amd64 docker compose -f docker-compose-gobuild.yml run --rm gobuild
# or/and
$ GOARCH=arm64 docker compose -f docker-compose-gobuild.yml run --rm gobuild

# Build kmuto:hotrod image
$ docker build -t kmuto:hotrod .
# (or for Docker push)
$ docker buildx build --platform linux/amd64,linux/arm64 -t ghcr.io/kmuto/hotrod . --push

# Set API Key. You can get API Key from Mackerel setting screen
$ export MACKEREL_APIKEY=...

# Up HotRod on http://localhost:8080
#              http://localhost:8081
#              http://localhost:8082
#              http://localhost:8083
$ docker compose up

 ...(play with HotRod and see traces on Mackerel APM screen)...

# Finished
$ docker compose down
```

- http://localhost:8080 : Default. SQL SELECT lock, route concurrent limit over
- http://localhost:8081 : SQL SELECT is solved, route concurrent limit over
- http://localhost:8082 : SQL SELECT is solved, route concurrent limit over is solved
- http://localhost:8083 : Amazing Coffee Roasters causes a bug
- http://localhost:16686 : Jaeger UI
