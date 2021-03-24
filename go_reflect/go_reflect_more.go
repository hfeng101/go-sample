package go_reflect

import (
	"reflect"
	"testing"
)

type Users struct {
	ID   int
	Name string
}

type TestInterface interface {
	GetName() string
}

func (u *Users) UpdateName(newName string) {
	u.Name = newName
}
func (u *Users) GetName() string {
	return u.Name
}

func TestReflect(t *testing.T) {
	u := Users{1, "mike"}
	//返回指定对象的Kind类型
	t.Log(reflect.TypeOf(32).Kind())
	t.Log(reflect.ValueOf(32).Kind())

	//根据方法名找方法
	t.Log(reflect.TypeOf(&u).MethodByName("UpdateName"))
	t.Log(reflect.ValueOf(&u).MethodByName("UpdateName"))

	//返回第i个方法
	t.Log(reflect.TypeOf(&u).Method(0))
	t.Log(reflect.ValueOf(&u).Method(0))

	//返回拥有的方法总数，包括unexported方法
	t.Log(reflect.TypeOf(&u).NumMethod())
	t.Log(reflect.ValueOf(&u).NumMethod())

	//取struct结构的第n个field
	t.Log(reflect.TypeOf(u).Field(0))
	t.Log(reflect.ValueOf(u).Field(1))

	//嵌套的方式取struct的field，比如v.FieldByIndex(1,2,3)等价于 v.field(1).field(2).field(3)
	t.Log(reflect.TypeOf(u).FieldByIndex([]int{0}))
	t.Log(reflect.ValueOf(u).FieldByIndex([]int{0}))

	//返回名称匹配match函数的field
	t.Log(reflect.TypeOf(u).FieldByName("ID"))
	t.Log(reflect.ValueOf(u).FieldByName("Name"))

	//返回struct所包含的field数量
	t.Log(reflect.TypeOf(u).NumField())
	t.Log(reflect.ValueOf(u).NumField())

	//分配内存时的内存对齐字节数
	t.Log(reflect.TypeOf(u).Align())
	//作为struct的field时内存对齐字节数
	t.Log(reflect.TypeOf(u).FieldAlign())
	//type名 string类型
	t.Log(reflect.TypeOf(u).Name())
	//包路径， "encoding/base64"， 内置类型返回empty string
	t.Log(reflect.TypeOf(u).PkgPath())
	//该类型变量占用字节数
	t.Log(reflect.TypeOf(u).Size())
	//type的string表示方式
	t.Log(reflect.TypeOf(u).String())
	//判断该类型是否实现了某个接口
	t.Log(reflect.TypeOf(u).Implements(reflect.TypeOf((*TestInterface)(nil)).Elem()))
	//判断该类型能否赋值给某个类型
	t.Log(reflect.TypeOf(u).AssignableTo(reflect.TypeOf(Users{})))
	//判断该类型能否转换为另外一种类型
	t.Log(reflect.TypeOf(u).ConvertibleTo(reflect.TypeOf(1)))
	//判断该类型变量是否可以比较
	t.Log(reflect.TypeOf(u).Comparable())
	//取该类型的元素，指针指向的结构
	t.Log(reflect.TypeOf(&u).Elem())

	//调用函数
	t.Log(reflect.ValueOf(&u).MethodByName("GetName").Call([]reflect.Value{}))
	//判断能否取地址
	t.Log(reflect.ValueOf(&u).CanAddr())
	//判断Interface方法能否使用
	t.Log(reflect.ValueOf(&u).CanInterface())
	//判断值能否改变
	t.Log(reflect.ValueOf(&u).CanSet())

	a := []int{0, 1}
	//获取容量 Array/Chan/Slice
	t.Log(reflect.ValueOf(a).Cap())
	c := make(chan int)
	//关闭channel
	reflect.ValueOf(c).Close()
	//返回指针实际的值
	t.Log(reflect.ValueOf(&u).Elem())
	//索引操作 Array/Slice/String
	t.Log(reflect.ValueOf(a).Index(0))
	//修改数组第一个索引的值
	reflect.ValueOf(a).Index(0).Set(reflect.ValueOf(1))
	t.Log(a[0])
	//将当前value以interface形式返回
	t.Log(reflect.ValueOf(&u).Interface())
	//判断是否为nil，chan, func, interface, map, pointer, or slice valu
	t.Log(reflect.ValueOf(&u).IsNil())
	//是否是可操作的Value，返回false表示为zero Value
	t.Log(reflect.ValueOf(&u).IsValid())
	//获取长度，适用于Array, Chan, Map, Slice, or String
	t.Log(reflect.ValueOf(a).Len())
	m := map[int]string{1: "Mike", 2: "Tom"}
	//对map类型按key取值
	t.Log(reflect.ValueOf(m).MapIndex(reflect.ValueOf(1)))
	//map类型的所有key的列表
	for index, key := range reflect.ValueOf(m).MapKeys() {
		t.Log("key=", key)
		t.Log("idnex=", index)
	}
	//返回value的Type
	t.Log(reflect.ValueOf(1).Type())
}
