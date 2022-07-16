import base64

secret = "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
decodedByte = base64.b64decode(secret)
print(decodedByte.decode("utf-8").replace(":"," ")[::-1]) #Join us at LINE MAN Wongnai