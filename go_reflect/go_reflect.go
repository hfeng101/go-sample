package go_reflect

import (
	"fmt"
	"reflect"
)

func Sample(){


	var x float32 = 1.0

	v := reflect.ValueOf(x)
	fmt.Println("type:",v.Type())

	v.Type()

	t := reflect.TypeOf(x)
	k := t.Kind()

	t.Elem()

	t.NumField()

	t.String()

	//Channel := reflect.ChanOf(x,int)

	//Array := reflect.ArrayOf(x, 5)

	if v.CanSet(){
		fmt.Println("v can set ")
		v.SetFloat(1.6)
	}else {
		fmt.Println("v can not set")
	}

	p := reflect.ValueOf(&x)

	if p.CanSet(){
		fmt.Println("p can set ")
		p.SetFloat(1.6)
		v2 := p.Float()
		fmt.Println("p:",p," v2:",v2," x:",v)
	}else {
		fmt.Println("p can not set ")
		v2 := p.Elem()
		if v2.CanSet(){
			fmt.Println("v2 can set ")
			v2.SetFloat(1.6)
			fmt.Println("v2:",v2.Float()," x:",x)
			fmt.Println("v2地址:",v2.Addr()," x地址:",&x)
		}else{
			fmt.Println("v2 can not set")
		}
	}

}
