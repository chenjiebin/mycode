<?php

namespace Our;

class Common {

    /**
     * ajax请求时输出统一格式的信息,并退出
     * 
     * @param string $message
     * @param int $err 错误编号
     */
    public function echoAjaxErrorMessage($message, $callback, $err = 1) {
        $return = array(
            'err' => $err,
            'errmsg' => $message,
            'result' => array()
        );
        echo $callback . '(' . json_encode($return) . ')';
        exit();
    }

    /**
     * ajax请求时输出成功的信息,并退出
     * 
     * @param string $message
     * @param array $result
     */
    public function echoAjaxSuccessMessage($message, $callback, $result = array()) {
        $return = array(
            'err' => 0,
            'errmsg' => $message,
            'result' => $result
        );
        echo $callback . '(' . json_encode($return) . ')';
        exit();
    }

    /**
     * 对cacheid进行md5加密
     * 
     * @param array $paramArray
     * @return string
     */
    public static function md5CacheId($paramArray) {
        if (!$paramArray || !is_array($paramArray)) {
            throw new Exception('md5 cacheId error!');
        }
        $cacheId = '';
        if ($paramArray) {
            foreach ($paramArray as $param) {
                if (is_bool($param)) {
                    $param = $param ? "1" : "0";
                } else if (is_array($param)) {
                    $param = md5(json_encode($param));
                } else if (is_object($param)) {
                    $param = md5(serialize($param));
                }
                $cacheId .= '_iceup_' . $param;
            }
        }

        $md5String = md5($cacheId);
        return $md5String;
    }

    /**
     * 判断一个数是否是正整数
     * 
     * @param string|int|float $num 
     * @return mixed
     */
    public static function isPositiveIntNumber($num) {
        $int_options = array("options" => array("min_range" => 1));
        return filter_var($num, FILTER_VALIDATE_INT, $int_options);
    }

    /**
     * 判断一个数是否是非负整数
     * 
     * @param string|int|float $num 
     * @return mixed
     */
    public static function isNonnegativeIntNumber($num) {
        $int_options = array("options" => array("min_range" => 0));
        return filter_var($num, FILTER_VALIDATE_INT, $int_options);
    }

    /**
     * 获取http状态码
     * 
     * @param int $num 
     * @return string
     */
    public static function getHttpStatusCode($num) {
        $httpStatusCodes = array(
            100 => "HTTP/1.1 100 Continue",
            101 => "HTTP/1.1 101 Switching Protocols",
            200 => "HTTP/1.1 200 OK",
            201 => "HTTP/1.1 201 Created",
            202 => "HTTP/1.1 202 Accepted",
            203 => "HTTP/1.1 203 Non-Authoritative Information",
            204 => "HTTP/1.1 204 No Content",
            205 => "HTTP/1.1 205 Reset Content",
            206 => "HTTP/1.1 206 Partial Content",
            300 => "HTTP/1.1 300 Multiple Choices",
            301 => "HTTP/1.1 301 Moved Permanently",
            302 => "HTTP/1.1 302 Found",
            303 => "HTTP/1.1 303 See Other",
            304 => "HTTP/1.1 304 Not Modified",
            305 => "HTTP/1.1 305 Use Proxy",
            307 => "HTTP/1.1 307 Temporary Redirect",
            400 => "HTTP/1.1 400 Bad Request",
            401 => "HTTP/1.1 401 Unauthorized",
            402 => "HTTP/1.1 402 Payment Required",
            403 => "HTTP/1.1 403 Forbidden",
            404 => "HTTP/1.1 404 Not Found",
            405 => "HTTP/1.1 405 Method Not Allowed",
            406 => "HTTP/1.1 406 Not Acceptable",
            407 => "HTTP/1.1 407 Proxy Authentication Required",
            408 => "HTTP/1.1 408 Request Time-out",
            409 => "HTTP/1.1 409 Conflict",
            410 => "HTTP/1.1 410 Gone",
            411 => "HTTP/1.1 411 Length Required",
            412 => "HTTP/1.1 412 Precondition Failed",
            413 => "HTTP/1.1 413 Request Entity Too Large",
            414 => "HTTP/1.1 414 Request-URI Too Large",
            415 => "HTTP/1.1 415 Unsupported Media Type",
            416 => "HTTP/1.1 416 Requested range not satisfiable",
            417 => "HTTP/1.1 417 Expectation Failed",
            500 => "HTTP/1.1 500 Internal Server Error",
            501 => "HTTP/1.1 501 Not Implemented",
            502 => "HTTP/1.1 502 Bad Gateway",
            503 => "HTTP/1.1 503 Service Unavailable",
            504 => "HTTP/1.1 504 Gateway Time-out"
        );

        return isset($httpStatusCodes[$num]) ? $httpStatusCodes[$num] : '';
    }

    /**
     * 获取客户端IP
     *
     * @param  boolean $checkProxy
     * @return string
     */
    public static function getClientIp($checkProxy = true) {
        if ($checkProxy && isset($_SERVER['HTTP_CLIENT_IP']) && $_SERVER['HTTP_CLIENT_IP'] != null) {
            $ip = $_SERVER['HTTP_CLIENT_IP'];
        } else if ($checkProxy && isset($_SERVER['HTTP_X_FORWARDED_FOR']) && $_SERVER['HTTP_X_FORWARDED_FOR'] != null) {
            $ip = $_SERVER['HTTP_X_FORWARDED_FOR'];
        } else {
            $ip = $_SERVER['REMOTE_ADDR'];
        }

        return $ip;
    }

    /**
     * 获取当前访问的url地址
     * 
     * @return string
     */
    public static function getRequestUrl() {
        return 'http://' . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];
    }

    /**
     * 调试方法
     * 
     * @param mixed $data
     */
    public static function dump($data) {
        echo "<pre>";
        var_dump($data);
        echo "</pre>";
    }

    /**
     * 检测Url中是否指明不启用读缓存
     */
    public static function isUrlSetNoReadCache() {
        if (stripos($_SERVER['REQUEST_URI'], 'NO_READ_CACHE') !== false) {
            return true;
        }
        return false;
    }

    /**
     * 检测Url中是否指明不启用缓存
     */
    public static function isUrlSetNoCache() {
        if (stripos($_SERVER['REQUEST_URI'], 'NO_CACHE') !== false) {
            return true;
        }
        return false;
    }

}
