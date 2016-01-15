<?php

class Foo {

    public $a = 1;

}

function Bar($foo) {
    $foo->a += 1;
}

$f = new Foo();
Bar($f);

echo $f->a;
