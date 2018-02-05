package main

import "fmt"
import "reflect"


/*Primitive data types
1. int
2. float
3. boolean
4. string
5. complex */

func main() {
    var a = 1;
    var b = 2.4;
    var c = true;
    var d = "d";
    var e = 0.867 + 0.5i;
    
    fmt.Println(a);
    fmt.Println(b);
    fmt.Println(c);
    fmt.Println(d);
    fmt.Println(e);
    
    fmt.Println(reflect.TypeOf(a));
    fmt.Println(reflect.TypeOf(b));
    fmt.Println(reflect.TypeOf(c));
    fmt.Println(reflect.TypeOf(d));
    fmt.Println(reflect.TypeOf(e));

}