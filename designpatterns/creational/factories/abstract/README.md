# Abstract factory

The Abstract Factory design pattern solves problems like:

* How can an application be independent of how its objects are created?
* How can a type be independent of how the objects it requires are created?
* How can families of related or dependent objects be created?

## Problem

* Let's suppose we are creating a game about motorcycle simulator. In the game there are many factories that make beautiful motorcycles and the idea is that you can ride them to learn the necessary skills to ride different types of motorcycles.
* In this version the game will only have the `BMW` and `Ducati` factories.
* Each factory will only make 3 types of motorcycles: `urban`, `sport` and `adventure`.
* For now, as a game user, you can buy a motorcycle and ask the factory to build the motorcycle you want.
* Once you have the motorcycle, the game will interact with the motorcycle through the behaviours of the motorbike.
    - speed up.
    - stop.
    - change gear.
    - etc.
* In the future, the game will add more motorcycle factories.

## Solution

* Since we are going to have different kind of motorcycles we can define interfaces for those products to define a specific behavior we want for the game.
* We created a package called `motorcycles`that will provide behavior for motorcycles.
* `motorcycles` will define the common behavior for motorcyle types.
* I decided to create the factory packages inside `motorcycles`, so you will see the `bmw` and `ducati` packages there.
* Each factory define the internal logic for their motorcycles.

## How to call it

* I added within the `motorcycles_test.go` file the logic about how clients of the package should use it.
* Or using the version that has generics in `factories` package

## Comments.

* I don't like that factories have to know about the categories behavior explicitly (When you don't have genercis).



## Sources

* Interesting [post](https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces) about generics


