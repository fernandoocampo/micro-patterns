# Adapter pattern

> "Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate." - [refactoring.guru](https://refactoring.guru/design-patterns/adapter)

## Problem

Let's suppose we defined a paying module and we need to integrate with multiples 3d party payment providers, but all of them provide different api contracts. 

## Solution

We want uniformity, so we are going to apply the adapter pattern, to convert the interface of one object so that another object can understand it.