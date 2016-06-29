'use strict';

var url = require('url');

// 分析当前访问的host
var getHost = function (req, res, next) {
    // 查看req对象
    console.log(req);
    console.log(req.protocol);
    console.log(req.headers.host);
    console.log(req.originalUrl);
    res.end(req.originalUrl);
};

module.exports = function (router) {
    var IndexModel = require('../models/index');
    var model = new IndexModel();
    router.get('/', function (req, res) {
        res.render('index', model);
    });
    router.get('/gethost', getHost);
};
