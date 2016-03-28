<?php

$str = '我是明文我是明文我是明文我是明文我是明文dsfdghgasdfasdddddddd';
$key = 'aSGJLGYEWERWRREW4567i8o';

$str1 = AES::encode($key, $str);
$str2 = AES::decode($key, $str1);


echo '<xmp>';
var_dump($str);
var_dump($str1);
var_dump($str2);
var_dump(rtrim($str2));

/**
 * 利用mcrypt做AES加密解密
 * @author ts24<tsxw24@gmail.com>
 */
abstract class AES {

    /**
     * 算法,另外还有192和256两种长度
     */
    const CIPHER = MCRYPT_RIJNDAEL_128;

    /**
     * 模式 
     */
    const MODE = MCRYPT_MODE_ECB;

    /**
     * 加密
     * @param string $key	密钥
     * @param string $str	需加密的字符串
     * @return type 
     */
    static public function encode($key, $str) {
        $iv = mcrypt_create_iv(mcrypt_get_iv_size(self::CIPHER, self::MODE), MCRYPT_RAND);
        return mcrypt_encrypt(self::CIPHER, $key, $str, self::MODE, $iv);
    }

    /**
     * 解密
     * @param type $key
     * @param type $str
     * @return type 
     */
    static public function decode($key, $str) {
        $iv = mcrypt_create_iv(mcrypt_get_iv_size(self::CIPHER, self::MODE), MCRYPT_RAND);
        return mcrypt_decrypt(self::CIPHER, $key, $str, self::MODE, $iv);
    }

}
