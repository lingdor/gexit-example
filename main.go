package main

import (
"context"
"fmt"
"github.com/lingdor/gexit"
"time"
)


func init(){
	//init
	ctx:=context.Background()
	gexit.InitGExit(ctx)
	gexit.AddPreExitFunc(func(){
		fmt.Println("pre exit invoked!")
	})

}

func main(){

	for i:=0;i<3;i++ {
		go func(num int){
			ctx:=gexit.GetContext()  //get exit context
			for{
				select{
				case <-ctx.Done(): //when exit event
					fmt.Printf("goroutine %d at done!\n",num)
					return
				default:break
				}
				//normal program
				time.Sleep(time.Second)
				fmt.Printf("goroutine %d at runing!\n",num)
			}
		}(i)
	}
	fmt.Println("Press Ctrl - c or run : kill -3 no. ")
	<- context.Background().Done()
}
