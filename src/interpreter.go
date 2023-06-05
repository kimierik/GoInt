package main

import "fmt"



type Variable interface{}

type Corefn func(a ...any)(n int, err error) //this can be changed to interface i think
 




type StackMember struct{}

type Interpreter struct{
    memory map[string]Variable          //all variables are stored here
    functions map[string]Function       //all function declaretions are stored here
    stack []StackMember                 //this is stackmachine so we need stack
}



//adds core fns
func (int*Interpreter) addCoreFns(){
    var lbod []Statement    
    var v Corefn
    v = fmt.Println
    lbod = append(lbod, v)

    int.functions["log"] = Function{body: lbod}
}



func (int*Interpreter) InterpretAst(stats []Statement){
    for i:=0; i<len(stats);i++{
        switch statement:= stats[i];statement.(type){
        case Corefn:
            fmt.Println("corefn type")
        case Function:
            fmt.Println("fun")
        case FuncCall:
            //now we need to figure out how do we membets from the funcall thing 
            str:=statement.(FuncCall).fName
            prms:=statement.(FuncCall).params
            //fmt.Println("funcall: ",str)
            callFuncion(int.functions[str], prms)


        default: 
            fmt.Println("interpre defaulted")
        }
    }

}



func callFuncion(fn Function, params []Expr){
    for i:=0; i<len(fn.body);i++{
        fn.body[i].(Corefn)(params)
        //prints  [{int}] 
        //[] since []expr 
        //{} since ilit struct
        //would need to spread params and then take the values themselves

    }
}




