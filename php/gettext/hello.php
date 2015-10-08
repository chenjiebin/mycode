<?php

//设置语言文件的存放目录
$dir    = __DIR__ . '/locale';
//文件的名称，如该例为hello.mo
$domain = 'hello';
//语言类型
$locale = 'zh_CN';


putenv('LANG=' . $locale);
setlocale(LC_MESSAGES, $locale);

bindtextdomain($domain, $dir);
bind_textdomain_codeset($domain, "UTF-8");
textdomain($domain);

//_是gettext函数的简写
echo _("Hello World!");
