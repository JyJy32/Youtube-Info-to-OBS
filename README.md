# youtube info to obs 
this is a tool to get the current youtube video playing in your chrome browser to obs

install the plugin manually by going to the extension managing page and enabling developer mode in the top right and then 
clicking load unpacked on the top left

No compiled binary will be shipped so you need to either use 
```
go run .
```
to run the apllication or 
```
go build
```
to get an executable

then you can add a browser source to OBS with url http://localhost:6969/video
