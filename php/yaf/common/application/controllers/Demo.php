<?php

class DemoController extends \Yaf\Controller_Abstract {

    public function testdbAction() {
        $userMapper = new \Mapper\UserModel();
        $user = $userMapper->find(1);
        
        $this->getView()->assign("user", $user);
    }

}