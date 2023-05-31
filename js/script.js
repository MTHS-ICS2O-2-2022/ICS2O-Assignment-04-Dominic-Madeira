// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {
  // input
  const TAX = 1.13
  const COST_SIX_INCH = 1
  const COST_TWELVE_INCH = 1.75
  const COST_STEAK = 7.0
  const COST_HAM = 5.0
  const COST_CHICKEN = 5.5
  const COST_TURKEY = 6.0
  let costLength = 0
  let cost = 0
  let price = 0

  // input
  const subMeat = parseInt(document.getElementById("sub-meat").value)
  const subLength = parseInt(document.getElementById("sub-length").value)

  // process



  if (subLength == 6) {
    costLength = COST_SIX_INCH
  } else {
    costLength = COST_TWELVE_INCH
  }


  if (subMeat == 1) {
    cost = COST_STEAK
  } else if (subMeat == 2) {
    cost = COST_HAM
  } else if (subMeat == 3) {
    cost = COST_CHICKEN
  } else {
    cost = COST_TURKEY
  }

  price = (costLength * cost) * TAX

  // output
  document.getElementById("answer").innerHTML = "Your total comes to $" + price.toFixed(2) + ". Thank you for eating at Subway!"
}
