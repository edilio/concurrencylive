/**
 * Created by edilio on 12/18/16.
 */

// Load the TCP Library
net = require('net');


function fib(n) {
    if ((n === 0) || (n === 1)) {
        return 1
    } else {
        return fib(n - 1) + fib(n - 2)
    }
}

// Start a TCP Server
net.createServer(function (socket) {

  // Identify this client
  socket.name = socket.remoteAddress + ":" + socket.remotePort;

  // Put this new client in the list
  // Send a nice welcome message and announce
  socket.write("Welcome " + socket.name + "\n");


  // Handle incoming messages from clients.
  socket.on('data', function (data) {
      var n = parseInt(data.toString());
      var result = fib(n);
      socket.write(result.toString() + '\n');
  });

  // Remove the client from the list when it leaves
  socket.on('end', function () {
    socket.write("good bye.\n");
  });



}).listen(25000);
