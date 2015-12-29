<?php

//Binary safe string comparison
//Returns < 0 if str1 is less than str2; > 0 if str1 is greater than str2, and 0 if they are equal.

$var1 = "Hello";
$var2 = "hello";

var_dump(strcmp($var1, $var2));
