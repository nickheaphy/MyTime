function validateFormData(formdataobject) {
    // check to make sure all the required fields have been set
    const requiredData = ['start','end','bgcolour','description','customer','primaryLogType','secondaryLogType']
    for (i in requiredData) {
        if (formdataobject.hasOwnProperty(requiredData[i]) != true) return false

    }
    return true
}

// change the colours
function LightenDarkenColor(col, amt) {
    col = parseInt(col, 16);
    return (((col & 0x0000FF) + amt) | ((((col >> 8) & 0x00FF) + amt) << 8) | (((col >> 16) + amt) << 16)).toString(16);
}

// https://stackoverflow.com/questions/4833651/javascript-array-sort-and-unique
function sort_unique(arr) {
    if (arr.length === 0) return arr;
    arr = arr.sort(function (a, b) { return a*1 - b*1; });
    var ret = [arr[0]];
    for (var i = 1; i < arr.length; i++) { //Start loop at 1: arr[0] can never be a duplicate
      if (arr[i-1] !== arr[i]) {
        ret.push(arr[i]);
      }
    }
    return ret;
}

// https://domhabersack.com/snippets/array-numbers-to-fractions
function toRelative(numbers) {
  const largestNumber = Math.max(...numbers)
  return numbers.map(number => number / largestNumber)
}

console.log("helperfunction Imported")