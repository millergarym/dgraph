package lexer

//go:generate java -jar $GOPATH/src/github.com/millergarym/antlr4-go/lib/antlr4-4.6-SNAPSHOT.jar -package parser1 -o parser1 -visitor -Dlanguage=Go GraphQLPMLexer1.g4

//go:generate java -jar $GOPATH/src/github.com/millergarym/antlr4-go/lib/antlr4-4.6-SNAPSHOT.jar -package parser1 -o parser1 -visitor -Dlanguage=Go GraphQLPMLexer1.g4
//go:generate sh -c "(cd parser1; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/millergarym/antlr4-go!' *.go )"

//go:generate java -jar $GOPATH/src/github.com/millergarym/antlr4-go/lib/antlr4-4.6-SNAPSHOT.jar -package parser2 -o parser2 -visitor -Dlanguage=Go GraphQLPMLexer2.g4
//go:generate sh -c "(cd parser2; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/millergarym/antlr4-go!' *.go )"
