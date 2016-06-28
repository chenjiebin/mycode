'use strict';

module.exports = function (router) {

    // session功能测试
    router.get('/views', function (req, res) {
        console.log("session:", req.session);
        console.log("session id:", req.session.id);
        console.log("session cookie:", req.session.cookie);
        var sess = req.session;
        if (!sess.views) {
            sess.views = 1;
            return res.end('welcome to the session demo. refresh!');
        }
        sess.views++;
        res.setHeader('Content-Type', 'text/html');
        res.write('<p>views: ' + sess.views + '</p>');
        res.write('<p>expires in: ' + (sess.cookie.maxAge / 1000) + 's</p>');
        return res.end();
    });
};
