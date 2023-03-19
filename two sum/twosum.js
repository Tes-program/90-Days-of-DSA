/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    let diff = undefined;

    for (let i = 0; i < nums.length; i++) {
        const curr = nums[i]
        diff = target - curr
        const i_diff = nums.indexOf(diff)

        if (i_diff != -1 && i_diff != i) return [i, i_diff] ;
    }
};