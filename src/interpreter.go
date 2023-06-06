package main

import (
	"fmt"
)



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

    int.functions["log"] = Function{body: lbod,name: "log"}
}



func (int*Interpreter) InterpretAst(stats []Statement){
    for i:=0; i<len(stats);i++{
        switch statement:= stats[i];statement.(type){
        case Corefn:
            fmt.Println("corefn type")

            //function statement
        case Function:
            fn:=statement.(Function)
            //fmt.Println("declaring fn: ",fn.name, ". params: ",fn.params, ". with body: ",fn.body)
            int.functions[fn.name]=fn
            //fmt.Println(int.functions)

        case FuncCall:
            //now we need to figure out how do we membets from the funcall thing 
            str:=statement.(FuncCall).fName
            prms:=statement.(FuncCall).params
            //fmt.Println("funcall: ",str)
            int.callFuncion(int.functions[str], prms)


        default: 
            fmt.Println("interpre defaulted")
        }
    }

}



func (int*Interpreter)callFuncion(fn Function, params []Expr){

    //fmt.Println("called: ",fn.name, "with params: ", fn.params, "has body : ",fn.body)   
    //loop statements in body
    for i:=0; i<len(fn.body);i++{

        switch fun:=fn.body[i]; fun.(type){
            case Corefn:
                fn.body[i].(Corefn)(params)

            //if function call 
            case FuncCall: 
                a:=int.functions[fun.(FuncCall).fName]
                //fmt.Println("a.params: ",fun.(FuncCall).params) 
                int.callFuncion(a,fun.(FuncCall).params)
        }
        //prints  [{int}] 
        //[] since []expr 
        //{} since ilit struct
        //would need to spread params and then take the values themselves
    }
}




