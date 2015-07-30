<?php

date_default_timezone_set("Asia/Shanghai");

echo getMicroSecond(true);

/**
 * 获取当前微妙数
 * 
 * @param bool $isWithLeadingZeros 是否需要补齐前导0
 * @return string
 */
function getMicroSecond($isWithLeadingZeros = false) {
    $microTime   = microtime();
    $a           = explode(" ", $microTime);
    $microSecond = $a[0] * 1000000;
    if ($isWithLeadingZeros) {
        while (strlen($microSecond) < 6) {
            $microSecond = "0" . $microSecond;
        }
    }
    return $microSecond;
}
