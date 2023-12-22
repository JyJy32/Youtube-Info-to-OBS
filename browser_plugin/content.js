console.log("PLUGIN: yt info to obs");
async function getVideoId() {
    while (true) {
        const title = document.getElementById('title');
        console.log(`PLUGIN: ${title.innerHTML}`);

        await new Promise(r => setTimeout(r, 1000));
    }
}

getVideoId();
