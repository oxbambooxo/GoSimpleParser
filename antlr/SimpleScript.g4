grammar SimpleScript;

import SimpleLexer;


programm: (intDeclaration | assignment | expression)+;

intDeclaration: Int Identifier (AssignmentOP additive)? SemiColon?;

assignment: Identifier AssignmentOP additive SemiColon?;

expression: additive SemiColon?;

additive: multiplicative (op=(Plus | Minus) multiplicative)*;

multiplicative: primary (op=(Star | Slash) primary)*;

primary: IntegerLiteral | Identifier | LeftParen expression RightParen;
