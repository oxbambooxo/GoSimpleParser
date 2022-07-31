grammar SimpleScript;

import SimpleLexer;


programm
    : statement
    | (statement ';')+
    ;

statement
    : intDeclaration
    | assignment
    | expression
    ;

intDeclaration
    : 'int' Identifier ('=' additive)?
    ;

assignment
    : Identifier '=' additive
    ;

expression
    : additive
    ;

additive
    : multiplicative
    | additive op=('+' | '-') multiplicative
    ;

multiplicative
    : primary
    | multiplicative op=('*' | '/') primary
    ;

primary
    : IntegerLiteral
    | Identifier
    | '(' expression ')'
    ;
