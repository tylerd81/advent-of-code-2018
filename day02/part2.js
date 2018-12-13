const fs = require("fs");

const checkForDifferences = (s1, s2) => {
  let numDiffs = 0;
  let sameChars = [];

  if (s1 === s2 || s1.length != s2.length) {
    return { num: s1.length, same: s1 };
  }

  for (let i = 0; i < s1.length; i++) {
    if (s1.charAt(i) !== s2.charAt(i)) {
      numDiffs++;
    } else {
      sameChars.push(s1.charAt(i));
    }
  }

  return { num: numDiffs, same: sameChars.join("") };
};

// let input = ["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"];
let input = fs
  .readFileSync("input.txt", "utf8")
  .trim()
  .split("\n");
// console.log(checkForDifferences(input[1], input[4]));
let found = false;
let same = "";

input.forEach(s1 => {
  input.forEach(s2 => {
    let diff = checkForDifferences(s1, s2);
    if (diff.num === 1) {
      found = true;
      same = diff.same;
      return false;
    }
  });
  if (found) {
    return false;
  }
});

if (found) {
  console.log(same);
}
