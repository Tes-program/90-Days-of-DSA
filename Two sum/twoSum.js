function twoSum(nums, target) {
  const map = new Map();
  
  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    
    if (map.has(diff)) {
      return [i, map.get(diff)];
    }
    
    map.set(nums[i], i);
  }
}
