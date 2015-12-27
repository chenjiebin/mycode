<?php

$url = "http://localhost:9090/";

$params = json_encode(array(
    'users' => array(
        array('user_id' => '1', 'name' => 'tony'),
        array('user_id' => '2', 'name' => 'andy')
        )));

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_HTTPHEADER, array(
    'Content-Type: application/json',
    'Content-Length: ' . strlen($params)
));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_BINARYTRANSFER, true);
//curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $params);

$res = curl_exec($ch);
curl_close($ch);
var_dump($res);
