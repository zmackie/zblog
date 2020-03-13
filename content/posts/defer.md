---
title: "Defer"
date: 2020-03-13T10:23:38-04:00
draft: false
---

Defer in Go is profoundly weird. In general Go does not do weird. But this appears to be an exception (oh, wait it doesn't have those either 😸). I won't go too much into gotchas and stuff but I wanted to document some things I learned the other day.

So what does the following code print? ([Playground](https://play.golang.org/p/Dk8WIc7tPKJ))
```go
package main

import "fmt"

func main() {
	a, b := test()

	fmt.Printf("%d, %s", a, b)
}

func test() (int, string) {
	var err string

	defer func() {
		err = "after"
		fmt.Println("after")
	}()
	err = "before"
	fmt.Println("before")
	return 1, err
}

```

It prints:
```bash
before
after
1, before
```

This is...perhaps suprising. Maybe it's not.

What about this code? ([Playground](https://play.golang.org/p/Sb3c9-MqaCB))
```go
package main

import "fmt"

func main() {
	a, b := test()

	fmt.Printf("%d, %s", a, b)
}

func test() (i int, err string) {
	defer func() {
		err = "after"
		fmt.Println("after")
	}()
	err = "before"
	fmt.Println("before")
	return 1, err
}

```

Given what I showed above, you might be surprised to learn that it prints:
```bash
before
after
1, after
```

https://golang.org/ref/spec#Defer_statements documents this, but I have to say I was quite taken aback.

Pasting here for posterity (emphasis mine):
> Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. That is, **if the surrounding function returns through an explicit return statement, deferred functions are executed after any result parameters are set by that return statement but before the function returns to its caller.** If a deferred function value evaluates to nil, execution panics when the function is invoked, not when the "defer" statement is executed.

So it looks something like:

```bash
<- evaulate defer statement and add to defer stack (but don't invoke) 
<- set return values
<- invoke stack of defer functions
<- return
```

In the case of the `name result parameter` example, I'm honestly still a bit confused *why* it behaves this way; I'm guessing the behavior is because the scope of the `err` variable in the first example is the `{}` of the function block and the named result parameters are outside that scope, and therefore available for assignment from within the defer. I'm sure I could get this with some piecing together of the spec (https://golang.org/ref/spec#Declarations_and_scope) but for the moment I need more ☕️.
