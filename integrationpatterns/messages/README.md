# message

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