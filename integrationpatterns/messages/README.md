# message

When two applications wish to exchange a piece of data, they do so by wrapping it in a message.

## Message

* Problem
> How can two applications connected by a message channel exchange a piece of information?
* Solution
> Package the information into a Message, a data record that the messaging system can transmit through a message channel.

## Command Message

An application needs to invoke functionality provided by other applications. It would typically use `Remote Procedure Invocation`, but would like to take advantage of the benefits of using Messaging.

* Problem
> How can messaging be used to invoke a procedure in another application?
* Solution
> Use a Command Message to reliably invoke a procedure in another application.

## Document Message

An application would like to transfer data to another application. It could do so using File Transfer or Shared Database, but those approaches have shortcomings. The transfer might work better using Messaging.

* Problem
> How can messaging be used to transfer data between applications?
* Solution
> Use a Document Message to reliably transfer a data structure between applications.

## Event Message

Several applications would like to use event-notification to coordinate their actions, and would like to use Messaging to communicate those events.

* Problem
> How can messaging be used to transmit events from one application to another?
* Solution
> Use an Event Message for reliable, asynchronous event notification between applications.

## Request/Reply

When two applications communicate via Messaging, the communication is one-way. The applications may want a two-way conversation.

* Problem

> When an application sends a message, how can it get a response from the receiver?
> How does a replier know where to send the reply?

* Solution

> Send a pair of Request-Reply messages, each on its own channel.
> The request message should contain a Return Address that indicates where to send the reply message.

## Return address

Similar to Request/Reply, we need to answer this question `How does a replier know where to send the reply?`. The answer `the request message should contain a Return Address that indicates where to send the reply message`.

## Correlation Identifier

* Problem

> How does a requestor that has received a reply know which request this is the reply for?

* Solution

> Each reply message should contain a Correlation Identifier, a unique identifier that indicates which request message this reply is for.

There are also proposals like [OpenTelemetry](https://opentelemetry.io) that you can use to achieve this.

* how to test?

```sh
go1.18.3 test -race -timeout 5s -count 1 -run ^TestCorrelationID$ github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/correlations
```

## Message Expiration.

My application is using Messaging. If a Message’s data or request is not received by a certain time, it is useless and should be ignored.

* Problem

How can a sender indicate when a message should be considered stale and thus shouldn’t be processed?

* Solution

Set the Message Expiration to specify a time limit how long the message is viable.

## Format Indicator

Several applications are communicating via Messages that follow an agreed upon data format, perhaps an enterprise-wide Canonical Data Model. However, that format may need to change over time.

* Problem

How can a message’s data format be designed to allow for possible future changes?

* Solution

Design a data format that includes a Format Indicator, so that the message specifies what format it is using.