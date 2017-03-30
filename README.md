# snips-platform-lambda-samples

## Golang sample

### Installation

```
$ cd ~/go/src
$ git clone git@github.com:snipsco/snips-platform-lambda-samples.git
$ go install snips-platform-lambda-samples/...
```

### Configuration

Edit `go/hue/conf.ini` to setup:
- Snips Platform host/port
- your Philips Hue router/username/bulbs

### Start

```
$ ~/go/bin/hue -c ~/go/src/snips-platform-lambda-samples/go/hue/conf.ini
```