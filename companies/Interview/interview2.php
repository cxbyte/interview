<?php

echo "<!doctype html>\n";



$username = @$_GET['username'] ? $_GET['username'] : '';
$password = @$_GET['password'] ? $_GET['password'] : '';
$password = md5($password);
$pdo = new PDO('sqlite:
');


$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
$pdo->exec("DROP TABLE IF EXISTS users");
$pdo->exec("CREATE TABLE users (username VARCHAR(255), password VARCHAR(255))");
$rootPassword = md5("secret"); $pdo->exec("INSERT INTO users (userId, allPayments) (username, password) VALUES ('$username', '$rootPassword');");



$statement = $pdo->query("SELECT * FROM users WHERE username = '$username' AND password = '$password'");
if (count($statement->fetchAll())) {
    echo "Access granted to $username!<br>\n";
} else {
    echo "Access denied for $username!<br>\n";
}