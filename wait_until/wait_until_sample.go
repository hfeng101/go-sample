package wait_until

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"

	//"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type stop struct {
}

func Sample() error{
	stopCh := make(chan struct{})

	for i := 0; i < 3; i++{
		go wait.Until(func(){
			fmt.Printf("---%v----\n", i)
		}, time.Second * 2, stopCh)
	}

	stopCh<-stop{}

	return nil
}
