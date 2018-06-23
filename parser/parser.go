// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import "github.com/tesujiro/exp_yacc/ast"

type yySymType struct {
	yys   int
	token ast.Token
	expr  ast.Expr
	stmt  ast.Stmt
	stmts []ast.Stmt
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57348
	yyEofCode = 57344
	IDENT     = 57346
	NUMBER    = 57347
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -18
)

var (
	yyPrec = map[int]int{
		';':   0,
		'=':   1,
		IDENT: 2,
		'+':   3,
		'-':   3,
	}

	yyXLAT = map[int]int{
		57344: 0,  // $end (18x)
		10:    1,  // '\n' (16x)
		57346: 2,  // IDENT (13x)
		57347: 3,  // NUMBER (13x)
		59:    4,  // ';' (10x)
		43:    5,  // '+' (6x)
		45:    6,  // '-' (6x)
		61:    7,  // '=' (5x)
		57349: 8,  // expr (5x)
		57350: 9,  // newLine (5x)
		57351: 10, // newLines (3x)
		57352: 11, // opt_term (2x)
		57354: 12, // stmt (2x)
		57356: 13, // term (2x)
		57353: 14, // program (1x)
		57355: 15, // stmts (1x)
		57348: 16, // $default (0x)
		57345: 17, // error (0x)
	}

	yySymNames = []string{
		"$end",
		"'\\n'",
		"IDENT",
		"NUMBER",
		"';'",
		"'+'",
		"'-'",
		"'='",
		"expr",
		"newLine",
		"newLines",
		"opt_term",
		"stmt",
		"term",
		"program",
		"stmts",
		"$default",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {14, 2},
		2:  {15, 2},
		3:  {15, 3},
		4:  {12, 3},
		5:  {12, 1},
		6:  {8, 1},
		7:  {8, 1},
		8:  {8, 3},
		9:  {8, 3},
		10: {11, 0},
		11: {11, 1},
		12: {13, 2},
		13: {13, 1},
		14: {13, 1},
		15: {10, 1},
		16: {10, 2},
		17: {9, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [24][]uint8{
		// 0
		{1: 26, 8, 8, 23, 9: 25, 24, 21, 13: 22, 19, 20},
		{18},
		{8, 26, 4: 23, 9: 25, 24, 39, 13: 40},
		{2: 31, 32, 8: 30, 12: 29},
		{2: 7, 7},
		// 5
		{4, 26, 4, 4, 9: 25, 28},
		{5, 26, 5, 5, 9: 27},
		{3, 3, 3, 3},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		// 10
		{6, 26, 6, 6, 9: 27},
		{16, 16, 4: 16},
		{13, 13, 4: 13, 34, 35, 33},
		{12, 12, 4: 12, 12, 12, 12},
		{11, 11, 4: 11, 11, 11, 11},
		// 15
		{2: 31, 32, 8: 38},
		{2: 31, 32, 8: 37},
		{2: 31, 32, 8: 36},
		{9, 9, 4: 9, 9, 9, 9},
		{10, 10, 4: 10, 10, 10, 10},
		// 20
		{14, 14, 4: 14, 34, 35},
		{17},
		{7, 2: 31, 32, 8: 30, 12: 41},
		{15, 15, 4: 15},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 17

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yyVAL.stmts = yyS[yypt-1].stmts
			yylex.(*Lexer).Result = yyVAL.stmts
		}
	case 2:
		{
			yyVAL.stmts = []ast.Stmt{yyS[yypt-0].stmt}
		}
	case 3:
		{
			yyVAL.stmts = append(yyS[yypt-2].stmts, yyS[yypt-0].stmt)
		}
	case 4:
		{
			yyVAL.stmt = ast.AssStmt{Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 5:
		{
			yyVAL.stmt = ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 6:
		{
			yyVAL.expr = ast.IdentExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 7:
		{
			yyVAL.expr = ast.NumExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 8:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: '+', Right: yyS[yypt-0].expr}
		}
	case 9:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: '-', Right: yyS[yypt-0].expr}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
