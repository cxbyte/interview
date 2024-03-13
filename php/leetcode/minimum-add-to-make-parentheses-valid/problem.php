<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function minAddToMakeValid($s)
    {
        $stack = [];
        $strlen = strlen($s);
        for ($i = 0; $i < $strlen; $i++) {
            $currentSymbol = $s[$i];
            if ($currentSymbol == '(') {
                array_push($stack, $currentSymbol);
            } else {
                // if opened parentheses on top of stack
                if (($stack[count($stack) - 1] ?? 0) == '(') {
                    array_pop($stack);
                } else {
                    array_push($stack, $currentSymbol);
                }
            }
        }

        return count($stack);
    }
}

$expected = 1;
$in = "())";
$out = (new \Solution())->minAddToMakeValid($in);
if ($expected != $out){
    var_dump( '$expected: ', $expected);
    var_dump('$out: ', $out);
} else {
    var_dump('Pass');
}