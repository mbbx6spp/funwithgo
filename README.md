# Fun With Go

This is just my "fun with Go" repository to experiment with ideas to get use
to the ideas that I could use in a larger scale, long running service.

# TODO

I plan to implement the following ideas so far:

* Fan-in examples
* Request-response examples
* Subscription cleanup examples
* Client/server interfaces for testability

The above techniques were discussed in Google I/O talks on Go Concurrency
from 2012 (Rob Pike's talk) and 2013.

# Notes

Really this is just for me, but don't mind sharing at all.

All of the techniques above take into account the idea of sequentializing
logic and access to data, which is nothing new (hint: see Erlang and
other languages before it). There is a great deal of symmetry between
goroutines communicating across channels and "processes" with incoming
mailboxes.

What intrigues me is the send-receive synchronization of non-buffered
channels in Go. I am still mulling around the implications of programming
model in my head, thus need this repository to play around with these ideas.

Feel free to look around, and if you know Go, I would be happy to receive
feedback on the code as well as how to improve parts as I am just figuring
this out on my own right now.

Cheers!
Susan
