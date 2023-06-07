package main

//import "fmt"

import (
	"fmt"
	"strconv"
)

//what it does and what are the left and rigth ones
//l and r could be t  or a node


type Node               interface{ }
type Expr               interface{ }
type Statement          interface{ }


//expression
type Iliteral struct{
    val int   
}

type Stringliteral struct{
    val string   
}




type OpType int
const(
    Mul OpType =iota
    Add
    Div
    Sub
)

type Operaton struct{
    l Expr
    r Expr
    opp  string
}


//function
type Function struct{
    name string
    body []Statement
    params []Expr     //params thtat this accepts
    returns Expr
}


type VariableRefrence struct{
    name string
}

//statement & exression
type FuncCall struct{
    fName string
    params []Expr
}



type Parser struct{
    currentToken int
    tokens []Token
}


//should we do parser on class and follow current i etc


//should be used for files and lines
//and this does not have a main fn
func (self*Parser)Parse()[]Statement{
    var stats []Statement
     
    for self.currentToken<len(self.tokens){
        

        stat:=self.parseStatement()
        if stat==nil{
            break
        }
        stats = append(stats, stat)
        
        self.currentToken++
        //fmt.Println(stats)

    }

    return stats
}


func (self*Parser) parseStatement()Statement{
    //could be call or decl
    if self.tokens[self.currentToken].tokenType==Identifier && self.tokens[self.currentToken+1].tokenVal=="(" { 
        statement:=self.parseFunction()
        return statement
    }
    if self.tokens[self.currentToken].tokenType==EOF{
        return nil
    }

    fmt.Println("no statement parsed from : ", self.tokens[self.currentToken], self.currentToken)
    panic("no statement parsed")

}







//handles function declaradions and calls
func (self *Parser) parseFunction()Statement{
    var stat Statement
    
    // go untill we hit the ) then if next is { then its a declt

    a:=self.currentToken
    for (self.tokens)[a].tokenVal!=")"{
        a++
    }
    if (self.tokens)[a+1].tokenVal=="{"{
        //parse decl
        stat=self.parseFnDecl()
    }else{
        //parse call
        stat=self.parseFnCall()
    }
    

    return stat
}



//assumes current token is id
func (self*Parser) parseFnDecl()Function{
    decl:= Function{}
    decl.name= self.tokens[self.currentToken].tokenVal
    self.currentToken++ //skip id
    
    

    for self.tokens[self.currentToken].tokenVal!=")"{
        expression:=self.parseExpression()
        decl.params = append(decl.params, expression)
        self.currentToken++
    }
    self.currentToken++ //skip )
    self.currentToken++ //skip {


    //this does not stop the thing something jups over it
    for self.tokens[self.currentToken].tokenVal!="}"{
        stat:=self.parseStatement()
        decl.body = append(decl.body, stat)
    }

    //fmt.Println("fn decl name: ",decl.name,". params: ", decl.params, ". body: ",decl.body )

    return decl
}


//assumes current token is id
func (self*Parser) parseFnCall()FuncCall{
    var call FuncCall
    call.fName= (self.tokens)[self.currentToken].tokenVal
    self.currentToken=self.currentToken+2// jump fn fname and (

    call.params=self.parseFuncParameters()

    self.currentToken++//jump last )
    return call
}




//assumes the first token is ?? TODO figure it out
func (self*Parser) parseFuncParameters() []Expr{
    var expressions []Expr
    
    //buffer for expressions that we have yet to evaluate 
    //should we instead of []token do []Expr and we collapse this into a op node when we evaluate it


    //TODO this needs to change if we want funcs in expressions


    var toEvalueate []Token
    for self.tokens[self.currentToken].tokenVal!=")"{
        //needs to go through al tokens untill )
        if self.tokens[self.currentToken].tokenType!=Comma {
            toEvalueate = append(toEvalueate, self.tokens[self.currentToken])
            self.currentToken++ 
        }else{
            if len(toEvalueate)>0{
                expressions= append(expressions,self.EvaluateExpression(toEvalueate))
                toEvalueate = toEvalueate[:0] //clear list
            }
            self.currentToken++ 
        }

    }


    if len(toEvalueate)>0{
        expressions= append(expressions,self.EvaluateExpression(toEvalueate))
    }

    
    //fmt.Println("expressions of parse func parameters")
    //fmt.Println(expressions)
    return expressions
}


