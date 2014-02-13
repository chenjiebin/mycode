<?php
/** 
 * 公钥加密 
 * 
 * @param string 明文 
 * @param string 证书文件（.crt） 
 * @return string 密文（base64编码） 
 */  
function publickey_encodeing($sourcestr, $fileName)  
{  
    $key_content = file_get_contents($fileName);  
    $pubkeyid    = openssl_get_publickey($key_content);  
      
    if (openssl_public_encrypt($sourcestr, $crypttext, $pubkeyid))  
    {  
        return base64_encode("".$crypttext);  
    }  
}

/** 
 * 私钥解密 
 * 
 * @param string 密文（二进制格式且base64编码） 
 * @param string 密钥文件（.pem / .key） 
 * @param string 密文是否来源于JS的RSA加密 
 * @return string 明文 
 */  
function privatekey_decodeing($crypttext, $fileName, $fromjs = FALSE)  
{  
    $key_content = file_get_contents($fileName);  
    $prikeyid    = openssl_get_privatekey($key_content);  
    $crypttext   = base64_decode($crypttext);  
    $padding = $fromjs ? OPENSSL_NO_PADDING : OPENSSL_PKCS1_PADDING;  
    if (openssl_private_decrypt($crypttext, $sourcestr, $prikeyid, $padding))  
    {  
        return $fromjs ? rtrim(strrev($sourcestr), "\0") : "".$sourcestr;  
    }  
    return ;  
}








/*

$txt_en = '98bb72fd087897cfca88b2d82876e34e';  
$txt_en = base64_encode(pack("H*", $txt_en));  
$file = './server.key';  
$txt_de = privatekey_decodeing($txt_en, $file, TRUE);
echo $txt_de;  

*/








//PHP->PHP 测试  
//$data = "汉字:1a2b3c";  

//$data='123456';
//$file1 = './id_rsa.pub';  
$file2 = './id_rsa';
//$a = publickey_encodeing($data, $file1);  

$b = privatekey_decodeing('Aj6QjtPeWGNEYqP07qK7XFZan5KpBsbBE0UcTgCuBMYLFaXp2XpnCl1Mn8Cwdi95CVDoZA2X4314lw3GwVO7UsRP06+3cRcumvk+78PaNEP8H7mY1nJhd8A+JcXhwhJeEDVBzsr6lttkNwm2PnBz+Ic/XxNeA/qiXM61MxBnnATvYU6NDS9lU4dv0knN/voE60bCXj54pYtqbdsZJrk47iPCW0gDzi3S64jPZjkVF0m+eaZEFzrVbqCZB82YsDD8+NJSULrsuKkeo66RsIBUvxUYkm0R6hSjCxD3GIHbRMaOVtZk3EqqQ6+Ig8qLq8SwWPwD8EUEK2OCtHwm1fOcVw==', $file2);  
var_dump($b);  

