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
			pos:  position{line: 15, col: 1, offset: 164},
			expr: &choiceExpr{
				pos: position{line: 15, col: 12, offset: 177},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 15, col: 12, offset: 177},
						run: (*parser).callonDocument2,
						expr: &seqExpr{
							pos: position{line: 15, col: 12, offset: 177},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 15, col: 12, offset: 177},
									label: "d",
									expr: &choiceExpr{
										pos: position{line: 15, col: 15, offset: 180},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 15, col: 15, offset: 180},
												name: "Schema",
											},
											&ruleRefExpr{
												pos:  position{line: 15, col: 24, offset: 189},
												name: "QueryDocument",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 39, offset: 204},
									name: "EOF",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 17, col: 5, offset: 229},
						run: (*parser).callonDocument9,
						expr: &anyMatcher{
							line: 17, col: 5, offset: 229,
						},
					},
				},
			},
		},
		{
			name: "QueryDocument",
			pos:  position{line: 21, col: 1, offset: 328},
			expr: &actionExpr{
				pos: position{line: 21, col: 17, offset: 346},
				run: (*parser).callonQueryDocument1,
				expr: &labeledExpr{
					pos:   position{line: 21, col: 17, offset: 346},
					label: "defs",
					expr: &oneOrMoreExpr{
						pos: position{line: 21, col: 22, offset: 351},
						expr: &ruleRefExpr{
							pos:  position{line: 21, col: 22, offset: 351},
							name: "Definition",
						},
					},
				},
			},
		},
		{
			name: "Schema",
			pos:  position{line: 47, col: 1, offset: 1230},
			expr: &actionExpr{
				pos: position{line: 47, col: 10, offset: 1241},
				run: (*parser).callonSchema1,
				expr: &labeledExpr{
					pos:   position{line: 47, col: 10, offset: 1241},
					label: "defs",
					expr: &oneOrMoreExpr{
						pos: position{line: 47, col: 15, offset: 1246},
						expr: &ruleRefExpr{
							pos:  position{line: 47, col: 15, offset: 1246},
							name: "SchemaDefinition",
						},
					},
				},
			},
		},
		{
			name: "SchemaDefinition",
			pos:  position{line: 69, col: 1, offset: 1888},
			expr: &actionExpr{
				pos: position{line: 69, col: 20, offset: 1909},
				run: (*parser).callonSchemaDefinition1,
				expr: &seqExpr{
					pos: position{line: 69, col: 20, offset: 1909},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 69, col: 20, offset: 1909},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 69, col: 22, offset: 1911},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 69, col: 25, offset: 1914},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 69, col: 25, offset: 1914},
										name: "TypeDefinition",
									},
									&ruleRefExpr{
										pos:  position{line: 69, col: 42, offset: 1931},
										name: "TypeExtension",
									},
									&ruleRefExpr{
										pos:  position{line: 69, col: 58, offset: 1947},
										name: "EnumDefinition",
									},
									&ruleRefExpr{
										pos:  position{line: 69, col: 75, offset: 1964},
										name: "Comment",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 84, offset: 1973},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Definition",
			pos:  position{line: 73, col: 1, offset: 1998},
			expr: &choiceExpr{
				pos: position{line: 73, col: 14, offset: 2013},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 73, col: 14, offset: 2013},
						run: (*parser).callonDefinition2,
						expr: &seqExpr{
							pos: position{line: 73, col: 14, offset: 2013},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 73, col: 14, offset: 2013},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 73, col: 16, offset: 2015},
									label: "d",
									expr: &choiceExpr{
										pos: position{line: 73, col: 19, offset: 2018},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 73, col: 19, offset: 2018},
												name: "OperationDefinition",
											},
											&ruleRefExpr{
												pos:  position{line: 73, col: 41, offset: 2040},
												name: "FragmentDefinition",
											},
											&ruleRefExpr{
												pos:  position{line: 73, col: 62, offset: 2061},
												name: "Comment",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 73, col: 71, offset: 2070},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 75, col: 5, offset: 2093},
						run: (*parser).callonDefinition11,
						expr: &anyMatcher{
							line: 75, col: 5, offset: 2093,
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 79, col: 1, offset: 2174},
			expr: &actionExpr{
				pos: position{line: 79, col: 11, offset: 2186},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 79, col: 11, offset: 2186},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 79, col: 11, offset: 2186},
							val:        "#",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 79, col: 15, offset: 2190},
							expr: &charClassMatcher{
								pos:        position{line: 79, col: 15, offset: 2190},
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
			name: "OperationDefinition",
			pos:  position{line: 81, col: 1, offset: 2229},
			expr: &choiceExpr{
				pos: position{line: 81, col: 23, offset: 2253},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 81, col: 23, offset: 2253},
						run: (*parser).callonOperationDefinition2,
						expr: &labeledExpr{
							pos:   position{line: 81, col: 23, offset: 2253},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 28, offset: 2258},
								name: "SelectionSet",
							},
						},
					},
					&actionExpr{
						pos: position{line: 88, col: 13, offset: 2402},
						run: (*parser).callonOperationDefinition5,
						expr: &seqExpr{
							pos: position{line: 88, col: 14, offset: 2403},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 88, col: 14, offset: 2403},
									label: "ot",
									expr: &ruleRefExpr{
										pos:  position{line: 88, col: 17, offset: 2406},
										name: "OperationType",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 31, offset: 2420},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 33, offset: 2422},
									label: "on",
									expr: &ruleRefExpr{
										pos:  position{line: 88, col: 36, offset: 2425},
										name: "OperationName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 50, offset: 2439},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 52, offset: 2441},
									label: "vds",
									expr: &zeroOrOneExpr{
										pos: position{line: 88, col: 56, offset: 2445},
										expr: &ruleRefExpr{
											pos:  position{line: 88, col: 56, offset: 2445},
											name: "VariableDefinitions",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 77, offset: 2466},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 79, offset: 2468},
									label: "ds",
									expr: &zeroOrOneExpr{
										pos: position{line: 88, col: 82, offset: 2471},
										expr: &ruleRefExpr{
											pos:  position{line: 88, col: 82, offset: 2471},
											name: "Directives",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 88, col: 94, offset: 2483},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 88, col: 96, offset: 2485},
									label: "sels",
									expr: &ruleRefExpr{
										pos:  position{line: 88, col: 101, offset: 2490},
										name: "SelectionSet",
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
			pos:  position{line: 108, col: 1, offset: 2911},
			expr: &choiceExpr{
				pos: position{line: 108, col: 17, offset: 2929},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 108, col: 17, offset: 2929},
						run: (*parser).callonOperationType2,
						expr: &litMatcher{
							pos:        position{line: 108, col: 17, offset: 2929},
							val:        "query",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 109, col: 18, offset: 2993},
						run: (*parser).callonOperationType4,
						expr: &litMatcher{
							pos:        position{line: 109, col: 18, offset: 2993},
							val:        "mutation",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperationName",
			pos:  position{line: 110, col: 1, offset: 3046},
			expr: &actionExpr{
				pos: position{line: 110, col: 17, offset: 3064},
				run: (*parser).callonOperationName1,
				expr: &ruleRefExpr{
					pos:  position{line: 110, col: 17, offset: 3064},
					name: "Name",
				},
			},
		},
		{
			name: "VariableDefinitions",
			pos:  position{line: 113, col: 1, offset: 3101},
			expr: &actionExpr{
				pos: position{line: 113, col: 23, offset: 3125},
				run: (*parser).callonVariableDefinitions1,
				expr: &seqExpr{
					pos: position{line: 113, col: 23, offset: 3125},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 113, col: 23, offset: 3125},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 113, col: 27, offset: 3129},
							label: "vds",
							expr: &oneOrMoreExpr{
								pos: position{line: 113, col: 31, offset: 3133},
								expr: &ruleRefExpr{
									pos:  position{line: 113, col: 31, offset: 3133},
									name: "VariableDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 113, col: 51, offset: 3153},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariableDefinition",
			pos:  position{line: 120, col: 1, offset: 3314},
			expr: &actionExpr{
				pos: position{line: 120, col: 22, offset: 3337},
				run: (*parser).callonVariableDefinition1,
				expr: &seqExpr{
					pos: position{line: 120, col: 22, offset: 3337},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 120, col: 22, offset: 3337},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 120, col: 24, offset: 3339},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 120, col: 26, offset: 3341},
								name: "Variable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 120, col: 35, offset: 3350},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 120, col: 37, offset: 3352},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 120, col: 41, offset: 3356},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 120, col: 43, offset: 3358},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 120, col: 45, offset: 3360},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 120, col: 50, offset: 3365},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 120, col: 52, offset: 3367},
							label: "d",
							expr: &zeroOrOneExpr{
								pos: position{line: 120, col: 54, offset: 3369},
								expr: &ruleRefExpr{
									pos:  position{line: 120, col: 54, offset: 3369},
									name: "DefaultValue",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 120, col: 68, offset: 3383},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DefaultValue",
			pos:  position{line: 132, col: 1, offset: 3619},
			expr: &actionExpr{
				pos: position{line: 132, col: 16, offset: 3636},
				run: (*parser).callonDefaultValue1,
				expr: &seqExpr{
					pos: position{line: 132, col: 16, offset: 3636},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 132, col: 16, offset: 3636},
							val:        "=",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 132, col: 20, offset: 3640},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 132, col: 22, offset: 3642},
								name: "Value",
							},
						},
					},
				},
			},
		},
		{
			name: "SelectionSet",
			pos:  position{line: 134, col: 1, offset: 3667},
			expr: &actionExpr{
				pos: position{line: 134, col: 16, offset: 3684},
				run: (*parser).callonSelectionSet1,
				expr: &seqExpr{
					pos: position{line: 134, col: 16, offset: 3684},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 134, col: 16, offset: 3684},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 134, col: 20, offset: 3688},
							label: "s",
							expr: &oneOrMoreExpr{
								pos: position{line: 134, col: 23, offset: 3691},
								expr: &ruleRefExpr{
									pos:  position{line: 134, col: 23, offset: 3691},
									name: "Selection",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 134, col: 35, offset: 3703},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Selection",
			pos:  position{line: 145, col: 1, offset: 3969},
			expr: &choiceExpr{
				pos: position{line: 145, col: 13, offset: 3983},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 145, col: 13, offset: 3983},
						run: (*parser).callonSelection2,
						expr: &seqExpr{
							pos: position{line: 145, col: 14, offset: 3984},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 145, col: 14, offset: 3984},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 145, col: 16, offset: 3986},
									label: "f",
									expr: &ruleRefExpr{
										pos:  position{line: 145, col: 18, offset: 3988},
										name: "Field",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 145, col: 24, offset: 3994},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 148, col: 5, offset: 4076},
						run: (*parser).callonSelection8,
						expr: &seqExpr{
							pos: position{line: 148, col: 6, offset: 4077},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 148, col: 6, offset: 4077},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 148, col: 8, offset: 4079},
									label: "fs",
									expr: &ruleRefExpr{
										pos:  position{line: 148, col: 11, offset: 4082},
										name: "FragmentSpread",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 148, col: 26, offset: 4097},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 151, col: 5, offset: 4216},
						run: (*parser).callonSelection14,
						expr: &seqExpr{
							pos: position{line: 151, col: 6, offset: 4217},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 151, col: 6, offset: 4217},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 151, col: 8, offset: 4219},
									label: "fs",
									expr: &ruleRefExpr{
										pos:  position{line: 151, col: 11, offset: 4222},
										name: "InlineFragment",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 151, col: 26, offset: 4237},
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
			pos:  position{line: 156, col: 1, offset: 4355},
			expr: &actionExpr{
				pos: position{line: 156, col: 9, offset: 4365},
				run: (*parser).callonField1,
				expr: &seqExpr{
					pos: position{line: 156, col: 9, offset: 4365},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 156, col: 9, offset: 4365},
							label: "fa",
							expr: &zeroOrOneExpr{
								pos: position{line: 156, col: 12, offset: 4368},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 12, offset: 4368},
									name: "FieldAlias",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 156, col: 24, offset: 4380},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 156, col: 26, offset: 4382},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 156, col: 29, offset: 4385},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 156, col: 39, offset: 4395},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 156, col: 41, offset: 4397},
							label: "as",
							expr: &zeroOrOneExpr{
								pos: position{line: 156, col: 44, offset: 4400},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 44, offset: 4400},
									name: "Arguments",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 156, col: 55, offset: 4411},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 156, col: 57, offset: 4413},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 156, col: 60, offset: 4416},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 60, offset: 4416},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 156, col: 72, offset: 4428},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 156, col: 74, offset: 4430},
							label: "sels",
							expr: &zeroOrOneExpr{
								pos: position{line: 156, col: 79, offset: 4435},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 79, offset: 4435},
									name: "SelectionSet",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FieldAlias",
			pos:  position{line: 183, col: 1, offset: 4958},
			expr: &actionExpr{
				pos: position{line: 183, col: 14, offset: 4973},
				run: (*parser).callonFieldAlias1,
				expr: &seqExpr{
					pos: position{line: 183, col: 14, offset: 4973},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 14, offset: 4973},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 16, offset: 4975},
								name: "Name",
							},
						},
						&litMatcher{
							pos:        position{line: 183, col: 21, offset: 4980},
							val:        ":",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "FieldName",
			pos:  position{line: 184, col: 1, offset: 5002},
			expr: &ruleRefExpr{
				pos:  position{line: 184, col: 13, offset: 5016},
				name: "Name",
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 185, col: 1, offset: 5021},
			expr: &actionExpr{
				pos: position{line: 185, col: 13, offset: 5035},
				run: (*parser).callonArguments1,
				expr: &seqExpr{
					pos: position{line: 185, col: 13, offset: 5035},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 185, col: 13, offset: 5035},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 185, col: 17, offset: 5039},
							label: "args",
							expr: &zeroOrMoreExpr{
								pos: position{line: 185, col: 23, offset: 5045},
								expr: &ruleRefExpr{
									pos:  position{line: 185, col: 23, offset: 5045},
									name: "Argument",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 185, col: 34, offset: 5056},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Argument",
			pos:  position{line: 196, col: 1, offset: 5301},
			expr: &actionExpr{
				pos: position{line: 196, col: 12, offset: 5314},
				run: (*parser).callonArgument1,
				expr: &seqExpr{
					pos: position{line: 196, col: 12, offset: 5314},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 196, col: 12, offset: 5314},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 196, col: 14, offset: 5316},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 17, offset: 5319},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 30, offset: 5332},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 196, col: 32, offset: 5334},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 36, offset: 5338},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 196, col: 38, offset: 5340},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 40, offset: 5342},
								name: "Value",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 46, offset: 5348},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentName",
			pos:  position{line: 202, col: 1, offset: 5421},
			expr: &ruleRefExpr{
				pos:  position{line: 202, col: 16, offset: 5438},
				name: "Name",
			},
		},
		{
			name: "Name",
			pos:  position{line: 204, col: 1, offset: 5444},
			expr: &actionExpr{
				pos: position{line: 204, col: 8, offset: 5453},
				run: (*parser).callonName1,
				expr: &seqExpr{
					pos: position{line: 204, col: 8, offset: 5453},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 204, col: 8, offset: 5453},
							val:        "[a-z_]i",
							chars:      []rune{'_'},
							ranges:     []rune{'a', 'z'},
							ignoreCase: true,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 204, col: 16, offset: 5461},
							expr: &charClassMatcher{
								pos:        position{line: 204, col: 16, offset: 5461},
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
			pos:  position{line: 208, col: 1, offset: 5506},
			expr: &actionExpr{
				pos: position{line: 208, col: 19, offset: 5526},
				run: (*parser).callonFragmentSpread1,
				expr: &seqExpr{
					pos: position{line: 208, col: 19, offset: 5526},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 208, col: 19, offset: 5526},
							val:        "...",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 208, col: 25, offset: 5532},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 208, col: 28, offset: 5535},
								name: "FragmentName",
							},
						},
						&labeledExpr{
							pos:   position{line: 208, col: 41, offset: 5548},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 208, col: 44, offset: 5551},
								expr: &ruleRefExpr{
									pos:  position{line: 208, col: 44, offset: 5551},
									name: "Directives",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineFragment",
			pos:  position{line: 219, col: 1, offset: 5751},
			expr: &actionExpr{
				pos: position{line: 219, col: 19, offset: 5771},
				run: (*parser).callonInlineFragment1,
				expr: &seqExpr{
					pos: position{line: 219, col: 19, offset: 5771},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 219, col: 19, offset: 5771},
							val:        "...",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 219, col: 25, offset: 5777},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 219, col: 27, offset: 5779},
							val:        "on",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 219, col: 32, offset: 5784},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 219, col: 34, offset: 5786},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 219, col: 37, offset: 5789},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 219, col: 46, offset: 5798},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 219, col: 48, offset: 5800},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 219, col: 51, offset: 5803},
								expr: &ruleRefExpr{
									pos:  position{line: 219, col: 51, offset: 5803},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 219, col: 63, offset: 5815},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 219, col: 65, offset: 5817},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 219, col: 70, offset: 5822},
								name: "SelectionSet",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 219, col: 83, offset: 5835},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FragmentDefinition",
			pos:  position{line: 231, col: 1, offset: 6079},
			expr: &actionExpr{
				pos: position{line: 231, col: 22, offset: 6102},
				run: (*parser).callonFragmentDefinition1,
				expr: &seqExpr{
					pos: position{line: 231, col: 22, offset: 6102},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 231, col: 22, offset: 6102},
							val:        "fragment",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 33, offset: 6113},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 231, col: 35, offset: 6115},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 231, col: 38, offset: 6118},
								name: "FragmentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 51, offset: 6131},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 231, col: 53, offset: 6133},
							val:        "on",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 58, offset: 6138},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 231, col: 60, offset: 6140},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 231, col: 63, offset: 6143},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 73, offset: 6153},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 231, col: 75, offset: 6155},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 231, col: 78, offset: 6158},
								expr: &ruleRefExpr{
									pos:  position{line: 231, col: 78, offset: 6158},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 90, offset: 6170},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 231, col: 92, offset: 6172},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 231, col: 97, offset: 6177},
								name: "SelectionSet",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 110, offset: 6190},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FragmentName",
			pos:  position{line: 243, col: 1, offset: 6457},
			expr: &actionExpr{
				pos: position{line: 243, col: 16, offset: 6474},
				run: (*parser).callonFragmentName1,
				expr: &labeledExpr{
					pos:   position{line: 243, col: 16, offset: 6474},
					label: "n",
					expr: &ruleRefExpr{
						pos:  position{line: 243, col: 18, offset: 6476},
						name: "Name",
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 245, col: 1, offset: 6500},
			expr: &actionExpr{
				pos: position{line: 245, col: 9, offset: 6510},
				run: (*parser).callonValue1,
				expr: &seqExpr{
					pos: position{line: 245, col: 9, offset: 6510},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 245, col: 9, offset: 6510},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 245, col: 11, offset: 6512},
							label: "v",
							expr: &choiceExpr{
								pos: position{line: 245, col: 14, offset: 6515},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 245, col: 14, offset: 6515},
										name: "Null",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 21, offset: 6522},
										name: "Boolean",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 31, offset: 6532},
										name: "Int",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 37, offset: 6538},
										name: "Float",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 45, offset: 6546},
										name: "String",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 54, offset: 6555},
										name: "EnumValue",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 66, offset: 6567},
										name: "Array",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 74, offset: 6575},
										name: "Object",
									},
									&ruleRefExpr{
										pos:  position{line: 245, col: 83, offset: 6584},
										name: "Variable",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 245, col: 93, offset: 6594},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Null",
			pos:  position{line: 249, col: 1, offset: 6616},
			expr: &actionExpr{
				pos: position{line: 249, col: 8, offset: 6625},
				run: (*parser).callonNull1,
				expr: &litMatcher{
					pos:        position{line: 249, col: 8, offset: 6625},
					val:        "null",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Boolean",
			pos:  position{line: 250, col: 1, offset: 6652},
			expr: &choiceExpr{
				pos: position{line: 250, col: 11, offset: 6664},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 250, col: 11, offset: 6664},
						run: (*parser).callonBoolean2,
						expr: &litMatcher{
							pos:        position{line: 250, col: 11, offset: 6664},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 250, col: 41, offset: 6694},
						run: (*parser).callonBoolean4,
						expr: &litMatcher{
							pos:        position{line: 250, col: 41, offset: 6694},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Int",
			pos:  position{line: 251, col: 1, offset: 6724},
			expr: &actionExpr{
				pos: position{line: 251, col: 7, offset: 6732},
				run: (*parser).callonInt1,
				expr: &seqExpr{
					pos: position{line: 251, col: 7, offset: 6732},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 251, col: 7, offset: 6732},
							expr: &ruleRefExpr{
								pos:  position{line: 251, col: 7, offset: 6732},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 251, col: 13, offset: 6738},
							name: "IntegerPart",
						},
					},
				},
			},
		},
		{
			name: "Float",
			pos:  position{line: 254, col: 1, offset: 6791},
			expr: &actionExpr{
				pos: position{line: 254, col: 9, offset: 6801},
				run: (*parser).callonFloat1,
				expr: &seqExpr{
					pos: position{line: 254, col: 9, offset: 6801},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 254, col: 9, offset: 6801},
							expr: &ruleRefExpr{
								pos:  position{line: 254, col: 9, offset: 6801},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 254, col: 15, offset: 6807},
							name: "IntegerPart",
						},
						&litMatcher{
							pos:        position{line: 254, col: 27, offset: 6819},
							val:        ".",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 254, col: 31, offset: 6823},
							expr: &ruleRefExpr{
								pos:  position{line: 254, col: 31, offset: 6823},
								name: "Digit",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 254, col: 38, offset: 6830},
							expr: &ruleRefExpr{
								pos:  position{line: 254, col: 38, offset: 6830},
								name: "ExponentPart",
							},
						},
					},
				},
			},
		},
		{
			name: "Sign",
			pos:  position{line: 257, col: 1, offset: 6895},
			expr: &litMatcher{
				pos:        position{line: 257, col: 8, offset: 6904},
				val:        "-",
				ignoreCase: false,
			},
		},
		{
			name: "IntegerPart",
			pos:  position{line: 258, col: 1, offset: 6908},
			expr: &choiceExpr{
				pos: position{line: 258, col: 15, offset: 6924},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 258, col: 15, offset: 6924},
						val:        "0",
						ignoreCase: false,
					},
					&seqExpr{
						pos: position{line: 258, col: 21, offset: 6930},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 258, col: 21, offset: 6930},
								name: "NonZeroDigit",
							},
							&zeroOrMoreExpr{
								pos: position{line: 258, col: 34, offset: 6943},
								expr: &ruleRefExpr{
									pos:  position{line: 258, col: 34, offset: 6943},
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
			pos:  position{line: 259, col: 1, offset: 6950},
			expr: &seqExpr{
				pos: position{line: 259, col: 16, offset: 6967},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 259, col: 16, offset: 6967},
						val:        "e",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 259, col: 20, offset: 6971},
						expr: &ruleRefExpr{
							pos:  position{line: 259, col: 20, offset: 6971},
							name: "Sign",
						},
					},
					&oneOrMoreExpr{
						pos: position{line: 259, col: 26, offset: 6977},
						expr: &ruleRefExpr{
							pos:  position{line: 259, col: 26, offset: 6977},
							name: "Digit",
						},
					},
				},
			},
		},
		{
			name: "Digit",
			pos:  position{line: 260, col: 1, offset: 6984},
			expr: &charClassMatcher{
				pos:        position{line: 260, col: 9, offset: 6994},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDigit",
			pos:  position{line: 261, col: 1, offset: 7000},
			expr: &charClassMatcher{
				pos:        position{line: 261, col: 16, offset: 7017},
				val:        "[123456789]",
				chars:      []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "String",
			pos:  position{line: 262, col: 1, offset: 7029},
			expr: &actionExpr{
				pos: position{line: 262, col: 10, offset: 7040},
				run: (*parser).callonString1,
				expr: &seqExpr{
					pos: position{line: 262, col: 10, offset: 7040},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 262, col: 10, offset: 7040},
							val:        "\"",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 262, col: 14, offset: 7044},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 262, col: 16, offset: 7046},
								name: "string",
							},
						},
						&litMatcher{
							pos:        position{line: 262, col: 23, offset: 7053},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "string",
			pos:  position{line: 265, col: 1, offset: 7085},
			expr: &actionExpr{
				pos: position{line: 265, col: 10, offset: 7096},
				run: (*parser).callonstring1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 265, col: 10, offset: 7096},
					expr: &ruleRefExpr{
						pos:  position{line: 265, col: 10, offset: 7096},
						name: "StringCharacter",
					},
				},
			},
		},
		{
			name: "StringCharacter",
			pos:  position{line: 268, col: 1, offset: 7145},
			expr: &choiceExpr{
				pos: position{line: 268, col: 19, offset: 7165},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 268, col: 19, offset: 7165},
						val:        "[^\\\\\"]",
						chars:      []rune{'\\', '"'},
						ignoreCase: false,
						inverted:   true,
					},
					&seqExpr{
						pos: position{line: 268, col: 28, offset: 7174},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 268, col: 28, offset: 7174},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 268, col: 33, offset: 7179},
								name: "EscapedCharacter",
							},
						},
					},
					&seqExpr{
						pos: position{line: 268, col: 52, offset: 7198},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 268, col: 52, offset: 7198},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 268, col: 57, offset: 7203},
								name: "EscapedUnicode",
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedUnicode",
			pos:  position{line: 269, col: 1, offset: 7218},
			expr: &seqExpr{
				pos: position{line: 269, col: 18, offset: 7237},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 269, col: 18, offset: 7237},
						val:        "u",
						ignoreCase: false,
					},
					&charClassMatcher{
						pos:        position{line: 269, col: 22, offset: 7241},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 269, col: 32, offset: 7251},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 269, col: 42, offset: 7261},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 269, col: 52, offset: 7271},
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
			pos:  position{line: 270, col: 1, offset: 7281},
			expr: &choiceExpr{
				pos: position{line: 270, col: 20, offset: 7302},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 270, col: 20, offset: 7302},
						val:        "[\"/bfnrt]",
						chars:      []rune{'"', '/', 'b', 'f', 'n', 'r', 't'},
						ignoreCase: false,
						inverted:   false,
					},
					&litMatcher{
						pos:        position{line: 270, col: 32, offset: 7314},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "EnumValue",
			pos:  position{line: 272, col: 1, offset: 7320},
			expr: &actionExpr{
				pos: position{line: 272, col: 13, offset: 7334},
				run: (*parser).callonEnumValue1,
				expr: &seqExpr{
					pos: position{line: 272, col: 13, offset: 7334},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 272, col: 13, offset: 7334},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 272, col: 16, offset: 7337},
								name: "TypeName",
							},
						},
						&litMatcher{
							pos:        position{line: 272, col: 25, offset: 7346},
							val:        ".",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 272, col: 29, offset: 7350},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 272, col: 31, offset: 7352},
								name: "EnumValueName",
							},
						},
					},
				},
			},
		},
		{
			name: "Array",
			pos:  position{line: 279, col: 1, offset: 7457},
			expr: &actionExpr{
				pos: position{line: 279, col: 9, offset: 7467},
				run: (*parser).callonArray1,
				expr: &seqExpr{
					pos: position{line: 279, col: 9, offset: 7467},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 279, col: 9, offset: 7467},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 279, col: 13, offset: 7471},
							label: "values",
							expr: &zeroOrMoreExpr{
								pos: position{line: 279, col: 20, offset: 7478},
								expr: &ruleRefExpr{
									pos:  position{line: 279, col: 20, offset: 7478},
									name: "Value",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 279, col: 27, offset: 7485},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Object",
			pos:  position{line: 288, col: 1, offset: 7646},
			expr: &actionExpr{
				pos: position{line: 288, col: 10, offset: 7657},
				run: (*parser).callonObject1,
				expr: &seqExpr{
					pos: position{line: 288, col: 10, offset: 7657},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 288, col: 10, offset: 7657},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 288, col: 14, offset: 7661},
							label: "ps",
							expr: &oneOrMoreExpr{
								pos: position{line: 288, col: 17, offset: 7664},
								expr: &ruleRefExpr{
									pos:  position{line: 288, col: 17, offset: 7664},
									name: "Property",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 288, col: 27, offset: 7674},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 300, col: 1, offset: 7927},
			expr: &choiceExpr{
				pos: position{line: 300, col: 12, offset: 7940},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 300, col: 12, offset: 7940},
						run: (*parser).callonVariable2,
						expr: &seqExpr{
							pos: position{line: 300, col: 12, offset: 7940},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 300, col: 12, offset: 7940},
									label: "vn",
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 15, offset: 7943},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 300, col: 28, offset: 7956},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 300, col: 32, offset: 7960},
									label: "pn",
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 35, offset: 7963},
										name: "PropertyName",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 305, col: 5, offset: 8096},
						run: (*parser).callonVariable9,
						expr: &labeledExpr{
							pos:   position{line: 305, col: 5, offset: 8096},
							label: "vn",
							expr: &ruleRefExpr{
								pos:  position{line: 305, col: 8, offset: 8099},
								name: "VariableName",
							},
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 310, col: 1, offset: 8172},
			expr: &actionExpr{
				pos: position{line: 310, col: 16, offset: 8189},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 310, col: 16, offset: 8189},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 310, col: 16, offset: 8189},
							val:        "$",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 310, col: 20, offset: 8193},
							expr: &charClassMatcher{
								pos:        position{line: 310, col: 20, offset: 8193},
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
			pos:  position{line: 318, col: 1, offset: 8516},
			expr: &actionExpr{
				pos: position{line: 318, col: 12, offset: 8529},
				run: (*parser).callonProperty1,
				expr: &seqExpr{
					pos: position{line: 318, col: 12, offset: 8529},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 318, col: 12, offset: 8529},
							label: "pn",
							expr: &ruleRefExpr{
								pos:  position{line: 318, col: 15, offset: 8532},
								name: "PropertyName",
							},
						},
						&litMatcher{
							pos:        position{line: 318, col: 28, offset: 8545},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 318, col: 32, offset: 8549},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 318, col: 34, offset: 8551},
								name: "Value",
							},
						},
					},
				},
			},
		},
		{
			name: "PropertyName",
			pos:  position{line: 321, col: 1, offset: 8621},
			expr: &actionExpr{
				pos: position{line: 321, col: 16, offset: 8638},
				run: (*parser).callonPropertyName1,
				expr: &ruleRefExpr{
					pos:  position{line: 321, col: 16, offset: 8638},
					name: "Name",
				},
			},
		},
		{
			name: "Directives",
			pos:  position{line: 323, col: 1, offset: 8674},
			expr: &actionExpr{
				pos: position{line: 323, col: 14, offset: 8689},
				run: (*parser).callonDirectives1,
				expr: &labeledExpr{
					pos:   position{line: 323, col: 14, offset: 8689},
					label: "ds",
					expr: &oneOrMoreExpr{
						pos: position{line: 323, col: 17, offset: 8692},
						expr: &ruleRefExpr{
							pos:  position{line: 323, col: 17, offset: 8692},
							name: "Directive",
						},
					},
				},
			},
		},
		{
			name: "Directive",
			pos:  position{line: 330, col: 1, offset: 8841},
			expr: &actionExpr{
				pos: position{line: 330, col: 13, offset: 8855},
				run: (*parser).callonDirective1,
				expr: &seqExpr{
					pos: position{line: 330, col: 13, offset: 8855},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 330, col: 13, offset: 8855},
							val:        "@",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 330, col: 17, offset: 8859},
							label: "d",
							expr: &choiceExpr{
								pos: position{line: 330, col: 20, offset: 8862},
								alternatives: []interface{}{
									&actionExpr{
										pos: position{line: 330, col: 20, offset: 8862},
										run: (*parser).callonDirective6,
										expr: &seqExpr{
											pos: position{line: 330, col: 21, offset: 8863},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 330, col: 21, offset: 8863},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 330, col: 24, offset: 8866},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 330, col: 38, offset: 8880},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 330, col: 42, offset: 8884},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 330, col: 44, offset: 8886},
													label: "v",
													expr: &ruleRefExpr{
														pos:  position{line: 330, col: 46, offset: 8888},
														name: "Value",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 336, col: 5, offset: 8998},
										run: (*parser).callonDirective14,
										expr: &seqExpr{
											pos: position{line: 336, col: 6, offset: 8999},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 336, col: 6, offset: 8999},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 336, col: 9, offset: 9002},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 336, col: 23, offset: 9016},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 336, col: 27, offset: 9020},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 336, col: 29, offset: 9022},
													label: "t",
													expr: &ruleRefExpr{
														pos:  position{line: 336, col: 31, offset: 9024},
														name: "Type",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 342, col: 5, offset: 9133},
										run: (*parser).callonDirective22,
										expr: &labeledExpr{
											pos:   position{line: 342, col: 5, offset: 9133},
											label: "dn",
											expr: &ruleRefExpr{
												pos:  position{line: 342, col: 8, offset: 9136},
												name: "DirectiveName",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 346, col: 5, offset: 9212},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DirectiveName",
			pos:  position{line: 350, col: 1, offset: 9234},
			expr: &ruleRefExpr{
				pos:  position{line: 350, col: 17, offset: 9252},
				name: "Name",
			},
		},
		{
			name: "Type",
			pos:  position{line: 352, col: 1, offset: 9258},
			expr: &actionExpr{
				pos: position{line: 352, col: 8, offset: 9267},
				run: (*parser).callonType1,
				expr: &labeledExpr{
					pos:   position{line: 352, col: 8, offset: 9267},
					label: "t",
					expr: &choiceExpr{
						pos: position{line: 352, col: 11, offset: 9270},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 352, col: 11, offset: 9270},
								name: "OptionalType",
							},
							&ruleRefExpr{
								pos:  position{line: 352, col: 26, offset: 9285},
								name: "GenericType",
							},
						},
					},
				},
			},
		},
		{
			name: "OptionalType",
			pos:  position{line: 353, col: 1, offset: 9316},
			expr: &actionExpr{
				pos: position{line: 353, col: 16, offset: 9333},
				run: (*parser).callonOptionalType1,
				expr: &seqExpr{
					pos: position{line: 353, col: 16, offset: 9333},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 353, col: 16, offset: 9333},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 353, col: 18, offset: 9335},
								name: "GenericType",
							},
						},
						&litMatcher{
							pos:        position{line: 353, col: 30, offset: 9347},
							val:        "?",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "GenericType",
			pos:  position{line: 358, col: 1, offset: 9418},
			expr: &actionExpr{
				pos: position{line: 358, col: 15, offset: 9434},
				run: (*parser).callonGenericType1,
				expr: &seqExpr{
					pos: position{line: 358, col: 15, offset: 9434},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 358, col: 15, offset: 9434},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 358, col: 18, offset: 9437},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 358, col: 27, offset: 9446},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 358, col: 29, offset: 9448},
							label: "tps",
							expr: &zeroOrOneExpr{
								pos: position{line: 358, col: 33, offset: 9452},
								expr: &ruleRefExpr{
									pos:  position{line: 358, col: 33, offset: 9452},
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
			pos:  position{line: 363, col: 1, offset: 9519},
			expr: &seqExpr{
				pos: position{line: 363, col: 14, offset: 9534},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 363, col: 14, offset: 9534},
						val:        ":",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 363, col: 18, offset: 9538},
						val:        "<",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 363, col: 22, offset: 9542},
						expr: &ruleRefExpr{
							pos:  position{line: 363, col: 22, offset: 9542},
							name: "Type",
						},
					},
					&litMatcher{
						pos:        position{line: 363, col: 28, offset: 9548},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TypeName",
			pos:  position{line: 364, col: 1, offset: 9552},
			expr: &ruleRefExpr{
				pos:  position{line: 364, col: 12, offset: 9565},
				name: "Name",
			},
		},
		{
			name: "TypeDefinition",
			pos:  position{line: 365, col: 1, offset: 9570},
			expr: &actionExpr{
				pos: position{line: 365, col: 18, offset: 9589},
				run: (*parser).callonTypeDefinition1,
				expr: &seqExpr{
					pos: position{line: 365, col: 18, offset: 9589},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 365, col: 18, offset: 9589},
							val:        "type",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 365, col: 25, offset: 9596},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 365, col: 27, offset: 9598},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 365, col: 30, offset: 9601},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 365, col: 39, offset: 9610},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 365, col: 41, offset: 9612},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 365, col: 44, offset: 9615},
								expr: &ruleRefExpr{
									pos:  position{line: 365, col: 44, offset: 9615},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 365, col: 56, offset: 9627},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 365, col: 58, offset: 9629},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 365, col: 62, offset: 9633},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 365, col: 66, offset: 9637},
								expr: &ruleRefExpr{
									pos:  position{line: 365, col: 66, offset: 9637},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 365, col: 83, offset: 9654},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeExtension",
			pos:  position{line: 385, col: 1, offset: 10135},
			expr: &actionExpr{
				pos: position{line: 385, col: 17, offset: 10153},
				run: (*parser).callonTypeExtension1,
				expr: &seqExpr{
					pos: position{line: 385, col: 17, offset: 10153},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 385, col: 17, offset: 10153},
							val:        "extend",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 385, col: 26, offset: 10162},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 385, col: 28, offset: 10164},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 385, col: 31, offset: 10167},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 385, col: 40, offset: 10176},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 385, col: 42, offset: 10178},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 385, col: 45, offset: 10181},
								expr: &ruleRefExpr{
									pos:  position{line: 385, col: 45, offset: 10181},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 385, col: 57, offset: 10193},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 385, col: 59, offset: 10195},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 385, col: 63, offset: 10199},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 385, col: 67, offset: 10203},
								expr: &ruleRefExpr{
									pos:  position{line: 385, col: 67, offset: 10203},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 385, col: 84, offset: 10220},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Interfaces",
			pos:  position{line: 405, col: 1, offset: 10700},
			expr: &actionExpr{
				pos: position{line: 405, col: 14, offset: 10715},
				run: (*parser).callonInterfaces1,
				expr: &oneOrMoreExpr{
					pos: position{line: 405, col: 14, offset: 10715},
					expr: &ruleRefExpr{
						pos:  position{line: 405, col: 14, offset: 10715},
						name: "GenericType",
					},
				},
			},
		},
		{
			name: "FieldDefinition",
			pos:  position{line: 409, col: 1, offset: 10822},
			expr: &actionExpr{
				pos: position{line: 409, col: 19, offset: 10842},
				run: (*parser).callonFieldDefinition1,
				expr: &seqExpr{
					pos: position{line: 409, col: 19, offset: 10842},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 409, col: 19, offset: 10842},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 409, col: 21, offset: 10844},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 409, col: 24, offset: 10847},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 409, col: 34, offset: 10857},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 409, col: 36, offset: 10859},
							label: "args",
							expr: &zeroOrOneExpr{
								pos: position{line: 409, col: 41, offset: 10864},
								expr: &ruleRefExpr{
									pos:  position{line: 409, col: 41, offset: 10864},
									name: "ArgumentDefinitions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 409, col: 62, offset: 10885},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 409, col: 64, offset: 10887},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 409, col: 68, offset: 10891},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 409, col: 70, offset: 10893},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 409, col: 72, offset: 10895},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 409, col: 77, offset: 10900},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinitions",
			pos:  position{line: 420, col: 1, offset: 11137},
			expr: &actionExpr{
				pos: position{line: 420, col: 23, offset: 11161},
				run: (*parser).callonArgumentDefinitions1,
				expr: &seqExpr{
					pos: position{line: 420, col: 23, offset: 11161},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 420, col: 23, offset: 11161},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 420, col: 27, offset: 11165},
							label: "args",
							expr: &oneOrMoreExpr{
								pos: position{line: 420, col: 32, offset: 11170},
								expr: &ruleRefExpr{
									pos:  position{line: 420, col: 32, offset: 11170},
									name: "ArgumentDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 420, col: 52, offset: 11190},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinition",
			pos:  position{line: 427, col: 1, offset: 11353},
			expr: &actionExpr{
				pos: position{line: 427, col: 22, offset: 11376},
				run: (*parser).callonArgumentDefinition1,
				expr: &seqExpr{
					pos: position{line: 427, col: 22, offset: 11376},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 427, col: 22, offset: 11376},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 427, col: 24, offset: 11378},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 427, col: 27, offset: 11381},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 427, col: 40, offset: 11394},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 427, col: 42, offset: 11396},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 427, col: 46, offset: 11400},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 427, col: 48, offset: 11402},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 427, col: 50, offset: 11404},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 427, col: 55, offset: 11409},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 427, col: 57, offset: 11411},
							label: "dv",
							expr: &zeroOrOneExpr{
								pos: position{line: 427, col: 60, offset: 11414},
								expr: &ruleRefExpr{
									pos:  position{line: 427, col: 60, offset: 11414},
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
			pos:  position{line: 439, col: 1, offset: 11645},
			expr: &actionExpr{
				pos: position{line: 439, col: 18, offset: 11664},
				run: (*parser).callonEnumDefinition1,
				expr: &seqExpr{
					pos: position{line: 439, col: 18, offset: 11664},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 439, col: 18, offset: 11664},
							val:        "enum",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 439, col: 25, offset: 11671},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 439, col: 27, offset: 11673},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 30, offset: 11676},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 439, col: 39, offset: 11685},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 439, col: 41, offset: 11687},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 439, col: 45, offset: 11691},
							label: "vals",
							expr: &oneOrMoreExpr{
								pos: position{line: 439, col: 50, offset: 11696},
								expr: &ruleRefExpr{
									pos:  position{line: 439, col: 50, offset: 11696},
									name: "EnumValueName",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 439, col: 65, offset: 11711},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EnumValueName",
			pos:  position{line: 449, col: 1, offset: 11892},
			expr: &actionExpr{
				pos: position{line: 449, col: 17, offset: 11910},
				run: (*parser).callonEnumValueName1,
				expr: &seqExpr{
					pos: position{line: 449, col: 17, offset: 11910},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 449, col: 17, offset: 11910},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 449, col: 19, offset: 11912},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 449, col: 21, offset: 11914},
								name: "Name",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 449, col: 26, offset: 11919},
							name: "_",
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"ignored\"",
			pos:         position{line: 451, col: 1, offset: 11940},
			expr: &actionExpr{
				pos: position{line: 451, col: 15, offset: 11956},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 451, col: 15, offset: 11956},
					expr: &choiceExpr{
						pos: position{line: 451, col: 16, offset: 11957},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 451, col: 16, offset: 11957},
								name: "whitespace",
							},
							&ruleRefExpr{
								pos:  position{line: 451, col: 29, offset: 11970},
								name: "Comment",
							},
							&litMatcher{
								pos:        position{line: 451, col: 39, offset: 11980},
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
			pos:  position{line: 452, col: 1, offset: 12006},
			expr: &charClassMatcher{
				pos:        position{line: 452, col: 14, offset: 12021},
				val:        "[ \\n\\t\\r]",
				chars:      []rune{' ', '\n', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 454, col: 1, offset: 12032},
			expr: &notExpr{
				pos: position{line: 454, col: 7, offset: 12040},
				expr: &anyMatcher{
					line: 454, col: 8, offset: 12041,
				},
			},
		},
	},
}

