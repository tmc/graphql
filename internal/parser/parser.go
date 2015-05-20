package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/tmc/graphql"
)

// helpers
func ifs(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 15, col: 1, offset: 179},
			expr: &actionExpr{
				pos: position{line: 15, col: 12, offset: 192},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 15, col: 12, offset: 192},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 15, col: 12, offset: 192},
							label: "stmts",
							expr: &oneOrMoreExpr{
								pos: position{line: 15, col: 18, offset: 198},
								expr: &ruleRefExpr{
									pos:  position{line: 15, col: 18, offset: 198},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 15, col: 29, offset: 209},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 41, col: 1, offset: 1081},
			expr: &actionExpr{
				pos: position{line: 41, col: 13, offset: 1095},
				run: (*parser).callonStatement1,
				expr: &seqExpr{
					pos: position{line: 41, col: 13, offset: 1095},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 41, col: 13, offset: 1095},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 41, col: 15, offset: 1097},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 41, col: 18, offset: 1100},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 41, col: 18, offset: 1100},
										name: "Operation",
									},
									&ruleRefExpr{
										pos:  position{line: 41, col: 30, offset: 1112},
										name: "FragmentDefinition",
									},
									&ruleRefExpr{
										pos:  position{line: 41, col: 51, offset: 1133},
										name: "TypeDefinition",
									},
									&ruleRefExpr{
										pos:  position{line: 41, col: 68, offset: 1150},
										name: "TypeExtension",
									},
									&ruleRefExpr{
										pos:  position{line: 41, col: 84, offset: 1166},
										name: "EnumDefinition",
									},
									&ruleRefExpr{
										pos:  position{line: 41, col: 101, offset: 1183},
										name: "Comment",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 41, col: 110, offset: 1192},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 45, col: 1, offset: 1214},
			expr: &actionExpr{
				pos: position{line: 45, col: 11, offset: 1226},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 45, col: 11, offset: 1226},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 45, col: 11, offset: 1226},
							val:        "#",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 45, col: 15, offset: 1230},
							expr: &charClassMatcher{
								pos:        position{line: 45, col: 15, offset: 1230},
								val:        "[^\\n]",
								chars:      []rune{'\n'},
								ignoreCase: false,
								inverted:   true,
							},
						},
					},
				},
			},
		},
		{
			name: "Operation",
			pos:  position{line: 47, col: 1, offset: 1269},
			expr: &choiceExpr{
				pos: position{line: 47, col: 13, offset: 1283},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 47, col: 13, offset: 1283},
						run: (*parser).callonOperation2,
						expr: &labeledExpr{
							pos:   position{line: 47, col: 13, offset: 1283},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 18, offset: 1288},
								name: "Selections",
							},
						},
					},
					&actionExpr{
						pos: position{line: 54, col: 13, offset: 1427},
						run: (*parser).callonOperation5,
						expr: &seqExpr{
							pos: position{line: 54, col: 14, offset: 1428},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 54, col: 14, offset: 1428},
									label: "ot",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 17, offset: 1431},
										name: "OperationType",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 31, offset: 1445},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 33, offset: 1447},
									label: "on",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 36, offset: 1450},
										name: "OperationName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 50, offset: 1464},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 52, offset: 1466},
									label: "vds",
									expr: &zeroOrOneExpr{
										pos: position{line: 54, col: 56, offset: 1470},
										expr: &ruleRefExpr{
											pos:  position{line: 54, col: 56, offset: 1470},
											name: "VariableDefinitions",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 77, offset: 1491},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 79, offset: 1493},
									label: "ds",
									expr: &zeroOrOneExpr{
										pos: position{line: 54, col: 82, offset: 1496},
										expr: &ruleRefExpr{
											pos:  position{line: 54, col: 82, offset: 1496},
											name: "Directives",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 54, col: 94, offset: 1508},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 54, col: 96, offset: 1510},
									label: "sels",
									expr: &ruleRefExpr{
										pos:  position{line: 54, col: 101, offset: 1515},
										name: "Selections",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "OperationType",
			pos:  position{line: 74, col: 1, offset: 1931},
			expr: &choiceExpr{
				pos: position{line: 74, col: 17, offset: 1949},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 74, col: 17, offset: 1949},
						run: (*parser).callonOperationType2,
						expr: &litMatcher{
							pos:        position{line: 74, col: 17, offset: 1949},
							val:        "query",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 75, col: 18, offset: 2013},
						run: (*parser).callonOperationType4,
						expr: &litMatcher{
							pos:        position{line: 75, col: 18, offset: 2013},
							val:        "mutation",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperationName",
			pos:  position{line: 76, col: 1, offset: 2066},
			expr: &actionExpr{
				pos: position{line: 76, col: 17, offset: 2084},
				run: (*parser).callonOperationName1,
				expr: &ruleRefExpr{
					pos:  position{line: 76, col: 17, offset: 2084},
					name: "Name",
				},
			},
		},
		{
			name: "VariableDefinitions",
			pos:  position{line: 79, col: 1, offset: 2121},
			expr: &actionExpr{
				pos: position{line: 79, col: 23, offset: 2145},
				run: (*parser).callonVariableDefinitions1,
				expr: &seqExpr{
					pos: position{line: 79, col: 23, offset: 2145},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 79, col: 23, offset: 2145},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 79, col: 27, offset: 2149},
							label: "vds",
							expr: &oneOrMoreExpr{
								pos: position{line: 79, col: 31, offset: 2153},
								expr: &ruleRefExpr{
									pos:  position{line: 79, col: 31, offset: 2153},
									name: "VariableDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 79, col: 51, offset: 2173},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariableDefinition",
			pos:  position{line: 86, col: 1, offset: 2334},
			expr: &actionExpr{
				pos: position{line: 86, col: 22, offset: 2357},
				run: (*parser).callonVariableDefinition1,
				expr: &seqExpr{
					pos: position{line: 86, col: 22, offset: 2357},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 86, col: 22, offset: 2357},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 86, col: 24, offset: 2359},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 86, col: 26, offset: 2361},
								name: "Variable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 35, offset: 2370},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 86, col: 37, offset: 2372},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 41, offset: 2376},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 86, col: 43, offset: 2378},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 86, col: 45, offset: 2380},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 50, offset: 2385},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 86, col: 52, offset: 2387},
							label: "d",
							expr: &zeroOrOneExpr{
								pos: position{line: 86, col: 54, offset: 2389},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 54, offset: 2389},
									name: "DefaultValue",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 68, offset: 2403},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DefaultValue",
			pos:  position{line: 98, col: 1, offset: 2639},
			expr: &actionExpr{
				pos: position{line: 98, col: 16, offset: 2656},
				run: (*parser).callonDefaultValue1,
				expr: &seqExpr{
					pos: position{line: 98, col: 16, offset: 2656},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 98, col: 16, offset: 2656},
							val:        "=",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 98, col: 20, offset: 2660},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 22, offset: 2662},
								name: "Value",
							},
						},
					},
				},
			},
		},
		{
			name: "Selections",
			pos:  position{line: 100, col: 1, offset: 2687},
			expr: &actionExpr{
				pos: position{line: 100, col: 14, offset: 2702},
				run: (*parser).callonSelections1,
				expr: &seqExpr{
					pos: position{line: 100, col: 14, offset: 2702},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 100, col: 14, offset: 2702},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 100, col: 18, offset: 2706},
							label: "s",
							expr: &oneOrMoreExpr{
								pos: position{line: 100, col: 21, offset: 2709},
								expr: &ruleRefExpr{
									pos:  position{line: 100, col: 21, offset: 2709},
									name: "Selection",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 100, col: 33, offset: 2721},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Selection",
			pos:  position{line: 111, col: 1, offset: 2986},
			expr: &choiceExpr{
				pos: position{line: 111, col: 13, offset: 3000},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 111, col: 13, offset: 3000},
						run: (*parser).callonSelection2,
						expr: &seqExpr{
							pos: position{line: 111, col: 14, offset: 3001},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 111, col: 14, offset: 3001},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 111, col: 16, offset: 3003},
									label: "f",
									expr: &ruleRefExpr{
										pos:  position{line: 111, col: 18, offset: 3005},
										name: "Field",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 111, col: 24, offset: 3011},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 114, col: 5, offset: 3093},
						run: (*parser).callonSelection8,
						expr: &seqExpr{
							pos: position{line: 114, col: 6, offset: 3094},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 114, col: 6, offset: 3094},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 114, col: 8, offset: 3096},
									label: "fs",
									expr: &ruleRefExpr{
										pos:  position{line: 114, col: 11, offset: 3099},
										name: "FragmentSpread",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 114, col: 26, offset: 3114},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Field",
			pos:  position{line: 119, col: 1, offset: 3232},
			expr: &actionExpr{
				pos: position{line: 119, col: 9, offset: 3242},
				run: (*parser).callonField1,
				expr: &seqExpr{
					pos: position{line: 119, col: 9, offset: 3242},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 119, col: 9, offset: 3242},
							label: "fa",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 12, offset: 3245},
								expr: &ruleRefExpr{
									pos:  position{line: 119, col: 12, offset: 3245},
									name: "FieldAlias",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 24, offset: 3257},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 26, offset: 3259},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 119, col: 29, offset: 3262},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 39, offset: 3272},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 41, offset: 3274},
							label: "as",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 44, offset: 3277},
								expr: &ruleRefExpr{
									pos:  position{line: 119, col: 44, offset: 3277},
									name: "Arguments",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 55, offset: 3288},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 57, offset: 3290},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 60, offset: 3293},
								expr: &ruleRefExpr{
									pos:  position{line: 119, col: 60, offset: 3293},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 72, offset: 3305},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 74, offset: 3307},
							label: "sels",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 79, offset: 3312},
								expr: &ruleRefExpr{
									pos:  position{line: 119, col: 79, offset: 3312},
									name: "Selections",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FieldAlias",
			pos:  position{line: 146, col: 1, offset: 3829},
			expr: &actionExpr{
				pos: position{line: 146, col: 14, offset: 3844},
				run: (*parser).callonFieldAlias1,
				expr: &seqExpr{
					pos: position{line: 146, col: 14, offset: 3844},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 146, col: 14, offset: 3844},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 16, offset: 3846},
								name: "Name",
							},
						},
						&litMatcher{
							pos:        position{line: 146, col: 21, offset: 3851},
							val:        ":",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "FieldName",
			pos:  position{line: 147, col: 1, offset: 3873},
			expr: &ruleRefExpr{
				pos:  position{line: 147, col: 13, offset: 3887},
				name: "Name",
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 148, col: 1, offset: 3892},
			expr: &actionExpr{
				pos: position{line: 148, col: 13, offset: 3906},
				run: (*parser).callonArguments1,
				expr: &seqExpr{
					pos: position{line: 148, col: 13, offset: 3906},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 148, col: 13, offset: 3906},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 148, col: 17, offset: 3910},
							label: "args",
							expr: &zeroOrMoreExpr{
								pos: position{line: 148, col: 23, offset: 3916},
								expr: &ruleRefExpr{
									pos:  position{line: 148, col: 23, offset: 3916},
									name: "Argument",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 148, col: 34, offset: 3927},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Argument",
			pos:  position{line: 159, col: 1, offset: 4172},
			expr: &actionExpr{
				pos: position{line: 159, col: 12, offset: 4185},
				run: (*parser).callonArgument1,
				expr: &seqExpr{
					pos: position{line: 159, col: 12, offset: 4185},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 159, col: 12, offset: 4185},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 14, offset: 4187},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 17, offset: 4190},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 30, offset: 4203},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 32, offset: 4205},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 36, offset: 4209},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 38, offset: 4211},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 40, offset: 4213},
								name: "Value",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 46, offset: 4219},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentName",
			pos:  position{line: 165, col: 1, offset: 4292},
			expr: &ruleRefExpr{
				pos:  position{line: 165, col: 16, offset: 4309},
				name: "Name",
			},
		},
		{
			name: "Name",
			pos:  position{line: 167, col: 1, offset: 4315},
			expr: &actionExpr{
				pos: position{line: 167, col: 8, offset: 4324},
				run: (*parser).callonName1,
				expr: &seqExpr{
					pos: position{line: 167, col: 8, offset: 4324},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 167, col: 8, offset: 4324},
							val:        "[a-z_]i",
							chars:      []rune{'_'},
							ranges:     []rune{'a', 'z'},
							ignoreCase: true,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 167, col: 16, offset: 4332},
							expr: &charClassMatcher{
								pos:        position{line: 167, col: 16, offset: 4332},
								val:        "[0-9a-z_]i",
								chars:      []rune{'_'},
								ranges:     []rune{'0', '9', 'a', 'z'},
								ignoreCase: true,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "FragmentSpread",
			pos:  position{line: 171, col: 1, offset: 4377},
			expr: &actionExpr{
				pos: position{line: 171, col: 19, offset: 4397},
				run: (*parser).callonFragmentSpread1,
				expr: &seqExpr{
					pos: position{line: 171, col: 19, offset: 4397},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 171, col: 19, offset: 4397},
							val:        "...",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 171, col: 25, offset: 4403},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 171, col: 28, offset: 4406},
								name: "FragmentName",
							},
						},
						&labeledExpr{
							pos:   position{line: 171, col: 41, offset: 4419},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 171, col: 44, offset: 4422},
								expr: &ruleRefExpr{
									pos:  position{line: 171, col: 44, offset: 4422},
									name: "Directives",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FragmentDefinition",
			pos:  position{line: 181, col: 1, offset: 4621},
			expr: &actionExpr{
				pos: position{line: 181, col: 22, offset: 4644},
				run: (*parser).callonFragmentDefinition1,
				expr: &seqExpr{
					pos: position{line: 181, col: 22, offset: 4644},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 181, col: 22, offset: 4644},
							val:        "fragment",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 33, offset: 4655},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 35, offset: 4657},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 37, offset: 4659},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 42, offset: 4664},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 44, offset: 4666},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 47, offset: 4669},
								name: "FragmentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 60, offset: 4682},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 62, offset: 4684},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 181, col: 65, offset: 4687},
								expr: &ruleRefExpr{
									pos:  position{line: 181, col: 65, offset: 4687},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 77, offset: 4699},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 79, offset: 4701},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 84, offset: 4706},
								name: "Selections",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 95, offset: 4717},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FragmentName",
			pos:  position{line: 193, col: 1, offset: 4977},
			expr: &actionExpr{
				pos: position{line: 193, col: 16, offset: 4994},
				run: (*parser).callonFragmentName1,
				expr: &labeledExpr{
					pos:   position{line: 193, col: 16, offset: 4994},
					label: "n",
					expr: &ruleRefExpr{
						pos:  position{line: 193, col: 18, offset: 4996},
						name: "Name",
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 195, col: 1, offset: 5020},
			expr: &actionExpr{
				pos: position{line: 195, col: 9, offset: 5030},
				run: (*parser).callonValue1,
				expr: &seqExpr{
					pos: position{line: 195, col: 9, offset: 5030},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 195, col: 9, offset: 5030},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 195, col: 11, offset: 5032},
							label: "v",
							expr: &choiceExpr{
								pos: position{line: 195, col: 14, offset: 5035},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 195, col: 14, offset: 5035},
										name: "Null",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 21, offset: 5042},
										name: "Boolean",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 31, offset: 5052},
										name: "Int",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 37, offset: 5058},
										name: "Float",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 45, offset: 5066},
										name: "String",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 54, offset: 5075},
										name: "EnumValue",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 66, offset: 5087},
										name: "Array",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 74, offset: 5095},
										name: "Object",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 83, offset: 5104},
										name: "Variable",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 195, col: 93, offset: 5114},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Null",
			pos:  position{line: 199, col: 1, offset: 5136},
			expr: &actionExpr{
				pos: position{line: 199, col: 8, offset: 5145},
				run: (*parser).callonNull1,
				expr: &litMatcher{
					pos:        position{line: 199, col: 8, offset: 5145},
					val:        "null",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Boolean",
			pos:  position{line: 200, col: 1, offset: 5172},
			expr: &choiceExpr{
				pos: position{line: 200, col: 11, offset: 5184},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 200, col: 11, offset: 5184},
						run: (*parser).callonBoolean2,
						expr: &litMatcher{
							pos:        position{line: 200, col: 11, offset: 5184},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 200, col: 41, offset: 5214},
						run: (*parser).callonBoolean4,
						expr: &litMatcher{
							pos:        position{line: 200, col: 41, offset: 5214},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Int",
			pos:  position{line: 201, col: 1, offset: 5244},
			expr: &actionExpr{
				pos: position{line: 201, col: 7, offset: 5252},
				run: (*parser).callonInt1,
				expr: &seqExpr{
					pos: position{line: 201, col: 7, offset: 5252},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 201, col: 7, offset: 5252},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 7, offset: 5252},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 13, offset: 5258},
							name: "IntegerPart",
						},
					},
				},
			},
		},
		{
			name: "Float",
			pos:  position{line: 204, col: 1, offset: 5311},
			expr: &actionExpr{
				pos: position{line: 204, col: 9, offset: 5321},
				run: (*parser).callonFloat1,
				expr: &seqExpr{
					pos: position{line: 204, col: 9, offset: 5321},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 204, col: 9, offset: 5321},
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 9, offset: 5321},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 204, col: 15, offset: 5327},
							name: "IntegerPart",
						},
						&litMatcher{
							pos:        position{line: 204, col: 27, offset: 5339},
							val:        ".",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 204, col: 31, offset: 5343},
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 31, offset: 5343},
								name: "Digit",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 204, col: 38, offset: 5350},
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 38, offset: 5350},
								name: "ExponentPart",
							},
						},
					},
				},
			},
		},
		{
			name: "Sign",
			pos:  position{line: 207, col: 1, offset: 5415},
			expr: &litMatcher{
				pos:        position{line: 207, col: 8, offset: 5424},
				val:        "-",
				ignoreCase: false,
			},
		},
		{
			name: "IntegerPart",
			pos:  position{line: 208, col: 1, offset: 5428},
			expr: &choiceExpr{
				pos: position{line: 208, col: 15, offset: 5444},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 208, col: 15, offset: 5444},
						val:        "0",
						ignoreCase: false,
					},
					&seqExpr{
						pos: position{line: 208, col: 21, offset: 5450},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 208, col: 21, offset: 5450},
								name: "NonZeroDigit",
							},
							&zeroOrMoreExpr{
								pos: position{line: 208, col: 34, offset: 5463},
								expr: &ruleRefExpr{
									pos:  position{line: 208, col: 34, offset: 5463},
									name: "Digit",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ExponentPart",
			pos:  position{line: 209, col: 1, offset: 5470},
			expr: &seqExpr{
				pos: position{line: 209, col: 16, offset: 5487},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 209, col: 16, offset: 5487},
						val:        "e",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 209, col: 20, offset: 5491},
						expr: &ruleRefExpr{
							pos:  position{line: 209, col: 20, offset: 5491},
							name: "Sign",
						},
					},
					&oneOrMoreExpr{
						pos: position{line: 209, col: 26, offset: 5497},
						expr: &ruleRefExpr{
							pos:  position{line: 209, col: 26, offset: 5497},
							name: "Digit",
						},
					},
				},
			},
		},
		{
			name: "Digit",
			pos:  position{line: 210, col: 1, offset: 5504},
			expr: &charClassMatcher{
				pos:        position{line: 210, col: 9, offset: 5514},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDigit",
			pos:  position{line: 211, col: 1, offset: 5520},
			expr: &charClassMatcher{
				pos:        position{line: 211, col: 16, offset: 5537},
				val:        "[123456789]",
				chars:      []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "String",
			pos:  position{line: 212, col: 1, offset: 5549},
			expr: &actionExpr{
				pos: position{line: 212, col: 10, offset: 5560},
				run: (*parser).callonString1,
				expr: &seqExpr{
					pos: position{line: 212, col: 10, offset: 5560},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 212, col: 10, offset: 5560},
							val:        "\"",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 212, col: 14, offset: 5564},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 212, col: 16, offset: 5566},
								name: "string",
							},
						},
						&litMatcher{
							pos:        position{line: 212, col: 23, offset: 5573},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "string",
			pos:  position{line: 215, col: 1, offset: 5605},
			expr: &actionExpr{
				pos: position{line: 215, col: 10, offset: 5616},
				run: (*parser).callonstring1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 215, col: 10, offset: 5616},
					expr: &ruleRefExpr{
						pos:  position{line: 215, col: 10, offset: 5616},
						name: "StringCharacter",
					},
				},
			},
		},
		{
			name: "StringCharacter",
			pos:  position{line: 218, col: 1, offset: 5665},
			expr: &choiceExpr{
				pos: position{line: 218, col: 19, offset: 5685},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 218, col: 19, offset: 5685},
						val:        "[^\\\\\"]",
						chars:      []rune{'\\', '"'},
						ignoreCase: false,
						inverted:   true,
					},
					&seqExpr{
						pos: position{line: 218, col: 28, offset: 5694},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 218, col: 28, offset: 5694},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 218, col: 33, offset: 5699},
								name: "EscapedCharacter",
							},
						},
					},
					&seqExpr{
						pos: position{line: 218, col: 52, offset: 5718},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 218, col: 52, offset: 5718},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 218, col: 57, offset: 5723},
								name: "EscapedUnicode",
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedUnicode",
			pos:  position{line: 219, col: 1, offset: 5738},
			expr: &seqExpr{
				pos: position{line: 219, col: 18, offset: 5757},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 219, col: 18, offset: 5757},
						val:        "u",
						ignoreCase: false,
					},
					&charClassMatcher{
						pos:        position{line: 219, col: 22, offset: 5761},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 219, col: 32, offset: 5771},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 219, col: 42, offset: 5781},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 219, col: 52, offset: 5791},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "EscapedCharacter",
			pos:  position{line: 220, col: 1, offset: 5801},
			expr: &choiceExpr{
				pos: position{line: 220, col: 20, offset: 5822},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 220, col: 20, offset: 5822},
						val:        "[\"/bfnrt]",
						chars:      []rune{'"', '/', 'b', 'f', 'n', 'r', 't'},
						ignoreCase: false,
						inverted:   false,
					},
					&litMatcher{
						pos:        position{line: 220, col: 32, offset: 5834},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "EnumValue",
			pos:  position{line: 222, col: 1, offset: 5840},
			expr: &actionExpr{
				pos: position{line: 222, col: 13, offset: 5854},
				run: (*parser).callonEnumValue1,
				expr: &seqExpr{
					pos: position{line: 222, col: 13, offset: 5854},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 222, col: 13, offset: 5854},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 16, offset: 5857},
								name: "TypeName",
							},
						},
						&litMatcher{
							pos:        position{line: 222, col: 25, offset: 5866},
							val:        ".",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 222, col: 29, offset: 5870},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 31, offset: 5872},
								name: "EnumValueName",
							},
						},
					},
				},
			},
		},
		{
			name: "Array",
			pos:  position{line: 228, col: 1, offset: 5976},
			expr: &seqExpr{
				pos: position{line: 228, col: 9, offset: 5986},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 228, col: 9, offset: 5986},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 228, col: 13, offset: 5990},
						expr: &ruleRefExpr{
							pos:  position{line: 228, col: 13, offset: 5990},
							name: "Value",
						},
					},
					&litMatcher{
						pos:        position{line: 228, col: 20, offset: 5997},
						val:        "]",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Object",
			pos:  position{line: 229, col: 1, offset: 6001},
			expr: &seqExpr{
				pos: position{line: 229, col: 10, offset: 6012},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 229, col: 10, offset: 6012},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 229, col: 14, offset: 6016},
						expr: &ruleRefExpr{
							pos:  position{line: 229, col: 14, offset: 6016},
							name: "Property",
						},
					},
					&litMatcher{
						pos:        position{line: 229, col: 24, offset: 6026},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 231, col: 1, offset: 6031},
			expr: &choiceExpr{
				pos: position{line: 231, col: 12, offset: 6044},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 231, col: 12, offset: 6044},
						run: (*parser).callonVariable2,
						expr: &seqExpr{
							pos: position{line: 231, col: 12, offset: 6044},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 231, col: 12, offset: 6044},
									label: "vn",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 15, offset: 6047},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 231, col: 28, offset: 6060},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 231, col: 32, offset: 6064},
									label: "pn",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 35, offset: 6067},
										name: "PropertyName",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 236, col: 5, offset: 6200},
						run: (*parser).callonVariable9,
						expr: &labeledExpr{
							pos:   position{line: 236, col: 5, offset: 6200},
							label: "vn",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 8, offset: 6203},
								name: "VariableName",
							},
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 241, col: 1, offset: 6276},
			expr: &actionExpr{
				pos: position{line: 241, col: 16, offset: 6293},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 241, col: 16, offset: 6293},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 241, col: 16, offset: 6293},
							val:        "$",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 241, col: 20, offset: 6297},
							expr: &charClassMatcher{
								pos:        position{line: 241, col: 20, offset: 6297},
								val:        "[0-9a-z_]i",
								chars:      []rune{'_'},
								ranges:     []rune{'0', '9', 'a', 'z'},
								ignoreCase: true,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Property",
			pos:  position{line: 249, col: 1, offset: 6620},
			expr: &seqExpr{
				pos: position{line: 249, col: 12, offset: 6633},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 249, col: 12, offset: 6633},
						name: "PropertyName",
					},
					&litMatcher{
						pos:        position{line: 249, col: 25, offset: 6646},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 249, col: 29, offset: 6650},
						name: "Value",
					},
				},
			},
		},
		{
			name: "PropertyName",
			pos:  position{line: 250, col: 1, offset: 6656},
			expr: &ruleRefExpr{
				pos:  position{line: 250, col: 16, offset: 6673},
				name: "Name",
			},
		},
		{
			name: "Directives",
			pos:  position{line: 252, col: 1, offset: 6679},
			expr: &actionExpr{
				pos: position{line: 252, col: 14, offset: 6694},
				run: (*parser).callonDirectives1,
				expr: &labeledExpr{
					pos:   position{line: 252, col: 14, offset: 6694},
					label: "ds",
					expr: &oneOrMoreExpr{
						pos: position{line: 252, col: 17, offset: 6697},
						expr: &ruleRefExpr{
							pos:  position{line: 252, col: 17, offset: 6697},
							name: "Directive",
						},
					},
				},
			},
		},
		{
			name: "Directive",
			pos:  position{line: 259, col: 1, offset: 6846},
			expr: &actionExpr{
				pos: position{line: 259, col: 13, offset: 6860},
				run: (*parser).callonDirective1,
				expr: &seqExpr{
					pos: position{line: 259, col: 13, offset: 6860},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 259, col: 13, offset: 6860},
							val:        "@",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 259, col: 17, offset: 6864},
							label: "d",
							expr: &choiceExpr{
								pos: position{line: 259, col: 20, offset: 6867},
								alternatives: []interface{}{
									&actionExpr{
										pos: position{line: 259, col: 20, offset: 6867},
										run: (*parser).callonDirective6,
										expr: &seqExpr{
											pos: position{line: 259, col: 21, offset: 6868},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 259, col: 21, offset: 6868},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 259, col: 24, offset: 6871},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 259, col: 38, offset: 6885},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 259, col: 42, offset: 6889},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 259, col: 44, offset: 6891},
													label: "v",
													expr: &ruleRefExpr{
														pos:  position{line: 259, col: 46, offset: 6893},
														name: "Value",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 265, col: 5, offset: 7003},
										run: (*parser).callonDirective14,
										expr: &seqExpr{
											pos: position{line: 265, col: 6, offset: 7004},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 265, col: 6, offset: 7004},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 265, col: 9, offset: 7007},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 265, col: 23, offset: 7021},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 265, col: 27, offset: 7025},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 265, col: 29, offset: 7027},
													label: "t",
													expr: &ruleRefExpr{
														pos:  position{line: 265, col: 31, offset: 7029},
														name: "Type",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 271, col: 5, offset: 7138},
										run: (*parser).callonDirective22,
										expr: &labeledExpr{
											pos:   position{line: 271, col: 5, offset: 7138},
											label: "dn",
											expr: &ruleRefExpr{
												pos:  position{line: 271, col: 8, offset: 7141},
												name: "DirectiveName",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 275, col: 5, offset: 7217},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DirectiveName",
			pos:  position{line: 279, col: 1, offset: 7239},
			expr: &ruleRefExpr{
				pos:  position{line: 279, col: 17, offset: 7257},
				name: "Name",
			},
		},
		{
			name: "Type",
			pos:  position{line: 281, col: 1, offset: 7263},
			expr: &actionExpr{
				pos: position{line: 281, col: 8, offset: 7272},
				run: (*parser).callonType1,
				expr: &labeledExpr{
					pos:   position{line: 281, col: 8, offset: 7272},
					label: "t",
					expr: &choiceExpr{
						pos: position{line: 281, col: 11, offset: 7275},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 281, col: 11, offset: 7275},
								name: "OptionalType",
							},
							&ruleRefExpr{
								pos:  position{line: 281, col: 26, offset: 7290},
								name: "GenericType",
							},
						},
					},
				},
			},
		},
		{
			name: "OptionalType",
			pos:  position{line: 282, col: 1, offset: 7321},
			expr: &actionExpr{
				pos: position{line: 282, col: 16, offset: 7338},
				run: (*parser).callonOptionalType1,
				expr: &seqExpr{
					pos: position{line: 282, col: 16, offset: 7338},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 282, col: 16, offset: 7338},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 18, offset: 7340},
								name: "GenericType",
							},
						},
						&litMatcher{
							pos:        position{line: 282, col: 30, offset: 7352},
							val:        "?",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "GenericType",
			pos:  position{line: 287, col: 1, offset: 7423},
			expr: &actionExpr{
				pos: position{line: 287, col: 15, offset: 7439},
				run: (*parser).callonGenericType1,
				expr: &seqExpr{
					pos: position{line: 287, col: 15, offset: 7439},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 287, col: 15, offset: 7439},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 287, col: 18, offset: 7442},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 287, col: 27, offset: 7451},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 287, col: 29, offset: 7453},
							label: "tps",
							expr: &zeroOrOneExpr{
								pos: position{line: 287, col: 33, offset: 7457},
								expr: &ruleRefExpr{
									pos:  position{line: 287, col: 33, offset: 7457},
									name: "TypeParams",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "TypeParams",
			pos:  position{line: 292, col: 1, offset: 7524},
			expr: &seqExpr{
				pos: position{line: 292, col: 14, offset: 7539},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 292, col: 14, offset: 7539},
						val:        ":",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 292, col: 18, offset: 7543},
						val:        "<",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 292, col: 22, offset: 7547},
						expr: &ruleRefExpr{
							pos:  position{line: 292, col: 22, offset: 7547},
							name: "Type",
						},
					},
					&litMatcher{
						pos:        position{line: 292, col: 28, offset: 7553},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TypeName",
			pos:  position{line: 293, col: 1, offset: 7557},
			expr: &ruleRefExpr{
				pos:  position{line: 293, col: 12, offset: 7570},
				name: "Name",
			},
		},
		{
			name: "TypeDefinition",
			pos:  position{line: 294, col: 1, offset: 7575},
			expr: &actionExpr{
				pos: position{line: 294, col: 18, offset: 7594},
				run: (*parser).callonTypeDefinition1,
				expr: &seqExpr{
					pos: position{line: 294, col: 18, offset: 7594},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 294, col: 18, offset: 7594},
							val:        "type",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 294, col: 25, offset: 7601},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 294, col: 27, offset: 7603},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 294, col: 30, offset: 7606},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 294, col: 39, offset: 7615},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 294, col: 41, offset: 7617},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 294, col: 44, offset: 7620},
								expr: &ruleRefExpr{
									pos:  position{line: 294, col: 44, offset: 7620},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 294, col: 56, offset: 7632},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 294, col: 58, offset: 7634},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 294, col: 62, offset: 7638},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 294, col: 66, offset: 7642},
								expr: &ruleRefExpr{
									pos:  position{line: 294, col: 66, offset: 7642},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 294, col: 83, offset: 7659},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeExtension",
			pos:  position{line: 314, col: 1, offset: 8140},
			expr: &actionExpr{
				pos: position{line: 314, col: 17, offset: 8158},
				run: (*parser).callonTypeExtension1,
				expr: &seqExpr{
					pos: position{line: 314, col: 17, offset: 8158},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 314, col: 17, offset: 8158},
							val:        "extend",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 26, offset: 8167},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 314, col: 28, offset: 8169},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 314, col: 31, offset: 8172},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 40, offset: 8181},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 314, col: 42, offset: 8183},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 314, col: 45, offset: 8186},
								expr: &ruleRefExpr{
									pos:  position{line: 314, col: 45, offset: 8186},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 57, offset: 8198},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 314, col: 59, offset: 8200},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 314, col: 63, offset: 8204},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 314, col: 67, offset: 8208},
								expr: &ruleRefExpr{
									pos:  position{line: 314, col: 67, offset: 8208},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 314, col: 84, offset: 8225},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Interfaces",
			pos:  position{line: 334, col: 1, offset: 8705},
			expr: &actionExpr{
				pos: position{line: 334, col: 14, offset: 8720},
				run: (*parser).callonInterfaces1,
				expr: &oneOrMoreExpr{
					pos: position{line: 334, col: 14, offset: 8720},
					expr: &ruleRefExpr{
						pos:  position{line: 334, col: 14, offset: 8720},
						name: "GenericType",
					},
				},
			},
		},
		{
			name: "FieldDefinition",
			pos:  position{line: 338, col: 1, offset: 8827},
			expr: &actionExpr{
				pos: position{line: 338, col: 19, offset: 8847},
				run: (*parser).callonFieldDefinition1,
				expr: &seqExpr{
					pos: position{line: 338, col: 19, offset: 8847},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 338, col: 19, offset: 8847},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 338, col: 21, offset: 8849},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 338, col: 24, offset: 8852},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 338, col: 34, offset: 8862},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 338, col: 36, offset: 8864},
							label: "args",
							expr: &zeroOrOneExpr{
								pos: position{line: 338, col: 41, offset: 8869},
								expr: &ruleRefExpr{
									pos:  position{line: 338, col: 41, offset: 8869},
									name: "ArgumentDefinitions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 338, col: 62, offset: 8890},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 338, col: 64, offset: 8892},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 338, col: 68, offset: 8896},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 338, col: 70, offset: 8898},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 338, col: 72, offset: 8900},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 338, col: 77, offset: 8905},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinitions",
			pos:  position{line: 349, col: 1, offset: 9142},
			expr: &actionExpr{
				pos: position{line: 349, col: 23, offset: 9166},
				run: (*parser).callonArgumentDefinitions1,
				expr: &seqExpr{
					pos: position{line: 349, col: 23, offset: 9166},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 349, col: 23, offset: 9166},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 349, col: 27, offset: 9170},
							label: "args",
							expr: &oneOrMoreExpr{
								pos: position{line: 349, col: 32, offset: 9175},
								expr: &ruleRefExpr{
									pos:  position{line: 349, col: 32, offset: 9175},
									name: "ArgumentDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 349, col: 52, offset: 9195},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinition",
			pos:  position{line: 356, col: 1, offset: 9358},
			expr: &actionExpr{
				pos: position{line: 356, col: 22, offset: 9381},
				run: (*parser).callonArgumentDefinition1,
				expr: &seqExpr{
					pos: position{line: 356, col: 22, offset: 9381},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 356, col: 22, offset: 9381},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 356, col: 24, offset: 9383},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 356, col: 27, offset: 9386},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 356, col: 40, offset: 9399},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 356, col: 42, offset: 9401},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 356, col: 46, offset: 9405},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 356, col: 48, offset: 9407},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 356, col: 50, offset: 9409},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 356, col: 55, offset: 9414},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 356, col: 57, offset: 9416},
							label: "dv",
							expr: &zeroOrOneExpr{
								pos: position{line: 356, col: 60, offset: 9419},
								expr: &ruleRefExpr{
									pos:  position{line: 356, col: 60, offset: 9419},
									name: "DefaultValue",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EnumDefinition",
			pos:  position{line: 368, col: 1, offset: 9650},
			expr: &actionExpr{
				pos: position{line: 368, col: 18, offset: 9669},
				run: (*parser).callonEnumDefinition1,
				expr: &seqExpr{
					pos: position{line: 368, col: 18, offset: 9669},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 368, col: 18, offset: 9669},
							val:        "enum",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 368, col: 25, offset: 9676},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 368, col: 27, offset: 9678},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 368, col: 30, offset: 9681},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 368, col: 39, offset: 9690},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 368, col: 41, offset: 9692},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 368, col: 45, offset: 9696},
							label: "vals",
							expr: &oneOrMoreExpr{
								pos: position{line: 368, col: 50, offset: 9701},
								expr: &ruleRefExpr{
									pos:  position{line: 368, col: 50, offset: 9701},
									name: "EnumValueName",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 368, col: 65, offset: 9716},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EnumValueName",
			pos:  position{line: 378, col: 1, offset: 9897},
			expr: &actionExpr{
				pos: position{line: 378, col: 17, offset: 9915},
				run: (*parser).callonEnumValueName1,
				expr: &seqExpr{
					pos: position{line: 378, col: 17, offset: 9915},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 378, col: 17, offset: 9915},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 378, col: 19, offset: 9917},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 378, col: 21, offset: 9919},
								name: "Name",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 378, col: 26, offset: 9924},
							name: "_",
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"ignored\"",
			pos:         position{line: 380, col: 1, offset: 9945},
			expr: &actionExpr{
				pos: position{line: 380, col: 15, offset: 9961},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 380, col: 15, offset: 9961},
					expr: &choiceExpr{
						pos: position{line: 380, col: 16, offset: 9962},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 380, col: 16, offset: 9962},
								name: "whitespace",
							},
							&ruleRefExpr{
								pos:  position{line: 380, col: 29, offset: 9975},
								name: "Comment",
							},
							&litMatcher{
								pos:        position{line: 380, col: 39, offset: 9985},
								val:        ",",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "whitespace",
			pos:  position{line: 381, col: 1, offset: 10011},
			expr: &charClassMatcher{
				pos:        position{line: 381, col: 14, offset: 10026},
				val:        "[ \\n\\t\\r]",
				chars:      []rune{' ', '\n', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 383, col: 1, offset: 10037},
			expr: &notExpr{
				pos: position{line: 383, col: 7, offset: 10045},
				expr: &anyMatcher{
					line: 383, col: 8, offset: 10046,
				},
			},
		},
	},
}

