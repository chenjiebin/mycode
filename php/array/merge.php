<?php

//合并二位数组的重复项
//来源：http://www.oschina.net/question/101542_2135921

$arr = array(
    array('id' => '1', 'uname' => 'b', 'password' => 'c'),
    array('id' => '2', 'uname' => 'b', 'password' => 'c'),
    array('id' => '3', 'uname' => 'd', 'password' => 'c')
);

$result = array();
foreach ($arr as $row) {
    $key = md5(json_encode(array($row['uname'], $row['password'])));
    if (array_key_exists($key, $result)) {
        $result[$key]['id'] = $result[$key]['id'] . ',' . $row['id'];
        continue;
    }
    $result[$key] = $row;
}
var_dump($result);

