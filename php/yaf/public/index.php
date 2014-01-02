<?php

define("APPLICATION_PATH", realpath(dirname(__FILE__) . '/../')); //指向public的上一级
$app = new \Yaf\Application(APPLICATION_PATH . "/conf/application.ini");
$app->bootstrap()->run();
