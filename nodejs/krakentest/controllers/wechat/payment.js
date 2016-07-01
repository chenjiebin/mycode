/**
 * 微信公众号支付功能
 * 公众号必须是服务号
 */
'use strict';
var fs = require("fs");
// 测试微信授权登录
// 测试之前需要先调用/wechat/user/oauth进行授权登录获取openid

// ; 微信公众号支付配置
// payment.weixin.pub.appId = 'wx5e6a0b67b76db0ac'
// payment.weixin.pub.appSecret = 'aad7585526490642a03936ff0e45157a'
// payment.weixin.pub.mchId = '1321639601'
// payment.weixin.pub.key = 'da34w4uhi32g45hjhr2g52j4hgj53453'
var initConfig = {
    partnerKey: "da34w4uhi32g45hjhr2g52j4hgj53453",
    appId: "wx5e6a0b67b76db0ac",
    mchId: "1321639601",
    notifyUrl: "/notify",
    pfx: fs.readFileSync(appRoot + "/controllers/wechat/key/apiclient_cert.p12")
};

var send = function (req, res, next) {
    console.log(appRoot);
    var Payment = require('wechat-pay').Payment;
    var payment = new Payment(initConfig);
    var order = {
        body: '吮指原味鸡 * 1',
        attach: '{"部位":"三角"}',
        out_trade_no: 'kfc' + (+new Date),
        total_fee: 10 * 100,
        spbill_create_ip: "222.79.82.102",
        openid: 'omX4nt7SS_HAiumR7fHIEPB0dS8M',
        trade_type: 'JSAPI'
    };

    console.log(req.headers['user-agent']);

    payment.getBrandWCPayRequestParams(order, function (err, payargs) {
        console.log(err, payargs);
        console.log(JSON.stringify(payargs));

        var content = '' + "\n"
            + '<script type="text/javascript">' + "\n"
            + 'function onBridgeReady(){' + "\n"
            + 'WeixinJSBridge.invoke("getBrandWCPayRequest", ' + JSON.stringify(payargs) + ', function(res){' + "\n"
            + ' if (res.err_msg == "get_brand_wcpay_request:ok") {' + "\n"
            + ' alert("success");' + "\n"
            + '} else {' + "\n"
            + 'alert("fail");' + "\n"
            + '}' + "\n"
            + '});' + "\n"
            + '}' + "\n"
            + '' + "\n"
            + 'if (typeof WeixinJSBridge == "undefined"){' + "\n"
            + ' if( document.addEventListener ){' + "\n"
            + ' document.addEventListener("WeixinJSBridgeReady", onBridgeReady, false);' + "\n"
            + '  }else if (document.attachEvent){' + "\n"
            + 'document.attachEvent("WeixinJSBridgeReady", onBridgeReady); ' + "\n"
            + ' document.attachEvent("onWeixinJSBridgeReady", onBridgeReady);' + "\n"
            + '}' + "\n"
            + '}else{' + "\n"
            + '  onBridgeReady();' + "\n"
            + '}' + "\n"
            + '' + "\n"
            + '</script>' + "\n";

        res.end(content);

    });
};

// 支付后回调
var notify = function (message, req, res, next) {
    console.log(message);
    var openid = message.openid;
    var order_id = message.out_trade_no;
    var attach = {};
    try {
        attach = JSON.parse(message.attach);
    } catch (e) {

    }

    /**
     * 查询订单，在自己系统里把订单标为已处理
     * 如果订单之前已经处理过了直接返回成功
     */
    // res.reply('success');
    res.end('error');

    /**
     * 有错误返回错误，不然微信会在一段时间里以一定频次请求你
     * res.reply(new Error('...'))
     */
};

module.exports = function (router) {
    router.get('/send', send);

    var middleware = require('wechat-pay').middleware;
    router.get('/notify', middleware(initConfig).getNotify().done(notify));
};
