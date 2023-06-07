
package main

func pop[T any](array *[]T)T{
    elem := (*array)[len(*array)-1]
    *array = (*array)[:len(*array)-1]
    return elem
}


