<?php

date_default_timezone_set("Asia/Shanghai");

echo getNextMonth(1448899200);

function getNextMonth($time) {
    $m = date('n', $time);
    if ($m < 12) {
        $nextM = $m + 1;
        $year  = date("Y", $time);
        return strtotime("$year-$nextM-01");
    }
    $nextM = 1;
    $year  = date("Y", $time) + 1;
    return strtotime("$year-$nextM-01");
}
