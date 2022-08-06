# prototype

Creational design pattern, It is used when the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects.

## Problem

Let's say we want to create an exact copy of your motorcycle in the game to share it with a friend. We can think this is easy and it's just copying the their field values to the new object, but what would happen with private attributes? on the other hand what if the motorcycle object we want to clone from is behind an interface?

## Solution

The pattern proposes that the clone logic be implemented by the objects that are being cloned. So here we can define methods in Go to clone a specific type, it is the responsibility of the client of the package to define the interface to abstract the cloneable behavior. In our case, you can look at the unit test and see that we defined a cloneable interface.