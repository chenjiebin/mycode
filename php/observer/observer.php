<?php

/**
 * 观察者
 */
abstract class Observer {

    abstract public function update();
}

/**
 * 具体的观察者
 */
class ConcreteObserver extends Observer {

    private $observerState;
    private $name;
    private $concreteSubject;

    public function __construct(subject $concreteSubject, $name) {
        $this->setConcreteSubject($concreteSubject);
        $this->setName($name);
    }

    public function update() {
        $this->setObserverState($this->concreteSubject->getSubjectState());
        echo $this->getObserverState();
    }

    public function getObserverState() {
        return $this->observerState;
    }

    public function getName() {
        return $this->name;
    }

    public function getConcreteSubject() {
        return $this->concreteSubject;
    }

    public function setObserverState($observerState) {
        $this->observerState = $observerState;
    }

    public function setName($name) {
        $this->name = $name;
    }

    public function setConcreteSubject($concreteSubject) {
        $this->concreteSubject = $concreteSubject;
    }

}

/**
 * 观察者模式抽象subject类
 */
abstract class subject {

    private $observers = array();

    public function attach($observer) {
        $this->observers[] = $observer;
    }

    public function detach($observer) {
        $this->observers[] = $observer;
    }

    public function notify() {
        foreach ($this->observers as $key => $value) {
            $value->update();
        }
    }

}

/**
 * 观察者模式
 */
class ConcreteSubject extends subject {

    private $subjectState;

    function getSubjectState() {
        return $this->subjectState;
    }

    function setSubjectState($subjectState) {
        $this->subjectState = $subjectState;
    }

}

$subject = new ConcreteSubject();
$subject->attach(new ConcreteObserver($subject, "observer 1"));
$subject->attach(new ConcreteObserver($subject, "observer 2"));
$subject->attach(new ConcreteObserver($subject, "observer 3"));

$subject->setSubjectState("1");
$subject->notify();
