<?php

session_start();

echo session_id();
echo "<br />";
var_dump(session_save_path());

$_SESSION['name']  = "tony";
$_SESSION['users'] = array("tony", "andy");

echo sys_get_temp_dir();