func (c *current) onDocument1(stmts interface{}) (interface{}, error) {
	result := graphql.Document{
		Operations: []graphql.Operation{},
	}
	sl := ifs(stmts)
	for _, s := range sl {
		if s == nil {
			continue
		}
		if op, ok := s.(graphql.Operation); ok {
			result.Operations = append(result.Operations, op)
		} else if fragDef, ok := s.(graphql.FragmentDefinition); ok {
			result.FragmentDefinitions = append(result.FragmentDefinitions, fragDef)
		} else if typeDef, ok := s.(graphql.TypeDefinition); ok {
			result.TypeDefinitions = append(result.TypeDefinitions, typeDef)
		} else if typeExt, ok := s.(graphql.TypeExtension); ok {
			result.TypeExtensions = append(result.TypeExtensions, typeExt)
		} else if enumDef, ok := s.(graphql.EnumDefinition); ok {
			result.EnumDefinitions = append(result.EnumDefinitions, enumDef)
		} else {
			return result, fmt.Errorf("unhandled statement type: %#v", s)
		}
	}
	return result, nil
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["stmts"])
}

func (c *current) onStatement1(s interface{}) (interface{}, error) {
	return s, nil
}

func (p *parser) callonStatement1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement1(stack["s"])
}

func (c *current) onComment1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment1()
}

