/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {

    let neg = false
    if ( x < 0){
        neg = true
        x = -x
    }

    let rev = 0
    let nex = x
    while (nex > 0){
        rev = rev * 10 + nex % 10;
        nex = parseInt(nex /10)
    }


    if (neg){
        rev = rev + "-"
    }


    return x === rev
};