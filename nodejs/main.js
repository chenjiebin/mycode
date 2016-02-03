var http = require('http');

http.createServer(function(request, response) {
    console.log(request.method);
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