func (c *current) onOperation2(sels interface{}) (interface{}, error) {
	return graphql.Operation{
		Type:       graphql.OperationQuery,
		Selections: sels.([]graphql.Selection),
	}, nil

}

func (p *parser) callonOperation2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperation2(stack["sels"])
}

func (c *current) onOperation5(ot, on, vds, ds, sels interface{}) (interface{}, error) {
	var (
		varDefs    []graphql.VariableDefinition
		directives []graphql.Directive
	)
	if vds != nil {
		varDefs = vds.([]graphql.VariableDefinition)
	}
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.Operation{
		Type:                ot.(graphql.OperationType),
		Name:                on.(string),
		Selections:          sels.([]graphql.Selection),
		Directives:          directives,
		VariableDefinitions: varDefs,
	}, nil
}

func (p *parser) callonOperation5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperation5(stack["ot"], stack["on"], stack["vds"], stack["ds"], stack["sels"])
}

func (c *current) onOperationType2() (interface{}, error) {
	return graphql.OperationQuery, nil
}

func (p *parser) callonOperationType2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperationType2()
}

func (c *current) onOperationType4() (interface{}, error) {
	return graphql.OperationMutation, nil
}

func (p *parser) callonOperationType4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperationType4()
}

func (c *current) onOperationName1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonOperationName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperationName1()
}

