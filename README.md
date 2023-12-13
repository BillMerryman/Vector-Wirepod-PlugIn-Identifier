Steps:

Setup account with Clarifai
Get your API key from Clarifai and replace {your Clarifai api key here} in the go file with the key
Place the go file in the chipper/plugins of your wirepod installation (for me it was /home/pi/wire-pod/chipper/plugins)
Open a terminal and cd to this directory
Run sudo /usr/local/go/bin/go build -buildmode=plugin identifier.go
At this point it may complain about missing libraries. Use the messages and the imports section at the beginning of the go file to get a handle on what needs to be installed.

Once compiled, you may need to restart wirepod, but assuming you got your key right and got it compiled, you should be able to ask Vector 'What is this thing', and he will give you the top five results.
