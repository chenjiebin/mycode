<?php

namespace Our;

/**
 * memcache操作类
 */
class Memcache {

    private $_memcache = null;
    private $_config = null;

    private function __construct() {
        $this->_initConnection();
    }

    /**
     * 初始化连接
     */
    private function _initConnection() {
        if (!$this->_getConfig()) {
            return false;
        }
        $this->_memcache = new \Memcache();
        $this->_memcache->connect($this->_getConfig()->get('host'), $this->_getConfig()->get('port'));
    }

    /**
     * 获取memcache的配置
     * 
     * @return \Yaf\Config\Ini
     */
    private function _getConfig() {
        if (!$this->_config) {
            $this->_config = \Yaf\Registry::get('config')->get('memcache.params');
        }

        return $this->_config;
    }

    /**
     * 增加缓存
     * 
     * @param type $key
     * @param type $value
     * @param type $timeout
     */
    public function add($key, $value, $timeout = -1) {
        if (!$this->_getConfig()) {
            return false;
        }
        if (!$this->_getConfig()->get('caching')) {
            return false;
        }
        if ($this->_isUrlSetNoCache()) {
            return false;
        }

        if ($timeout == -1) {
            $timeout = $this->_getConfig()->get('timeout');
            if (!$timeout) {
                $timeout = 1800;
            }
        }

        return $this->_memcache->set($this->_getConfig()->get('prefix') . '_' . $key, $value, false, $timeout);
    }

    /**
     * 根据key值获取缓存数据
     * 
     * @param string $key
     * @return mixed
     */
    public function get($key) {
        if (!$this->_getConfig()) {
            return false;
        }
        if (!$this->_getConfig()->get('caching')) {
            return false;
        }
        if ($this->_isUrlSetNoReadCache()) {
            return false;
        }
        if ($this->_isUrlSetNoCache()) {
            return false;
        }

        return $this->_memcache->get($this->_getConfig()->get('prefix') . '_' . $key);
    }

    /**
     * 检测Url中是否指明不启用读缓存
     */
    protected function _isUrlSetNoReadCache() {
        if (stripos($_SERVER['REQUEST_URI'], 'NO_READ_CACHE') !== false) {
            return true;
        }
        return false;
    }

    /**
     * 检测Url中是否指明不启用缓存
     */
    protected function _isUrlSetNoCache() {
        if (stripos($_SERVER['REQUEST_URI'], 'NO_CACHE') !== false) {
            return true;
        }
        return false;
    }

    private static $_instance = null;

    /**
     * 获取实例
     * 
     * @return \Youba\Memcache
     */
    public static function getInstance() {
        if (is_null(self::$_instance)) {
            self::$_instance = new self();
        }
        return self::$_instance;
    }

}