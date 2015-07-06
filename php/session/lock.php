<?php

ini_set("session.save_handler", "redis");
ini_set("session.save_path", "tcp://192.168.3.233:6380");

session_start();

$_SESSION['name']  = "tony";
$_SESSION['users'] = array("tony", "andy", "tom");

sleep(10);
