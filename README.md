Steps:

1. Setup account with Clarifai
2. Get your API key from Clarifai and replace {your Clarifai api key here} in the go file with the key
3. Place the go file in the chipper/plugins directory of your wirepod installation (for me it was /home/pi/wire-pod/chipper/plugins)
4. Open a terminal and cd to this directory
5. Run sudo /usr/local/go/bin/go build -buildmode=plugin identifier.go
6. At this point it may complain about missing libraries. Use the messages and the imports section at the beginning of the go file to get a handle on what needs to be installed.
7. Follow the directions to install the Vector Python SDK (https://github.com/cyb3rdog/vector-python-sdk/tree/master)
8. From a command prompt: python3 -m anki_vector.configure_pod -n {your vector name} -i {your vector ip address} -s {your vector serial number}
9. Copy the generated folder .anki_vector to the chipper directory of your wirepod installation (for me it was /home/pi/wire-pod/chipper)
10. Update the sdk_config.ini in this directory so that the path on the 'cert' line is the path of the cert file in this directory


You may need to restart wirepod, but assuming you got your key right, got it compiled, and have your cert installed correctly, you should be able to ask Vector 'What is this thing', and he will give you the top five results.

You can see it in action here:

https://youtu.be/kiu00_PvZfc