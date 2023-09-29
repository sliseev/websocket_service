var socket = new WebSocket("ws://localhost:8080/ws");

const Names = ["Sergey", "Ivan", "Vitaly", "Oleg", "Daniel", "Anna", "Nina", "Oxana", "Alex", "Bot"];

let connect = cb => {
  console.log("connecting...");

  socket.onopen = () => {
    console.log("Successfully Connected");
    var name = Names[Math.floor(Math.random() * Names.length)];
    socket.send(name);
  };

  socket.onmessage = msg => {
    console.log("Received message: (see the next)");
    console.log(msg)
    cb(msg)
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
