<?php

class computer extends Thread {

    public $id;
    public $runing = false;
    public $params = null;

    public function __construct($id) {
        $this->id     = $id;
        $this->runing = true;
    }

    public function run() {
        while ($this->runing) {
            if (is_null($this->params)) {
                echo "线程({$this->id})等待任务...\n";
            } else {
                echo "线程({$this->id}) 收到任务参数::{$this->params}.\n";
                $this->params = null;
            }
            sleep(1);
        }
    }

}

//这里创建线程池.
$pool = array(new computer('a'), new computer('b'), new computer('c'));

//启动所有线程,使其处于工作状态
foreach ($pool as $w) {
    $w->start();
}

//派发任务给线程
for ($i = 0; $i < 10; $i++) {
    $params = rand(10, 99);
    while (true) {
        foreach ($pool as $worker) {
            //参数为空则说明线程空闲
            if (is_null($worker->params)) {
                $worker->params = $params;
                echo "({$worker->id})线程空闲,放入参数{$params}.\n";
                break 2;
            }
        }
        sleep(1);
    }
}

//关闭线程
while (count($pool)) {
    //遍历检查线程组运行结束
    foreach ($pool as $key => $worker) {
        if ($worker->params == '') {
            echo "({$worker->id})线程运行完成,退出.\n";
            //设置结束标志
            $worker->runing = false;
            unset($pool[$key]);
        }
    }
    echo "等待退出中...\n";
    sleep(1);
}

echo "退出成功\n";
