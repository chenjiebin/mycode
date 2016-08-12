<?php

header("Content-type: text/html; charset=utf-8");

$str = '<h2>Middlewares</h2>
 Middlewares allow you easily plugin/unplugin features for your Macaron applications. 
 There are already many  <a href="https://github.com/go-macaron">middlewares</a>  to simplify your work: 
<ul>
  <li> render - Go template engine </li>
  <li> static - Serves static files </li>
  <li> <a href="https://github.com/go-macaron/gzip">gzip</a>  - Gzip compression to all responses </li>
  <li> <a href="https://github.com/go-macaron/binding">binding</a>  - Request data binding and validation </li>
  <li> <a href="https://github.com/go-macaron/i18n">i18n</a>  - Internationalization and Localization </li>
  <li> <a href="https://github.com/go-macaron/cache">cache</a>  - Cache manager </li>
  <li> <a href="https://github.com/go-macaron/session">session</a>  - Session manager </li>
  <li> <a href="https://github.com/go-macaron/csrf">csrf</a>  - Generates and validates csrf tokens </li>
  <li> <a href="https://github.com/go-macaron/captcha">captcha</a>  - Captcha service </li>
  <li> <a href="https://github.com/go-macaron/pongo2">pongo2</a>  - Pongo2 template engine support </li>
  <li> <a href="https://github.com/go-macaron/sockets">sockets</a>  - WebSockets channels binding </li>
  <li> <a href="https://github.com/go-macaron/bindata">bindata</a>  - Embed binary data as static and template files </li>
  <li> <a href="https://github.com/go-macaron/toolbox">toolbox</a>  - Health check, pprof, profile and statistic services </li>
  <li> <a href="https://github.com/go-macaron/oauth2">oauth2</a>  - OAuth 2.0 backend </li>
  <li> <a href="https://github.com/go-macaron/switcher">switcher</a>  - Multiple-site support </li>
  <li> <a href="https://github.com/go-macaron/method">method</a>  - HTTP method override </li>
  <li> <a href="https://github.com/xyproto/permissions2">permissions2</a>  - Cookies, users and permissions 
  </li>
  <li> <a href="https://github.com/go-macaron/renders">renders</a>  - Beego-like render engine(Macaron has built-in template engine, this is another option) </li>
</ul>';


//preg_match_all('/href="([\w\W]+?)"/', $str, $matchs);
//var_dump($matchs);
//exit();
//$result = preg_replace('/href="([\w\W]+?)"/', 'href="http://www.01happy.com/redirect.php?u=' . rawurlencode('\1') . '"', $str);
$result = preg_replace_callback('/href="([\w\W]+?)"/', function($matches) {
    return 'href="http://www.01happy.com/redirect.php?u=' . rawurlencode($matches[1]) . '" target="_blank"';
}, $str);
var_dump($result);
exit();

