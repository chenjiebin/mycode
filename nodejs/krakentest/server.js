'use strict';

var app = require('./index');
var http = require('http');


var server;

/*
 * Create and start HTTP server.
 */

server = http.createServer(app);
server.listen(process.env.PORT || 8001);
// server.listen(process.env.PORT || 80); // 微信授权登录回调接口要是80端口
server.on('listening', function () {
    console.log('Server listening on http://localhost:%d', this.address().port);
});
