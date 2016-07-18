'use strict';


module.exports = function (router) {
    // md5功能测试
    router.get('/', function (req, res) {
        var crypto = require('crypto');
        var text = "hello";
        var textmd5 = crypto.createHash('md5').update(text).digest('hex');
        var textmd5utf8 = crypto.createHash('md5').update(text, 'utf-8').digest('hex');
        console.log(textmd5);
        console.log(textmd5utf8);
        // output:
        // 5d41402abc4b2a76b9719d911017c592
        // 5d41402abc4b2a76b9719d911017c592

        var textcn = "hello 世界";
        var textcnmd5 = crypto.createHash('md5').update(textcn).digest('hex');
        var textcnmd5utf8 = crypto.createHash('md5').update(textcn, 'utf-8').digest('hex');
        console.log(textcnmd5);
        console.log(textcnmd5utf8);
        // output:
        // 99b8fa1c74ade005bfa1442b649647f4
        // 1aaa8e8010645fe4e3d44ad9745bb94e



        return res.end(textmd5);
    });
};
