package main

import "fmt"

type Emp struct {
    name string
    id int
    bill_paid bool
}
func main() {
    var user_one Emp
    fmt.Println(user_one)
    
    user_one.name = "Vijay"
    user_one.id = 4556
    user_one.bill_paid = false
    
    fmt.Println("Details of the first user -> ", user_one)
    
    user_two := Emp {
        bill_paid : false,
        id : 4890,
        name : "Anthony",
    }
    
    fmt.Println("Details of the second user -> ", user_two)
    
    user_three := Emp{"Shankar", 4367, true}
    fmt.Println("Third user details -> ", user_three)
    
}
