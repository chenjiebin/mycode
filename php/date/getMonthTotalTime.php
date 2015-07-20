<?php

date_default_timezone_set("Asia/Shanghai");

echo getMonthTotalTime(1425139200);

//根据时间戳获取当月有多少秒

function getMonthTotalTime($time) {
    $year  = date("Y", $time);
    $month = date("n", $time);
    $day   = 0;
    switch ($month) {
        case 4 :
        case 6 :
        case 9 :
        case 11 :
            $days = 30;
            break;
        case 2 :
            if ($year % 4 == 0) {
                if ($year % 100 == 0) {
                    $days = $year % 400 == 0 ? 29 : 28;
                } else {
                    $days = 29;
                }
            } else {
                $days = 28;
            }
            break;

        default :
            $days = 31;
            break;
    }
    return $days;
}

