<?php

readFileFromDir(__DIR__);

var_dump(glob(__DIR__ . "/*"));

function readFileFromDir($dir) {
    if (!is_dir($dir)) {
        return false;
    }
    //打开目录
    $handle = opendir($dir);
    while (($file   = readdir($handle)) !== false) {
        //排除掉当前目录和上一个目录
        if ($file == "." || $file == "..") {
            continue;
        }
        $file = $dir . DIRECTORY_SEPARATOR . $file;
        //如果是文件就打印出来，否则递归调用
        if (is_file($file)) {
            print $file . '<br />';
        } elseif (is_dir($file)) {
            readFileFromDir($file);
        }
    }
}
