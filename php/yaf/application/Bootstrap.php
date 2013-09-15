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
    protected $config = null;

    /**
     * 把配置存到注册表
     */
    public function _initConfig() {
        $config = \Yaf\Application::app()->getConfig();

        $this->config = $config;
        \Yaf\Registry::set('config', $config);
    }

    /**
     * 设置默认值
     *
     * @param \Yaf\Dispatcher $dispatcher
     */
    public function _initDefaultName(\Yaf\Dispatcher $dispatcher) {
        $dispatcher->setDefaultModule('index')
                ->setDefaultController('index')
                ->setDefaultAction('index');
    }

    /**
     * 处理会话
     * Session 直接交由Yaf自带的类去处理.
     */
    public function _initSession() {
        $session = \Yaf\Session::getInstance();
        \Yaf\Registry::set('session', $session);
    }

    /**
     * 注册本地类
     */
    public function _initRegisterLocalNamespace() {
        //申明本地类
        \Yaf\Loader::getInstance()->registerLocalNamespace(array('Zend', 'Ku'));
    }

    /**
     * Redis数据库
     *
     * @throws Exception 'Redis is need redis Extension!
     */
    public function _initRedis() {
        if (!extension_loaded('redis')) {
            throw new Exception('Redis is need redis Extension!');
        }

        $conf = $this->config->get('resources.redis');

        if (!$conf)
            return false;

        $redis = new \Redis();

        /*
         * 连接Redis
         *
         * 当没有定义 port 时, 可以支持 sock.
         * 但是, 需要注意: 如果host是IP或者主机名时, port 的默认值是 6379
         */
        if ($conf->get('port'))
            $status = $redis->pconnect($conf->get('host'), $conf->get('port'));
        else
            $status = $redis->pconnect($conf->get('host'));

        if (!$status)
            throw new \Exception('Unable to connect to the redis:' . $conf->get('host'));

        // 是否有密码
        if ($conf->get('auth'))
            $redis->auth($conf->get('auth'));

        // 是否要切换Db
        if ($conf->get('db'))
            $redis->select($conf->get('db'));

        // Key前缀
        if ($conf->get('options.prefix'))
            $redis->setOption(\Redis::OPT_PREFIX, $conf->get('options.prefix'));

        \Yaf\Registry::set('redis', $redis);
    }

    /**
     * 通过派遣器得到默认的路由器
     * 主要有以下几种路由协议
     *
     * Yaf\Route_Simple
     * Yaf\Route_Supervar
     * Yaf\Route_Static
     * Yaf\Route_Map
     * Yaf\Route_Rewrite
     * Yaf\Route_Regex
     *
     * @param \Yaf\Dispatcher $dispatcher
     */
    public function _initRoute(\Yaf\Dispatcher $dispatcher) {
        // 通过派遣器得到默认的路由器
        $router = $dispatcher->getRouter();

        if ($this->config->routes)
            $router->addConfig($this->config->routes);

        // 添加一个以 Module\Controller\Acation 方式优先的路由.
        $mcaRoute = new \Ku\Route();
        $router->addRoute('Kumca', $mcaRoute);
    }

    /**
     * 连接 MySQL
     */
    public function _initMySQL() {
        $conf = $this->config->get('resources.database.params');

        if (!$conf)
            return false;

        $dbAdapter = new \Zend\Db\Adapter\Adapter($conf->toArray());
        \Yaf\Registry::set('dbAdapter', $dbAdapter);
    }

    /**
     * 设置Layout
     *
     * @param \Yaf\Dispatcher $dispatcher
     */
    public function _initLayout(\Yaf\Dispatcher $dispatcher) {
        $layout = new \Ku\Layout($this->config->get('application.layout.directory'));
        $dispatcher->setView($layout);
        \Yaf\Registry::set('layout', $layout);
    }

}
