'use strict';

var test = function (req, res, next) {
    console.log("test");
    res.send('you can find me at /users/test');
}

module.exports = function (router) {
    // notice that my route is '/' but I respond to '/users'
    router.get('/', function (req, res) {
        res.send('you can find me at /users');
    });
    router.get('/test', test);
};
