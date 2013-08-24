<?php
echo pack('C3', 80, 72, 80);

echo PHP_EOL;


print_r(unpack("C*", "PHP"));
