# youtube info to obs

this is a tool to get the current youtube video playing in your chrome browser to obs

install the plugin manually by going to the extension managing page and enabling developer mode in the top right and then
clicking load unpacked on the top left

No compiled binary will be shipped.
go into the server folder and then
you need to either use

```
go run .
```

to run the apllication or

```
go build
```

to get an executable

you will also need an api key for the [youtube data api](https://developers.google.com/youtube/v3) and either put it as a env variable or use a .env file in the [server](/server) folder

then you can add a browser source to OBS with url http://localhost:6969/current_video

This is a very beta build atm and needs a lot of pollish