func (c *current) onDocument2(d interface{}) (interface{}, error) {
	return d, nil
}

func (p *parser) callonDocument2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument2(stack["d"])
}

func (c *current) onDocument9() (interface{}, error) {
	return nil, errors.New("no graphql document found. expected a Query Document or a Schema")
}

func (p *parser) callonDocument9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument9()
}

func (c *current) onQueryDocument1(defs interface{}) (interface{}, error) {
	result := graphql.Document{
		Operations: []graphql.Operation{},
	}
	sl := ifs(defs)
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

func (p *parser) callonQueryDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryDocument1(stack["defs"])
}

func (c *current) onSchema1(defs interface{}) (interface{}, error) {
	result := graphql.Document{
		Operations: []graphql.Operation{},
	}
	sl := ifs(defs)
	for _, s := range sl {
		if s == nil {
			continue
		}
		if typeDef, ok := s.(graphql.TypeDefinition); ok {
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

func (p *parser) callonSchema1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSchema1(stack["defs"])
}

func (c *current) onSchemaDefinition1(s interface{}) (interface{}, error) {
	return s, nil
}

func (p *parser) callonSchemaDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSchemaDefinition1(stack["s"])
}

func (c *current) onDefinition2(d interface{}) (interface{}, error) {
	return d, nil
}

func (p *parser) callonDefinition2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDefinition2(stack["d"])
}

