'use strict';
var URL = require('url');

module.exports = function (router) {
    router.get('/', function (req, res) {
        var testUrl = 'http://cdn.firstlinkapp.com/real/i_pic/20160310/20160310/45c92b42-6d1e-4e5e-bf5e-66ae3d35cfda.png';
        var p = URL.parse(testUrl);

        var tmpArray = p.path.split('/');
        tmpArray.pop();
        console.log(tmpArray.join('/'));
        console.log(p.href); //取到的值是：http://localhost:8888/select?aa=001&bb=002
        console.log(p.protocol); //取到的值是：http: 
        console.log(p.hostname);//取到的值是：locahost
        console.log(p.host);//取到的值是：localhost:8888
        console.log(p.port);//取到的值是：8888
        console.log(p.path);//取到的值是：/select?aa=001&bb=002
        console.log(p.hash);//取到的值是：null 
        console.log(p.query);// 取到的值是：aa=001
        res.end(p.path);
    });
};
