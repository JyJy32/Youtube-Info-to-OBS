document.addEventListener("DOMContentLoaded", function() {
    checkServerStatus();
    fetchVideo();
})
function checkServerStatus() {
    fetch("http://localhost:6969/status")
        .then(response => {
            document.getElementById("server_status").innerHTML = "offline"
            document.getElementById("server_status").style.color = "red"
            if (response.status === 200) {
                document.getElementById("server_status").innerHTML = "online"
                document.getElementById("server_status").style.color = "green"
            }
        })
        .catch(error => {
            document.getElementById("server_status").innerHTML = "offline" 
            document.getElementById("server_status").style.color = "red"
        });
}

function fetchVideo() {
    fetch("http://localhost:6969/video")
        .then(response => response.text())
        .then(html => {
            document.getElementById("video_content").innerHTML = html
        })
        .catch(error => {
            document.getElementById("video_content").innerHTML = "Error loading video: " + error
        });
}
