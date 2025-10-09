//https://pkg.go.dev/context
// context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
//----------------------------
//1. Controlling timeouts
//2. Cancelling go routines
//3. Passing metadata across your go application

//
/*
   The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc. Calling the CancelFunc directly cancels the child and its children, removes the parent's reference to the child, and stops any associated timers. Failing to call the CancelFunc leaks the child and its children until the parent is canceled. The go vet tool checks that CancelFuncs are used on all control-flow paths.
*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
     
	//initialize a context that is used my main  
	ctx := context.Background();

	//context with timeouts
	//cancel -> cancel
	// context.withTimeout(context.Backgorund(),3*time.Second)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second) 
	defer cancel();

	done := make(chan bool);

	go func(){
      time.Sleep( 4*time.Second)
	  done<-true
	}()

	select {
	case <-done:
		fmt.Println("The temp work is done!")
	case <-ctxWithTimeout.Done():
		fmt.Println("TIME OUT!",ctxWithTimeout.Err())

	}
	//IN CONSOLE:  TIMEOUT! Deadline exceeded  (cuz the temp work taking too much time(more than 3 sec defined in the context timeout))
     //----------------------------
	// with values
    
	// ctxWithValue := context.WithValue(ctx,"user_id",10);
	//warning:should not use built-in type int as key for value; define your own type to avoid collisions (SA1029)go-staticcheck

	type key int;
	var user_id key;
	ctxWithValue := context.WithValue(ctx,user_id,10);

	value,ok := ctxWithValue.Value("user_id").(int);
	if ok {
		fmt.Println("The value is ",value)
	}else {
		fmt.Println("nope:",value)
	}


	//  real use case
    //in controller
	Handler()
	

}
func Handler (w http.ResponseWriter, r * http.Request){
     //context in request 
      
	 ctx ,cancel := context.WithTimeout(r.Context(),2*time.Second)
      defer cancel()
	 select {
	 case <-time.After(3 * time.Second):
		fmt.Println("API RESPONSE!!!")
	 case <- ctx.Done():
		fmt.Println("The context expired")
		http.Error(w,"Request context time out.took too much time",http.StatusRequestTimeout)
	 }

	}