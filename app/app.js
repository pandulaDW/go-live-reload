const http = require("http");
const { generateRandomNum } = require("./helper1");

const requestListener = (_, res) => {
  res.writeHead(200, {
    "Content-Type": "application/json",
  });
  const response = {
    num1: generateRandomNum(),
    message: "it works!!",
  };
  res.end(JSON.stringify(response));
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
