import 'dart:async';
import 'dart:io';

void main() {
  print('Advent Of Code - Day 05');

  //String testInput = "dabAcCaCBAcCcaDA";
  // String input = testInput; // replace with data from file

  var file = new File('input.txt');
  var input = file.readAsStringSync().trim();
  // var input = "VvbBfpPFrRyRrYNpYyPDlLdVvNnMmnOCcosOoSoOfkKKkFJ";
  //var input = "fkKKkFJ";
  // var input = "dabAcCaCBAcCcaDA";
  // var input = "CCCAaAaAaAaAaAaAaAaAaAaAaAabB";

  String letters = "bcdefghijklmnopqrstuvwxyz";
  int smallestReactionCount = reactions(input, "a");
  String smallestReactionLetter = 'a';

  letters.split('').forEach((letter) {
    var len = reactions(input, letter);
    print("Reacting $letter");
    if (len < smallestReactionCount) {
      smallestReactionCount = len;
      smallestReactionLetter = letter;
    }
  });

  print(
      "The unit $smallestReactionLetter has the smallest reaction with $smallestReactionCount units.");
  // reactions(input, "a");
}

// work through the reactions using a list that keeps track of the previous letter
// each time a reaction happens remove the letter from the previous letter
// or when there is no reaction add the letter to the list
int reactions(String input, [String ignore]) {
  // add() and removeLast()
  var previousChars = [];
  bool done = false;
  int nextIndex = 0;
  String nextChar = null;

  while (!done) {
    if (nextIndex < input.length) {
      nextChar = input[nextIndex++];
      if (ignore != null) {
        if (nextChar.toUpperCase() == ignore.toUpperCase()) {
          continue;
        }
      }
      // compre previous character and the next char
      if (previousChars.length != 0 &&
          isSameChar(previousChars.last, nextChar)) {
        var prev = previousChars.last;
        bool reaction = false;

        if (isUpperCase(prev) && isLowerCase(nextChar)) {
          // if prev is uppercase, next must be lowercase for their
          // to be a reaction
          reaction = true;
        } else if (isLowerCase(prev) && isUpperCase(nextChar)) {
          // if prev is lower case, next must be uppercase
          reaction = true;
        } else {
          // same character but both are upper or lower case
          previousChars.add(nextChar);
        }

        if (reaction) {
          //previous and next are gone
          previousChars.removeLast();
        }
      } else {
        previousChars.add(nextChar);
      }
    } else {
      done = true;
    }
  }

  // print(previousChars.join(""));
  print('There are ${previousChars.length} units left');
  return previousChars.length;
}

// case insensitive character comparison
bool isSameChar(String a, String b) {
  if (a.toLowerCase() == b.toLowerCase()) {
    return true;
  } else {
    return false;
  }
}

bool isUpperCase(String ch) {
  if (ch == ch.toUpperCase()) {
    return true;
  } else {
    return false;
  }
}

bool isLowerCase(String ch) {
  if (ch == ch.toLowerCase()) {
    return true;
  } else {
    return false;
  }
}
