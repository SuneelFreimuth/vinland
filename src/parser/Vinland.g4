grammar Vinland;

Whitespace: [ \r\n\t]+ -> skip;
Identifier: [A-Za-z][A-Za-z0-9]*;

program: declList EOF;

declList: decl*;

decl
  : functionDefinition
  ;

Comment
  : '#' ~[\r\n]* -> skip
  ;

stmtBlock: '{' stmtList '}';

stmtList: (stmt ';')* expr0?;

stmt
  : binding
  | expr0
  ;

binding: Identifier '=' expr0;

functionDefinition: 'fn' Identifier '(' paramList ')' stmtBlock;

paramList: (param (',' param)*)?;

param: Identifier;

// Operators and expression precedence taken from Crux, a language for teaching 
// compilers and interpreters by Prof. Brian Demsky, UC Irvine.
op0
  : '>='
  | '<='
  | '!='
  | '=='
  | '>'
  | '<'
  ;
op1
  : '+'
  | '-'
  | '||'
  ;
op2
  : '*'
  | '/'
  | '&&'
  ;

expr0
  : expr1
  | ifExpr;
expr1
  : expr2 (op0 expr2)?
  ;
expr2
  : expr3 op1 expr3
  | expr3
  ;
expr3
  : expr4
  | expr3 op2 expr4
  ;
expr4
  : '!' expr4
  | '(' expr0 ')'
  | name
  | callExpr
  | literal
  ;

name: Identifier;
callExpr: Identifier expr0+;
ifExpr: 'if' expr0 'then' expr0 'else' expr0;

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