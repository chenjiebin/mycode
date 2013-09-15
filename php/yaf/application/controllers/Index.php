<?php

class IndexController extends \Yaf\Controller_Abstract {

    public function indexAction() {//默认Action
        $this->getView()->assign("content", "Hello World");
    }
    
    public $config = null;
    
    public function testdbAction() {
        
         $config = \Yaf\Application::app()->getConfig();

        $this->config = $config;
        
        $conf = $this->config->get('resources.database.params');

        if (!$conf)
            return false;

        $dbAdapter = new \Zend\Db\Adapter\Adapter($conf->toArray());
        
        \Yaf\Registry::set('dbAdapter', $dbAdapter);
        
        $userMapper  = new \Mapper\UserModel();
        $model = $userMapper->find(1);
        
        
        $this->getView()->assign("content", "test db");
    }

}
