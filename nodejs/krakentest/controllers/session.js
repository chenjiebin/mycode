'use strict';

module.exports = function (router) {
    router.get('/views', function (req, res) {
        console.log(req.session);
        console.log(req.session.id);
        var sess = req.session;
        if (sess.views) {
            sess.views++;
            res.setHeader('Content-Type', 'text/html');
            res.write('<p>views: ' + sess.views + '</p>');
            res.write('<p>expires in: ' + (sess.cookie.maxAge / 1000) + 's</p>');
            res.end();
        } else {
            sess.views = 1;
            res.end('welcome to the session demo. refresh!')
        }
    });
};
