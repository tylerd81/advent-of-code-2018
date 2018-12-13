const fs = require('fs');

const input = fs.readFileSync('input.txt', 'utf-8').trim().split('\n');

let freqCount = {};
let found;
//const input = ['+1','-2','+3','+1'];
let sum = 0;

while(!found) {
    for(let i = 0;i < input.length; i++) {
        const num = Number.parseInt(input[i]);
        sum += num;

        if(typeof freqCount[sum] === 'undefined') {
            freqCount[sum] = 1;
        } else {
            freqCount[sum] += 1;
            if(freqCount[sum] === 2) {
                found = sum;
                break;
            }
        }
    }
}
console.log(`Found: ${found}`);
