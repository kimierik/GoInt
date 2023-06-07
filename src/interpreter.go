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
    stack []Expr                 //this is stackmachine so we need stack
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
            str:=statement.(FuncCall).fName
            prms:=statement.(FuncCall).params
            //fmt.Println("funcall: ",str)
            int.callFuncion(int.functions[str], prms)


        default: 
            fmt.Println("interpre defaulted")
        }
    }

}


//pushes it onto the int.stack?
// TODO rn the stack is filled with random ints and strings this could be turned into tokens or somethign
//maybe make the stack member some actual value that can be used here
func (int* Interpreter) solveOperation(opp Operaton){

    switch opp.l.(type){
    case Operaton:
        int.solveOperation(opp.l.(Operaton))
    case Iliteral:
        int.stack = append(int.stack,opp.l.(Iliteral).val )
    case Token:
        //fmt.Println("token")

    default:
        //fmt.Println("opp.l defaulted")
        //fmt.Println(opp.l)
    }

    switch opp.r.(type){
    case Operaton:
        int.solveOperation(opp.r.(Operaton))
    case Iliteral:
        int.stack = append(int.stack,opp.r.(Iliteral).val )
    }

    int.stack = append(int.stack, opp.opp)
}




//TODO make this into a real function 
func solveOpp(l int, r int ,opp string)int{

    switch opp{
    case "+":
        return l+r;
    case "*":
        return l*r;
    case "-":
        return l-r;
    case "/":
        return l/r;
    default:
        return 0
    }

}

//TODO
//should we have a function for just resolving the stack
//we could then have var refrences and fuinction calls in order etc
func (intr * Interpreter) resolveOppTree(opp Operaton)Iliteral{
    //we have a tree and we will start calculationg how the fuck that is calculated
    //we need to convert this into postfix agains?

    intr.solveOperation(opp)

    //fmt.Println("solve")
    //fmt.Println(int.stack)

    //solve the postfix notation
    var rStack []Expr

    //i quess push all lits on to a stack iteratively then when we hit opp we pop 2 and calculate
    for i:=0;i<len( intr.stack);i++{ 
        switch intr.stack[i].(type){
        case int:
            rStack = append(rStack, intr.stack[i].(int))
        case string: //this is operator
            l:= pop(&rStack)      
            r:= pop(&rStack)      
            rStack = append(rStack, solveOpp(l.(int),r.(int) , intr.stack[i].(string) ))

        default:
            fmt.Println("resolving postfix defaulted")

        }

    }



    return Iliteral{rStack[0].(int)}
}




func (int*Interpreter)callFuncion(fn Function, params []Expr){

    var ResolvedParameters []Expr
    //its a list?

    for i:=0; i<len(params);i++{
        switch param:=params[i];param.(type){
        case Operaton:
            ResolvedParameters = append(ResolvedParameters, int.resolveOppTree(params[i].(Operaton)).val)

        case Iliteral:
            ResolvedParameters = append(ResolvedParameters,param.(Iliteral).val )
        default:
            //fmt.Println("call fn defaulted params[i] is ")
            //fmt.Println(params[i])
            ResolvedParameters = append(ResolvedParameters, params[i])
        }

    }

    //we need to resolve expressions before calling

    //fmt.Println("called: ",fn.name, "with params: ", fn.params, "has body : ",fn.body)   
    //loop statements in body
    for i:=0; i<len(fn.body);i++{

        switch fun:=fn.body[i]; fun.(type){
            case Corefn:
                fn.body[i].(Corefn)(ResolvedParameters)

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