func (c *current) onDefinition11() (interface{}, error) {
	panic(errors.New("expected top-level operation or fragment definition"))
}

func (p *parser) callonDefinition11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDefinition11()
}

func (c *current) onComment1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment1()
}

func (c *current) onOperationDefinition2(sels interface{}) (interface{}, error) {
	return graphql.Operation{
		Type:         graphql.OperationQuery,
		SelectionSet: sels.(graphql.SelectionSet),
	}, nil

}

func (p *parser) callonOperationDefinition2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperationDefinition2(stack["sels"])
}

func (c *current) onOperationDefinition5(ot, on, vds, ds, sels interface{}) (interface{}, error) {
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
		SelectionSet:        sels.(graphql.SelectionSet),
		Directives:          directives,
		VariableDefinitions: varDefs,
	}, nil
}

func (p *parser) callonOperationDefinition5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperationDefinition5(stack["ot"], stack["on"], stack["vds"], stack["ds"], stack["sels"])
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

func (c *current) onSelectionSet1(s interface{}) (interface{}, error) {
	result := graphql.SelectionSet{}
	for _, sel := range ifs(s) {
		if sel, ok := sel.(graphql.Selection); ok {
			result = append(result, sel)
		} else {
			return result, fmt.Errorf("got unexpected (non-statement) type: %#v", sel)
		}
	}
	return result, nil
}

