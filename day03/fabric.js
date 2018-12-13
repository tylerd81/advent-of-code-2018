const fs = require("fs");

const stringToClaimObject = claimString => {
  const s1 = claimString.split(" ");
  const id = s1[0];
  const startx = Number.parseInt(s1[2].split(",")[0]);
  const starty = Number.parseInt(s1[2].split(",")[1]);
  const width = Number.parseInt(s1[3].split("x")[0]);
  const height = Number.parseInt(s1[3].split("x")[1]);
  return { id, startx, starty, width, height };
};

const markFabric = (fabric, claims) => {
  claims.forEach(claim => {
    const { startx, starty, width, height } = claim;

    for (let x = 0; x < width; x++) {
      let currX = startx + x;
      for (let y = 0; y < height; y++) {
        let currY = starty + y;

        if (typeof fabric[currX] === "undefined") {
          fabric[currX] = {};
        }
        if (typeof fabric[currX][currY] === "undefined") {
          fabric[currX][currY] = [];
        }
        fabric[currX][currY].push(claim);
      }
    }
  });
};

const countOverlaps = fabric => {
  const xValues = Object.keys(fabric);
  let overlaps = 0;

  xValues.forEach(x => {
    const yValues = Object.keys(fabric[x]);
    yValues.forEach(y => {
      if (fabric[x][y].length > 1) {
        overlaps++;
      }
    });
  });
  return overlaps;
};

const findLonelyClaim = fabric => {
  // find a coordinate that has only one claim on it
  // check if that is the starting x of the claim
  // if it is use the width and height to check the other squares
  // if none of them have claims then it is the lonelyclaim
  let lonelyClaim = {};
  let found = false;
  const xValues = Object.keys(fabric);
  console.log(typeof xValues[0]);
  let possible = 0;

  xValues.forEach(x => {
    const yValues = Object.keys(fabric[x]);

    yValues.forEach(y => {
      if (fabric[x][y].length === 1) {
        if (fabric[x][y][0].startx === Number.parseInt(x)) {
          possible++;
          if (checkForCollisions(fabric, fabric[x][y]) === false) {
            lonelyClaim = fabric[x][y];
            found = true;
            return false;
          }
        }
      }
    });
    if (found) {
      return false;
    }
  });

  // return lonelyClaim;
  return { possible };
};

const checkForCollisions = (fabric, claim) => true;
const filename = "input.txt";
const input = fs
  .readFileSync(filename, "utf-8")
  .trim()
  .split("\n");
const fabric = {};
const claims = [];

input.forEach(claim => {
  claims.push(stringToClaimObject(claim));
});

markFabric(fabric, claims);

const overlaps = countOverlaps(fabric);
console.log(`There are ${overlaps} square inches within 2 or more claims.`);
console.log(findLonelyClaim(fabric));
