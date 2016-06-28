'use strict';

module.exports = function (router) {

    // mongoose功能测试
    router.get('/', function (req, res) {
        var mongoose = require('mongoose');
        mongoose.connect('mongodb://192.168.1.119/test');

        var Cat = mongoose.model('Cat', {name: String});

        var kitty = new Cat({name: 'Zildjian'});
        kitty.save(function (err) {
            if (err) {
                res.send(err);
            } else {
                res.send('meow');
            }
        });
    });
};
