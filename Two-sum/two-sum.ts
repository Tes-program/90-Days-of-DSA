function twoSum(nums: number[], target: number): number[] {
    /*
        We are looking for what two unknowns would give the target number
        Let the two unknowns be a and b, 
        Hence, a + b = target
        From my code, if the hashmap checks for the value of "b" i.e
        b = target - a.
        if "b" is in the hashmap, the code returns [b, a]
    */
    // Initialize a map to track the indexes and their values
    const map = new Map()

    // Result array for indexes of numbers
    const arr = []

    // Run a loop through the "nums" array
    for (let ind = 0; ind <  nums.length; ind++){

        if (map.has(target - nums[ind])){
            return [map.get(target - nums[ind]) ,  ind]
        }else{
            map.set(nums[ind], ind)
        }
    }
};