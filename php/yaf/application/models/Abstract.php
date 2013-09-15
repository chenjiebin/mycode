<?php

namespace Base\Model;

/**
 * 数据模型的抽象类, 普通的数据模型都要基于此类,
 */
abstract class AbstractModel {

    public function __construct($options = null) {
        if ($options)
            $this->setOptions($options);
    }

    /**
     *
     * @param string $name
     * @param mixed $value
     * @throws \Exception
     */
    public function __set($name, $value) {
        $method = 'set' . ucfirst($name);
        if (!method_exists($this, $method))
            throw new \Exception('Invalid model property');

        $this->$method($value);
    }

    /**
     *
     * @param string $name
     * @return mixed
     * @throws \Exception
     */
    public function __get($name) {

        $method = 'get' . ucfirst($name);
        if (!method_exists($this, $method))
            throw new \Exception('Invalid model property');

        return $this->$method();
    }

    /**
     * 通用设置方法
     *
     * @param array $options    参数. 如果是类, 必需实现了toArray(), 或者Traversabl接口的类.
     * @return \Base\Model\AbstractModel
     */
    public function setOptions($options) {
        if (is_object($options)) {
            if (method_exists($options, 'toArray')) {
                $options = $options->toArray();
            }
            else if (!($options instanceof \Traversable)) {
                return $this;
            }
        }
        else if (!is_array($options)) {
            return $this;
        }

        foreach ($options as $key => $value) {
            $method = 'set' . ucfirst($key);
            if (method_exists($this, $method))
                $this->$method($value);
        }

        return $this;
    }

    /**
     * Returna a array of model properties
     *
     * @return array
     */
    abstract public function toArray();
}