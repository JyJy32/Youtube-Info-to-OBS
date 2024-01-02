package main

var index = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>GoHTML Example</title>
        <style>
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    align-items: center;
    justify-content: center;
    height: 100vh;
}
    .container {
        display: grid;
        grid-template-columns: auto 1fr;
    }
    #image-container {
        grid-column: 1 / 2;
        margin-right: 20px;
        display: flex;
        justify-content: center;
        align-items: center;
        overflow: hidden;
    }

    #image-container img {
        flex-shrink: 0;
        min-width: 100%;
        min-height: 100%;
    }

    #text-container {
        grid-column: 2 / 3;
        text-align: left;
    }

    img {
        max-width: 100%;
        height: auto;
    }
        </style>
    
        <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    </head>
    <body>
        <div hx-get="/video" hx-trigger="every 1s" class="container"></div>
    </body>
</html>
`   

var video = `<div id="image-container">
    <img src="{{.Thumbnail}}" alt="Album Cover">
</div>
<div id="text-container">
    <h1>{{.Title}}</h1>
    <h2>{{.Channel}}</h2>
</div>`
