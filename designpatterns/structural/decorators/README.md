# Decorator pattern

> "Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors." - [refactoring.guru](https://refactoring.guru/design-patterns/decorator)

## Other names

* Wrapper.

## Problem

Let's imagine that our motorcycle game simulates motorcyclist crashes. Therefore, users must wear clothing to protect themselves. For now, we are going to protect only the rider head with a helmet. So when a crash occurs, each part within the helmet will protect the rider head from injury.

## Solution

Let's use decorator pattern to simulate an impact over the helmet that goes through all of its internal layers.