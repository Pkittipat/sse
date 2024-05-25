console.log("start");
// when close the window or tab the source will be closed.
const source = new EventSource("/sse");
source.onmessage = function (event) {
  const data = JSON.parse(event.data);
  document.getElementById("result").innerHTML += data.time + "<br>";
};
