grammar Cipher;

ASSIGN: '=';
DOT: '.';
COMMA: ',';
QUESTION: '?';
NOT: '!' | 'not';
PREDTWO: ('+' | '-' | '^');
PREDONE: ('*' | '/' | '%');
COMPARATIVE: ('==' | '!=' | '>' | '<' | '<=' | '>=' | '&&' | '||');
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
CONST: 'const';
WS: [ \t\r\n]+ -> skip;
COMMENT: ('//' | '<--' | '#') ~[\r\n]* -> skip;
MULTILINECOMMENT: '/*' .*? '*/' -> skip;
BOOL: 'true' | 'false';
NULL: 'null';
ID: [a-zA-Z_] [a-zA-Z_0-9]*;
INT: '-'? [0-9]+;
FLOAT: '-'? [0-9]* '.' [0-9]+;
STRING: '"' .*? '"' | '\'' .*? '\'';

parse: stmt* EOF;

block: '{' stmt* '}';

stmt: expr | assignments | allStmts | keywordStmts;

keywordStmts: BREAK | CONTINUE | RETURN expr;

allStmts: ifStmt | whileStmt | useStmt;

useStmt: 'use' STRING;

ifStmt: IF condition block (ELSE IF condition block)* (ELSE block)?;

whileStmt: 'while' condition block;

condition: expr;

args: expr (',' expr)*;
params: ID (',' ID)*;

call: ID '(' args? ')';

assignments: varAssign | funcAssign;

varAssign: (PUBLIC | PRIVATE)? CONST? ID (PREDONE | PREDTWO)? ASSIGN expr;
funcAssign: (PUBLIC | PRIVATE)? 'func' OVERRIDE? ID '(' params? ')' block;

getAttributes: (STRING | ID) '.' ID '(' args? ')';

memoryAddress: '&' ID;

cast: atom '.' '(' ID ')';

expr: call | atom | '(' expr ')' | op=NOT expr | expr op=PREDONE expr | expr op=PREDTWO expr | expr op=COMPARATIVE expr | cast | getAttributes | memoryAddress;

array: '[' args? ']';

atom: array | ID | INT | FLOAT | STRING | NULL | BOOL;