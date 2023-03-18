/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    const upp_lim = Math.pow(2,31)-1;
    const low_lim = Math.pow(-2,31);

    let neg = false

    if (x < 0){
        neg = true
        x = -x
    }

    let rev = 0

    while (x > 0){
        rev = rev * 10 + x % 10;
        x = parseInt(x /10)
    }

   if (rev > upp_lim || rev < low_lim) return 0;

    return neg ? -rev : rev
    
};