func (c *current) onVariableDefinitions1(vds interface{}) (interface{}, error) {
	result := []graphql.VariableDefinition{}
	for _, d := range ifs(vds) {
		result = append(result, d.(graphql.VariableDefinition))
	}
	return result, nil
}

func (p *parser) callonVariableDefinitions1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableDefinitions1(stack["vds"])
}

func (c *current) onVariableDefinition1(v, t, d interface{}) (interface{}, error) {
	var defaultValue *graphql.Value
	if d != nil {
		v := d.(graphql.Value)
		defaultValue = &v
	}
	return graphql.VariableDefinition{
		Variable:     v.(graphql.Variable),
		Type:         t.(graphql.Type),
		DefaultValue: defaultValue,
	}, nil
}

func (p *parser) callonVariableDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableDefinition1(stack["v"], stack["t"], stack["d"])
}

func (c *current) onDefaultValue1(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonDefaultValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDefaultValue1(stack["v"])
}

func (c *current) onSelections1(s interface{}) (interface{}, error) {
	result := []graphql.Selection{}
	for _, sel := range ifs(s) {
		if sel, ok := sel.(graphql.Selection); ok {
			result = append(result, sel)
		} else {
			return result, fmt.Errorf("got unexpected (non-statement) type: %#v", sel)
		}
	}
	return result, nil
}

