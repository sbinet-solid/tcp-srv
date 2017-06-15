# solid-mon-rpi

`solid-mon-rpi` is an HTTP server publishing data from `I2C` and `SMBus` sensors.

## Example

```sh
$> solid-mon-rpi -addr=:80
tcp-srv starting up server on: :10000
tcp-srv daq: {Timestamp:2017-01-12 20:05:41.512823703 +0000 UTC Tsl:{Lux:7.017600000000001 Full:5 IR:2} Sht31:{Temp:19.309529259174482 Hum:23.302052338445105} Si7021:[{Temp:19.087900390625002 Hum:23.521942138671875} {Temp:19.034274902343746 Hum:22.682708740234375}] Bme:{Temp:20.034765625 Hum:34.294829763189966 Pres:976.813259481733}}
tcp-srv daq: {Timestamp:2017-01-12 20:05:44.111935763 +0000 UTC Tsl:{Lux:7.017600000000001 Full:5 IR:2} Sht31:{Temp:19.280155642023345 Hum:23.300526436255435} Si7021:[{Temp:19.06645019531249 Hum:23.529571533203125} {Temp:19.034274902343746 Hum:22.659820556640625}] Bme:{Temp:20.044921875 Hum:34.23593420511301 Pres:976.8052497789915}}
[...]
```

### client

One can inspect what `tcp-srv` serves like so:

```sh
$> netcat clrmedaq02.in2p3.fr 10000
^{"timestamp":"2017-01-12T19:31:25.705277264Z","tsl":{"lux":7.017600000000001,"full":5,"ir":2},"sht31":{"temp":19.293507286182958,"hum":23.44701304646372},"si7021":[{"temp":19.06645019531249,"hum":23.499053955078125},{"temp":19.01282470703125,"hum":22.606414794921875}],"bme280":{"temp":20.012109375,"hum":34.75959760679188,"pres":976.9939283323138}}
^{"timestamp":"2017-01-12T19:31:28.318336683Z","tsl":{"lux":7.017600000000001,"full":5,"ir":2},"sht31":{"temp":19.293507286182958,"hum":23.48973830777447},"si7021":[{"temp":19.055725097656243,"hum":23.506683349609375},{"temp":19.002099609375,"hum":22.598785400390625}],"bme280":{"temp":20.0138671875,"hum":34.76515944214093,"pres":977.0379350899934}}
]{"timestamp":"2017-01-12T19:31:30.929595755Z","tsl":{"lux":7.017600000000001,"full":5,"ir":2},"sht31":{"temp":19.280155642023345,"hum":23.460746166170747},"si7021":[{"temp":19.044999999999995,"hum":23.521942138671875},{"temp":19.002099609375,"hum":22.606414794921875}],"bme280":{"temp":20.01328125,"hum":34.82477987227392,"pres":977.0232661389944}}
```

Or, via the `tcp-cli` Go command:
```sh
$> tcp-cli
tcp-cli dialing: clrmedaq02.in2p3.fr:10000
{"timestamp":"2017-01-12T19:33:00.99704373Z","tsl":{"lux":7.017600000000001,"full":5,"ir":2},"sht31":{"temp":19.266803997863732,"hum":23.471427481498434},"si7021":[{"temp":19.023549804687498,"hum":23.552459716796875},{"temp":19.002099609375,"hum":22.644561767578125}],"bme280":{"temp":20.001953125,"hum":34.72083367101912,"pres":977.0775195381289}}
{"timestamp":"2017-01-12T19:33:03.607593279Z","tsl":{"lux":7.017600000000001,"full":5,"ir":2},"sht31":{"temp":19.253452353704134,"hum":23.50194552529183},"si7021":[{"temp":19.055725097656243,"hum":23.552459716796875},{"temp":18.991374511718753,"hum":22.659820556640625}],"bme280":{"temp":20.0025390625,"hum":34.726303666877925,"pres":977.0750095726227}}
```
