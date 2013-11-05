<?php

class UserModel extends \Base\Model\AbstractModel {

    protected $_id = null;
    protected $_name = null;

    public function toArray() {
        return array(
            'id' => $this->_id,
            'name' => $this->_name
        );
    }

    public function getId() {
        return $this->_id;
    }

    public function setId($_id) {
        $this->_id = $_id;
    }

    public function getName() {
        return $this->_name;
    }

    public function setName($_name) {
        $this->_name = $_name;
    }

}
