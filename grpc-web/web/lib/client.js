const {HelloRequest, HelloReply} = require('./helloworld_pb.js');
const {GreeterClient} = require('./helloworld_grpc_web_pb.js');

var client = new GreeterClient('http://localhost:9000');

var btn = document.getElementById('btn');
btn.onclick = function() {
    var msgDom = document.getElementById('msg');
    var msg = msgDom.value;
    if (!msg) {
        return;
    }

    var request = new HelloRequest();
    request.setName(msg);

    client.sayHello(request, {}, (err, response) => {
        var replyMsg = response.getMessage()
        var divDom = document.getElementById('show');
        divDom.innerHTML = replyMsg;
    });
}

