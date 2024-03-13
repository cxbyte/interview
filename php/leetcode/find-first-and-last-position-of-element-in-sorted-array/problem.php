<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function searchRange($nums, $target)
    {
        $count = count($nums) - 1;
        $first = 0;
        $last = $count;
        $return = 0;
        $out = [];
        $currentIndex = floor(($last - $first) / 2);
        while (!$return) {
            if ($nums[$currentIndex] > $target) {
                $last = $currentIndex;
            } elseif ($nums[$currentIndex] < $target) {
                $first += $currentIndex;
            } else {
                var_dump('!');
                $out[] = $currentIndex;
            }

            $currentIndex = floor(($last - $first) / 2);

            var_dump($first, $last);

            sleep(1);



            if ($last - $first == 1) {
                $return = 1;
            }
        }

        return $out;
    }
}

$expected = [3,4];
$nums = [5,7,7,8,8,10];
//0, 5
//7
//2,5



$target = 8;
$out = (new \Solution())->searchRange($nums, $target);
if ($expected != $out){
    var_dump( '$expected: ', $expected);
    var_dump('$out: ', $out);
} else {
    var_dump('Pass');
}