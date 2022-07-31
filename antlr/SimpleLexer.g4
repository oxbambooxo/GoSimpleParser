lexer grammar SimpleLexer;


//关键字
Int:                'int';

//常量
IntegerLiteral:     [0-9]+;

//操作符
AssignmentOP:       '=' ;
Plus:               '+';
Minus:              '-';
Star:               '*';
Slash:              '/';
Comma:              ',';
SemiColon:          ';';
LeftParen:          '(';
RightParen:         ')';

//标识符
Identifier:         [a-zA-Z_] ([a-zA-Z_] | [0-9])*;

//空白字符，抛弃
Whitespace:         [ \t]+ -> skip;
Newline:            ( '\r' '\n'?|'\n')-> skip;