func (p *parser) callonSelections1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelections1(stack["s"])
}

func (c *current) onSelection2(f interface{}) (interface{}, error) {
	field := f.(graphql.Field)
	return graphql.Selection{Field: &field}, nil
}

func (p *parser) callonSelection2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelection2(stack["f"])
}

func (c *current) onSelection8(fs interface{}) (interface{}, error) {
	fragmentSpread := fs.(graphql.FragmentSpread)
	return graphql.Selection{FragmentSpread: &fragmentSpread}, nil
}

func (p *parser) callonSelection8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelection8(stack["fs"])
}

func (c *current) onField1(fa, fn, as, ds, sels interface{}) (interface{}, error) {
	var (
		selections []graphql.Selection
		arguments  []graphql.Argument
		directives []graphql.Directive
		fieldAlias string
	)
	if fa != nil {
		fieldAlias = fa.(string)
	}
	if sels != nil {
		selections = sels.([]graphql.Selection)
	}
	if as != nil {
		arguments = as.([]graphql.Argument)
	}
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.Field{
		Name:       fn.(string),
		Alias:      fieldAlias,
		Arguments:  arguments,
		Selections: selections,
		Directives: directives,
	}, nil
}

func (p *parser) callonField1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onField1(stack["fa"], stack["fn"], stack["as"], stack["ds"], stack["sels"])
}

