const fs = require("fs");

const lines = fs
  .readFileSync("input.txt", "utf-8")
  .trim()
  .split("\n");
// const lines = [
//   "[1518-04-15 23:58] Guard #373 begins shift",
//   "[1518-04-15 23:58] Guard #373 begins shift"
//   //"[1518-09-14 00:54] wakes up"
// ];

function compareDates(d1, d2) {
  //get the dates
  let first = new Date(d1.substring(1, 17));
  let second = new Date(d2.substring(1, 17));

  if (first > second) {
    return 1;
  } else if (first < second) {
    return -1;
  } else {
    return 0;
  }
}

lines.sort(compareDates);
fs.writeFileSync("sorted.txt", lines.join("\n"));
console.log("done");
