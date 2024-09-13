function play(event) {
    let player = document.getElementById("music");
    player.src = event.target.getAttribute("pod-src");
    player.play();
    updateTimeline();
    document.getElementById("pButton").style.display = "none";
    document.getElementById("pauseButton").style.display = "inline";
}

function pause() {
    let player = document.getElementById("music");
    if (player.paused) {
        player.play()
        return;
    }
    player.pause();
}

function updateTimeline() {
    const player = document.getElementById("music");
    const timeline = document.getElementById("timeline");
    const playhead = document.getElementById("playhead");

    player.addEventListener('timeupdate', () => {
        const percent = (player.currentTime / player.duration) * 100;
        playhead.style.left = `calc(${percent}% - ${playhead.offsetWidth / 2}px)`;
    });

    timeline.addEventListener('click', (event) => {
        const rect = timeline.getBoundingClientRect();
        const clickX = event.clientX - rect.left;
        const percent = (clickX / timeline.offsetWidth) * 100;
        playhead.style.left = `calc(${percent}% - ${playhead.offsetWidth / 2}px)`;
        player.currentTime = (percent / 100) * player.duration;
    });
}

document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll('[pod-src]').forEach(button => {
        button.addEventListener('click', play);
    });

    document.getElementById("pauseButton").addEventListener('click', pause);
});