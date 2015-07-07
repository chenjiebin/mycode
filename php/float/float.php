<?php

for ($i = 1; $i < 10; $i ++) {
    echo $i, " ";
    $f = ($i / 10);
    var_dump($f * 10 == $i);
    echo "<br />";
}
echo "<br />";

for ($i = 1; $i < 100; $i ++) {
    echo $i, " ";
    $f = ($i / 100);
    var_dump($f * 100 == $i);
    echo "<br />";
}


$f = 0.07;
var_dump($f * 100 == 7);
