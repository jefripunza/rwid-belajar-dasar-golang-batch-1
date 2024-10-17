const human: boolean = true;

if (human) {
  console.log("I'm a Human");
} else {
  console.log("I'm not a Human");
}

// switch
switch (human) {
  case true:
    console.log("I'm a Human");
    break;
  case false:
    console.log("I'm not a Human");
    break;
  default:
    console.log("I don't know");
}

[1, 2, 3, 4, 5].forEach((num) => {
  console.log(num);
});

// for
for (let i = 0; i < 10; i++) {
  console.log(i);
}

// while
let i = 0;
while (i < 10) {
  console.log(i);
  i++;
}

Array(10).filter();
Object.keys({});
String().toLowerCase();
Number(0.3333).toFixed(2);
Boolean().valueOf();

enum Color {
  Red = "RED",
  Green = "GREEN",
  Blue = "BLUE",
}

const warna: Color = Color.Blue;
