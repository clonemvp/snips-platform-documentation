# Snips Voice Platform Documentation

![Snips Voice Platform](https://s3.amazonaws.com/get.docs.snips.ai/static/images/wiki/snips_banner_prod.png)

### About the Platform
The Snips Voice Platform allows anyone to integrate AI powered voice interaction in their devices with ease. The end-to-end pipeline - Hotword detection, Automatic Speech Recognition (ASR) and Natural Language Understanding (NLU) - runs fully on device, powered by state of the art deep learning. By using Snips, you can avoid cloud provider costs, cloud latency, and protect user's privacy.

Your voice assistant in English, French, German, Spanish or Korean (more to come) can be configured easily via a web console. You can select pre-built assistants, or create completely custom ones. Today, the assistant can be deployed to a Raspberry Pi 3. More platforms are available for enterprise clients, contact us at contact@snips.ai.


### Getting Started
You will find everything you need in the [wiki](https://github.com/snipsco/snips-platform-documentation/wiki). You can get started with the [overview](https://github.com/snipsco/snips-platform-documentation/wiki), or alternatively, navigate the doc from here:

0. [Overview](https://github.com/snipsco/snips-platform-documentation/wiki)
    * [Hardware requirements](https://github.com/snipsco/snips-platform-documentation/wiki#hardware-requirements)
    * [Building your assistant via the console](https://github.com/snipsco/snips-platform-documentation/wiki#building-your-assistant-via-the-console)
    * [On-device processing](https://github.com/snipsco/snips-platform-documentation/wiki#building-your-assistant-via-the-console)

1. [Setup the Snips Voice Platform on your Raspberry Pi](https://github.com/snipsco/snips-platform-documentation/wiki/1.-Setup-the-Snips-Voice-Platform-on-your-Raspberry-Pi)
    * [Snips Platform Installation](https://github.com/snipsco/snips-platform-documentation/wiki/1.-Setup-the-Snips-Voice-Platform-on-your-Raspberry-Pi#1-snips-platform-installation)
    * [2. Snips Platform Configuration](https://github.com/snipsco/snips-platform-documentation/wiki/1.-Setup-the-Snips-Voice-Platform-on-your-Raspberry-Pi#2-snips-platform-configuration)

2. [Running your first end to end assistant](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant)
    * [Scope of the tutorial](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant#scope-of-the-tutorial)
    * [1. Downloading the assistant from the web console](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant#1-downloading-your-first-assistant-from-the-web-console)
    * [2. Running your assistant on device](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant#2-running-your-assistant-on-device)
    * [3. Responding to user requests with your own handler](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant#3-responding-to-user-requests-with-your-own-handler)
3. [Building an end to end assistant with a custom intent](https://github.com/snipsco/snips-platform-documentation/wiki/3.-Building-an-end-to-end-assistant-with-a-custom-intent)
    * [Scope of the tutorial](https://github.com/snipsco/snips-platform-documentation/wiki/3.-Building-an-end-to-end-assistant-with-a-custom-intent)
    * [Diving in](https://github.com/snipsco/snips-platform-documentation/wiki/3.-Building-an-end-to-end-assistant-with-a-custom-intent#diving-in)
4. [Create an assistant with custom slot types (advanced)](https://github.com/snipsco/snips-platform-documentation/wiki/4.-Create-an-assistant-with-custom-slot-types-(advanced))
5. [Learn more: Key Concepts](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts)
    1. [Hotword Detection](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#1-hotword-detection)
    2. [Automatic Speech Recognition](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#2-automatic-speech-recognition)
        * [Snips ASR  (English, with French and more coming soon)](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#snips-asr)
        * [Using Google’s Cloud service for other languages](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#using-googles-cloud-service-for-other-languages)
    3. [Natural Language Understanding](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#3-natural-language-understanding)
       * [Intent](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#intent)
       * [Slot](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#slot)
       * [Confirmation](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#confirmation)
       * [Bundle](https://github.com/snipsco/snips-platform-documentation/wiki/5.-Learn-more:-Key-Concepts#bundle)
6. [Learn more: Platform Commands](https://github.com/snipsco/snips-platform-documentation/wiki/6.--Learn-more:-Platform-Commands)
    * [Restarting the Snips Voice Platform](https://github.com/snipsco/snips-platform-documentation/wiki/6.--Learn-more:-Platform-Commands#restarting-the-snips-voice-platform)
    * [Platform and assistant updates](https://github.com/snipsco/snips-platform-documentation/wiki/6.--Learn-more:-Platform-Commands#platform-and-assistant-updates)
    * [Logger settings](https://github.com/snipsco/snips-platform-documentation/wiki/6.--Learn-more:-Platform-Commands#logger-settings)

7. [Learn more: built in resources](https://github.com/snipsco/snips-platform-documentation/wiki/7.-Learn-more:-built-in-resources)
    1. [Built-in slot types](https://github.com/snipsco/snips-platform-documentation/wiki/7.-Learn-more:-built-in-resources#1-built-in-slot-types)
    2. [Built-in intents](https://github.com/snipsco/snips-platform-documentation/wiki/7.-Learn-more:-built-in-resources#2-built-in-intents)
    3. [Bundles](https://github.com/snipsco/snips-platform-documentation/wiki/7.-Learn-more:-built-in-resources#3-bundles)

8. [Learn more : Snips Internals](https://github.com/snipsco/snips-platform-documentation/wiki/8.-Learn-more-:-Snips-Internals)

[Contact us](https://github.com/snipsco/snips-platform-documentation/wiki/Contact-us)

[FAQ](https://github.com/snipsco/snips-platform-documentation/wiki/FAQ)

### Handler code example for your Snips assistant

When you have built an assistant on <https://snips.ai> by following [the tutorial](https://github.com/snipsco/snips-platform-documentation/wiki), you can write your own handler to act on the behalf of the user.

This repository contains a few examples in various languages

#### Display messages

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


#### IoT Assistant

If you followed [this tutorial](https://github.com/snipsco/snips-platform-documentation/wiki/2.-Running-your-first-end-to-end-assistant) you have built an assistant to handle a few interesting intents for IoT.

The code is a handler which shows how to get the intents as they are sent by the natural language understanding, and do an action by sending a helpful message on the text-to-speech channel