func (c *current) onFieldAlias1(n interface{}) (interface{}, error) {
	return n, nil
}

func (p *parser) callonFieldAlias1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFieldAlias1(stack["n"])
}

func (c *current) onArguments1(args interface{}) (interface{}, error) {
	results := []graphql.Argument{}
	for _, a := range ifs(args) {
		if a, ok := a.(graphql.Argument); ok {
			results = append(results, a)
		} else {
			return results, fmt.Errorf("got unexpected type: %#v", a)
		}
	}
	return results, nil
}

func (p *parser) callonArguments1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments1(stack["args"])
}

func (c *current) onArgument1(an, v interface{}) (interface{}, error) {
	return graphql.Argument{
		Name:  an.(string),
		Value: v,
	}, nil
}

func (p *parser) callonArgument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgument1(stack["an"], stack["v"])
}

func (c *current) onName1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onName1()
}

func (c *current) onFragmentSpread1(fn, ds interface{}) (interface{}, error) {
	var directives []graphql.Directive
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.FragmentSpread{
		Name:       fn.(string),
		Directives: directives,
	}, nil
}

func (p *parser) callonFragmentSpread1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFragmentSpread1(stack["fn"], stack["ds"])
}

func (c *current) onFragmentDefinition1(t, fn, ds, sels interface{}) (interface{}, error) {
	var directives []graphql.Directive
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.FragmentDefinition{
		Name:       fn.(string),
		Type:       t.(graphql.Type),
		Selections: sels.([]graphql.Selection),
		Directives: directives,
	}, nil
}

