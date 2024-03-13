<?php

class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function islandPerimeter($grid) {
        $perimeter = 0;

        foreach ($grid as $i => $row) {
            foreach ($row as $j => $cell) {
                if ($cell === 0) {
                    continue;
                }
                $perimeter += 4
                    - ($grid[$i - 1][$j] ?? 0) // left
                    - ($grid[$i + 1][$j] ?? 0) // right
                    - ($grid[$i][$j - 1] ?? 0) // up
                    - ($grid[$i][$j + 1] ?? 0); // down
            }
        }

        return $perimeter;
    }
}

$expected = 16;
$in = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]];
$out = (new \Solution())->islandPerimeter($in);
if ($expected != $out){
    var_dump( '$expected: ', $expected);
    var_dump('$out: ', $out);
} else {
    var_dump('Pass');
}