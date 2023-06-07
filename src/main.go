package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)






func main(){


    
   intr:=Interpreter{memory:make(map[string]Variable), functions: make(map[string]Function) }
   intr.addCoreFns()

    if len(os.Args)>1{
        fmt.Println("reading from file")

        fp:=os.Args[1]
        data,err:=ioutil.ReadFile(fp)
        if err!=nil{
            log.Panic("cannot read file : ",fp)
        }

        tokens:=lexInput(string(data))
        //fmt.Println(tokens)

        parser:=Parser{tokens:tokens}
        stats:=parser.Parse()
        //fmt.Println(stats)
        //the above does nothing

        //stats:=genTestParse()
        intr.InterpretAst(stats)

    } else{
        cliInterpret(&intr)
    }




}




func cliInterpret(intr* Interpreter){

    fmt.Println("reading commands from stdin")

    for{
        fmt.Print(">")
        reader:= bufio.NewReader(os.Stdin)
        input, _ :=reader.ReadString('\n')
        tokens:=lexInput(input)

        parser:=Parser{tokens:tokens}
        stats:=parser.Parse()
        //fmt.Println(stats)

        intr.InterpretAst(stats)

    }
}

