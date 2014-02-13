<?php

//签名
$privateKeyFile = './id_rsa';
$keyContent = file_get_contents($privateKeyFile);  
$prikeyid = openssl_get_privatekey($keyContent);  

if (openssl_sign('hello', $out, $prikeyid))
$sign = base64_encode($out);

echo 'sign:'.$sign.'<br />';


//验证签名
$publicKeyFile = './id_rsa.pub';
$key_content = file_get_contents($publicKeyFile);  
$pubkeyid    = openssl_get_publickey($key_content);  


$sign = base64_decode($sign);
var_dump(openssl_verify('hello', $sign, $pubkeyid));