package main

import "fmt"

func main() {

 test1 := sleepIn(false, false) 
 test2 := sleepIn(true, false) 
 test3 := sleepIn(false, true) 

 fmt.Println(test1, test2, test3);

}


/**
The parameter weekday is true if it is a weekday, and the parameter vacation is true if we are on vacation. We sleep in if it is not a weekday or we're on vacation. Return true if we sleep in.
*/
func sleepIn(weekday , vacation bool) bool {

	if !weekday || vacation {
        return true;
	}
	return false;
}