const sequence = (start: number, end: number): number[] => {
  const result: number[] = [];
  for (let i = start; i <= end; i++) {
    result.push(i);
  }
  return result;
};

const getFizzBuzzString = (i: number): string => {
  let out = '';

  if (i % 3 == 0) {
    out += 'Fizz';
  }

  if (i % 5 == 0) {
    out += 'Buzz';
  }
  return out || i.toString();
};

for (const i of sequence(1, 50)) {
  const message = getFizzBuzzString(i);
  console.log(message);
}

