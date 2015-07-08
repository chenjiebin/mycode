<?php

//这里主要要完成浮点数的误差演示
$f = 0.07;
echo (number_format($f * 100, 20));
//var_dump($f * 100 == 7);
//var_dump(bccomp($f * 100, 7, 3));


//for ($i = 1; $i < 10; $i ++) {
//    echo $i, " ";
//    $f = ($i / 10);
//    var_dump($f * 10 == $i);
//    echo "<br />";
//}
//echo "<br />";
//
//for ($i = 1; $i < 100; $i ++) {
//    echo $i, " ";
//    $f = ($i / 100);
//    var_dump($f * 100 == $i);
//    echo "<br />";
//}
//
//
//$f = 0.07;
//var_dump($f * 100 == 7);
