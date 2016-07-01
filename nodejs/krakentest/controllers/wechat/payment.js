/**
 * 微信公众号支付功能
 * 公众号必须是服务号
 */
'use strict';
var fs = require("fs");
// 测试微信授权登录
// ; 微信公众号支付配置
// payment.weixin.pub.appId = 'wx5e6a0b67b76db0ac'
// payment.weixin.pub.appSecret = 'aad7585526490642a03936ff0e45157a'
// payment.weixin.pub.mchId = '1321639601'
// payment.weixin.pub.key = 'da34w4uhi32g45hjhr2g52j4hgj53453'
var send = function (req, res, next) {
    console.log(appRoot);
    var Payment = require('wechat-pay').Payment;
    var initConfig = {
        partnerKey: "da34w4uhi32g45hjhr2g52j4hgj53453",
        appId: "wx5e6a0b67b76db0ac",
        mchId: "1321639601",
        notifyUrl: "/notify",
        pfx: fs.readFileSync(appRoot + "/controllers/wechat/key/apiclient_cert.p12")
    };
    var payment = new Payment(initConfig);

    var order = {
        body: '吮指原味鸡 * 1',
        attach: '{"部位":"三角"}',
        out_trade_no: 'kfc' + (+new Date),
        total_fee: 10 * 100,
        spbill_create_ip: req.ip,
        openid: req.user.openid,
        trade_type: 'JSAPI'
    };

    payment.getBrandWCPayRequestParams(order, function (err, payargs) {
        res.json(payargs);
    });
};

// 支付后回调
var notify = function (req, res, next) {

};

module.exports = function (router) {
    router.get('/send', send);
    router.get('/notify', notify);
};
