const {HelloRequest, HelloReply} = require('./helloworld_pb.js');
const {GreeterClient} = require('./helloworld_grpc_web_pb.js');

var client = new GreeterClient('http://localhost:9000');

var request = new HelloRequest();
request.setName('luoji');

client.sayHello(request, {}, (err, response) => {
  console.log(response.getMessage());
});
