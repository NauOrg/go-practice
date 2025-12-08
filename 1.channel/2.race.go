package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch1 <- 1
	}()
	go func() {
		ch1 <- 2
	}()

	select {
	case val := <-ch1:
		fmt.Println("Got from ch1:", val)
	case val := <-ch2:
		fmt.Println("Got from ch2:", val)
	case <-time.After(time.Second):
		fmt.Println("Timeout!")
	}

}

/***
function sendToChannel(value, delay = 0) {
  return new Promise((resolve) => {
    setTimeout(() => resolve(value), delay);
  });
}

async function main() {
  const ch1_1 = sendToChannel(1); // simulates: go routine sending 1
  const ch1_2 = sendToChannel(2); // simulates: another go routine sending 2

  const ch2 = new Promise(() => {});
  // never resolves → same as a goroutine that never sends

  const timeout = new Promise((resolve) =>
    setTimeout(() => resolve("timeout"), 1000)
  );

  const result = await Promise.race([ch1_1, ch1_2, ch2, timeout]);

  if (result === "timeout") {
    console.log("Timeout!");
  } else {
    console.log("Got from channel:", result);
  }
}

main();

***/
