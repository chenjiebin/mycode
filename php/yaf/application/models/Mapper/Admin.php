<?php

namespace Mapper;

/**
 * model mapper.
 *
 * 行为分析
 *
 * @package Mapper
 */
class AdminModel extends \Mapper\AbstractModel {

    protected $_tableName = 'admin';
    protected $_dbModelClass = '\AdminModel';
    private static $_instance = null;

    public static function getInstance() {
        if (!(self::$_instance instanceof self)) {
            self::$_instance = new self();
        }

        return self::$_instance;
    }

    public function __clone() {
        trigger_error('Clone is not allow!', E_USER_ERROR);
    }

}