<?php

class Solution {

    /**
     * @param Integer $target
     * @param Integer[] $nums
     * @return Integer
     */
    function minSubArrayLen($target, $nums) {
        if (!$nums) {
            return 0;
        }

        $firstPointer = 0;
        $secondPointer = 0;
        $sum = 0;
        $lastNum = count($nums) - 1;

        while (1) {
            $sum += $nums[$secondPointer];
            if ($sum >= $target || $secondPointer >= $lastNum) break;

            $secondPointer++;
        }

        $windowSize = $secondPointer - $firstPointer + 1;

//        while (1) {
//
//        }

        var_dump($secondPointer - $firstPointer + 1);
    }
}

$expected = 2;
$target = 7;
$nums = [2,3,1,2,4,3];

$out = (new \Solution())->minSubArrayLen($target, $nums);
if ($expected != $out){
    var_dump( '$expected: ', $expected);
    var_dump('$out: ', $out);
} else {
    var_dump('Pass');
}
