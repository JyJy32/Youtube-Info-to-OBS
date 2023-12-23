chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    console.log(`PLUGIN: videoId: ${request.videoId}`);
    fetch(`http://localhost:6969/${request.videoId}`, {
        method: "POST"
    }).then(response => console.log(response.status))
    .catch(error => console.log(error));
});

