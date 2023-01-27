package main

import (
    "fmt"
    "container/list"
)

func main() {
    // Create the first linked list
    list1 := list.New()
    list1.PushBack(4)
    list1.PushBack(5)
    list1.PushBack(6)
    list1.PushBack(7)
    list1.PushBack(8)
    list1.PushBack(9)

    // Create the second linked list
    list2 := list.New()
    list2.PushBack(11)
    list2.PushBack(12)
    list2.PushBack(13)
    list2.PushBack(6)
    list2.PushBack(7)
    list2.PushBack(8)
    list2.PushBack(9)

    // Create a hash set to store the elements of the first list
    set := make(map[interface{}]bool)

    // Iterate through the first list and add elements to the hash set
    for e := list1.Front(); e != nil; e = e.Next() {
        set[e.Value] = true
    }

    // Create a list to store the common elements
    commonElements := list.New()

    // Iterate through the second list and check for elements in the hash set
    for e := list2.Front(); e != nil; e = e.Next() {
        if set[e.Value] {
            commonElements.PushBack(e.Value)
        }
    }

    // Print the common elements
    for e := commonElements.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}
