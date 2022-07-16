let secret = "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
let decodedString = atob(secret)
let arr = decodedString.split("")
decodedString = arr.reverse().join("")
console.log(decodedString.replaceAll(":"," ")) //Join us at LINE MAN Wongnai