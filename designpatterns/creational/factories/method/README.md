# Factory method

Creational design pattern that creates objects through a defined operation. This operation can be altered by subtypes that needs to provide different objects.

Factory Method solves the following problems:

* How can an object be created so that subclasses can redefine which class to instantiate?
* How can a class defer instantiation to subclasses?

## Problem

* Let's suppose we are creating a game about motorcycle simulator. In the game there are many factories that make beautiful motorcycles and the same factories make their own engines. All the engines are not equal and some of them must be built different and we don't want to couple our code to a specific engine.

Now the idea is that when you (as a user) order a new motorcycle, the motorcycle factory will create a new engine based on the given parameters and category.

## Solution

We need an option to avoid instantiating specific engines directly and to avoid adding multiple conditions that choose the engine to create. The factory method pattern uses factory methods to deal with the problem of creating objects without having to specify the exact type.

* Since we are going to have different type of engines we need to define a common behavior for them. We use an interface to specify such a behavior and we called it `EngineBehavior`. Remember interfaces provide behavior not data.
* We created a package called `engines`that will provide behavior to build engines.
* `engines` will define the common behavior for motorcyle engine types.
* I decided to create the factory packages inside the specific brand owner of the engines, so you will see the brands there.
* Each brand can provide different engine lines: urban, sport and adventure.
* Each factory define the internal logic for their engines.