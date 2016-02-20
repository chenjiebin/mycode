'use strict';

module.exports = function (router) {
    // notice that my route is '/' but I respond to '/users'
    router.get('/', function (req, res) {
        res.send('you can find me at /users');
    });
};