func (p *parser) callonFragmentDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFragmentDefinition1(stack["t"], stack["fn"], stack["ds"], stack["sels"])
}

func (c *current) onFragmentName1(n interface{}) (interface{}, error) {
	return n, nil
}

func (p *parser) callonFragmentName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFragmentName1(stack["n"])
}

func (c *current) onValue1(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue1(stack["v"])
}

func (c *current) onNull1() (interface{}, error) {
	return nil, nil
}

func (p *parser) callonNull1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNull1()
}

func (c *current) onBoolean2() (interface{}, error) {
	return true, nil
}

func (p *parser) callonBoolean2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoolean2()
}

func (c *current) onBoolean4() (interface{}, error) {
	return false, nil
}

func (p *parser) callonBoolean4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoolean4()
}

func (c *current) onInt1() (interface{}, error) {
	return strconv.Atoi(string(c.text))
}

func (p *parser) callonInt1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInt1()
}

func (c *current) onFloat1() (interface{}, error) {
	return strconv.ParseFloat(string(c.text), 64)
}

func (p *parser) callonFloat1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFloat1()
}

func (c *current) onString1(s interface{}) (interface{}, error) {
	return s.(string), nil
}

func (p *parser) callonString1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onString1(stack["s"])
}

func (c *current) onstring1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonstring1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstring1()
}

