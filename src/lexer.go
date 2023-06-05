package main

import (
	"fmt"
	"go/doc/comment"
	"os"
)



type Lexer struct{
    thing int
    tokens []Token
}


type TokenType int

const (
    Identifier TokenType = iota
    Operator
    StringLiteral
    IntLiteral

    Paren
    Curly
    Comma
)


//token
type Token struct{
    tokenType TokenType
    tokenVal string
}


// this is how you do functions in structs etc
//this would be a copy not a ref
//func (self Lexer ) hehe(){ }




func lexInput(input string)[]Token{
    var tokens []Token

    for i:=0; i<len(input);{
        
        //used for storing literals and identifiers
        char:=input[i]


        //do we just drop the switch fo whiles etc
        if char>='0'&&char<='9'{
            lit:=handleIntLiteral(&i,&input)
            tokens = append(tokens, Token{IntLiteral,lit})
            continue
        }

        //is identifier uses the same function as StringLiteral for simplicity
        if char>='a'&&char<='z' || char>='A'&&char<='Z'{
            id:=handleId(&i,&input)
            tokens = append(tokens, Token{Identifier,id})
            continue
        }


        //string literal
        if char=='"'{
            i++; //skip quote
            str:=handleStringLiteral(&i,&input)
            tokens = append(tokens, Token{StringLiteral,str})
            i++;//skip last quote
            continue
        }



        if char =='(' || char==')'{
            i++; 
            tokens = append(tokens, Token{Paren,string(char)})
            continue
        }

        if char =='{' || char=='}'{
            i++; 
            tokens = append(tokens, Token{Curly,string(char)})
            continue
        }


        if char ==',' {
            i++; 
            tokens = append(tokens, Token{Comma,string(char)})
            continue
        }


        //if whitespace 
        if char==' ' || char=='\n' ||char=='\t'{
            i++;
            continue
        }





        //if we do not know what the charactter is
        fmt.Println("lexing error unidentified character : ",char)
        os.Exit(1)//maybe we dont need to exit just discard tokenistaion and return to main function


    }
    return tokens

}




func handleIntLiteral(index *int, input *string) string{
    var literalbuffer []byte

    for (*input)[*index] >='0' && (*input)[*index]<='9'{
        literalbuffer=append(literalbuffer,(*input)[*index])
        (*index)++
    }
    return string(literalbuffer[:])
}

func handleId(index *int, input *string) string{
    var literalbuffer []byte

    for (*input)[*index] >='a' && (*input)[*index]<='z' || (*input)[*index] >='A' && (*input)[*index]<='Z' {
        literalbuffer=append(literalbuffer,(*input)[*index])
        (*index)++
    }
    return string(literalbuffer[:])
}

func handleStringLiteral(index *int, input *string) string{
    var literalbuffer []byte
    //untill

    for (*input)[*index]!= '"'{
        literalbuffer=append(literalbuffer,(*input)[*index])
        (*index)++
    }
    return string(literalbuffer[:])
}


