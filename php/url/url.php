<?php

// 演示一些url的操作

$newUrl = addQueryParams("http://120.55.205.125:32000/createwatermark", "status=1");
var_dump($newUrl);
exit();

function addQueryParams($url, $params) {
    $result = parse_url($url);
    // 如果有get参数
    if (isset($result["query"])) {
        $result["new_query"] = $result["query"] . "&" . $params;
        return str_replace("?" . $result["query"], "?" . $result["new_query"], $url);
    }

    $result["new_query"] = $result["query"] . "&" . $params;
    return str_replace("?" . $result["query"], "?" . $result["new_query"], $url);
}
