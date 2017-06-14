var mqtt = require('mqtt');
var mqtt_client = mqtt.connect('mqtt://localhost:9898');

// Subscribe to the important messages
mqtt_client.on('connect', function () {
    mqtt_client.subscribe('hermes/intent/ActivateObject');
    mqtt_client.subscribe('hermes/intent/DeactivateObject');
    mqtt_client.subscribe('hermes/intent/MuteObject');
    mqtt_client.subscribe('hermes/intent/ActivateLightColor');
});

mqtt_client.on('message', function (topic, message) {
    if (topic == 'hermes/intent/ActivateObject') {
        var slots = parseSlots(message.toString());
        if (!('objectType' in slots)) {
            say('I don\'t know what object I should activate');
        } else {
            var objectName = slots['objectType'];
            var datetime = slots['startDateTime'] || 'now';
            say('Activating ' + objectName + ' ' + datetime);
        }
    } else if (topic == 'hermes/intent/DeactivateObject') {
        var slots = parseSlots(message.toString());
        if (!('objectType' in slots)) {
            say('I don\'t know what object I should deactivate');
        } else {
            var objectName = slots['objectType'];
            var datetime = slots['startDateTime'] || 'now';
            say('Deactivating ' + objectName + ' ' + datetime);
        }
    } else if (topic == 'hermes/intent/MuteObject') {
        var slots = parseSlots(message.toString());
        if (!('deviceType' in slots)) {
            say('I don\'t know what object I should mute');
        } else {
            var objectName = slots['objectType'];
            say('Mute ' + objectName);
        }
    } else if (topic == 'hermes/intent/ActivateLightColor') {
        var slots = parseSlots(message.toString());
        var color = (slots['objectColor'] || 'white').toLowerCase();
        say('Turning the light ' + color);
    }
});

function parseSlots(message) {
    var data = JSON.parse(message);
    return data.slots.reduce((acc, {slotName, rawValue}) => {
      acc[slotName] =  rawValue
      return acc
    }, {});
}

function say(text) {
    console.log(text);
    mqtt_client.publish('hermes/tts/say', JSON.stringify({'text': text}));
}
