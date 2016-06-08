<?php

header('Access-Control-Allow-Origin:*');
$data = json_encode(array("id" => "1", "name" => "tom"));
echo $data;
