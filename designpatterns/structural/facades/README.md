# Facade design pattern

This is a simple interface to a complex component that provides lots of functionalities.

## Problem

Let's imagine that we are building the core functionality within the backend of our motorcycle game, but to start a new game the front should call many endpoints to create it. For instance it needs to create the world, generate characters, assign initial avatar set up, assign an initial motorcycle to the user and so forth.

## Solution

Let's implement the facade design pattern to simplify what start a new game is.