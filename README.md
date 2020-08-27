# mirbft-visualizer

This project is provides a WebAssembly (WASM) based application capable of executing one or more [MirBFT](https://github.com/IBM/mirbft) state machines internally, while rendering a visualization of each state machine state as events are generated and received.

Visualization can be done in one of two modes and is selected on the first page form.

1. The WASM application can generate a mocked network, comprised of a set of Mir nodes connected via a virtual network.  These nodes operate just like a real network, events like network messages and ticks are fed into the state machine, and the internal 'application' consumers of these nodes perform hashing, enqueue broadcasts, etc.  In this mode, the network will operate indefinitely, heartbeating, and rotating epoch, as the events of client submission and ticks are injected externally.
2. The WASM application can replay a log from a real (or test) network.  There is a [recording format](https://github.com/IBM/mirbft/tree/recorder/recorderpb) that this mode can process.  These recording files can be produced via an `EventInterceptor` like [this one](https://github.com/IBM/mirbft/tree/recorder/interceptor.go), or via tests, like the integration tests used by the MirBFT library.

## Technology

Naturally, this project depends on the [MirBFT library](https://github.com/IBM/mirbft), as well as the [Vugu Framework](https://github.com/vugu/vugu) which implements a Golang based web component framework (like VueJS).

## Running the visualizer

It is possible to build a deployable WASM application using the [instructions on the Vugu site](https://www.vugu.org/doc/build-and-dist).  At some point, I'll add a link to a built one that can be demoed via the Internet in your browser here.  However, at this stage of rapid development, it's usually easier to simply run the development server yourself.  You can do this like so:

```sh
$ go get -u github.com/vugu/vgrun
$ vgrun devserver.go
2020/08/27 12:13:43 Starting HTTP Server at "127.0.0.1:8844"
```

At this point, you should be able to navigate to the address printed at startup in your web browser to play with the application.

## Current limitations

There are a number of outstanding TODO items which are visible inside of the application.  For instance, when first starting, there are `*pb.StateEvent_Initialize` events which are simply printed as unknown.  There are other similar cases where messages are not displayed with any sort of detail.  Finally, there are off by one errors in some places, and likely many other bugs that could/should be addressed.

The [MirBFT library](https://github.com/IBM/mirbft) is still under heavy development, including structural changes, so occasional syncs with this upstream project will require fixing compilation and other erros, as the API is not yet stable.

## Desired (but unimplemented) features

* The ability to only display events of a certain type.
* The ability to skip to a certain time in the log.
* The ability to skip to the next instance of a certain event type in the log.
* Predefined profiles for different interesting configurations (e.g. Classical PBFT, Leader Spinning, etc.)
* General prettification.

## Contributing

Feel free to submit pull requests, but please make sure that your commit includes a DCO sign-off. To add a DCO to your commit, simply create your commit with `git commit -s` (and make sure that you have completed the [first time identity setup](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup#_your_identity) for your git environment.
