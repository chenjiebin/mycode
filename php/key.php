<?php

// 生成随机字符串的key
$str    = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST";
$strLen = strlen($str);
$len    = isset($_GET['l']) ? $_GET['l'] : "32";

$key = "";
while (strlen($key) < $len) {
    $key .= $str[rand(0, $strLen - 1)];
}
echo $key;
