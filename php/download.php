<?php
$fileName = 'qr.png';
$content = '图片内容，从图片文件读取出来的内容';
$fileSize = strlen($content);

header("Content-type: application/txt");
header("Content-Disposition: attachment;filename=$fileName");
header("Content-Length: " . $fileSize);
exit($content);
