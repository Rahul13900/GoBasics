Channel is one way communication pipe, things go in one end comes out the other end in the same order
Multiple reader and write can share it safely

Goroutine is a unit of independent execution (coroutine)
its easy to start a goroutine just put go infront of the function call

The trick is knowing how to stop a goroutine
1. you have well defined loop termination or
2. you have completion signal through channel or context or
3. you let it run unit the program stops


Select statement allows you to take response from the channel as they are ready to be read instead of following a particular pattern. 