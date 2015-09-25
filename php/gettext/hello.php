<?php

$dir    = __DIR__ . '/locale';
$domain = 'hello';
$locale = 'zh_CN';

putenv('LANG=' . $locale);
setlocale(LC_MESSAGES, $locale);

bindtextdomain($domain, $dir);
bind_textdomain_codeset($domain, "UTF-8");
textdomain($domain);

echo _("Hello World!");
