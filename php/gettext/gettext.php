<?php

define('PACKAGE', 'hello');

putenv('LANG=zh_CN');
setlocale(LC_ALL, 'zh_CN');

bindtextdomain(PACKAGE, __DIR__ . '/languages');
textdomain(PACKAGE);

echo _("Hello World!");
