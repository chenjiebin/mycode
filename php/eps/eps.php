<?php

$epsString = '%!PS-Adobe EPSF-3.0
%%Creator: Zend_Matrixcode_Qrcode
%%Title: QRcode
%%CreationDate: 2013-05-04
%%DocumentData: Clean7Bit
%%LanguageLevel: 2
%%Pages: 1
%%BoundingBox: 0 0 264 264
8 8 scale
2 2 translate
/F { rectfill } def
0 0 0 setrgbcolor
';



$im = imagecreatefrompng('liantu.png');
$x = imagesx($im);
$y = imagesy($im);

for ($i = 0; $i < $x; $i++) {
    for ($j = 0; $j < $y; $j++) {
        $rgb = imagecolorat($im, $i, $j);
        if ($rgb == 0) {
            $epsString .= $i . ' ' . ($y - $j - 1) . " 1 1 F\n";
        }
    }
}

$epsString .= '%%EOF';
file_put_contents('liantu.eps', $epsString);