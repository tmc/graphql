package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 5, col: 1, offset: 20},
			expr: &oneOrMoreExpr{
				pos: position{line: 5, col: 12, offset: 33},
				expr: &ruleRefExpr{
					pos:  position{line: 5, col: 12, offset: 33},
					name: "Statement",
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 7, col: 1, offset: 45},
			expr: &choiceExpr{
				pos: position{line: 7, col: 13, offset: 59},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 7, col: 13, offset: 59},
						name: "Operation",
					},
					&ruleRefExpr{
						pos:  position{line: 7, col: 25, offset: 71},
						name: "FragmentDefinition",
					},
					&ruleRefExpr{
						pos:  position{line: 7, col: 46, offset: 92},
						name: "TypeDefinition",
					},
					&ruleRefExpr{
						pos:  position{line: 7, col: 63, offset: 109},
						name: "TypeExtension",
					},
					&ruleRefExpr{
						pos:  position{line: 7, col: 79, offset: 125},
						name: "EnumDefinition",
					},
					&ruleRefExpr{
						pos:  position{line: 7, col: 96, offset: 142},
						name: "Comment",
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 9, col: 1, offset: 151},
			expr: &seqExpr{
				pos: position{line: 9, col: 11, offset: 163},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 9, col: 11, offset: 163},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 9, col: 15, offset: 167},
						expr: &charClassMatcher{
							pos:        position{line: 9, col: 15, offset: 167},
							val:        "[^\\n]",
							chars:      []rune{'\n'},
							ignoreCase: false,
							inverted:   true,
						},
					},
				},
			},
		},
		{
			name: "Operation",
			pos:  position{line: 11, col: 1, offset: 175},
			expr: &choiceExpr{
				pos: position{line: 11, col: 13, offset: 189},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 11, col: 13, offset: 189},
						name: "Selections",
					},
					&seqExpr{
						pos: position{line: 12, col: 14, offset: 215},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 12, col: 14, offset: 215},
								name: "OperationType",
							},
							&ruleRefExpr{
								pos:  position{line: 12, col: 28, offset: 229},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 12, col: 30, offset: 231},
								name: "OperationName",
							},
							&ruleRefExpr{
								pos:  position{line: 12, col: 44, offset: 245},
								name: "_",
							},
							&zeroOrOneExpr{
								pos: position{line: 12, col: 46, offset: 247},
								expr: &ruleRefExpr{
									pos:  position{line: 12, col: 46, offset: 247},
									name: "VariableDefinitions",
								},
							},
							&ruleRefExpr{
								pos:  position{line: 12, col: 67, offset: 268},
								name: "_",
							},
							&zeroOrOneExpr{
								pos: position{line: 12, col: 69, offset: 270},
								expr: &ruleRefExpr{
									pos:  position{line: 12, col: 69, offset: 270},
									name: "Directives",
								},
							},
							&ruleRefExpr{
								pos:  position{line: 12, col: 81, offset: 282},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 12, col: 83, offset: 284},
								name: "Selections",
							},
						},
					},
				},
			},
		},
		{
			name: "OperationType",
			pos:  position{line: 14, col: 1, offset: 297},
			expr: &choiceExpr{
				pos: position{line: 14, col: 17, offset: 315},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 14, col: 17, offset: 315},
						val:        "query",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 14, col: 27, offset: 325},
						val:        "mutation",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "OperationName",
			pos:  position{line: 15, col: 1, offset: 336},
			expr: &ruleRefExpr{
				pos:  position{line: 15, col: 17, offset: 354},
				name: "Name",
			},
		},
		{
			name: "VariableDefinitions",
			pos:  position{line: 16, col: 1, offset: 359},
			expr: &seqExpr{
				pos: position{line: 16, col: 23, offset: 383},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 16, col: 23, offset: 383},
						val:        "(",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 16, col: 27, offset: 387},
						expr: &ruleRefExpr{
							pos:  position{line: 16, col: 27, offset: 387},
							name: "VariableDefinition",
						},
					},
					&litMatcher{
						pos:        position{line: 16, col: 47, offset: 407},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "VariableDefinition",
			pos:  position{line: 17, col: 1, offset: 411},
			expr: &seqExpr{
				pos: position{line: 17, col: 22, offset: 434},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 22, offset: 434},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 24, offset: 436},
						name: "Variable",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 33, offset: 445},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 35, offset: 447},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 39, offset: 451},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 41, offset: 453},
						name: "Type",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 46, offset: 458},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 17, col: 48, offset: 460},
						expr: &ruleRefExpr{
							pos:  position{line: 17, col: 48, offset: 460},
							name: "DefaultValue",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 62, offset: 474},
						name: "_",
					},
				},
			},
		},
		{
			name: "DefaultValue",
			pos:  position{line: 18, col: 1, offset: 476},
			expr: &seqExpr{
				pos: position{line: 18, col: 16, offset: 493},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 18, col: 16, offset: 493},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 18, col: 20, offset: 497},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Selections",
			pos:  position{line: 20, col: 1, offset: 504},
			expr: &seqExpr{
				pos: position{line: 20, col: 14, offset: 519},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 20, col: 14, offset: 519},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 20, col: 18, offset: 523},
						expr: &ruleRefExpr{
							pos:  position{line: 20, col: 18, offset: 523},
							name: "Selection",
						},
					},
					&litMatcher{
						pos:        position{line: 20, col: 29, offset: 534},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Selection",
			pos:  position{line: 21, col: 1, offset: 538},
			expr: &choiceExpr{
				pos: position{line: 21, col: 13, offset: 552},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 21, col: 13, offset: 552},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 21, col: 13, offset: 552},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 21, col: 15, offset: 554},
								name: "Field",
							},
							&ruleRefExpr{
								pos:  position{line: 21, col: 21, offset: 560},
								name: "_",
							},
						},
					},
					&seqExpr{
						pos: position{line: 21, col: 25, offset: 564},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 21, col: 25, offset: 564},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 21, col: 27, offset: 566},
								name: "FragmentSpread",
							},
							&ruleRefExpr{
								pos:  position{line: 21, col: 42, offset: 581},
								name: "_",
							},
						},
					},
				},
			},
		},
		{
			name: "Field",
			pos:  position{line: 23, col: 1, offset: 584},
			expr: &seqExpr{
				pos: position{line: 23, col: 9, offset: 594},
				exprs: []interface{}{
					&zeroOrOneExpr{
						pos: position{line: 23, col: 9, offset: 594},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 9, offset: 594},
							name: "FieldAlias",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 21, offset: 606},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 23, offset: 608},
						name: "FieldName",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 33, offset: 618},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 23, col: 35, offset: 620},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 35, offset: 620},
							name: "Arguments",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 46, offset: 631},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 23, col: 48, offset: 633},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 48, offset: 633},
							name: "Directives",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 60, offset: 645},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 23, col: 62, offset: 647},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 62, offset: 647},
							name: "Selections",
						},
					},
				},
			},
		},
		{
			name: "FieldAlias",
			pos:  position{line: 24, col: 1, offset: 659},
			expr: &seqExpr{
				pos: position{line: 24, col: 14, offset: 674},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 24, col: 14, offset: 674},
						name: "Name",
					},
					&litMatcher{
						pos:        position{line: 24, col: 19, offset: 679},
						val:        ":",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "FieldName",
			pos:  position{line: 25, col: 1, offset: 683},
			expr: &ruleRefExpr{
				pos:  position{line: 25, col: 13, offset: 697},
				name: "Name",
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 26, col: 1, offset: 702},
			expr: &seqExpr{
				pos: position{line: 26, col: 13, offset: 716},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 26, col: 13, offset: 716},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 26, col: 17, offset: 720},
						expr: &ruleRefExpr{
							pos:  position{line: 26, col: 17, offset: 720},
							name: "Argument",
						},
					},
					&litMatcher{
						pos:        position{line: 26, col: 27, offset: 730},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Argument",
			pos:  position{line: 27, col: 1, offset: 734},
			expr: &seqExpr{
				pos: position{line: 27, col: 12, offset: 747},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 12, offset: 747},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 14, offset: 749},
						name: "ArgumentName",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 27, offset: 762},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 27, col: 29, offset: 764},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 33, offset: 768},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 35, offset: 770},
						name: "Value",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 41, offset: 776},
						name: "_",
					},
				},
			},
		},
		{
			name: "ArgumentName",
			pos:  position{line: 28, col: 1, offset: 778},
			expr: &ruleRefExpr{
				pos:  position{line: 28, col: 16, offset: 795},
				name: "Name",
			},
		},
		{
			name: "Name",
			pos:  position{line: 30, col: 1, offset: 801},
			expr: &seqExpr{
				pos: position{line: 30, col: 8, offset: 810},
				exprs: []interface{}{
					&charClassMatcher{
						pos:        position{line: 30, col: 8, offset: 810},
						val:        "[a-z_]i",
						chars:      []rune{'_'},
						ranges:     []rune{'a', 'z'},
						ignoreCase: true,
						inverted:   false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 30, col: 16, offset: 818},
						expr: &charClassMatcher{
							pos:        position{line: 30, col: 16, offset: 818},
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
		{
			name: "FragmentSpread",
			pos:  position{line: 32, col: 1, offset: 831},
			expr: &seqExpr{
				pos: position{line: 32, col: 19, offset: 851},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 32, col: 19, offset: 851},
						val:        "...",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 32, col: 25, offset: 857},
						name: "FragmentName",
					},
					&zeroOrOneExpr{
						pos: position{line: 32, col: 38, offset: 870},
						expr: &ruleRefExpr{
							pos:  position{line: 32, col: 38, offset: 870},
							name: "Directives",
						},
					},
				},
			},
		},
		{
			name: "FragmentDefinition",
			pos:  position{line: 33, col: 1, offset: 882},
			expr: &seqExpr{
				pos: position{line: 33, col: 22, offset: 905},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 33, col: 22, offset: 905},
						val:        "fragment",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 33, offset: 916},
						name: "Type",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 38, offset: 921},
						name: "FragmentName",
					},
					&zeroOrOneExpr{
						pos: position{line: 33, col: 51, offset: 934},
						expr: &ruleRefExpr{
							pos:  position{line: 33, col: 51, offset: 934},
							name: "Directives",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 63, offset: 946},
						name: "Selections",
					},
				},
			},
		},
		{
			name: "FragmentName",
			pos:  position{line: 34, col: 1, offset: 957},
			expr: &ruleRefExpr{
				pos:  position{line: 34, col: 16, offset: 974},
				name: "Name",
			},
		},
		{
			name: "Value",
			pos:  position{line: 36, col: 1, offset: 980},
			expr: &seqExpr{
				pos: position{line: 36, col: 9, offset: 990},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 36, col: 9, offset: 990},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 36, col: 12, offset: 993},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 36, col: 12, offset: 993},
								name: "Null",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 19, offset: 1000},
								name: "Boolean",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 29, offset: 1010},
								name: "Int",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 35, offset: 1016},
								name: "Float",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 43, offset: 1024},
								name: "String",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 52, offset: 1033},
								name: "Array",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 60, offset: 1041},
								name: "Object",
							},
							&ruleRefExpr{
								pos:  position{line: 36, col: 69, offset: 1050},
								name: "Variable",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 36, col: 79, offset: 1060},
						name: "_",
					},
				},
			},
		},
		{
			name: "Null",
			pos:  position{line: 39, col: 1, offset: 1110},
			expr: &litMatcher{
				pos:        position{line: 39, col: 8, offset: 1119},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "Boolean",
			pos:  position{line: 40, col: 1, offset: 1126},
			expr: &choiceExpr{
				pos: position{line: 40, col: 11, offset: 1138},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 40, col: 11, offset: 1138},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 40, col: 20, offset: 1147},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Int",
			pos:  position{line: 41, col: 1, offset: 1155},
			expr: &seqExpr{
				pos: position{line: 41, col: 7, offset: 1163},
				exprs: []interface{}{
					&zeroOrOneExpr{
						pos: position{line: 41, col: 7, offset: 1163},
						expr: &ruleRefExpr{
							pos:  position{line: 41, col: 7, offset: 1163},
							name: "Sign",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 13, offset: 1169},
						name: "IntegerPart",
					},
				},
			},
		},
		{
			name: "Float",
			pos:  position{line: 42, col: 1, offset: 1181},
			expr: &seqExpr{
				pos: position{line: 42, col: 9, offset: 1191},
				exprs: []interface{}{
					&zeroOrOneExpr{
						pos: position{line: 42, col: 9, offset: 1191},
						expr: &ruleRefExpr{
							pos:  position{line: 42, col: 9, offset: 1191},
							name: "Sign",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 42, col: 15, offset: 1197},
						name: "IntegerPart",
					},
					&litMatcher{
						pos:        position{line: 42, col: 27, offset: 1209},
						val:        ".",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 42, col: 31, offset: 1213},
						expr: &ruleRefExpr{
							pos:  position{line: 42, col: 31, offset: 1213},
							name: "Digit",
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 42, col: 38, offset: 1220},
						expr: &ruleRefExpr{
							pos:  position{line: 42, col: 38, offset: 1220},
							name: "ExponentPart",
						},
					},
				},
			},
		},
		{
			name: "Sign",
			pos:  position{line: 43, col: 1, offset: 1234},
			expr: &litMatcher{
				pos:        position{line: 43, col: 8, offset: 1243},
				val:        "-",
				ignoreCase: false,
			},
		},
		{
			name: "IntegerPart",
			pos:  position{line: 44, col: 1, offset: 1247},
			expr: &choiceExpr{
				pos: position{line: 44, col: 15, offset: 1263},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 44, col: 15, offset: 1263},
						val:        "0",
						ignoreCase: false,
					},
					&seqExpr{
						pos: position{line: 44, col: 21, offset: 1269},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 44, col: 21, offset: 1269},
								name: "NonZeroDigit",
							},
							&zeroOrMoreExpr{
								pos: position{line: 44, col: 34, offset: 1282},
								expr: &ruleRefExpr{
									pos:  position{line: 44, col: 34, offset: 1282},
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
			pos:  position{line: 45, col: 1, offset: 1289},
			expr: &seqExpr{
				pos: position{line: 45, col: 16, offset: 1306},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 45, col: 16, offset: 1306},
						val:        "e",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 45, col: 20, offset: 1310},
						expr: &ruleRefExpr{
							pos:  position{line: 45, col: 20, offset: 1310},
							name: "Sign",
						},
					},
					&oneOrMoreExpr{
						pos: position{line: 45, col: 26, offset: 1316},
						expr: &ruleRefExpr{
							pos:  position{line: 45, col: 26, offset: 1316},
							name: "Digit",
						},
					},
				},
			},
		},
		{
			name: "Digit",
			pos:  position{line: 46, col: 1, offset: 1323},
			expr: &charClassMatcher{
				pos:        position{line: 46, col: 9, offset: 1333},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDigit",
			pos:  position{line: 47, col: 1, offset: 1339},
			expr: &charClassMatcher{
				pos:        position{line: 47, col: 16, offset: 1356},
				val:        "[123456789]",
				chars:      []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "String",
			pos:  position{line: 48, col: 1, offset: 1368},
			expr: &seqExpr{
				pos: position{line: 48, col: 10, offset: 1379},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 48, col: 10, offset: 1379},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 48, col: 14, offset: 1383},
						expr: &ruleRefExpr{
							pos:  position{line: 48, col: 14, offset: 1383},
							name: "StringCharacter",
						},
					},
					&litMatcher{
						pos:        position{line: 48, col: 31, offset: 1400},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "StringCharacter",
			pos:  position{line: 49, col: 1, offset: 1404},
			expr: &choiceExpr{
				pos: position{line: 49, col: 19, offset: 1424},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 49, col: 19, offset: 1424},
						val:        "[^\\\\\"]",
						chars:      []rune{'\\', '"'},
						ignoreCase: false,
						inverted:   true,
					},
					&seqExpr{
						pos: position{line: 49, col: 28, offset: 1433},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 49, col: 28, offset: 1433},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 33, offset: 1438},
								name: "EscapedCharacter",
							},
						},
					},
					&seqExpr{
						pos: position{line: 49, col: 52, offset: 1457},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 49, col: 52, offset: 1457},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 57, offset: 1462},
								name: "EscapedUnicode",
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedUnicode",
			pos:  position{line: 50, col: 1, offset: 1477},
			expr: &seqExpr{
				pos: position{line: 50, col: 18, offset: 1496},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 50, col: 18, offset: 1496},
						val:        "u",
						ignoreCase: false,
					},
					&charClassMatcher{
						pos:        position{line: 50, col: 22, offset: 1500},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 50, col: 32, offset: 1510},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 50, col: 42, offset: 1520},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 50, col: 52, offset: 1530},
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
			pos:  position{line: 51, col: 1, offset: 1540},
			expr: &choiceExpr{
				pos: position{line: 51, col: 20, offset: 1561},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 51, col: 20, offset: 1561},
						val:        "[\"/bfnrt]",
						chars:      []rune{'"', '/', 'b', 'f', 'n', 'r', 't'},
						ignoreCase: false,
						inverted:   false,
					},
					&litMatcher{
						pos:        position{line: 51, col: 32, offset: 1573},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Array",
			pos:  position{line: 53, col: 1, offset: 1580},
			expr: &seqExpr{
				pos: position{line: 53, col: 9, offset: 1590},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 53, col: 9, offset: 1590},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 53, col: 13, offset: 1594},
						expr: &ruleRefExpr{
							pos:  position{line: 53, col: 13, offset: 1594},
							name: "Value",
						},
					},
					&litMatcher{
						pos:        position{line: 53, col: 20, offset: 1601},
						val:        "]",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Object",
			pos:  position{line: 54, col: 1, offset: 1605},
			expr: &seqExpr{
				pos: position{line: 54, col: 10, offset: 1616},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 54, col: 10, offset: 1616},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 54, col: 14, offset: 1620},
						expr: &ruleRefExpr{
							pos:  position{line: 54, col: 14, offset: 1620},
							name: "Property",
						},
					},
					&litMatcher{
						pos:        position{line: 54, col: 24, offset: 1630},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 56, col: 1, offset: 1635},
			expr: &ruleRefExpr{
				pos:  position{line: 56, col: 12, offset: 1648},
				name: "VariableName",
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 58, col: 1, offset: 1750},
			expr: &seqExpr{
				pos: position{line: 58, col: 16, offset: 1767},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 58, col: 16, offset: 1767},
						val:        "$",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 58, col: 20, offset: 1771},
						expr: &charClassMatcher{
							pos:        position{line: 58, col: 20, offset: 1771},
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
		{
			name: "VariablePropertySelection",
			pos:  position{line: 59, col: 1, offset: 1783},
			expr: &seqExpr{
				pos: position{line: 59, col: 29, offset: 1813},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 59, col: 29, offset: 1813},
						name: "Variable",
					},
					&litMatcher{
						pos:        position{line: 59, col: 38, offset: 1822},
						val:        ".",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 42, offset: 1826},
						name: "PropertyName",
					},
				},
			},
		},
		{
			name: "Property",
			pos:  position{line: 61, col: 1, offset: 1840},
			expr: &seqExpr{
				pos: position{line: 61, col: 12, offset: 1853},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 61, col: 12, offset: 1853},
						name: "PropertyName",
					},
					&litMatcher{
						pos:        position{line: 61, col: 25, offset: 1866},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 29, offset: 1870},
						name: "Value",
					},
				},
			},
		},
		{
			name: "PropertyName",
			pos:  position{line: 62, col: 1, offset: 1876},
			expr: &ruleRefExpr{
				pos:  position{line: 62, col: 16, offset: 1893},
				name: "Name",
			},
		},
		{
			name: "Directives",
			pos:  position{line: 64, col: 1, offset: 1899},
			expr: &oneOrMoreExpr{
				pos: position{line: 64, col: 14, offset: 1914},
				expr: &ruleRefExpr{
					pos:  position{line: 64, col: 14, offset: 1914},
					name: "Directive",
				},
			},
		},
		{
			name: "Directive",
			pos:  position{line: 65, col: 1, offset: 1925},
			expr: &seqExpr{
				pos: position{line: 65, col: 13, offset: 1939},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 65, col: 13, offset: 1939},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 65, col: 15, offset: 1941},
						val:        "@",
						ignoreCase: false,
					},
					&choiceExpr{
						pos: position{line: 65, col: 20, offset: 1946},
						alternatives: []interface{}{
							&seqExpr{
								pos: position{line: 65, col: 20, offset: 1946},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 65, col: 20, offset: 1946},
										name: "DirectiveName",
									},
									&litMatcher{
										pos:        position{line: 65, col: 34, offset: 1960},
										val:        ":",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 65, col: 38, offset: 1964},
										name: "_",
									},
									&ruleRefExpr{
										pos:  position{line: 65, col: 40, offset: 1966},
										name: "Value",
									},
								},
							},
							&seqExpr{
								pos: position{line: 66, col: 17, offset: 1990},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 66, col: 17, offset: 1990},
										name: "DirectiveName",
									},
									&litMatcher{
										pos:        position{line: 66, col: 31, offset: 2004},
										val:        ":",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 66, col: 35, offset: 2008},
										name: "_",
									},
									&ruleRefExpr{
										pos:  position{line: 66, col: 37, offset: 2010},
										name: "Type",
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 67, col: 20, offset: 2036},
								name: "DirectiveName",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 35, offset: 2051},
						name: "_",
					},
				},
			},
		},
		{
			name: "DirectiveName",
			pos:  position{line: 68, col: 1, offset: 2053},
			expr: &ruleRefExpr{
				pos:  position{line: 68, col: 17, offset: 2071},
				name: "Name",
			},
		},
		{
			name: "Type",
			pos:  position{line: 70, col: 1, offset: 2077},
			expr: &choiceExpr{
				pos: position{line: 70, col: 8, offset: 2086},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 70, col: 8, offset: 2086},
						name: "OptionalType",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 23, offset: 2101},
						name: "GenericType",
					},
				},
			},
		},
		{
			name: "OptionalType",
			pos:  position{line: 71, col: 1, offset: 2113},
			expr: &zeroOrOneExpr{
				pos: position{line: 71, col: 16, offset: 2130},
				expr: &ruleRefExpr{
					pos:  position{line: 71, col: 16, offset: 2130},
					name: "GenericType",
				},
			},
		},
		{
			name: "GenericType",
			pos:  position{line: 72, col: 1, offset: 2143},
			expr: &seqExpr{
				pos: position{line: 72, col: 15, offset: 2159},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 72, col: 15, offset: 2159},
						name: "TypeName",
					},
					&zeroOrOneExpr{
						pos: position{line: 72, col: 24, offset: 2168},
						expr: &ruleRefExpr{
							pos:  position{line: 72, col: 24, offset: 2168},
							name: "TypeParams",
						},
					},
				},
			},
		},
		{
			name: "TypeParams",
			pos:  position{line: 73, col: 1, offset: 2180},
			expr: &seqExpr{
				pos: position{line: 73, col: 14, offset: 2195},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 73, col: 14, offset: 2195},
						val:        ":",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 73, col: 18, offset: 2199},
						val:        "<",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 73, col: 22, offset: 2203},
						expr: &ruleRefExpr{
							pos:  position{line: 73, col: 22, offset: 2203},
							name: "Type",
						},
					},
					&litMatcher{
						pos:        position{line: 73, col: 28, offset: 2209},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TypeName",
			pos:  position{line: 74, col: 1, offset: 2213},
			expr: &ruleRefExpr{
				pos:  position{line: 74, col: 12, offset: 2226},
				name: "Name",
			},
		},
		{
			name: "TypeDefinition",
			pos:  position{line: 75, col: 1, offset: 2231},
			expr: &seqExpr{
				pos: position{line: 75, col: 18, offset: 2250},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 75, col: 18, offset: 2250},
						val:        "type",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 25, offset: 2257},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 27, offset: 2259},
						name: "TypeName",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 36, offset: 2268},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 75, col: 38, offset: 2270},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 38, offset: 2270},
							name: "Interfaces",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 50, offset: 2282},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 75, col: 52, offset: 2284},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 75, col: 56, offset: 2288},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 56, offset: 2288},
							name: "FieldDefinition",
						},
					},
					&litMatcher{
						pos:        position{line: 75, col: 73, offset: 2305},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TypeExtension",
			pos:  position{line: 76, col: 1, offset: 2309},
			expr: &seqExpr{
				pos: position{line: 76, col: 17, offset: 2327},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 76, col: 17, offset: 2327},
						val:        "extend",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 26, offset: 2336},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 28, offset: 2338},
						name: "TypeName",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 37, offset: 2347},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 76, col: 39, offset: 2349},
						expr: &ruleRefExpr{
							pos:  position{line: 76, col: 39, offset: 2349},
							name: "Interfaces",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 51, offset: 2361},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 76, col: 53, offset: 2363},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 76, col: 57, offset: 2367},
						expr: &ruleRefExpr{
							pos:  position{line: 76, col: 57, offset: 2367},
							name: "FieldDefinition",
						},
					},
					&litMatcher{
						pos:        position{line: 76, col: 74, offset: 2384},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Interfaces",
			pos:  position{line: 77, col: 1, offset: 2388},
			expr: &oneOrMoreExpr{
				pos: position{line: 77, col: 14, offset: 2403},
				expr: &ruleRefExpr{
					pos:  position{line: 77, col: 14, offset: 2403},
					name: "GenericType",
				},
			},
		},
		{
			name: "FieldDefinition",
			pos:  position{line: 78, col: 1, offset: 2416},
			expr: &seqExpr{
				pos: position{line: 78, col: 19, offset: 2436},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 78, col: 19, offset: 2436},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 21, offset: 2438},
						name: "FieldName",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 31, offset: 2448},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 78, col: 33, offset: 2450},
						expr: &ruleRefExpr{
							pos:  position{line: 78, col: 33, offset: 2450},
							name: "ArgumentDefinitions",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 54, offset: 2471},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 78, col: 56, offset: 2473},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 60, offset: 2477},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 62, offset: 2479},
						name: "Type",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 67, offset: 2484},
						name: "_",
					},
				},
			},
		},
		{
			name: "ArgumentDefinitions",
			pos:  position{line: 79, col: 1, offset: 2486},
			expr: &seqExpr{
				pos: position{line: 79, col: 23, offset: 2510},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 79, col: 23, offset: 2510},
						val:        "(",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 79, col: 27, offset: 2514},
						expr: &ruleRefExpr{
							pos:  position{line: 79, col: 27, offset: 2514},
							name: "ArgumentDefinition",
						},
					},
					&litMatcher{
						pos:        position{line: 79, col: 47, offset: 2534},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ArgumentDefinition",
			pos:  position{line: 80, col: 1, offset: 2538},
			expr: &seqExpr{
				pos: position{line: 80, col: 22, offset: 2561},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 80, col: 22, offset: 2561},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 24, offset: 2563},
						name: "ArgumentName",
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 37, offset: 2576},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 80, col: 39, offset: 2578},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 43, offset: 2582},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 45, offset: 2584},
						name: "Type",
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 50, offset: 2589},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 80, col: 52, offset: 2591},
						expr: &ruleRefExpr{
							pos:  position{line: 80, col: 52, offset: 2591},
							name: "DefaultValue",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 66, offset: 2605},
						name: "_",
					},
				},
			},
		},
		{
			name: "EnumDefinition",
			pos:  position{line: 81, col: 1, offset: 2607},
			expr: &seqExpr{
				pos: position{line: 81, col: 18, offset: 2626},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 81, col: 18, offset: 2626},
						val:        "enum",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 25, offset: 2633},
						name: "TypeName",
					},
					&litMatcher{
						pos:        position{line: 81, col: 34, offset: 2642},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 81, col: 38, offset: 2646},
						expr: &ruleRefExpr{
							pos:  position{line: 81, col: 38, offset: 2646},
							name: "EnumValueName",
						},
					},
					&litMatcher{
						pos:        position{line: 81, col: 53, offset: 2661},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "EnumValueName",
			pos:  position{line: 82, col: 1, offset: 2665},
			expr: &ruleRefExpr{
				pos:  position{line: 82, col: 17, offset: 2683},
				name: "Name",
			},
		},
		{
			name:        "_",
			displayName: "\"ignored\"",
			pos:         position{line: 84, col: 1, offset: 2689},
			expr: &zeroOrMoreExpr{
				pos: position{line: 84, col: 15, offset: 2705},
				expr: &choiceExpr{
					pos: position{line: 84, col: 16, offset: 2706},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 84, col: 16, offset: 2706},
							name: "whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 84, col: 29, offset: 2719},
							name: "Comment",
						},
						&litMatcher{
							pos:        position{line: 84, col: 39, offset: 2729},
							val:        ",",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "whitespace",
			pos:  position{line: 85, col: 1, offset: 2735},
			expr: &oneOrMoreExpr{
				pos: position{line: 85, col: 14, offset: 2750},
				expr: &charClassMatcher{
					pos:        position{line: 85, col: 14, offset: 2750},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 87, col: 1, offset: 2762},
			expr: &notExpr{
				pos: position{line: 87, col: 7, offset: 2770},
				expr: &anyMatcher{
					line: 87, col: 8, offset: 2771,
				},
			},
		},
	},
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
