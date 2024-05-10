grammar Vinland;

Whitespace: [ \r\n\t]+ -> skip;
Identifier: [A-Za-z][A-Za-z0-9]*;

program: declList EOF;

declList: decl*;

decl
  : functionDefinition
  ;

stmtBlock: '{' stmtList '}';

stmtList: stmt*;

stmt
  : binding
  | expr
  | call
  ;

binding: Identifier '=' expr;

functionDefinition: 'fn' Identifier '(' paramList ')' stmtBlock;

paramList: (param (',' param)*)?;

param: Identifier;

expr
    : '(' expr ')'             # Parenthesized
    | expr op=('+'|'-') expr   # AddSub
    | expr op=('*'|'/') expr   # MulDiv
    | call                     # CallExpr
    | literal                  # Number
    ;

call: Identifier expr+;

literal
  : IntLiteral
  | FloatLiteral
  | BooleanLiteral
  | StringLiteral
  ;

IntLiteral
  : DecimalNumber
  | HexadecimalNumber
  | BinaryNumber
  ;
DecimalNumber: ('-')? [0-9]+;
HexadecimalNumber: '0x' [0-9A-F]+;
BinaryNumber: '0b' [01]+;

FloatLiteral: [0-9]+ '.' [0-9]* ;

StringLiteral: '"' StringElement* '"';

// Adapted from https://github.com/antlr/grammars-v4/blob/master/scala/Scala.g4
fragment StringElement
    : '\u0020'
    | '\u0021'
    | '\u0023' .. '\u007F'
    | CharEscapeSeq
    ;

CharEscapeSeq: '\\' ('t' | 'n' | 'r' | '\\' | '"');

BooleanLiteral
  : True
  | False
  ;
True: 'true';
False: 'false';