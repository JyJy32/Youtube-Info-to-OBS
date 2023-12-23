console.log("PLUGIN: yt info to obs 0.1.3");


async function getVideoId() {
    while (true) {
        let videoId = window.location.search.split('v=')[1];
        if (videoId) {
            const ampersandPosition = videoId.indexOf('&');
            if (ampersandPosition !== -1) {
                videoId = videoId.substring(0, ampersandPosition);
            }
        }
        //post to localhost
        chrome.runtime.sendMessage({videoId: videoId}, function(response) {
            console.log(response);
        });
        //wait 1 second
        await new Promise(r => setTimeout(r, 1000));
    }
}

getVideoId();

