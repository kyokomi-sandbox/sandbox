function map(array, callback) {
    var result = [];
    for (var _i = 0, array_1 = array; _i < array_1.length; _i++) {
        var v = array_1[_i];
        result.push(callback(v));
    }
    return result;
}
var data = [1, 1, 2, 3, 5, 8, 13];
var result = map(data, function (x) { return x * 10; });
console.log(result);
