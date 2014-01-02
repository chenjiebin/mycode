<?php

namespace Our;

/**
 * 控制器抽象类
 */
abstract class Controller_Abstract extends \Yaf\Controller_Abstract {

    /**
     * 校验各个参数
     * 
     * @param array $params
     * @param boolean $isRedirect
     * @param string $redirectUrl
     */
    protected function _validParams($params, $isRedirect = true, $redirectUrl = '/') {
        $isValid = true;
        foreach ($params as $param) {
            switch ($param['type']) {
                case 1://判断是否为正整数
                    if (!\Youba\Common::isPositiveIntNumber($param['value'])) {
                        $isValid = false;
                    }
                    break;
                case 2://判断是否为正整数或者是0
                    if (\Youba\Common::isNonnegativeIntNumber($param['value']) === false) {
                        $isValid = false;
                    }
                    break;
                case 3://不为空字符串
                    if (!(mb_strlen($param['value']) > 0)) {
                        $isValid = false;
                    }
                    break;
                case 4://定值范围
                    if (!in_array($param['value'], $param['set'])) {
                        $isValid = false;
                    }
                    break;
                default:
                    break;
            }
        }

        if (!$isValid && $isRedirect) {
            $this->redirect($redirectUrl);
        }

        return $isValid;
    }

}