# Display messages

This is a simple Python example showing how to display the messages received on the MQTT bus.

This allows you to see how the Snips engine works internally!


```
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
