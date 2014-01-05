<?php
/* Connect to an ODBC database using driver invocation */
$dsn = 'mysql:dbname=my;host=192.168.1.201';
$user = 'root';
$password = '100301';

$slave = 'mysql:dbname=my;host=192.168.1.202';

try {
    $dbh = new PDO($dsn, $user, $password);
	$slavedbh = new PDO($slave, $user, $password);
	
	$sql = 'insert into user (username) values(\''.rand(1,100).'\')';
	$dbh->query($sql);

	$result = $slavedbh->query('select * from user');
	foreach ($result as $row) {
		print_r($row);
	}
} catch (PDOException $e) {
    echo 'Connection failed: ' . $e->getMessage();
}

?>