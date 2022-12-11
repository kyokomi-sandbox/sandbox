type User = {
  name: string;
  age: number;
  premiumUser: boolean;
}

const data: string = `
kyokomi,25,1
John Smith,17,0
Mary Sue,14,1
`;

const users: User[] = [];
const lines = data.split('\n');
for (const line of lines) {
  if (line === '') {
    continue;
  }

  const [name, ageString, premiumUserString] = line.split(',');
  users.push({
    name: name,
    age: Number(ageString),
    premiumUser: Boolean(premiumUserString),
  });
}

for (const user of users) {
  if (user.premiumUser) {
    console.log(`${user.name} (${user.age}) はプレミアムユーザーです。`);
  } else {
    console.log(`${user.name} (${user.age}) はプレミアムユーザーではありません。`);
  }
}
