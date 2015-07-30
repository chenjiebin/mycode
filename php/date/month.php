<?php

date_default_timezone_set("Asia/Shanghai");
$time = time();
echo getMonthStart($time), "<br />";
echo getMonthDay($time), "<br />";
echo getMonthEnd($time), "<br />";


/**
 * 月初时间戳计算
 * 
 * @param int $time 时间戳
 * @return int 月初时间戳
 */
function getMonthStart($time) {
    return strtotime(date("Y-m-01 00:00:00", $time));
}

/**
 * 计算每月有几天
 * 
 * @param int $time 时间戳
 * @return int 
 */
function getMonthDay($time) {
    return date("t", $time);
    return cal_days_in_month(CAL_GREGORIAN, date("m", $time), date("Y", $time));
}

$time = time();

/**
 * 计算每月有几天
 * 
 * @param int $time 时间戳
 * @return int 
 */
function getMonthEnd($time) {
    $monthStart  = strtotime(date("Y-m-01 00:00:00", $time));
    $monthSecond = 86400 * cal_days_in_month(CAL_GREGORIAN, date("m", $time), date("Y", $time));
    return $monthStart + $monthSecond - 1;
}

