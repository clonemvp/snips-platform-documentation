# snips-platform-lambda-samples

## Golang sample

### Installation

```
$ cd ~/go/src
$ git clone git@github.com:snipsco/snips-platform-lambda-samples.git
$ go get github.com/eclipse/paho.mqtt.golang
$ go install snips-platform-lambda-samples/...
```

### Hue sample

#### Configuration

Edit `go/hue/conf.ini` to setup:
- Snips Platform host/port
- your Philips Hue router/username/bulbs

#### Start

```
$ ~/go/bin/hue -c go/hue/conf.ini
```

### Weather sample

#### Configuration

Edit `go/weather/conf.ini` to setup:
- Snips Platform host/port
- your openweathermap.org api key

#### Start

```
$ ~/go/bin/weather -c go/weather/conf.ini
```