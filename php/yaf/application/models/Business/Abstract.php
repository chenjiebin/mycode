<?php

namespace Business;

/**
 * 业务层的抽象类
 */
abstract class AbstractModel {

    public function __clone() {
        trigger_error('Clone is not allow!', E_USER_ERROR);
    }

}