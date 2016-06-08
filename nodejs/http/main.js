var http = require('http');

http.createServer(function(request, response) {
    console.log(request.method);
    console.log(2.1 / 3 == 0.7);
    console.log(parseFloat(2.1 / 3).toFixed(2) == parseFloat(0.7).toFixed(2));
    if (request.method === 'GET' && request.url === '/echo') {
        var body = [];
        request.on('data', function(chunk) {
            body.push(chunk);
        }).on('end', function() {
            body = Buffer.concat(body).toString();
            response.end(body);
        })
    } else {
        response.statusCode = 404;
        response.end();
    }
}).listen(8080);