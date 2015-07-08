<?php

//将小数转换成二进制表示方式
//function floatToBin($f) {
//    
//}
//
//$r = "";
//$f = 0.07;
//for ($i = 0; $i < 27; $i++) {
//    $f = number_format($f * 2, 2);
//    echo $i, " ", $f;
//    if ($f > 1) {
//        $r .= "1";
//        $f = $f - 1;
//        echo " 1";
//    } else {
//        $r .= "0";
//        echo " 0";
//    }
//    echo "<br />";
//}
//
//echo $r;

$r = 0;
$s = "000100011110101110000101000";
for ($i = 0; $i < 27; $i++) {
    if ($s[$i] == 1) {
        $r = $r + pow(2, ($i + 1) * - 1);
    }
}
echo $r;exit();


//echo decbin(123);
