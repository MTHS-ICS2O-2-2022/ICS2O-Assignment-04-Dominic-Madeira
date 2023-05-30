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
  const TAX = 0.13
  const subMeat = parseInt(document.getElementById('sub-meat').value)
  const subLength = parseInt(document.getElementById('sub-length').value)
  let cost = 0

  // process
    if (subLength == 6) {
    if (subMeat == 1) {
      cost = 7.00
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else if (subMeat == 2) {
      cost = 5.00
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else if (subMeat == 3) {
      cost = 5.50
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else if (subMeat == 4) {
      cost = 6.00
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else {
      document.getElementById('answer').innerHTML = 'Please fill out the form completely.'
    }
    let price = TAX * cost + cost
  } else  if (subLength == 12) {

    if (subMeat == 1) {
      const cost = 7.00 * 1.75
      const TAX = 0.13
      const price = TAX * cost + cost
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else if (subMeat == 2) {
      const cost = 5.00 * 1.75
      const TAX = 0.13
      const price = TAX * cost + cost
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else if (subMeat == 3) {
      const cost = 5.50 * 1.75
      const TAX = 0.13
      const price = TAX * cost + cost
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else if (subMeat == 4) {
      const cost = 6.00 *1.75
      const TAX = 0.13
      const price = TAX * cost + cost
      document.getElementById('answer').innerHTML = 'Your total comes to $' + price.toFixed(2) + '. Thank you for eating at Subway!'
    } else {
      document.getElementById('answer').innerHTML = 'Please fill out the form completely.'
    }
  } else {
    document.getElementById('answer').innerHTML = 'Please fill out the form completely.'
  }
}
