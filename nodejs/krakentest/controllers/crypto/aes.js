'use strict';

var aes128cbc = function (req, res, next) {
    var crypto = require('crypto');

    /**
     * 加密方法
     * @param cryptKey 加密key
     * @param iv       向量
     * @param data     需要加密的数据
     * @returns string
     */
    var encrypt = function (cryptKey, iv, data) {
        var encipher = crypto.createCipheriv('aes-128-cbc', cryptKey, iv);
        var encryptData = encipher.update(data, 'utf8', 'binary');
        encryptData += encipher.final('binary');
        encryptData = new Buffer(encryptData, 'binary').toString('base64');
        return encryptData;
    };

    /**
     * 解密方法
     * @param cryptkey 解密的key
     * @param iv       向量
     * @param encryptData   密文
     * @returns string
     */
    var decrypt = function (cryptkey, iv, encryptData) {
        encryptData = new Buffer(encryptData, 'base64').toString('binary');
        var decipher = crypto.createDecipheriv('aes-128-cbc', cryptkey, iv);
        var decoded = decipher.update(encryptData, 'binary', 'utf8');
        decoded += decipher.final('utf8');
        return decoded;
    };

    var cryptKey = '53eb59a44da52b9072601e3b3bff23eb';
    console.log('加密的key:', cryptKey);
    var iv = 'a2xhcgAAAAAAAAAA';
    var data = "Hello, nodejs. 演示aes-128-cbc加密和解密";
    console.log("需要加密的数据:", data);
    var encryptData = encrypt(cryptKey, iv, data);
    console.log("数据加密后:", encryptData);
    var dec = decrypt(cryptKey, iv, encryptData);
    console.log("数据解密后:", dec);


};

module.exports = function (router) {
    router.get('/aes128cbc', aes128cbc);
};
