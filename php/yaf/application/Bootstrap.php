<?php

/**
 * 所有在Bootstrap类中, 以_init开头的方法, 都会被Yaf调用,
 * 这些方法, 都接受一个参数:\Yaf\Dispatcher $dispatcher
 * 调用的顺序, 和声明的顺序相同
 *
 */
class Bootstrap extends \Yaf\Bootstrap_Abstract {

    /**
     * Yaf的配置缓存
     *
     * @var \Yaf\Config\Ini
     */
    protected $_config = null;

    /**
     * 把配置存到注册表
     */
    public function _initConfig() {
        $config = \Yaf\Application::app()->getConfig();

        $this->_config = $config;
        \Yaf\Registry::set('config', $config);
    }

    /**
     * 注册本地类
     */
    public function _initRegisterLocalNamespace() {
        //申明本地类
        \Yaf\Loader::getInstance()->registerLocalNamespace(array('Zend', 'Our'));
    }

    /**
     * 连接 MySQL
     */
    public function _initMySQL() {
        $conf = $this->_config->get('resources.database.params');

        if (!$conf)
            return false;

        $dbAdapter = new \Zend\Db\Adapter\Adapter($conf->toArray());
        \Yaf\Registry::set('dbAdapter', $dbAdapter);
    }

    /**
     * 自定义路由规则
     * 
     * @param Yaf_Dispatcher $dispatcher
     */
    public function _initRoute(\Yaf\Dispatcher $dispatcher) {
        $router = \Yaf\Dispatcher::getInstance()->getRouter();

        $config = new \Yaf\Config\Ini(APPLICATION_PATH . '/conf/route.ini', 'common');
        if ($config->routes)
            $router->addConfig($config->routes);
    }

    /**
     * 使用自定义的View
     * 
     * @param \Yaf\Dispatcher $dispatcher
     */
    public function _initView(\Yaf\Dispatcher $dispatcher) {
        $dispatcher->setView(new \Our\View(self::getViewPath()));
    }

    /**
     * 获取视图目录
     * 
     * @return string
     */
    public static function getViewPath() {
        return APPLICATION_PATH . '/application/views';
    }

    /**
     * 获取自定义的网络地址
     * 
     * @param string $name
     * @return string 
     */
    public static function getConfigUrl($name) {
        return \Yaf\Registry::get('config')->get('config.url.' . $name);
    }

}
