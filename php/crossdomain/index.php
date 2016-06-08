<?php

$data     = json_encode(array("id" => "1", "name" => "tom"));
$callback = $_GET["callback"];
echo $callback . "(" . $data . ")";
