const fs = require('fs');

const input = fs.readFileSync('input.txt', 'utf-8').trim().split('\n');
let sum = 0;

for(let i = 0; i < input.length; i++) {
    let n = Number.parseInt(input[i]);
    if(n) {
        sum += n;
    }
}

console.log(sum);


