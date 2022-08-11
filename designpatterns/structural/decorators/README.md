# Decorator pattern

> "Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors." - [refactoring.guru](https://refactoring.guru/design-patterns/decorator)

## Other names

* Wrapper.

## Problem

Let's imagine that our motorcycle game simulates motorcyclist crashes. Therefore, users must wear clothing to protect themselves. For now, we are going to protect only the rider head with a helmet. So when a crash occurs, each part within the helmet will protect the rider head from injury.

## Solution

Let's use decorator pattern to simulate an impact over the helmet that goes through all of its internal layers.

## Other recommendations

* Use the Decorator pattern when you need to be able to assign extra behaviors to objects at runtime without breaking the code that uses these objects.

* The Decorator lets you structure your business logic into layers, create a decorator for each layer and compose objects with various combinations of this logic at runtime. The client code can treat all these objects in the same way, since they all follow a common interface.

* Use the pattern when it’s awkward or not possible to extend an object’s behavior using inheritance.

* Many programming languages have the final keyword that can be used to prevent further extension of a class. For a final class, the only way to reuse the existing behavior would be to wrap the class with your own wrapper, using the Decorator pattern.