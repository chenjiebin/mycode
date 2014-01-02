<?php

/**
 * 当有未捕获的异常, 则控制流会流到这里
 */
class ErrorController extends \Youba\Controller_Abstract {

    public function init() {
        \Yaf\Dispatcher::getInstance()->disableView();
    }

    public function errorAction($exception) {
        switch ($exception->getCode()) {
            case 404://404
            case 515:
            case 516:
            case 517:
                //记录404到本地日志
                \Youba\Log::getInstance()->write($exception->getMessage());
                //输出404
                header(\Youba\Common::getHttpStatusCode(404));
                echo '404';
                exit();
                break;
            default :
//                header(\Youba\Common::getHttpStatusCode(500));
                \Youba\Log::getInstance()->write($exception->getMessage());
                if (ini_get('display_errors')) {
                    echo '<!DOCTYPE html>
<html lang="zh-CN">
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <title>页面出现错误</title>
        </head>
        <body>';

                    echo $exception->getMessage() . "\n<br /><br />";
                    echo str_replace("\n", "\n<br />", $exception->getTraceAsString());

                    echo '</body></html>';
                }
                break;
        }
    }

}