<?php

class Solution
{
    /**
     * @param Integer $x
     * @return Integer
     */
    function reverse(int $x): int {
        $result = 0;
        $string_x = (string) $x;
        $reversed_string_x = strrev($string_x);

        if ($string_x[0] == "-") $reversed_value = 0 - (int) $reversed_string_x;
        else $reversed_value = (int) $reversed_string_x;

        if ($reversed_value < pow(2, 31) - 1 && $reversed_value > pow(-2, 31)) $result = $reversed_value;

        return $result;
    }
}
