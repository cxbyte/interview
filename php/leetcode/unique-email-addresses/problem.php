<?php

class Solution {

    /**
     * @param String[] $emails
     * @return Integer
     */
    function numUniqueEmails($emails) {
        return count(array_unique(array_map(
            function ($email) {
                [$local, $domain] = explode('@', $email);
                [$local] = explode('+', $local);

                return str_replace('.', '', $local) . '@' . $domain;
            },
            $emails
        )));
    }
}

$expected = 3;
$in = ["a@leetcode.com","b@leetcode.com","c@leetcode.com"];
$out = (new \Solution())->numUniqueEmails($in);
if ($expected != $out){
    var_dump( '$expected: ', $expected);
    var_dump('$out: ', $out);
} else {
    var_dump('Pass');
}


/**
 * [1,1,2]
 */
