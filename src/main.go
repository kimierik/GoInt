package main


//lexer
import (
    "bufio"
    "fmt"
    "os"
)






func main(){

    fmt.Println("starting interpreter")

    //wait for user input 
    //lex it
    //parse it  or error
    //execute it  (send to interpreter) //user input could be listened in the interpreter
    

    //mem and stack as 2 different variables
    //mem has functions and variables
    //stack is used in ops



    
   intr:=Interpreter{memory:make(map[string]Variable), functions: make(map[string]Function) }
   intr.addCoreFns()

    for{
        fmt.Print(">")
        reader:= bufio.NewReader(os.Stdin)
        input, _ :=reader.ReadString('\n')
        tokens:=lexInput(input)
        //fmt.Println(tokens)
        stats:=Parse(tokens)
        //fmt.Println(stats)
        //the above does nothing

        //stats:=genTestParse()
        intr.InterpretAst(stats)


        //fmt.Println(tokenizer.tokens)
    }


}

