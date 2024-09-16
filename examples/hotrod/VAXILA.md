# Send traces to Vaxila

```
$ cd examples/hotrod
# Make hotrod-linux binary
$ docker compose -f docker-compose-gobuild.yml run --rm gobuild
# Build kmuto:hotrod image
$ docker build -t kmuto:hotrod .
# Set API Key. You can get API Key from Vaxila setting screen
$ export MACKEREL_VAXILA_APIKEY=...
# Up HotRod on http://localhost:8080
$ docker compose up
 ...(play with HotRod and see traces on Vaxila screen)...
# Finished
$ docker compose down
```
