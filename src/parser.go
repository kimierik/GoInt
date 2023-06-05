package main

//import "fmt"

//what it does and what are the left and rigth ones
//l and r could be t  or a node


type Node interface{ }
type Expr interface{ }
type Statement interface{ }


//expression
type Iliteral struct{
    val int   
}




//function
type Function struct{
    body []Statement
    params []Expr
    returns Expr
}




//statement & exression
type FuncCall struct{
    fName string
    params []Expr
}



//lets assume the tokens are
// [log, ( , 1 , ), ]


//should we do parser on class and follow current i etc


//should be used for files and lines
//and this does not have a main fn
//so what we need to be able to do
func Parse(tokens []Token)[]Statement{
    var stats []Statement
     
    i:=0
    for i<len(tokens){
        
        //could be call or decl
        if tokens[i].tokenType==Identifier && tokens[i+1].tokenVal=="(" { 
            statement:=parseFunction(&tokens,&i)
            stats = append(stats, statement)
        }

        
        i++

    }

    return stats
}






//handles function declaradions and calls
func parseFunction(tokens *[]Token, i *int)Statement{
    var stat Statement
    
    // go untill we hit the ) then if next is { then its a declt

    a:=*i
    for (*tokens)[a].tokenVal!=")"{
        a++
    }
    if (*tokens)[a+1].tokenType==Curly{
        //parse decl
    }else{
        //parse call
        stat=parseFnCall(tokens,i)
    }
    

    return stat
}



func parseFnCall(tokens* []Token, i*int)FuncCall{
    var call FuncCall
    call.fName= (*tokens)[*i].tokenVal
    var params []Expr
    *i=*i+2// jump fn fname and (

    for (*tokens)[*i].tokenVal!=")"{
        params = append(params, (*tokens)[*i].tokenVal)
        *i++
    }

    call.params=params

    return call
}




//test funddtion on manually generating ast
func genTestParse()[]Statement{
    var stats []Statement
    var thing  Statement
    var parms []Expr
    parms = append(parms, Iliteral{1})
    thing = FuncCall{fName: "log",params:parms }

    stats = append(stats, thing)
    
    return stats
}


