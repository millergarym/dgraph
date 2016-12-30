lexer grammar GraphQLPMLexer2;

FILTER      : '@filter';
LP          : '(';
RP          : ')';
OR          : '||';
AND         : '&&';
ANYOF       : 'anyof';
ALLOF       : 'allof';
COMMA       : ','   ;
COLON       : ':' ;
QUERY       : 'query';
LC          : '{';
RC          : '}';
NAME        : [_A-Za-z] [._0-9A-Za-z]*   ;
STRING      : '"' ~["]* '"'   ;
WS          : [ \t\n\r]+ -> skip;
