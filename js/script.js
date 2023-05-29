// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: May 2023
// This file contains the JS functions for index.html

//    Note: decimals do not work due to the counter incrementing by a whole one.
// They do, however, work if the only decimal is typed into the first field,
// so a solution could involve switching which of the two numbers is counted to using the counter variable to the number that is NOT a decimal.
// Decimals would still require a better solution, as if both numbers are decimals, the program would simply not work as intended.
// It may require completely changing the counting and incrementing system, which I will not do.

"use strict"

function calculatePressed() {
  //This function takes two user-given numbers and multiplies them together using only addition and loops
  //Input through Textfields
  const firstNumber = parseInt(document.getElementById("first-number").value)
  const secondNumber = parseInt(document.getElementById("second-number").value)
  let counter = null
  let answer = null

  //Process
  if (secondNumber >= 0) {
    // for all positive or firstNumber negative mixes
    while (counter < secondNumber) {
      answer = answer + firstNumber
      counter++
    }
  } else {
    // for everything else: all negative or secondNumber negative mixes
    while (counter > secondNumber) {
      answer = answer - firstNumber
      counter--
    }
  }

  //Output
  document.getElementById("answer").innerHTML = "The answer is: " + answer
}
