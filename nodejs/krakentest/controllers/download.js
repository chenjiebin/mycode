'use strict';
var URL = require('url');
var request = require('request');
var fs = require('fs');
var mkdirp = require('mkdirp');

// 尝试下载页面
module.exports = function (router) {
    router.get('/', function (req, res) {
        var testUrl = 'http://cdn.firstlinkapp.com/real/i_pic/20160310/20160310/45c92b42-6d1e-4e5e-bf5e-66ae3d35cfda.png';
        var p = URL.parse(testUrl);

        var tmpArray = p.path.split('/');
        tmpArray.pop();
        var saveDir = appRoot + '/upload' + tmpArray.join('/');
        mkdirp(saveDir, function (err) {
            var savePath = appRoot + '/upload' + p.path;
            var stream = request(testUrl).pipe(fs.createWriteStream(savePath));
            stream.on('finish', function () {
                res.end(savePath);
            });
        });
    });
};
