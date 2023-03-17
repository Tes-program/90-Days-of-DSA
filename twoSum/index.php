<?php

class Solution
{
    function twoSum(array $nums, int $target): array
    {
        $solution = [];

        foreach ($nums as $index => $value) {
            foreach ($nums as $index2 => $value2) {
                if ($index == $index2) continue;
                else {
                    if ($value + $value2 == $target) $solution = [$index, $index2];
                }
            }
        }

        return $solution;
    }
}
