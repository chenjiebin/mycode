<?php

date_default_timezone_set("Asia/Shanghai");
$time = time();
echo getMonthStart($time), "<br />";
echo getMonthDay($time), "<br />";
echo getMonthEnd($time), "<br />";
var_dump(checkDateIsValid("2013-09-10")); //输出true
var_dump(checkDateIsValid("2013-09-ha")); //输出false
var_dump(checkDateIsValid("2012-02-29")); //输出true
var_dump(checkDateIsValid("2013-02-29")); //输出false

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

/**
 * 校验日期格式是否正确
 *
 * @param string $date 日期
 * @param string $formats 需要检验的格式数组
 * @return boolean
 */
function checkDateIsValid($date, $formats = array("Y-m-d", "Y/m/d")) {
    $unixTime = strtotime($date);
    if (!$unixTime) { //strtotime转换不对，日期格式显然不对。
        return false;
    }
    //校验日期的有效性，只要满足其中一个格式就OK
    foreach ($formats as $format) {
        if (date($format, $unixTime) == $date) {
            return true;
        }
    }
    return false;
}
