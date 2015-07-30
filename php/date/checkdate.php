<?php

var_dump(checkDateIsValid("2013-09-10")); //输出true
var_dump(checkDateIsValid("2013-09-ha")); //输出false
var_dump(checkDateIsValid("2012-02-29")); //输出true
var_dump(checkDateIsValid("2013-02-29")); //输出false
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
