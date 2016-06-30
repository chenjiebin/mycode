/**
 * 微信生成二维码功能
 * 公众号必须是服务号
 * 参考地址: https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1443433542&token=&lang=zh_CN
 */
'use strict';
var WechatAPI = require('wechat-api');
// 测试微信授权登录
var create = function (req, res, next) {
    var api = new WechatAPI('wx5e6a0b67b76db0ac', 'aad7585526490642a03936ff0e45157a');
    api.createTmpQRCode(1000001, 30, function (err, result) {
        console.log(err, result);
        var imgurl = api.showQRCodeURL(result.ticket);
        res.end(imgurl);
    });
};

module.exports = function (router) {
    router.get('/create', create);
};
