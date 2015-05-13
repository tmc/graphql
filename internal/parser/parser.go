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
			pos:  position{line: 141, col: 1, offset: 3742},
			expr: &seqExpr{
				pos: position{line: 141, col: 14, offset: 3757},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 141, col: 14, offset: 3757},
						name: "Name",
					},
					&litMatcher{
						pos:        position{line: 141, col: 19, offset: 3762},
						val:        ":",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "FieldName",
			pos:  position{line: 142, col: 1, offset: 3766},
			expr: &ruleRefExpr{
				pos:  position{line: 142, col: 13, offset: 3780},
				name: "Name",
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 143, col: 1, offset: 3785},
			expr: &actionExpr{
				pos: position{line: 143, col: 13, offset: 3799},
				run: (*parser).callonArguments1,
				expr: &seqExpr{
					pos: position{line: 143, col: 13, offset: 3799},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 143, col: 13, offset: 3799},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 143, col: 17, offset: 3803},
							label: "args",
							expr: &zeroOrMoreExpr{
								pos: position{line: 143, col: 23, offset: 3809},
								expr: &ruleRefExpr{
									pos:  position{line: 143, col: 23, offset: 3809},
									name: "Argument",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 143, col: 34, offset: 3820},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Argument",
			pos:  position{line: 154, col: 1, offset: 4065},
			expr: &actionExpr{
				pos: position{line: 154, col: 12, offset: 4078},
				run: (*parser).callonArgument1,
				expr: &seqExpr{
					pos: position{line: 154, col: 12, offset: 4078},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 154, col: 12, offset: 4078},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 154, col: 14, offset: 4080},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 17, offset: 4083},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 154, col: 30, offset: 4096},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 154, col: 32, offset: 4098},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 154, col: 36, offset: 4102},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 154, col: 38, offset: 4104},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 40, offset: 4106},
								name: "Value",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 154, col: 46, offset: 4112},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentName",
			pos:  position{line: 160, col: 1, offset: 4185},
			expr: &ruleRefExpr{
				pos:  position{line: 160, col: 16, offset: 4202},
				name: "Name",
			},
		},
		{
			name: "Name",
			pos:  position{line: 162, col: 1, offset: 4208},
			expr: &actionExpr{
				pos: position{line: 162, col: 8, offset: 4217},
				run: (*parser).callonName1,
				expr: &seqExpr{
					pos: position{line: 162, col: 8, offset: 4217},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 162, col: 8, offset: 4217},
							val:        "[a-z_]i",
							chars:      []rune{'_'},
							ranges:     []rune{'a', 'z'},
							ignoreCase: true,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 162, col: 16, offset: 4225},
							expr: &charClassMatcher{
								pos:        position{line: 162, col: 16, offset: 4225},
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
			pos:  position{line: 166, col: 1, offset: 4270},
			expr: &actionExpr{
				pos: position{line: 166, col: 19, offset: 4290},
				run: (*parser).callonFragmentSpread1,
				expr: &seqExpr{
					pos: position{line: 166, col: 19, offset: 4290},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 166, col: 19, offset: 4290},
							val:        "...",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 166, col: 25, offset: 4296},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 28, offset: 4299},
								name: "FragmentName",
							},
						},
						&labeledExpr{
							pos:   position{line: 166, col: 41, offset: 4312},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 166, col: 44, offset: 4315},
								expr: &ruleRefExpr{
									pos:  position{line: 166, col: 44, offset: 4315},
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
			pos:  position{line: 176, col: 1, offset: 4514},
			expr: &actionExpr{
				pos: position{line: 176, col: 22, offset: 4537},
				run: (*parser).callonFragmentDefinition1,
				expr: &seqExpr{
					pos: position{line: 176, col: 22, offset: 4537},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 176, col: 22, offset: 4537},
							val:        "fragment",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 33, offset: 4548},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 176, col: 35, offset: 4550},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 37, offset: 4552},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 42, offset: 4557},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 176, col: 44, offset: 4559},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 47, offset: 4562},
								name: "FragmentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 60, offset: 4575},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 176, col: 62, offset: 4577},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 176, col: 65, offset: 4580},
								expr: &ruleRefExpr{
									pos:  position{line: 176, col: 65, offset: 4580},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 77, offset: 4592},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 176, col: 79, offset: 4594},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 84, offset: 4599},
								name: "Selections",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 95, offset: 4610},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FragmentName",
			pos:  position{line: 188, col: 1, offset: 4870},
			expr: &actionExpr{
				pos: position{line: 188, col: 16, offset: 4887},
				run: (*parser).callonFragmentName1,
				expr: &labeledExpr{
					pos:   position{line: 188, col: 16, offset: 4887},
					label: "n",
					expr: &ruleRefExpr{
						pos:  position{line: 188, col: 18, offset: 4889},
						name: "Name",
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 190, col: 1, offset: 4913},
			expr: &actionExpr{
				pos: position{line: 190, col: 9, offset: 4923},
				run: (*parser).callonValue1,
				expr: &seqExpr{
					pos: position{line: 190, col: 9, offset: 4923},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 190, col: 9, offset: 4923},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 190, col: 11, offset: 4925},
							label: "v",
							expr: &choiceExpr{
								pos: position{line: 190, col: 14, offset: 4928},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 190, col: 14, offset: 4928},
										name: "Null",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 21, offset: 4935},
										name: "Boolean",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 31, offset: 4945},
										name: "Int",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 37, offset: 4951},
										name: "Float",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 45, offset: 4959},
										name: "String",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 54, offset: 4968},
										name: "EnumValue",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 66, offset: 4980},
										name: "Array",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 74, offset: 4988},
										name: "Object",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 83, offset: 4997},
										name: "Variable",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 190, col: 93, offset: 5007},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Null",
			pos:  position{line: 194, col: 1, offset: 5029},
			expr: &actionExpr{
				pos: position{line: 194, col: 8, offset: 5038},
				run: (*parser).callonNull1,
				expr: &litMatcher{
					pos:        position{line: 194, col: 8, offset: 5038},
					val:        "null",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Boolean",
			pos:  position{line: 195, col: 1, offset: 5065},
			expr: &choiceExpr{
				pos: position{line: 195, col: 11, offset: 5077},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 195, col: 11, offset: 5077},
						run: (*parser).callonBoolean2,
						expr: &litMatcher{
							pos:        position{line: 195, col: 11, offset: 5077},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 195, col: 41, offset: 5107},
						run: (*parser).callonBoolean4,
						expr: &litMatcher{
							pos:        position{line: 195, col: 41, offset: 5107},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Int",
			pos:  position{line: 196, col: 1, offset: 5137},
			expr: &actionExpr{
				pos: position{line: 196, col: 7, offset: 5145},
				run: (*parser).callonInt1,
				expr: &seqExpr{
					pos: position{line: 196, col: 7, offset: 5145},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 196, col: 7, offset: 5145},
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 7, offset: 5145},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 13, offset: 5151},
							name: "IntegerPart",
						},
					},
				},
			},
		},
		{
			name: "Float",
			pos:  position{line: 199, col: 1, offset: 5204},
			expr: &actionExpr{
				pos: position{line: 199, col: 9, offset: 5214},
				run: (*parser).callonFloat1,
				expr: &seqExpr{
					pos: position{line: 199, col: 9, offset: 5214},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 199, col: 9, offset: 5214},
							expr: &ruleRefExpr{
								pos:  position{line: 199, col: 9, offset: 5214},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 199, col: 15, offset: 5220},
							name: "IntegerPart",
						},
						&litMatcher{
							pos:        position{line: 199, col: 27, offset: 5232},
							val:        ".",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 199, col: 31, offset: 5236},
							expr: &ruleRefExpr{
								pos:  position{line: 199, col: 31, offset: 5236},
								name: "Digit",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 199, col: 38, offset: 5243},
							expr: &ruleRefExpr{
								pos:  position{line: 199, col: 38, offset: 5243},
								name: "ExponentPart",
							},
						},
					},
				},
			},
		},
		{
			name: "Sign",
			pos:  position{line: 202, col: 1, offset: 5308},
			expr: &litMatcher{
				pos:        position{line: 202, col: 8, offset: 5317},
				val:        "-",
				ignoreCase: false,
			},
		},
		{
			name: "IntegerPart",
			pos:  position{line: 203, col: 1, offset: 5321},
			expr: &choiceExpr{
				pos: position{line: 203, col: 15, offset: 5337},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 203, col: 15, offset: 5337},
						val:        "0",
						ignoreCase: false,
					},
					&seqExpr{
						pos: position{line: 203, col: 21, offset: 5343},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 203, col: 21, offset: 5343},
								name: "NonZeroDigit",
							},
							&zeroOrMoreExpr{
								pos: position{line: 203, col: 34, offset: 5356},
								expr: &ruleRefExpr{
									pos:  position{line: 203, col: 34, offset: 5356},
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
			pos:  position{line: 204, col: 1, offset: 5363},
			expr: &seqExpr{
				pos: position{line: 204, col: 16, offset: 5380},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 204, col: 16, offset: 5380},
						val:        "e",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 204, col: 20, offset: 5384},
						expr: &ruleRefExpr{
							pos:  position{line: 204, col: 20, offset: 5384},
							name: "Sign",
						},
					},
					&oneOrMoreExpr{
						pos: position{line: 204, col: 26, offset: 5390},
						expr: &ruleRefExpr{
							pos:  position{line: 204, col: 26, offset: 5390},
							name: "Digit",
						},
					},
				},
			},
		},
		{
			name: "Digit",
			pos:  position{line: 205, col: 1, offset: 5397},
			expr: &charClassMatcher{
				pos:        position{line: 205, col: 9, offset: 5407},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDigit",
			pos:  position{line: 206, col: 1, offset: 5413},
			expr: &charClassMatcher{
				pos:        position{line: 206, col: 16, offset: 5430},
				val:        "[123456789]",
				chars:      []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "String",
			pos:  position{line: 207, col: 1, offset: 5442},
			expr: &actionExpr{
				pos: position{line: 207, col: 10, offset: 5453},
				run: (*parser).callonString1,
				expr: &seqExpr{
					pos: position{line: 207, col: 10, offset: 5453},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 207, col: 10, offset: 5453},
							val:        "\"",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 207, col: 14, offset: 5457},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 207, col: 16, offset: 5459},
								name: "string",
							},
						},
						&litMatcher{
							pos:        position{line: 207, col: 23, offset: 5466},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "string",
			pos:  position{line: 210, col: 1, offset: 5498},
			expr: &actionExpr{
				pos: position{line: 210, col: 10, offset: 5509},
				run: (*parser).callonstring1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 210, col: 10, offset: 5509},
					expr: &ruleRefExpr{
						pos:  position{line: 210, col: 10, offset: 5509},
						name: "StringCharacter",
					},
				},
			},
		},
		{
			name: "StringCharacter",
			pos:  position{line: 213, col: 1, offset: 5558},
			expr: &choiceExpr{
				pos: position{line: 213, col: 19, offset: 5578},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 213, col: 19, offset: 5578},
						val:        "[^\\\\\"]",
						chars:      []rune{'\\', '"'},
						ignoreCase: false,
						inverted:   true,
					},
					&seqExpr{
						pos: position{line: 213, col: 28, offset: 5587},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 213, col: 28, offset: 5587},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 213, col: 33, offset: 5592},
								name: "EscapedCharacter",
							},
						},
					},
					&seqExpr{
						pos: position{line: 213, col: 52, offset: 5611},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 213, col: 52, offset: 5611},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 213, col: 57, offset: 5616},
								name: "EscapedUnicode",
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedUnicode",
			pos:  position{line: 214, col: 1, offset: 5631},
			expr: &seqExpr{
				pos: position{line: 214, col: 18, offset: 5650},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 214, col: 18, offset: 5650},
						val:        "u",
						ignoreCase: false,
					},
					&charClassMatcher{
						pos:        position{line: 214, col: 22, offset: 5654},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 214, col: 32, offset: 5664},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 214, col: 42, offset: 5674},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 214, col: 52, offset: 5684},
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
			pos:  position{line: 215, col: 1, offset: 5694},
			expr: &choiceExpr{
				pos: position{line: 215, col: 20, offset: 5715},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 215, col: 20, offset: 5715},
						val:        "[\"/bfnrt]",
						chars:      []rune{'"', '/', 'b', 'f', 'n', 'r', 't'},
						ignoreCase: false,
						inverted:   false,
					},
					&litMatcher{
						pos:        position{line: 215, col: 32, offset: 5727},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "EnumValue",
			pos:  position{line: 217, col: 1, offset: 5733},
			expr: &actionExpr{
				pos: position{line: 217, col: 13, offset: 5747},
				run: (*parser).callonEnumValue1,
				expr: &seqExpr{
					pos: position{line: 217, col: 13, offset: 5747},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 217, col: 13, offset: 5747},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 16, offset: 5750},
								name: "TypeName",
							},
						},
						&litMatcher{
							pos:        position{line: 217, col: 25, offset: 5759},
							val:        ".",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 217, col: 29, offset: 5763},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 31, offset: 5765},
								name: "EnumValueName",
							},
						},
					},
				},
			},
		},
		{
			name: "Array",
			pos:  position{line: 223, col: 1, offset: 5869},
			expr: &seqExpr{
				pos: position{line: 223, col: 9, offset: 5879},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 223, col: 9, offset: 5879},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 223, col: 13, offset: 5883},
						expr: &ruleRefExpr{
							pos:  position{line: 223, col: 13, offset: 5883},
							name: "Value",
						},
					},
					&litMatcher{
						pos:        position{line: 223, col: 20, offset: 5890},
						val:        "]",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Object",
			pos:  position{line: 224, col: 1, offset: 5894},
			expr: &seqExpr{
				pos: position{line: 224, col: 10, offset: 5905},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 224, col: 10, offset: 5905},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 224, col: 14, offset: 5909},
						expr: &ruleRefExpr{
							pos:  position{line: 224, col: 14, offset: 5909},
							name: "Property",
						},
					},
					&litMatcher{
						pos:        position{line: 224, col: 24, offset: 5919},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 226, col: 1, offset: 5924},
			expr: &choiceExpr{
				pos: position{line: 226, col: 12, offset: 5937},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 226, col: 12, offset: 5937},
						run: (*parser).callonVariable2,
						expr: &seqExpr{
							pos: position{line: 226, col: 12, offset: 5937},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 226, col: 12, offset: 5937},
									label: "vn",
									expr: &ruleRefExpr{
										pos:  position{line: 226, col: 15, offset: 5940},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 226, col: 28, offset: 5953},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 226, col: 32, offset: 5957},
									label: "pn",
									expr: &ruleRefExpr{
										pos:  position{line: 226, col: 35, offset: 5960},
										name: "PropertyName",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 231, col: 5, offset: 6093},
						run: (*parser).callonVariable9,
						expr: &labeledExpr{
							pos:   position{line: 231, col: 5, offset: 6093},
							label: "vn",
							expr: &ruleRefExpr{
								pos:  position{line: 231, col: 8, offset: 6096},
								name: "VariableName",
							},
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 236, col: 1, offset: 6169},
			expr: &actionExpr{
				pos: position{line: 236, col: 16, offset: 6186},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 236, col: 16, offset: 6186},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 236, col: 16, offset: 6186},
							val:        "$",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 236, col: 20, offset: 6190},
							expr: &charClassMatcher{
								pos:        position{line: 236, col: 20, offset: 6190},
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
			pos:  position{line: 244, col: 1, offset: 6513},
			expr: &seqExpr{
				pos: position{line: 244, col: 12, offset: 6526},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 244, col: 12, offset: 6526},
						name: "PropertyName",
					},
					&litMatcher{
						pos:        position{line: 244, col: 25, offset: 6539},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 244, col: 29, offset: 6543},
						name: "Value",
					},
				},
			},
		},
		{
			name: "PropertyName",
			pos:  position{line: 245, col: 1, offset: 6549},
			expr: &ruleRefExpr{
				pos:  position{line: 245, col: 16, offset: 6566},
				name: "Name",
			},
		},
		{
			name: "Directives",
			pos:  position{line: 247, col: 1, offset: 6572},
			expr: &actionExpr{
				pos: position{line: 247, col: 14, offset: 6587},
				run: (*parser).callonDirectives1,
				expr: &labeledExpr{
					pos:   position{line: 247, col: 14, offset: 6587},
					label: "ds",
					expr: &oneOrMoreExpr{
						pos: position{line: 247, col: 17, offset: 6590},
						expr: &ruleRefExpr{
							pos:  position{line: 247, col: 17, offset: 6590},
							name: "Directive",
						},
					},
				},
			},
		},
		{
			name: "Directive",
			pos:  position{line: 254, col: 1, offset: 6739},
			expr: &actionExpr{
				pos: position{line: 254, col: 13, offset: 6753},
				run: (*parser).callonDirective1,
				expr: &seqExpr{
					pos: position{line: 254, col: 13, offset: 6753},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 254, col: 13, offset: 6753},
							val:        "@",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 254, col: 17, offset: 6757},
							label: "d",
							expr: &choiceExpr{
								pos: position{line: 254, col: 20, offset: 6760},
								alternatives: []interface{}{
									&actionExpr{
										pos: position{line: 254, col: 20, offset: 6760},
										run: (*parser).callonDirective6,
										expr: &seqExpr{
											pos: position{line: 254, col: 21, offset: 6761},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 254, col: 21, offset: 6761},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 254, col: 24, offset: 6764},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 254, col: 38, offset: 6778},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 254, col: 42, offset: 6782},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 254, col: 44, offset: 6784},
													label: "v",
													expr: &ruleRefExpr{
														pos:  position{line: 254, col: 46, offset: 6786},
														name: "Value",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 260, col: 5, offset: 6896},
										run: (*parser).callonDirective14,
										expr: &seqExpr{
											pos: position{line: 260, col: 6, offset: 6897},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 260, col: 6, offset: 6897},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 260, col: 9, offset: 6900},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 260, col: 23, offset: 6914},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 260, col: 27, offset: 6918},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 260, col: 29, offset: 6920},
													label: "t",
													expr: &ruleRefExpr{
														pos:  position{line: 260, col: 31, offset: 6922},
														name: "Type",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 266, col: 5, offset: 7031},
										run: (*parser).callonDirective22,
										expr: &labeledExpr{
											pos:   position{line: 266, col: 5, offset: 7031},
											label: "dn",
											expr: &ruleRefExpr{
												pos:  position{line: 266, col: 8, offset: 7034},
												name: "DirectiveName",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 270, col: 5, offset: 7110},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DirectiveName",
			pos:  position{line: 274, col: 1, offset: 7132},
			expr: &ruleRefExpr{
				pos:  position{line: 274, col: 17, offset: 7150},
				name: "Name",
			},
		},
		{
			name: "Type",
			pos:  position{line: 276, col: 1, offset: 7156},
			expr: &actionExpr{
				pos: position{line: 276, col: 8, offset: 7165},
				run: (*parser).callonType1,
				expr: &labeledExpr{
					pos:   position{line: 276, col: 8, offset: 7165},
					label: "t",
					expr: &choiceExpr{
						pos: position{line: 276, col: 11, offset: 7168},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 276, col: 11, offset: 7168},
								name: "OptionalType",
							},
							&ruleRefExpr{
								pos:  position{line: 276, col: 26, offset: 7183},
								name: "GenericType",
							},
						},
					},
				},
			},
		},
		{
			name: "OptionalType",
			pos:  position{line: 277, col: 1, offset: 7214},
			expr: &actionExpr{
				pos: position{line: 277, col: 16, offset: 7231},
				run: (*parser).callonOptionalType1,
				expr: &seqExpr{
					pos: position{line: 277, col: 16, offset: 7231},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 277, col: 16, offset: 7231},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 277, col: 18, offset: 7233},
								name: "GenericType",
							},
						},
						&litMatcher{
							pos:        position{line: 277, col: 30, offset: 7245},
							val:        "?",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "GenericType",
			pos:  position{line: 282, col: 1, offset: 7316},
			expr: &actionExpr{
				pos: position{line: 282, col: 15, offset: 7332},
				run: (*parser).callonGenericType1,
				expr: &seqExpr{
					pos: position{line: 282, col: 15, offset: 7332},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 282, col: 15, offset: 7332},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 18, offset: 7335},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 282, col: 27, offset: 7344},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 282, col: 29, offset: 7346},
							label: "tps",
							expr: &zeroOrOneExpr{
								pos: position{line: 282, col: 33, offset: 7350},
								expr: &ruleRefExpr{
									pos:  position{line: 282, col: 33, offset: 7350},
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
			pos:  position{line: 287, col: 1, offset: 7417},
			expr: &seqExpr{
				pos: position{line: 287, col: 14, offset: 7432},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 287, col: 14, offset: 7432},
						val:        ":",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 287, col: 18, offset: 7436},
						val:        "<",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 287, col: 22, offset: 7440},
						expr: &ruleRefExpr{
							pos:  position{line: 287, col: 22, offset: 7440},
							name: "Type",
						},
					},
					&litMatcher{
						pos:        position{line: 287, col: 28, offset: 7446},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TypeName",
			pos:  position{line: 288, col: 1, offset: 7450},
			expr: &ruleRefExpr{
				pos:  position{line: 288, col: 12, offset: 7463},
				name: "Name",
			},
		},
		{
			name: "TypeDefinition",
			pos:  position{line: 289, col: 1, offset: 7468},
			expr: &actionExpr{
				pos: position{line: 289, col: 18, offset: 7487},
				run: (*parser).callonTypeDefinition1,
				expr: &seqExpr{
					pos: position{line: 289, col: 18, offset: 7487},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 289, col: 18, offset: 7487},
							val:        "type",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 25, offset: 7494},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 27, offset: 7496},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 30, offset: 7499},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 39, offset: 7508},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 41, offset: 7510},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 289, col: 44, offset: 7513},
								expr: &ruleRefExpr{
									pos:  position{line: 289, col: 44, offset: 7513},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 56, offset: 7525},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 289, col: 58, offset: 7527},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 289, col: 62, offset: 7531},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 289, col: 66, offset: 7535},
								expr: &ruleRefExpr{
									pos:  position{line: 289, col: 66, offset: 7535},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 289, col: 83, offset: 7552},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeExtension",
			pos:  position{line: 309, col: 1, offset: 8033},
			expr: &actionExpr{
				pos: position{line: 309, col: 17, offset: 8051},
				run: (*parser).callonTypeExtension1,
				expr: &seqExpr{
					pos: position{line: 309, col: 17, offset: 8051},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 309, col: 17, offset: 8051},
							val:        "extend",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 309, col: 26, offset: 8060},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 309, col: 28, offset: 8062},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 309, col: 31, offset: 8065},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 309, col: 40, offset: 8074},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 309, col: 42, offset: 8076},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 309, col: 45, offset: 8079},
								expr: &ruleRefExpr{
									pos:  position{line: 309, col: 45, offset: 8079},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 309, col: 57, offset: 8091},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 309, col: 59, offset: 8093},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 309, col: 63, offset: 8097},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 309, col: 67, offset: 8101},
								expr: &ruleRefExpr{
									pos:  position{line: 309, col: 67, offset: 8101},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 309, col: 84, offset: 8118},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Interfaces",
			pos:  position{line: 329, col: 1, offset: 8598},
			expr: &actionExpr{
				pos: position{line: 329, col: 14, offset: 8613},
				run: (*parser).callonInterfaces1,
				expr: &oneOrMoreExpr{
					pos: position{line: 329, col: 14, offset: 8613},
					expr: &ruleRefExpr{
						pos:  position{line: 329, col: 14, offset: 8613},
						name: "GenericType",
					},
				},
			},
		},
		{
			name: "FieldDefinition",
			pos:  position{line: 333, col: 1, offset: 8720},
			expr: &actionExpr{
				pos: position{line: 333, col: 19, offset: 8740},
				run: (*parser).callonFieldDefinition1,
				expr: &seqExpr{
					pos: position{line: 333, col: 19, offset: 8740},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 333, col: 19, offset: 8740},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 333, col: 21, offset: 8742},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 333, col: 24, offset: 8745},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 333, col: 34, offset: 8755},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 333, col: 36, offset: 8757},
							label: "args",
							expr: &zeroOrOneExpr{
								pos: position{line: 333, col: 41, offset: 8762},
								expr: &ruleRefExpr{
									pos:  position{line: 333, col: 41, offset: 8762},
									name: "ArgumentDefinitions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 333, col: 62, offset: 8783},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 333, col: 64, offset: 8785},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 333, col: 68, offset: 8789},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 333, col: 70, offset: 8791},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 333, col: 72, offset: 8793},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 333, col: 77, offset: 8798},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinitions",
			pos:  position{line: 344, col: 1, offset: 9035},
			expr: &actionExpr{
				pos: position{line: 344, col: 23, offset: 9059},
				run: (*parser).callonArgumentDefinitions1,
				expr: &seqExpr{
					pos: position{line: 344, col: 23, offset: 9059},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 344, col: 23, offset: 9059},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 344, col: 27, offset: 9063},
							label: "args",
							expr: &oneOrMoreExpr{
								pos: position{line: 344, col: 32, offset: 9068},
								expr: &ruleRefExpr{
									pos:  position{line: 344, col: 32, offset: 9068},
									name: "ArgumentDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 344, col: 52, offset: 9088},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinition",
			pos:  position{line: 351, col: 1, offset: 9251},
			expr: &actionExpr{
				pos: position{line: 351, col: 22, offset: 9274},
				run: (*parser).callonArgumentDefinition1,
				expr: &seqExpr{
					pos: position{line: 351, col: 22, offset: 9274},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 351, col: 22, offset: 9274},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 351, col: 24, offset: 9276},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 351, col: 27, offset: 9279},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 351, col: 40, offset: 9292},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 351, col: 42, offset: 9294},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 351, col: 46, offset: 9298},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 351, col: 48, offset: 9300},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 351, col: 50, offset: 9302},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 351, col: 55, offset: 9307},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 351, col: 57, offset: 9309},
							label: "dv",
							expr: &zeroOrOneExpr{
								pos: position{line: 351, col: 60, offset: 9312},
								expr: &ruleRefExpr{
									pos:  position{line: 351, col: 60, offset: 9312},
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
			pos:  position{line: 363, col: 1, offset: 9543},
			expr: &actionExpr{
				pos: position{line: 363, col: 18, offset: 9562},
				run: (*parser).callonEnumDefinition1,
				expr: &seqExpr{
					pos: position{line: 363, col: 18, offset: 9562},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 363, col: 18, offset: 9562},
							val:        "enum",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 363, col: 25, offset: 9569},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 363, col: 27, offset: 9571},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 363, col: 30, offset: 9574},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 363, col: 39, offset: 9583},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 363, col: 41, offset: 9585},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 363, col: 45, offset: 9589},
							label: "vals",
							expr: &oneOrMoreExpr{
								pos: position{line: 363, col: 50, offset: 9594},
								expr: &ruleRefExpr{
									pos:  position{line: 363, col: 50, offset: 9594},
									name: "EnumValueName",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 363, col: 65, offset: 9609},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EnumValueName",
			pos:  position{line: 373, col: 1, offset: 9790},
			expr: &actionExpr{
				pos: position{line: 373, col: 17, offset: 9808},
				run: (*parser).callonEnumValueName1,
				expr: &seqExpr{
					pos: position{line: 373, col: 17, offset: 9808},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 373, col: 17, offset: 9808},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 373, col: 19, offset: 9810},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 373, col: 21, offset: 9812},
								name: "Name",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 373, col: 26, offset: 9817},
							name: "_",
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"ignored\"",
			pos:         position{line: 375, col: 1, offset: 9838},
			expr: &actionExpr{
				pos: position{line: 375, col: 15, offset: 9854},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 375, col: 15, offset: 9854},
					expr: &choiceExpr{
						pos: position{line: 375, col: 16, offset: 9855},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 375, col: 16, offset: 9855},
								name: "whitespace",
							},
							&ruleRefExpr{
								pos:  position{line: 375, col: 29, offset: 9868},
								name: "Comment",
							},
							&litMatcher{
								pos:        position{line: 375, col: 39, offset: 9878},
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
			pos:  position{line: 376, col: 1, offset: 9904},
			expr: &charClassMatcher{
				pos:        position{line: 376, col: 14, offset: 9919},
				val:        "[ \\n\\t\\r]",
				chars:      []rune{' ', '\n', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 378, col: 1, offset: 9930},
			expr: &notExpr{
				pos: position{line: 378, col: 7, offset: 9938},
				expr: &anyMatcher{
					line: 378, col: 8, offset: 9939,
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
	)
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
