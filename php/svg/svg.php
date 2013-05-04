<?php

$svgString = '<?xml version="1.0" encoding="utf-8"?>
<svg version="1.1" baseProfile="full"  width="232" height="232" viewBox="0 0 232 232" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:ev="http://www.w3.org/2001/xml-events">
<rect width="232" height="232" fill="#fff" cx="0" cy="0" />
<defs>
<rect id="p" width="8" height="8" />
</defs>
<g fill="#000">';

$im = imagecreatefrompng('liantu.png');
$x = imagesx($im);
$y = imagesy($im);

for ($i = 0; $i < $x; $i++) {
    for ($j = 0; $j < $y; $j++) {
        $rgb = imagecolorat($im, $i, $j);
        if ($rgb == 0) {
            $svgString .= '<use x="' . ($i * 8) . '" y="' . ($j * 8) . '" xlink:href="#p" />' . "\n";
        }
    }
}

$svgString .= '</g>
</svg>';

file_put_contents('liantu.svg', $svgString);