func (c *current) onEnumValue1(tn, v interface{}) (interface{}, error) {
	return graphql.EnumValue{
		EnumTypeName: tn.(string),
		Value:        v.(string),
	}, nil
}

func (p *parser) callonEnumValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEnumValue1(stack["tn"], stack["v"])
}

func (c *current) onVariable2(vn, pn interface{}) (interface{}, error) {
	return graphql.Variable{
		Name:              vn.(string),
		PropertySelection: &graphql.Variable{Name: pn.(string)},
	}, nil
}

func (p *parser) callonVariable2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariable2(stack["vn"], stack["pn"])
}

func (c *current) onVariable9(vn interface{}) (interface{}, error) {
	return graphql.Variable{
		Name: vn.(string),
	}, nil
}

func (p *parser) callonVariable9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariable9(stack["vn"])
}

func (c *current) onVariableName1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonVariableName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableName1()
}

func (c *current) onDirectives1(ds interface{}) (interface{}, error) {
	result := []graphql.Directive{}
	for _, d := range ifs(ds) {
		result = append(result, d.(graphql.Directive))
	}
	return result, nil
}

func (p *parser) callonDirectives1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDirectives1(stack["ds"])
}

func (c *current) onDirective6(dn, v interface{}) (interface{}, error) {
	val := v.(graphql.Value)
	return graphql.Directive{
		Name:  dn.(string),
		Value: &val,
	}, nil
}

func (p *parser) callonDirective6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDirective6(stack["dn"], stack["v"])
}

func (c *current) onDirective14(dn, t interface{}) (interface{}, error) {
	typ := t.(graphql.Value)
	return graphql.Directive{
		Name:  dn.(string),
		Value: &typ,
	}, nil
}

func (p *parser) callonDirective14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDirective14(stack["dn"], stack["t"])
}

func (c *current) onDirective22(dn interface{}) (interface{}, error) {
	return graphql.Directive{
		Name: dn.(string),
	}, nil
}

func (p *parser) callonDirective22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDirective22(stack["dn"])
}

func (c *current) onDirective1(d interface{}) (interface{}, error) {
	return d, nil
}

func (p *parser) callonDirective1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDirective1(stack["d"])
}

func (c *current) onType1(t interface{}) (interface{}, error) {
	return t, nil
}

func (p *parser) callonType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onType1(stack["t"])
}

func (c *current) onOptionalType1(t interface{}) (interface{}, error) {
	typ := t.(graphql.Type)
	typ.Optional = true
	return typ, nil
}

func (p *parser) callonOptionalType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOptionalType1(stack["t"])
}

func (c *current) onGenericType1(tn, tps interface{}) (interface{}, error) {
	return graphql.Type{
		Name: tn.(string),
	}, nil
}

func (p *parser) callonGenericType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGenericType1(stack["tn"], stack["tps"])
}

func (c *current) onTypeDefinition1(tn, is, fds interface{}) (interface{}, error) {
	var (
		interfaces       []graphql.Interface
		fieldDefinitions []graphql.FieldDefinition
	)
	if is != nil {
		interfaces = is.([]graphql.Interface)
	}
	if fds != nil {
		fieldDefinitions = make([]graphql.FieldDefinition, 0, len(ifs(fds)))
	}
	for _, fd := range ifs(fds) {
		fieldDefinitions = append(fieldDefinitions, fd.(graphql.FieldDefinition))
	}
	return graphql.TypeDefinition{
		Name:             tn.(string),
		Interfaces:       interfaces,
		FieldDefinitions: fieldDefinitions,
	}, nil
}

func (p *parser) callonTypeDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefinition1(stack["tn"], stack["is"], stack["fds"])
}

func (c *current) onTypeExtension1(tn, is, fds interface{}) (interface{}, error) {
	var (
		interfaces       []graphql.Interface
		fieldDefinitions []graphql.FieldDefinition
	)
	if is != nil {
		interfaces = is.([]graphql.Interface)
	}
	if fds != nil {
		fieldDefinitions = make([]graphql.FieldDefinition, 0, len(ifs(fds)))
	}
	for _, fd := range ifs(fds) {
		fieldDefinitions = append(fieldDefinitions, fd.(graphql.FieldDefinition))
	}
	return graphql.TypeExtension{
		Name:             tn.(string),
		Interfaces:       interfaces,
		FieldDefinitions: fieldDefinitions,
	}, nil
}

func (p *parser) callonTypeExtension1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeExtension1(stack["tn"], stack["is"], stack["fds"])
}

func (c *current) onInterfaces1() (interface{}, error) {
	//return []graphql.Interface{}, nil
	return nil, fmt.Errorf("TODO: not yet implemented")
}

func (p *parser) callonInterfaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInterfaces1()
}

func (c *current) onFieldDefinition1(fn, args, t interface{}) (interface{}, error) {
	var argDefs []graphql.ArgumentDefinition
	if args != nil {
		argDefs = args.([]graphql.ArgumentDefinition)
	}
	return graphql.FieldDefinition{
		Name:                fn.(string),
		Type:                t.(graphql.Type),
		ArgumentDefinitions: argDefs,
	}, nil
}

func (p *parser) callonFieldDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFieldDefinition1(stack["fn"], stack["args"], stack["t"])
}

func (c *current) onArgumentDefinitions1(args interface{}) (interface{}, error) {
	result := []graphql.ArgumentDefinition{}
	for _, a := range ifs(args) {
		result = append(result, a.(graphql.ArgumentDefinition))
	}
	return result, nil
}

func (p *parser) callonArgumentDefinitions1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgumentDefinitions1(stack["args"])
}

func (c *current) onArgumentDefinition1(an, t, dv interface{}) (interface{}, error) {
	var defaultVal *graphql.Value
	if dv != nil {
		v := dv.(graphql.Value)
		defaultVal = &v
	}
	return graphql.ArgumentDefinition{
		Name:         an.(string),
		Type:         t.(graphql.Type),
		DefaultValue: defaultVal,
	}, nil
}

func (p *parser) callonArgumentDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgumentDefinition1(stack["an"], stack["t"], stack["dv"])
}

func (c *current) onEnumDefinition1(tn, vals interface{}) (interface{}, error) {
	values := []string{}
	for _, v := range ifs(vals) {
		values = append(values, v.(string))
	}
	return graphql.EnumDefinition{
		Name:   tn.(string),
		Values: values,
	}, nil
}

func (p *parser) callonEnumDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEnumDefinition1(stack["tn"], stack["vals"])
}

func (c *current) onEnumValueName1(n interface{}) (interface{}, error) {
	return n, nil
}

func (p *parser) callonEnumValueName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEnumValueName1(stack["n"])
}

func (c *current) on_1() (interface{}, error) {
	return nil, nil
}

func (p *parser) callon_1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.on_1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n > 0 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
