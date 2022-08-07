# builder pattern

"Builder is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code." - [refactoring.guru](https://refactoring.guru/design-patterns/builder)

## Problem

Let's imagine in our game that the motorcycle factories are going to build these wonderful motorcycles step by step. So, obviously we want to write clean code and avoid a spaghetti of logic to simulate the building process. On the other hand factories are in charge of their own building process, not the client.

Now let's imagine how to build a motorcycle, we would need to build the engine, the fairings, seats, suspensions, brakes, lights, throttle, central computer and so on.

## Solution

The builder patterns say that we should avoid letting the motorcycle itself define the build logic and let a builder do the work. So let's define a motorcycle builder object and let's build the motorcycle as the factory would need it.