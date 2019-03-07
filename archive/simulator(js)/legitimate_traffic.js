


// # this file will be used to have a notion of legitimate traffic
// # how much legtimate traffic is received at any given time
// # what is the latency observed and all those things will be modeled here
const packet = require("./packet")
const network = require("./network")

const EventEmitter = require('events');

class flowGen extends EventEmitter {}

const flowGen = new flowGen();

let i = 0

flowGen.on('send', (pkt,n) => {
  console.log('packet received to send on network!');
  network.sendtoNetwork(pkt);
  i++;
  if(i < n) {
  	self.emit("send")
  }
});


flowGen.emit('send',pkt,n);

// def FlowGenSimple(pkt_no):
//   for i in range(1,pkt_no):
//     sendtoNetwork(pkt)
  