func (p *parser) callonSelectionSet1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelectionSet1(stack["s"])
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

func (c *current) onSelection14(fs interface{}) (interface{}, error) {
	inlineFragment := fs.(graphql.InlineFragment)
	return graphql.Selection{InlineFragment: &inlineFragment}, nil
}

func (p *parser) callonSelection14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelection14(stack["fs"])
}

func (c *current) onField1(fa, fn, as, ds, sels interface{}) (interface{}, error) {
	var (
		selections graphql.SelectionSet
		arguments  []graphql.Argument
		directives []graphql.Directive
		fieldAlias string
	)
	if fa != nil {
		fieldAlias = fa.(string)
	}
	if sels != nil {
		selections = sels.(graphql.SelectionSet)
	}
	if as != nil {
		arguments = as.([]graphql.Argument)
	}
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.Field{
		Name:         fn.(string),
		Alias:        fieldAlias,
		Arguments:    arguments,
		SelectionSet: selections,
		Directives:   directives,
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

func (c *current) onInlineFragment1(tn, ds, sels interface{}) (interface{}, error) {
	var directives []graphql.Directive
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.InlineFragment{
		TypeCondition: tn.(string),
		Directives:    directives,
		SelectionSet:  sels.(graphql.SelectionSet),
	}, nil
}

func (p *parser) callonInlineFragment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineFragment1(stack["tn"], stack["ds"], stack["sels"])
}

func (c *current) onFragmentDefinition1(fn, tn, ds, sels interface{}) (interface{}, error) {
	var directives []graphql.Directive
	if ds != nil {
		directives = ds.([]graphql.Directive)
	}
	return graphql.FragmentDefinition{
		Name:          fn.(string),
		TypeCondition: tn.(string),
		SelectionSet:  sels.(graphql.SelectionSet),
		Directives:    directives,
	}, nil
}

func (p *parser) callonFragmentDefinition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFragmentDefinition1(stack["fn"], stack["tn"], stack["ds"], stack["sels"])
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

func (c *current) onArray1(values interface{}) (interface{}, error) {
	sl := ifs(values)
	result := make([]interface{}, 0, len(sl))
	for _, p := range sl {
		result = append(result, p)
	}
	return result, nil
}

func (p *parser) callonArray1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArray1(stack["values"])
}

func (c *current) onObject1(ps interface{}) (interface{}, error) {
	result := make(map[string]interface{})
	for _, p := range ifs(ps) {
		prop, ok := p.(graphql.Argument)
		if !ok {
			return nil, fmt.Errorf("expected Property, got %#v", p)
		}
		result[prop.Name] = prop.Value
	}
	return result, nil
}

func (p *parser) callonObject1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onObject1(stack["ps"])
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

func (c *current) onProperty1(pn, v interface{}) (interface{}, error) {
	return graphql.Argument{Name: pn.(string), Value: v}, nil
}

func (p *parser) callonProperty1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onProperty1(stack["pn"], stack["v"])
}

func (c *current) onPropertyName1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonPropertyName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPropertyName1()
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
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
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
