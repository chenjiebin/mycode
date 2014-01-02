<?php

class IndexController extends \Our\Controller_Abstract {

    public function init() {
        \Yaf\Dispatcher::getInstance()->disableView();
    }

    public function indexAction() {
        $this->getView()->display('index/index.phtml');
    }

}