//gets a small list of tokens and evaluates it as an expression
//token list like [ 1, +,  1, *, func, (, ), /, 2 ]
//returns binary expression
func (self*Parser) EvaluateExpression(toEvalueate []Token)Expr{
    //fmt.Println("evaluationg")
    //fmt.Println(toEvalueate)

    if(len(toEvalueate)==0){
        return nil
    }

    if(len(toEvalueate)==1){
        switch tok:=toEvalueate[0];tok.tokenType{
        case IntLiteral:
            i,_:=strconv.ParseInt( tok.tokenVal,0,32)
            return Iliteral{int(i)}
        case StringLiteral:
            return Stringliteral{tok.tokenVal}
        }
    }

    postfix:=InfixToPostfix(toEvalueate)

    //fmt.Println("post fix notation of eval is")
    //fmt.Println(postfix)
    //var exprs []Token
    var expressions []Expr


    for i:=0;i<len(postfix);i++{
        if postfix[i].tokenType!=Operator{
            //switch and match what the token is?
            //right now all happens to be ints TODO change
            val,_:=strconv.ParseInt(  postfix[i].tokenVal,0,32)
            expressions = append(expressions, Iliteral{ int(val)})
        }else{
            fmt.Println(expressions)
            l:=pop(&expressions)
            r:=pop(&expressions)
            expressions=append(expressions,  Operaton{ l:l, r: r, opp: postfix[i].tokenVal})
        }
    }

    
    //fmt.Println("this is the expression tree")
    //fmt.Println(expressions)

    return expressions[0]
}

func getOperatorPrecidence(op string)int{
    switch op{

    case "*":
        return 2
    case "/":
        return 2

    case "+":
        return 1
    case "-":
        return 1

    default:
        fmt.Println("token default")
        return 0
    }

}


func InfixToPostfix(tokens []Token)[]Token{
    var postfix []Token
    var stack []Token

    for i:=0; i<len(tokens);i++{

        if tokens[i].tokenType!=Operator{
            postfix = append(postfix, tokens[i])
        }else{
            //the thing is a variable or literal
            //if the precidence of the current opp is more than the precidence of the opp on the top of the stack
            //push it 

            for len(stack)>0&& getOperatorPrecidence(tokens[i].tokenVal) <= getOperatorPrecidence( stack[len(stack)-1].tokenVal ) {
                postfix = append(postfix, pop(&stack))
            }
            stack = append(stack, tokens[i])

        }

    }
    for len(stack)>0{
        postfix = append(postfix, pop(&stack))
    }



    return postfix
}



//does parses expression from current token
//currently only does expressions in function parameters
func (self*Parser) parseExpression()Expr{

    if (self.tokens[self.currentToken].tokenType==Comma){
        self.currentToken++
    }

    if (self.tokens[self.currentToken].tokenVal=="(" && self.tokens[self.currentToken+1].tokenVal==")" ){ 
        //empty parameters
        return nil
    }

    //this is a single thing and does not need to be an oper expression
    if (self.tokens[self.currentToken+1].tokenType==Comma || self.tokens[self.currentToken+1].tokenVal==")" ){ 


        if (self.tokens[self.currentToken].tokenType==IntLiteral){
            i,_:=strconv.ParseInt(self.tokens[self.currentToken].tokenVal,0,32)
            return Iliteral{int(i)}
        }


        if (self.tokens[self.currentToken].tokenType==StringLiteral){
            return Stringliteral{self.tokens[self.currentToken].tokenVal}
        }



        //expression is variable
        if (self.tokens[self.currentToken].tokenType==Identifier){ 
            return VariableRefrence{self.tokens[self.currentToken].tokenVal}
        }

    }
    //function call
    if (self.tokens[self.currentToken].tokenType==Identifier && self.tokens[self.currentToken+1].tokenVal=="(" ){
        return self.parseFnCall()
    }



    return nil
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


