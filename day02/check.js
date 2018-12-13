//const input = ['abcdef','bababc','abbcde','abcccd','aabcdd','abcdee','ababab'];
const fs = require('fs');
const input = fs.readFileSync('input.txt', 'utf8').trim().split('\n');

function createHistogram(idString) {
    let histogram = {};

    for(let i = 0; i < idString.length; i++) {
        let ch = idString.toLowerCase().charAt(i);
        if(typeof histogram[ch] === 'undefined') {
            histogram[ch] = 1;
        } else {
            histogram[ch] += 1;
        }
    }
    return histogram;
}

function calcScore(histogram, currScore) {
    const values = Object.values(histogram);
    if(typeof currScore === 'undefined') {
        currScore = {2:0, 3:0};
    }

    if(values.includes(2)) {
        currScore[2] += 1;
    }

    if(values.includes(3)) {
        currScore[3] += 1;
    }

    return currScore;
}

function calcCheckSum(ids) {
    let score;

    input.forEach(id => {
        score = calcScore(createHistogram(id), score);
    });

    return score[2] * score[3];
}

//let h = createHistogram(input[1]);
//let score = calcScore(h);
//console.log(score);
//score = calcScore(createHistogram(input[2]), score);
console.log(calcCheckSum(input));
