# Handler code examples for your Snips assistant


When you have built an assistant on <https://snips.ai> by following [the tutorial](https://github.com/snipsco/snips-platform-documentation/wiki), you can write your own handler to act on the behalf of the user.

This repository contains a few examples in Python

## Install requirements

The main requirement of the handler is to watch the MQTT bus with the Hermes messages from the Snips AI.

All you need is a MQTT client:

```
sudo apt-get install python python-pip
pip install paho-mqtt
```

## Display messages

The components of the AI are communicating and coordinating through a MQTT bus, a good way to understand what your on-device Snips assistant is doing is to look at the messages on the bus

```
python display_messages/display_messages.py
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

If you followed [this tutorial](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant) you have built an assistant to handle a few interesting intents for IoT.

The code is a handler which shows how to get the intents as they are sent by the natural language understanding, and do an action by sending a helpful message on the text-to-speech channel
