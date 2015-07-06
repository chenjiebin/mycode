<?php

ini_set("session.save_handler", "redis");
ini_set("session.save_path", "tcp://192.168.3.233:6380");

session_start();

echo session_id();
echo "<br />";
var_dump(session_save_path());

$_SESSION['name']  = "tony";
$_SESSION['users'] = array("tony", "andy", "anny");


echo sys_get_temp_dir();
