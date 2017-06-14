import paho.mqtt.client as mqtt
import json

# MQTT client to connect to the bus
mqtt_client = mqtt.Client()

# Subscribe to the important messages
def on_connect(client, userdata, flags, rc):
    mqtt_client.subscribe('hermes/intent/ActivateObject')
    mqtt_client.subscribe('hermes/intent/DeactivateObject')
    mqtt_client.subscribe('hermes/intent/MuteObject')
    mqtt_client.subscribe('hermes/intent/ActivateLightColor')

# Process a message as it arrives
def on_message(client, userdata, msg):
    if msg.topic == 'hermes/intent/ActivateObject':
        slots = parse_slots(msg)
        if 'objectType' not in slots:
            say('I don\'t know what object I should activate')
        else:
            object_name = slots['objectType']
            datetime = slots.get('startDateTime', 'now')
            say('Activating {} {}'.format(object_name, datetime))
    elif msg.topic == 'hermes/intent/DeactivateObject':
        slots = parse_slots(msg)
        if 'objectType' not in slots:
            say('I don\'t know what object I should deactivate')
        else:
            object_name = slots['objectType']
            datetime = slots.get('startDateTime', 'now')
            say('Deactivating {} {}'.format(object_name, datetime))
    elif msg.topic == 'hermes/intent/MuteObject':
        slots = parse_slots(msg)
        if 'deviceType' not in slots:
            say('I don\'t know what object I should mute')
        else:
            object_name = slots['deviceType']
            say('Mute {}'.format(object_name))
    elif msg.topic == 'hermes/intent/ActivateLightColor':
        slots = parse_slots(msg)
        color = slots.get('objectColor', 'white').lower()
        say('Turning the light {}'.format(color))

def parse_slots(msg):
    '''
    We extract the slots as a dict
    '''
    data = json.loads(msg.payload)
    return dict((slot['slotName'], slot['rawValue']) for slot in data['slots'])

def say(text):
    '''
    Print the output to the console and to the TTS engine
    '''
    print(text)
    mqtt_client.publish('hermes/tts/say', json.dumps({'text': text}))

if __name__ == '__main__':
    mqtt_client.on_connect = on_connect
    mqtt_client.on_message = on_message
    mqtt_client.connect('localhost', 9898)
    mqtt_client.loop_forever()
