function map<T, U>(array: T[], callback: (v: T) => U): U[] {
  const result: U[] = [];
  for (const v of array) {
    result.push(callback(v));
  }
  return result;
}

const data = [1, 1, 2, 3, 5, 8, 13];
const result = map(data, (x) => x * 10);
console.log(result);

