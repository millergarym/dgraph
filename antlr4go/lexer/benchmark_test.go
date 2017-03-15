package lexer

import (
	"testing"

	"github.com/dgraph-io/dgraph/antlr4go/lexer/parser1"
	"github.com/dgraph-io/dgraph/antlr4go/lexer/parser2"
	antlr "github.com/antlr/antlr4/runtime/Go/antlr"
)

var q1 = `
{
	al(xid: "alice") {
		status
		_xid_
		follows {
			status
			_xid_
			follows {
				status
				_xid_
				follows {
					_xid_
					status
				}
			}
		}
		status
		_xid_
	}
}
`
var q2 = `query queryName {
		me(uid : "0x0a") {
			friends {
				name
			}
			gender,age
			hometown
		}
	}
`

var q3 = `
{
  debug(xid: "m.0bxtg") {
    type.object.name.en
    film.actor.film {
      film.performance.film {
        film.film.directed_by {
          type.object.name.en
        }
      }
    }
  }
}
`

var q4 = `
{
  debug(_xid_: "m.06pj8") {
    type.object.name.en
    film.director.film {
      type.object.name.en
      film.film.initial_release_date
      film.film.country
      film.film.starring {
        film.performance.actor {
          type.object.name.en
        }
        film.performance.character {
          type.object.name.en
        }
      }
      film.film.genre {
        type.object.name.en
      }
    }
  }
}
`
var q5 = `{
   debug(_xid_: "m.06pj8") {
    type.object.name.en
    film.director.film (first: "2", offset:"10") @filter(anyof("type.object.name.en" , "war spies")
                                                         || allof("type.object.name.en", "hello world")
                                                         || allof("type.object.name.en", "wonder land"))
    {
      _uid_
      type.object.name.en
      film.film.initial_release_date
      film.film.country
      film.film.starring {
        film.performance.actor @filter(anyof("type.object.name.en", "dicaprio matt")) {
          type.object.name.en
        }
        film.performance.character {
          type.object.name.en
        }
      }
      film.film.genre {
        type.object.name.en
      }
    }
  }
}`

func BenchmarkQuery(b *testing.B) {
	b.Run("q1.1", func(b *testing.B) { runLexer1(q1, b) })
	b.Run("q1.2", func(b *testing.B) { runLexer2(q1, b) })

	b.Run("q2.1", func(b *testing.B) { runLexer1(q2, b) })
	b.Run("q2.2", func(b *testing.B) { runLexer2(q2, b) })

	b.Run("q3.1", func(b *testing.B) { runLexer1(q3, b) })
	b.Run("q3.2", func(b *testing.B) { runLexer2(q3, b) })

	b.Run("q4.1", func(b *testing.B) { runLexer1(q4, b) })
	b.Run("q4.2", func(b *testing.B) { runLexer2(q4, b) })

	b.Run("q5.1", func(b *testing.B) { runLexer1(q4, b) })
	b.Run("q5.2", func(b *testing.B) { runLexer2(q4, b) })

}

func TestQuery(t *testing.T) {
	// for i := 0; i < 1000; i++ {
	input := antlr.NewInputStream(q4)
	lexer := parser2.NewGraphQLPMLexer2(input)
	// stream := antlr.NewCommonTokenStream(lexer, 0)
	to := lexer.NextToken()
	for to.GetTokenType() != antlr.TokenEOF {
		to = lexer.NextToken()
	}

	// }
}

func runLexer1(q string, b *testing.B) {
	// input := antlr.NewInputStream(q)
	// lexer := parser1.NewGraphQLPMLexer1(input)
	// // stream := antlr.NewCommonTokenStream(lexer, 0)
	// t := lexer.NextToken()
	// for t.GetTokenType() != antlr.TokenEOF {
	// 	t = lexer.NextToken()
	// 	fmt.Printf("%v\n", t)
	// }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := antlr.NewInputStream(q)
		lexer := parser1.NewGraphQLPMLexer1(input)
		// stream := antlr.NewCommonTokenStream(lexer, 0)
		t := lexer.NextToken()
		for t.GetTokenType() != antlr.TokenEOF {
			t = lexer.NextToken()
		}
		// p := NewGraphQLPMParser(stream)
		// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		// p.BuildParseTrees = true
		// uptill here we have a cost of : 19000 for q1
		// next call makes it 100 times more costly to : 1800000
		// _ = p.Document()
	}
}

func runLexer2(q string, b *testing.B) {
	// input := antlr.NewInputStream(q)
	// lexer := parser2.NewGraphQLPMLexer2(input)
	// // stream := antlr.NewCommonTokenStream(lexer, 0)
	// t := lexer.NextToken()
	// for t.GetTokenType() != antlr.TokenEOF {
	// 	t = lexer.NextToken()
	// 	fmt.Printf("%v\n", t)
	// }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := antlr.NewInputStream(q)
		lexer := parser2.NewGraphQLPMLexer2(input)
		// stream := antlr.NewCommonTokenStream(lexer, 0)
		t := lexer.NextToken()
		for t.GetTokenType() != antlr.TokenEOF {
			t = lexer.NextToken()
		}
		// p := NewGraphQLPMParser(stream)
		// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		// p.BuildParseTrees = true
		// // uptill here we have a cost of : 19000 for q1
		// // next call makes it 100 times more costly to : 1800000
		// _ = p.Document()
	}
}