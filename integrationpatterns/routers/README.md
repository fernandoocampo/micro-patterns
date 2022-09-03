# Message Routing Patterns

They are used to decouple a message source from the ultimate destination of the message.

## Pipes And Filters

Let's imagine that a new event arrives and we need to execute some operations over the event's message. Transform, Encryption, Enrichment, so on and so forth.

* Problem

> How can we perform complex processing on a message while maintaining independence and flexibility?

* Solution

> Use the Pipes and Filters architectural style [1] to divide a larger processing task into a sequence of smaller, independent processing steps (Filters) that are connected by channels (Pipes).


## Message router

* Problem

> How can you decouple individual processing steps so that messages can be passed to different filters depending on a set of conditions?

* Solution

> Insert a special filter, a Message Router, which consumes a Message from one Message Channel and republishes it to a different Message Channel channel depending on a set of conditions.

here pipes and filters connects to multiple output channels. A key property of the Message Router is that it does not modify the message contents. It only concerns itself with the destination of the message.


## References

[1] Also known as pipeline. The topology of the pipeline architecture consists of pipes and filters. Pipes form the communication channel between filters. On the other hand filters perform one task only.