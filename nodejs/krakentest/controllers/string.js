'use strict';

module.exports = function (router) {

    // session功能测试
    router.get('/split', function (req, res) {
    	console.log('2014'.split('-'));
       	console.log('course-2014'.split('-'));
        return res.end('split');
    });
};
