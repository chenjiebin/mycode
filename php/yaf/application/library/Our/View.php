<?php

namespace Our;

class View extends \Yaf\View\Simple {

    /**
     * 获取自定义的网络地址
     * 
     * @param string $name
     * @return string 
     */
    public function getConfigUrl($name) {
        return \Bootstrap::getConfigUrl($name);
    }

    /**
     * 输出自定义的网络地址
     * 
     * @param string $name
     * @return string 
     */
    public function echoConfigUrl($name) {
        echo \Bootstrap::getConfigUrl($name);
    }

}