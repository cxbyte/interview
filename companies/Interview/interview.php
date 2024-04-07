<?php

/**
 *
 * Take below input data for number of births in different city in different years.
Calculate average birth rate of each city (total birth/number of years) and print it in ascending order on HTML page in tabular format using PHP program
 */

$birth_info = array(
    2021 =>
        array("Edmonton"=> 2989, "Toronto"=>25670, "New York" => 56685,
            "Washington" => 98056, "Chigaco" => 7890),

    2022 =>
        array("Edmonton"=> 2708, "Toronto"=>25503, "New York" => 50634,
            "Washington" => 98021, "Chigaco" => 7230),

    2023 =>
        array("Edmonton"=> 2182, "Toronto"=>24630, "New York" => 56681,
            "Washington" => 98012, "Chigaco" => 7821)
);


$birthesByCity = [];
$birthesYears = [];
foreach ($birth_info as $year => $birth){
    foreach ($birth as $city => $value){
        $birthesByCity[$city] += $value;
        $birthesYears[$city]++;
    }
}

foreach ($birthesByCity as $city => $value){
    $birthesByCity[$city] = $birthesByCity[$city] / $birthesYears[$city];
}

asort($birthesByCity);

foreach ($birthesByCity as $city => $value){
    echo sprintf('City:%s, Value: %s', $city, $value) . PHP_EOL;
}


class Singleton {
    private static $object;

    public static function getEntity(){
        if (self::$object === null ){
            self::$object = new self();
        };

        return self::$object;
    }

    private function __construct()
    {

    }

[2:28 AM] Umar Farook
emp table
[2:28 AM] Umar Farook
empid empName MgrId
[2:29 AM] Umar Farook
mgrid as well employee id
[2:29 AM] Umar Farook
write a query to display each emp name and their manager name
}


select
emp.empName, emp2.empName
from emp
left join emp as emp2 on emp.MgrId = emp2.empid