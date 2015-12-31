<?php

function my_range($start, $end, $step = 1) {
    for ($i = $start; $i <= $end; $i += $step) {
        yield $i;
    }
}

foreach (my_range(1, 1000) as $num) {
    echo $num, "\n";
}