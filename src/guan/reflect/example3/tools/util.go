//最近在尝试用Go写一些东西，发现Go不支持泛型确实是一件比较蛋疼的事，同样功能的一个类，只有底层数据结构有一点点差异，需要实现N遍。
//特别像我这种在C++世界挣扎也纠结了很多年，用惯了模板编程思想的程序猿。
//好在Golang提供了reflect机制，可以在一定程度上满足对泛型编程的一些需求。 想实现的一些需求：
//1.通过类型字符串动态创建类型对象
//2.动态的调用一些预定义函数，而不需要依赖实现该函数的package
// 3.实现一些通用的数据结构，比如像C++ STL那样的泛型容器
//4.一些特定类型的检查和操作 如chan Type，[]Type
//new object same the type as sample
package tools

import (
	"fmt"
	"reflect"
)

func New(sample interface{}) interface{} {
	t := reflect.ValueOf(sample).Type()
	v := reflect.New(t).Interface()
	return v
}

//---------------------------------check type of aninterface
func CheckType(val interface{}, kind reflect.Kind) bool {
	v := reflect.ValueOf(val)
	return kind == v.Kind()
} //---------------------------------if _func is not a functionor para num and type not match,it will cause panic
func Call(_func interface{}, params ...interface{}) (result []interface{}, err error) {
	f := reflect.ValueOf(_func)
	if len(params) != f.Type().NumIn() {
		ss := fmt.Sprintf("The number of params is not adapted.%s", f.String())
		panic(ss)
		return
	}
	var in []reflect.Value
	if len(params) > 0 { //prepare in paras
		in = make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
	}
	out := f.Call(in)
	if len(out) > 0 { //prepare out paras
		result = make([]interface{}, len(out), len(out))
		for i, v := range out {
			result[i] = v.Interface()
		}
	}
	return
} //---------------------------------if ch is not channel,itwill panic
func ChanRecv(ch interface{}) (r interface{}) {
	v := reflect.ValueOf(ch)
	if x, ok := v.Recv(); ok {
		r = x.Interface()
	}
	return
} //---------------------------------reflect fields of astruct
func reflect_struct_info(it interface{}) {
	t := reflect.TypeOf(it)
	fmt.Printf("interface info:%s %s %s %s\n", t.Kind(), t.PkgPath(), t.Name(), t)
	if t.Kind() == reflect.Ptr { //if it is pointer, get it element type
		tt := t.Elem()
		if t.Kind() == reflect.Interface {
			fmt.Println(t.PkgPath(), t.Name())
			for i := 0; i < tt.NumMethod(); i++ {
				f := tt.Method(i)
				fmt.Println(i, f)
			}
		}
	}
	v := reflect.ValueOf(it)
	k := t.Kind()
	if k == reflect.Ptr {
		v = v.Elem() //指针转换为对应的结构
		t = v.Type()
		k = t.Kind()
	}
	fmt.Printf("value type info:%s %s %s\n", t.Kind(), t.PkgPath(), t.Name())
	if k == reflect.Struct { //反射结构体成员信息
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Printf("%s %v\n", i, f)
		}
		for i := 0; i < t.NumMethod(); i++ {
			f := t.Method(i)
			fmt.Println(i, f)
		}
		fmt.Printf("Fileds:\n")
		f := v.MethodByName("func_name")
		if f.IsValid() { //执行某个成员函数
			arg := []reflect.Value{reflect.ValueOf(int(2))}
			f.Call(arg)
		}
	}
}

//PeekField 查看 一个结构体的所有
func PeekField(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var line string
	for i := 0; i < t.NumField(); i++ {
		// 返回 FieldStruct  , 如果不是该类型，会直接panic
		f := t.Field(i)
		value := v.Field(i).Interface()

		line = fmt.Sprintf(" %s name : %s \t type: %s \t value : %v \n", line, f.Name, f.Type, value)
	}
	return line

}

func PeekMethod(obj interface{}) string {
	t := reflect.TypeOf(obj)
	//	v := reflect.ValueOf(obj)
	var line string
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		line = fmt.Sprintf(" %s \t methodName: %s \t methodType %s \n ", line, method.Name, method.Type)
	}
	return line
}

// 反射给字段赋值
// SetValueForField 接受一个指针
func SetValueForField(obj interface{}, fieldName string, target int64) {
	// 得到反射类型
	t := reflect.TypeOf(obj)
	// 输出接口类型
	fmt.Println(t.Kind())
	// fieldValue, err := t.FieldByName(fieldName)
	// if err != nil {
	// 	panic(err)
	// }

	v := reflect.ValueOf(obj)
	vv := v.Elem()
	fmt.Println(vv.Kind())
	fieldValue2 := vv.FieldByName(fieldName)
	// 输出以前的值

	fmt.Println(fieldValue2.Interface())
	if fieldValue2.CanSet() {
		fieldValue2.SetInt(target)
	}

}

//反射调用方法
func CallMethod(obj interface{}, methodName string, param1 int, param2 int) int {
	v := reflect.ValueOf(obj)
	vv := v.Elem()
	fmt.Println(vv.Kind())
	method := vv.MethodByName(methodName)
	fmt.Println(method.Kind())
	args := []reflect.Value{reflect.ValueOf(param1), reflect.ValueOf(param2)}
	result := method.Call(args)
	return result[0].Interface().(int)

}
