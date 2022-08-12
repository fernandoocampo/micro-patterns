# flyweight design pattern.

This design pattern is an object that minimizes memory usage by holding data related to objects that share similarities with each other.

## Problem

Let's imagine that in our motorcycle game we constantly need to know where the player is and what things they have. For example, position, motorcycles, engines, tires, money, jackets, boots, helmets, etc. Each of the items mentioned could use thousands of bytes and having them loaded all the time would be inefficient.

## Solution

We can use the flyweight design pattern, where an instance of the player state has references to every item it owns or some lightweight attributes that we can keep loaded. The reference can be an ID that identifies the element and the attributes can be primitives. only in case it is necessary to load the element information we use the reference id to read the element with all its weighted attributes.

