'use strict';

module.exports = function (router) {
    router.get('/', function (req, res) {
        var t1 = new Date();
        res.write(t1.toString(), "\n");
        t1.setDate(t1.getDate() + 5);
        res.write(t1.toString(), "\n");

        var str = '20160701151811';
        res.write(str.substr(0, 4) + "-"
            + str.substr(4, 2) + "-"
            + str.substr(6, 2) + " "
            + str.substr(8, 2) + ":"
            + str.substr(10, 2) + ":"
            + str.substr(12, 2) + "");

        res.write('\n');
        var d = Date.parse("2016-07-01 15:18:11");

        var today = new Date("2016-07-01 15:18:11");
        res.write(today.toString());

        res.end(d.toString());


    });
};
