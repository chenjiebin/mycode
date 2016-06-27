/**
 * 测试微信授权登录功能
 * 公众号必须是服务号
 * 参考地址: https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842&token=&lang=zh_CN
 */
'use strict';

var OAuth = require('wechat-oauth');
var client = new OAuth('wx5e6a0b67b76db0ac', 'aad7585526490642a03936ff0e45157a');

// 测试微信授权登录
var oauth = function (req, res, next) {
    var url = client.getAuthorizeURL('http://www.kuiyinapp.com/wechat/user/oauthcallback', 'state', 'snsapi_userinfo');
    console.log(url);
    // https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx5e6a0b67b76db0ac&redirect_uri=http%3A%2F%2Fwww.kuiyinapp.com%2Fwechat%2Fuser%2Foauthcallback&response_type=code&scope=snsapi_userinfo&state=state#wechat_redirect
    
    // res.send(url);
    res.redirect(url);
};

// 用户授权后回调
var oauthCallback = function (req, res, next) {
    console.log(req.url);
    console.log(req.query);
    // { code: '001Loasu0Zp29q1wBMtu0locsu0Loasn', state: 'state' }

    client.getAccessToken(req.query.code, function (err, result) {
        console.log(err, result);
        // null { data:
        // { access_token: 'LCUjXN-yeNIcBVSjaSPsvteeT3KPdqQzfrJgbJKYN-Ix2GFEW-4egD3vxgz_Re20uskEpG3brBP73yzFoyP87OiP7t6kCv-qjslp0xBZJwk',
        //     expires_in: 7200,
        //     refresh_token: 'FaueH5jeTluWfsZe_bFHwrUlB_BTmZV3GNUWoKL3tsD7eR8zxS7-ePfK0svmsH2ScQ5PjrHrwuisvQfNsGzVpVysJ7fB4WgTTg8l7q0pXtA',
        //     openid: 'omX4nt7SS_HAiumR7fHIEPB0dS8M',
        //     scope: 'snsapi_userinfo',
        //     create_at: 1467000238398 } }

        client.getUser(result.data.openid, function (err, result) {
            console.log(err, result);
            // null { openid: 'omX4nt7SS_HAiumR7fHIEPB0dS8M',
            //     nickname: '陈杰斌',
            //     sex: 1,
            //     language: 'zh_CN',
            //     city: 'Xiamen',
            //     province: 'Fujian',
            //     country: 'China',
            //     headimgurl: 'http://wx.qlogo.cn/mmopen/ajNVdqHZLLBia65RdX7Lx7faML4kO5Gt7r9iarsMB3kn76Sgj73qicbhNWAa3IVt7kc32qvviaV2H6G9CwQgZxNj6w/0',
            //     privilege: [] }
        });
    });

    res.send(req.url);
};


module.exports = function (router) {
    router.get('/oauth', oauth);
    router.get('/oauthcallback', oauthCallback);
};
