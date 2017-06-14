var mqtt = require('mqtt');
var mqtt_client = mqtt.connect('mqtt://localhost:9898');

// Subscribe to the important messages
mqtt_client.on('connect', function () {
    mqtt_client.subscribe('#');
});

mqtt_client.on('message', function (topic, message) {
    console.log(' -  ' + topic + ': ' + message.toString());
});
