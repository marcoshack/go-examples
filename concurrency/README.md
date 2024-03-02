
A goroutine periodically change data, and multiple other go routines read that data very frequently.

This could be for example a remote configuration data that is periodically reloaded, or reloaded based on events received by the program.

How to make the read operations non-blocking while keeping it safe and consistent?
