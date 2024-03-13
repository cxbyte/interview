<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums) {
        $countNums = count($nums);
        if ($countNums == 1) {
            return 1;
        }

        $readPointer = 0;
        $writePointer = 0;
        while ($readPointer < $countNums) {
            // If the same value just move $readPointer
            if ($nums[$readPointer] == $nums[$writePointer]) {
                $readPointer++;
                continue;
            }

            // If not the same value rewrite $writePointer with $readPointer value
            $writePointer++;
            $nums[$writePointer] = $nums[$readPointer];
            $readPointer++;
        }

        return ++$writePointer;
    }
}

$expected = 5;
$in = [0,0,1,1,1,2,2,3,3,4];
$out = (new \Solution())->removeDuplicates($in);
if ($expected != $out){
    //var_dump( '$expected: ', $expected);
    var_dump('$out: ', $out);
} else {
    var_dump('Pass');
}


/**
 * [1,1,2]
 */
