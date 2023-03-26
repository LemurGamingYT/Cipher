grammar Cipher;

LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LLIST: '[';
RLIST: ']';
ASSIGN: '=';
DOT: '.';
COMMA: ',';
QUESTION: '?';
NOT: '!' | 'not';
PREDTWO: '+' | '-' | '^';
PREDONE: '*' | '/' | '%';
COMPARATIVE: '==' | '!=' | '>' | '<' | '<=' | '>=' | 'and' | '&&' | 'or' | '||';
FUNC: 'func';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
USE: 'use';
OVERRIDE: 'override';
PUBLIC: 'public';
PRIVATE: 'private';
RETURN: 'return';
BREAK: 'break';
CONTINUE: 'continue';
UNDEFINE: 'undefine';
CONST: 'const';
WS: [ \t\r\n]+ -> skip;
COMMENT: ('//' | '<--' | '#') ~[\r\n]* -> skip;
MULTILINECOMMENT: '/*' .*? '*/' -> skip;
BOOL: 'true' | 'false';
NULL: 'null';
ID: [a-zA-Z_] [a-zA-Z_0-9]*;
INT: '-'? [0-9]+;
FLOAT: '-'? [0-9]* '.' [0-9]+;
STRING: '"' ('\\' . | ~[\\"\r\n])* '"' | '\'' ('\\' . | ~[\\'\r\n])* '\'';

parse: stmt* EOF;

block: LBRACE stmt* RBRACE;

stmt: expr | assignments | allStmts | keywordStmts;

keywordStmts: BREAK | CONTINUE | RETURN expr;

allStmts: ifStmt | whileStmt | useStmt;

useList: STRING (COMMA STRING)*;
useStmt: USE useList;

ifStmt: IF condition block (ELSE IF condition block)* (ELSE block)?;

whileStmt: WHILE condition block;

condition: expr;

args: expr (COMMA expr)*;
params: ID (COMMA ID)*;

call: ID LPAREN args? RPAREN;

assignments: varAssign | funcAssign;

varAssign: (PUBLIC | PRIVATE)? CONST? ID ASSIGN expr;
funcAssign: (PUBLIC | PRIVATE)? FUNC OVERRIDE? ID LPAREN params? RPAREN block;

getAttributes: atom DOT ID LPAREN args? RPAREN;

expr: call | atom | LPAREN expr RPAREN | op=NOT expr | expr op=PREDONE expr | expr op=PREDTWO expr | expr op=COMPARATIVE expr | getAttributes;

array: LLIST args? RLIST;

atom: array | ID | INT | FLOAT | STRING | NULL | BOOL;