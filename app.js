const http = require("http");

const requestListener = (req, res) => {
  res.writeHead(200);
  res.end("My first server!");
};

const server = http.createServer(requestListener);

server.addListener("connection", () => console.log("server made a connection"));
server.addListener("listening", () => console.log("server listening..."));

server.on("close", function () {
  console.log("server stopping ...");
});

process.on("SIGINT", function () {
  server.close();
});

server.listen(5000);
