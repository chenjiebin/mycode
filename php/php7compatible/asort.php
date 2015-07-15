<?php

$array = array(0 => 0, 1 => 0);
asort($array);
//PHP5: array(1=>0, 0=>0); 
//PHP7: array(0=>1, 1=>0);