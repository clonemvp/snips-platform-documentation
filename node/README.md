# Handler code examples for your Snips assistant


When you have built an assistant on <https://snips.ai> by following [the tutorial](https://github.com/snipsco/snips-platform-documentation/wiki), you can write your own handler to act on the behalf of the user.

This repository contains a few examples in Python

## Install requirements

The main requirement of the handler is to watch the MQTT bus with the Hermes messages from the Snips AI.

You can install node on your Raspberry Pi from the packages of <https://nodejs.org/en/download>

```
wget https://nodejs.org/dist/v6.11.0/node-v6.11.0-linux-armv7l.tar.xz
tar -xf node-v6.11.0-linux-armv7l.tar.xz
cd node-v6.11.0-linux-armv7l
sudo cp -R * /usr/local
```

Then all you need is a MQTT client:

```
npm install mqtt --save
```

You also need to have [setup your platform with the Snips assistant](https://github.com/snipsco/snips-platform-documentation/wiki/1.-Setup-the-Snips-Voice-Platform-on-your-Raspberry-Pi) and copied an assistant model (you can use the [demo IoT assistant](https://github.com/snipsco/snips-platform-documentation/raw/master/resources/iot_assistant.zip)) to `/opt/snips/config` on your device

## Display messages

The components of the AI are communicating and coordinating through a MQTT bus, a good way to understand what your on-device Snips assistant is doing is to look at the messages on the bus

```
cd display_messages
npm install

node index.js
-  hermes/hotword/detected:
-  hermes/hotword/wait:
-  hermes/asr/toggleOn:
-  hermes/audioServer/playFile: {"filePath":"/usr/share/snips/dialogue/sound/start_of_input.wav"}
-  hermes/asr/textCaptured: {"text":"turn the lights blue","likelihood":0.0034613716,"seconds":3462000.0}
-  hermes/asr/toggleOff:
-  hermes/audioServer/playFile: {"filePath":"/usr/share/snips/dialogue/sound/end_of_input.wav"}
-  hermes/nlu/query: {"text":"turn the lights blue","likelihood":0.0034613716,"seconds":3462000.0}
-  hermes/nlu/intentParsed: {"input":"turn the lights blue","intent":{"intentName":"ActivateLightColor","probability":0.9836065},"slots":[{"value":{"kind":"Custom","value":"blue"},"rawValue":"blue","range":{"start":16,"end":20},"entity":"objectColor","slotName":"objectColor"}]}
-  hermes/intent/ActivateLightColor: {"input":"turn the lights blue","intent":{"intentName":"ActivateLightColor","probability":0.9836065},"slots":[{"value":{"kind":"Custom","value":"blue"},"rawValue":"blue","range":{"start":16,"end":20},"entity":"objectColor","slotName":"objectColor"}]}
```


## IoT Assistant

If you followed [this tutorial](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant) you have built an assistant to handle a few interesting intents for IoT. You can also install a prebuilt [IoT assistant ](https://github.com/snipsco/snips-platform-documentation/raw/master/resources/iot_assistant.zip) and uncompress it on your device in `/opt/snips/config`

The code is a handler which shows how to get the intents as they are sent by the natural language understanding, and do an action by sending a helpful message on the text-to-speech channel
