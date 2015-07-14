<?php

//将小数0.07转化成二进制
$bin  = "";
$int  = 7;
$base = 100;
echo "<table border='1'>";
echo "<td width='50'>位数</td>";
echo "<td width='50'>x2</td>";
echo "<td width='50'>位值</td>";
for ($i = 0; $i <= 60; $i++) {
    echo "<tr>";
    echo "<td>$i</td>";
    $int = $int * 2;
    echo "<td>$int</td>";
    if ($int == 100) {
        $bin.="1";
        echo "<td>1</td>";
        break;
    }
    if ($int > 100) {
        $bin.="1";
        $int = $int - $base;
        echo "<td>1</td>";
    } else {
        $bin .= "0";
        echo "<td>0</td>";
    }
    echo "</td>";
    echo "</tr>";
}
echo "</table>";
echo $bin;


