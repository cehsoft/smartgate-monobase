var mediaSource = new MediaSource();
var sourceBuffer;
var queue = [];
var ws;

window.onload = function () {
  function pushPacket(data) {
    console.log(videoElement.error);
    queue.push(new Uint8Array(data));
    loadPacket();
  }

  function loadPacket() {
    // get package data from queue FIFO
    const data = queue.shift();

    if (!data || sourceBuffer.updating) {
      // sourceBuffer is busy or no packet data to display
      return; // nothings to load
    }

    // add to sourceBuffer for displaying
    sourceBuffer.appendBuffer(data);
  }

  function handleSourceOpen() {
    ws = new WebSocket("ws://127.0.0.1:3000/ws");
    // ws = new WebSocket("ws://127.0.0.1:8083/ws/live?suuid=H264_AAC");
    ws.binaryType = "arraybuffer";

    ws.onopen = function () {
      // initial sourceBuffer for format video/mp4; codecs="avc1.42C01E"
      sourceBuffer = mediaSource.addSourceBuffer(
        'video/mp4; codecs="avc1.42C01E"'
      );
      sourceBuffer.mode = "segments";

      sourceBuffer.addEventListener("updateend", () => {
        // load packet data when sourceBuffer idle
        loadPacket();
      });
    };

    ws.onmessage = function (event) {
      pushPacket(event.data);
    };
  }

  var videoElement = document.querySelector("video");
  mediaSource.addEventListener("sourceopen", handleSourceOpen, false);
  videoElement.src = window.URL.createObjectURL(mediaSource);
};
