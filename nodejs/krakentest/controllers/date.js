'use strict';

module.exports = function (router) {
    router.get('/', function (req, res) {
        var d = Date.parse("2016-07-01 15:18:11");

        var today = new Date("2016-07-01 15:18:11");
        res.write(today.toString());

        res.end(d.toString());
    });
};
