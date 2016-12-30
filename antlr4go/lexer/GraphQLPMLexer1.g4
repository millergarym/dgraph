lexer grammar GraphQLPMLexer1;

AT          : '@';
//FILTER : 'f' 'i' 'l' 't' 'e' 'r';
LP          : '(';
RP          : ')';
OR          : '||';
AND         : '&&';
//   : 'anyof' | 'allof'
COMMA       : ','   ;
COLON       : ':' ;
//   : 'query'
LC          : '{';
RC          : '}';
NAME        : [_A-Za-z] [._0-9A-Za-z]*   ;
STRING      : '"' ~["]* '"'   ;
WS          : [ \t\n\r]+ -> skip;
