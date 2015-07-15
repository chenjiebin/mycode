<?php

class Foo {

    public $name = "tom";

}

$keys = array("0" => "name");

$foo = new Foo();
echo $foo->$keys[0];

//上面的语句在php5.6如同：
echo $foo->{$keys[0]};

//在php7中如同：
//echo {$foo->$keys}[0];
