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
									label: "stmts",
									expr: &oneOrMoreExpr{
										pos: position{line: 15, col: 18, offset: 183},
										expr: &ruleRefExpr{
											pos:  position{line: 15, col: 18, offset: 183},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 15, col: 29, offset: 194},
									name: "EOF",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 39, col: 5, offset: 1067},
						run: (*parser).callonDocument8,
						expr: &anyMatcher{
							line: 39, col: 5, offset: 1067,
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 43, col: 1, offset: 1130},
			expr: &choiceExpr{
				pos: position{line: 43, col: 13, offset: 1144},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 43, col: 13, offset: 1144},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 43, col: 13, offset: 1144},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 43, col: 13, offset: 1144},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 43, col: 15, offset: 1146},
									label: "s",
									expr: &choiceExpr{
										pos: position{line: 43, col: 18, offset: 1149},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 43, col: 18, offset: 1149},
												name: "Operation",
											},
											&ruleRefExpr{
												pos:  position{line: 43, col: 30, offset: 1161},
												name: "FragmentDefinition",
											},
											&ruleRefExpr{
												pos:  position{line: 43, col: 51, offset: 1182},
												name: "TypeDefinition",
											},
											&ruleRefExpr{
												pos:  position{line: 43, col: 68, offset: 1199},
												name: "TypeExtension",
											},
											&ruleRefExpr{
												pos:  position{line: 43, col: 84, offset: 1215},
												name: "EnumDefinition",
											},
											&ruleRefExpr{
												pos:  position{line: 43, col: 101, offset: 1232},
												name: "Comment",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 110, offset: 1241},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 45, col: 5, offset: 1264},
						run: (*parser).callonStatement14,
						expr: &anyMatcher{
							line: 45, col: 5, offset: 1264,
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 49, col: 1, offset: 1325},
			expr: &actionExpr{
				pos: position{line: 49, col: 11, offset: 1337},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 49, col: 11, offset: 1337},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 49, col: 11, offset: 1337},
							val:        "#",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 49, col: 15, offset: 1341},
							expr: &charClassMatcher{
								pos:        position{line: 49, col: 15, offset: 1341},
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
			pos:  position{line: 51, col: 1, offset: 1380},
			expr: &choiceExpr{
				pos: position{line: 51, col: 13, offset: 1394},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 51, col: 13, offset: 1394},
						run: (*parser).callonOperation2,
						expr: &labeledExpr{
							pos:   position{line: 51, col: 13, offset: 1394},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 18, offset: 1399},
								name: "SelectionSet",
							},
						},
					},
					&actionExpr{
						pos: position{line: 58, col: 13, offset: 1543},
						run: (*parser).callonOperation5,
						expr: &seqExpr{
							pos: position{line: 58, col: 14, offset: 1544},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 58, col: 14, offset: 1544},
									label: "ot",
									expr: &ruleRefExpr{
										pos:  position{line: 58, col: 17, offset: 1547},
										name: "OperationType",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 58, col: 31, offset: 1561},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 58, col: 33, offset: 1563},
									label: "on",
									expr: &ruleRefExpr{
										pos:  position{line: 58, col: 36, offset: 1566},
										name: "OperationName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 58, col: 50, offset: 1580},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 58, col: 52, offset: 1582},
									label: "vds",
									expr: &zeroOrOneExpr{
										pos: position{line: 58, col: 56, offset: 1586},
										expr: &ruleRefExpr{
											pos:  position{line: 58, col: 56, offset: 1586},
											name: "VariableDefinitions",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 58, col: 77, offset: 1607},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 58, col: 79, offset: 1609},
									label: "ds",
									expr: &zeroOrOneExpr{
										pos: position{line: 58, col: 82, offset: 1612},
										expr: &ruleRefExpr{
											pos:  position{line: 58, col: 82, offset: 1612},
											name: "Directives",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 58, col: 94, offset: 1624},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 58, col: 96, offset: 1626},
									label: "sels",
									expr: &ruleRefExpr{
										pos:  position{line: 58, col: 101, offset: 1631},
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
			pos:  position{line: 78, col: 1, offset: 2052},
			expr: &choiceExpr{
				pos: position{line: 78, col: 17, offset: 2070},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 78, col: 17, offset: 2070},
						run: (*parser).callonOperationType2,
						expr: &litMatcher{
							pos:        position{line: 78, col: 17, offset: 2070},
							val:        "query",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 18, offset: 2134},
						run: (*parser).callonOperationType4,
						expr: &litMatcher{
							pos:        position{line: 79, col: 18, offset: 2134},
							val:        "mutation",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperationName",
			pos:  position{line: 80, col: 1, offset: 2187},
			expr: &actionExpr{
				pos: position{line: 80, col: 17, offset: 2205},
				run: (*parser).callonOperationName1,
				expr: &ruleRefExpr{
					pos:  position{line: 80, col: 17, offset: 2205},
					name: "Name",
				},
			},
		},
		{
			name: "VariableDefinitions",
			pos:  position{line: 83, col: 1, offset: 2242},
			expr: &actionExpr{
				pos: position{line: 83, col: 23, offset: 2266},
				run: (*parser).callonVariableDefinitions1,
				expr: &seqExpr{
					pos: position{line: 83, col: 23, offset: 2266},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 83, col: 23, offset: 2266},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 83, col: 27, offset: 2270},
							label: "vds",
							expr: &oneOrMoreExpr{
								pos: position{line: 83, col: 31, offset: 2274},
								expr: &ruleRefExpr{
									pos:  position{line: 83, col: 31, offset: 2274},
									name: "VariableDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 83, col: 51, offset: 2294},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariableDefinition",
			pos:  position{line: 90, col: 1, offset: 2455},
			expr: &actionExpr{
				pos: position{line: 90, col: 22, offset: 2478},
				run: (*parser).callonVariableDefinition1,
				expr: &seqExpr{
					pos: position{line: 90, col: 22, offset: 2478},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 90, col: 22, offset: 2478},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 90, col: 24, offset: 2480},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 26, offset: 2482},
								name: "Variable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 35, offset: 2491},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 90, col: 37, offset: 2493},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 41, offset: 2497},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 90, col: 43, offset: 2499},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 45, offset: 2501},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 50, offset: 2506},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 90, col: 52, offset: 2508},
							label: "d",
							expr: &zeroOrOneExpr{
								pos: position{line: 90, col: 54, offset: 2510},
								expr: &ruleRefExpr{
									pos:  position{line: 90, col: 54, offset: 2510},
									name: "DefaultValue",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 68, offset: 2524},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DefaultValue",
			pos:  position{line: 102, col: 1, offset: 2760},
			expr: &actionExpr{
				pos: position{line: 102, col: 16, offset: 2777},
				run: (*parser).callonDefaultValue1,
				expr: &seqExpr{
					pos: position{line: 102, col: 16, offset: 2777},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 16, offset: 2777},
							val:        "=",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 102, col: 20, offset: 2781},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 22, offset: 2783},
								name: "Value",
							},
						},
					},
				},
			},
		},
		{
			name: "SelectionSet",
			pos:  position{line: 104, col: 1, offset: 2808},
			expr: &actionExpr{
				pos: position{line: 104, col: 16, offset: 2825},
				run: (*parser).callonSelectionSet1,
				expr: &seqExpr{
					pos: position{line: 104, col: 16, offset: 2825},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 104, col: 16, offset: 2825},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 104, col: 20, offset: 2829},
							label: "s",
							expr: &oneOrMoreExpr{
								pos: position{line: 104, col: 23, offset: 2832},
								expr: &ruleRefExpr{
									pos:  position{line: 104, col: 23, offset: 2832},
									name: "Selection",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 104, col: 35, offset: 2844},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Selection",
			pos:  position{line: 115, col: 1, offset: 3110},
			expr: &choiceExpr{
				pos: position{line: 115, col: 13, offset: 3124},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 115, col: 13, offset: 3124},
						run: (*parser).callonSelection2,
						expr: &seqExpr{
							pos: position{line: 115, col: 14, offset: 3125},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 115, col: 14, offset: 3125},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 115, col: 16, offset: 3127},
									label: "f",
									expr: &ruleRefExpr{
										pos:  position{line: 115, col: 18, offset: 3129},
										name: "Field",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 115, col: 24, offset: 3135},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 118, col: 5, offset: 3217},
						run: (*parser).callonSelection8,
						expr: &seqExpr{
							pos: position{line: 118, col: 6, offset: 3218},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 118, col: 6, offset: 3218},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 118, col: 8, offset: 3220},
									label: "fs",
									expr: &ruleRefExpr{
										pos:  position{line: 118, col: 11, offset: 3223},
										name: "FragmentSpread",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 118, col: 26, offset: 3238},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 121, col: 5, offset: 3357},
						run: (*parser).callonSelection14,
						expr: &seqExpr{
							pos: position{line: 121, col: 6, offset: 3358},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 121, col: 6, offset: 3358},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 121, col: 8, offset: 3360},
									label: "fs",
									expr: &ruleRefExpr{
										pos:  position{line: 121, col: 11, offset: 3363},
										name: "InlineFragment",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 121, col: 26, offset: 3378},
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
			pos:  position{line: 126, col: 1, offset: 3496},
			expr: &actionExpr{
				pos: position{line: 126, col: 9, offset: 3506},
				run: (*parser).callonField1,
				expr: &seqExpr{
					pos: position{line: 126, col: 9, offset: 3506},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 126, col: 9, offset: 3506},
							label: "fa",
							expr: &zeroOrOneExpr{
								pos: position{line: 126, col: 12, offset: 3509},
								expr: &ruleRefExpr{
									pos:  position{line: 126, col: 12, offset: 3509},
									name: "FieldAlias",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 24, offset: 3521},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 26, offset: 3523},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 29, offset: 3526},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 39, offset: 3536},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 41, offset: 3538},
							label: "as",
							expr: &zeroOrOneExpr{
								pos: position{line: 126, col: 44, offset: 3541},
								expr: &ruleRefExpr{
									pos:  position{line: 126, col: 44, offset: 3541},
									name: "Arguments",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 55, offset: 3552},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 57, offset: 3554},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 126, col: 60, offset: 3557},
								expr: &ruleRefExpr{
									pos:  position{line: 126, col: 60, offset: 3557},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 72, offset: 3569},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 74, offset: 3571},
							label: "sels",
							expr: &zeroOrOneExpr{
								pos: position{line: 126, col: 79, offset: 3576},
								expr: &ruleRefExpr{
									pos:  position{line: 126, col: 79, offset: 3576},
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
			pos:  position{line: 153, col: 1, offset: 4099},
			expr: &actionExpr{
				pos: position{line: 153, col: 14, offset: 4114},
				run: (*parser).callonFieldAlias1,
				expr: &seqExpr{
					pos: position{line: 153, col: 14, offset: 4114},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 153, col: 14, offset: 4114},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 16, offset: 4116},
								name: "Name",
							},
						},
						&litMatcher{
							pos:        position{line: 153, col: 21, offset: 4121},
							val:        ":",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "FieldName",
			pos:  position{line: 154, col: 1, offset: 4143},
			expr: &ruleRefExpr{
				pos:  position{line: 154, col: 13, offset: 4157},
				name: "Name",
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 155, col: 1, offset: 4162},
			expr: &actionExpr{
				pos: position{line: 155, col: 13, offset: 4176},
				run: (*parser).callonArguments1,
				expr: &seqExpr{
					pos: position{line: 155, col: 13, offset: 4176},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 155, col: 13, offset: 4176},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 155, col: 17, offset: 4180},
							label: "args",
							expr: &zeroOrMoreExpr{
								pos: position{line: 155, col: 23, offset: 4186},
								expr: &ruleRefExpr{
									pos:  position{line: 155, col: 23, offset: 4186},
									name: "Argument",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 155, col: 34, offset: 4197},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Argument",
			pos:  position{line: 166, col: 1, offset: 4442},
			expr: &actionExpr{
				pos: position{line: 166, col: 12, offset: 4455},
				run: (*parser).callonArgument1,
				expr: &seqExpr{
					pos: position{line: 166, col: 12, offset: 4455},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 166, col: 12, offset: 4455},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 166, col: 14, offset: 4457},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 17, offset: 4460},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 166, col: 30, offset: 4473},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 166, col: 32, offset: 4475},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 166, col: 36, offset: 4479},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 166, col: 38, offset: 4481},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 40, offset: 4483},
								name: "Value",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 166, col: 46, offset: 4489},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentName",
			pos:  position{line: 172, col: 1, offset: 4562},
			expr: &ruleRefExpr{
				pos:  position{line: 172, col: 16, offset: 4579},
				name: "Name",
			},
		},
		{
			name: "Name",
			pos:  position{line: 174, col: 1, offset: 4585},
			expr: &actionExpr{
				pos: position{line: 174, col: 8, offset: 4594},
				run: (*parser).callonName1,
				expr: &seqExpr{
					pos: position{line: 174, col: 8, offset: 4594},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 174, col: 8, offset: 4594},
							val:        "[a-z_]i",
							chars:      []rune{'_'},
							ranges:     []rune{'a', 'z'},
							ignoreCase: true,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 174, col: 16, offset: 4602},
							expr: &charClassMatcher{
								pos:        position{line: 174, col: 16, offset: 4602},
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
			pos:  position{line: 178, col: 1, offset: 4647},
			expr: &actionExpr{
				pos: position{line: 178, col: 19, offset: 4667},
				run: (*parser).callonFragmentSpread1,
				expr: &seqExpr{
					pos: position{line: 178, col: 19, offset: 4667},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 178, col: 19, offset: 4667},
							val:        "...",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 178, col: 25, offset: 4673},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 178, col: 28, offset: 4676},
								name: "FragmentName",
							},
						},
						&labeledExpr{
							pos:   position{line: 178, col: 41, offset: 4689},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 178, col: 44, offset: 4692},
								expr: &ruleRefExpr{
									pos:  position{line: 178, col: 44, offset: 4692},
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
			pos:  position{line: 189, col: 1, offset: 4892},
			expr: &actionExpr{
				pos: position{line: 189, col: 19, offset: 4912},
				run: (*parser).callonInlineFragment1,
				expr: &seqExpr{
					pos: position{line: 189, col: 19, offset: 4912},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 189, col: 19, offset: 4912},
							val:        "...",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 25, offset: 4918},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 189, col: 27, offset: 4920},
							val:        "on",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 32, offset: 4925},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 34, offset: 4927},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 37, offset: 4930},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 46, offset: 4939},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 48, offset: 4941},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 51, offset: 4944},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 51, offset: 4944},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 63, offset: 4956},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 189, col: 65, offset: 4958},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 70, offset: 4963},
								name: "SelectionSet",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 83, offset: 4976},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FragmentDefinition",
			pos:  position{line: 201, col: 1, offset: 5220},
			expr: &actionExpr{
				pos: position{line: 201, col: 22, offset: 5243},
				run: (*parser).callonFragmentDefinition1,
				expr: &seqExpr{
					pos: position{line: 201, col: 22, offset: 5243},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 201, col: 22, offset: 5243},
							val:        "fragment",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 33, offset: 5254},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 201, col: 35, offset: 5256},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 38, offset: 5259},
								name: "FragmentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 51, offset: 5272},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 201, col: 53, offset: 5274},
							val:        "on",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 58, offset: 5279},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 201, col: 60, offset: 5281},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 63, offset: 5284},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 73, offset: 5294},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 201, col: 75, offset: 5296},
							label: "ds",
							expr: &zeroOrOneExpr{
								pos: position{line: 201, col: 78, offset: 5299},
								expr: &ruleRefExpr{
									pos:  position{line: 201, col: 78, offset: 5299},
									name: "Directives",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 90, offset: 5311},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 201, col: 92, offset: 5313},
							label: "sels",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 97, offset: 5318},
								name: "SelectionSet",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 110, offset: 5331},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FragmentName",
			pos:  position{line: 213, col: 1, offset: 5598},
			expr: &actionExpr{
				pos: position{line: 213, col: 16, offset: 5615},
				run: (*parser).callonFragmentName1,
				expr: &labeledExpr{
					pos:   position{line: 213, col: 16, offset: 5615},
					label: "n",
					expr: &ruleRefExpr{
						pos:  position{line: 213, col: 18, offset: 5617},
						name: "Name",
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 215, col: 1, offset: 5641},
			expr: &actionExpr{
				pos: position{line: 215, col: 9, offset: 5651},
				run: (*parser).callonValue1,
				expr: &seqExpr{
					pos: position{line: 215, col: 9, offset: 5651},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 215, col: 9, offset: 5651},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 215, col: 11, offset: 5653},
							label: "v",
							expr: &choiceExpr{
								pos: position{line: 215, col: 14, offset: 5656},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 215, col: 14, offset: 5656},
										name: "Null",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 21, offset: 5663},
										name: "Boolean",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 31, offset: 5673},
										name: "Int",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 37, offset: 5679},
										name: "Float",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 45, offset: 5687},
										name: "String",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 54, offset: 5696},
										name: "EnumValue",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 66, offset: 5708},
										name: "Array",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 74, offset: 5716},
										name: "Object",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 83, offset: 5725},
										name: "Variable",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 215, col: 93, offset: 5735},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Null",
			pos:  position{line: 219, col: 1, offset: 5757},
			expr: &actionExpr{
				pos: position{line: 219, col: 8, offset: 5766},
				run: (*parser).callonNull1,
				expr: &litMatcher{
					pos:        position{line: 219, col: 8, offset: 5766},
					val:        "null",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Boolean",
			pos:  position{line: 220, col: 1, offset: 5793},
			expr: &choiceExpr{
				pos: position{line: 220, col: 11, offset: 5805},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 220, col: 11, offset: 5805},
						run: (*parser).callonBoolean2,
						expr: &litMatcher{
							pos:        position{line: 220, col: 11, offset: 5805},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 220, col: 41, offset: 5835},
						run: (*parser).callonBoolean4,
						expr: &litMatcher{
							pos:        position{line: 220, col: 41, offset: 5835},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Int",
			pos:  position{line: 221, col: 1, offset: 5865},
			expr: &actionExpr{
				pos: position{line: 221, col: 7, offset: 5873},
				run: (*parser).callonInt1,
				expr: &seqExpr{
					pos: position{line: 221, col: 7, offset: 5873},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 221, col: 7, offset: 5873},
							expr: &ruleRefExpr{
								pos:  position{line: 221, col: 7, offset: 5873},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 221, col: 13, offset: 5879},
							name: "IntegerPart",
						},
					},
				},
			},
		},
		{
			name: "Float",
			pos:  position{line: 224, col: 1, offset: 5932},
			expr: &actionExpr{
				pos: position{line: 224, col: 9, offset: 5942},
				run: (*parser).callonFloat1,
				expr: &seqExpr{
					pos: position{line: 224, col: 9, offset: 5942},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 224, col: 9, offset: 5942},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 9, offset: 5942},
								name: "Sign",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 224, col: 15, offset: 5948},
							name: "IntegerPart",
						},
						&litMatcher{
							pos:        position{line: 224, col: 27, offset: 5960},
							val:        ".",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 224, col: 31, offset: 5964},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 31, offset: 5964},
								name: "Digit",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 224, col: 38, offset: 5971},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 38, offset: 5971},
								name: "ExponentPart",
							},
						},
					},
				},
			},
		},
		{
			name: "Sign",
			pos:  position{line: 227, col: 1, offset: 6036},
			expr: &litMatcher{
				pos:        position{line: 227, col: 8, offset: 6045},
				val:        "-",
				ignoreCase: false,
			},
		},
		{
			name: "IntegerPart",
			pos:  position{line: 228, col: 1, offset: 6049},
			expr: &choiceExpr{
				pos: position{line: 228, col: 15, offset: 6065},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 228, col: 15, offset: 6065},
						val:        "0",
						ignoreCase: false,
					},
					&seqExpr{
						pos: position{line: 228, col: 21, offset: 6071},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 228, col: 21, offset: 6071},
								name: "NonZeroDigit",
							},
							&zeroOrMoreExpr{
								pos: position{line: 228, col: 34, offset: 6084},
								expr: &ruleRefExpr{
									pos:  position{line: 228, col: 34, offset: 6084},
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
			pos:  position{line: 229, col: 1, offset: 6091},
			expr: &seqExpr{
				pos: position{line: 229, col: 16, offset: 6108},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 229, col: 16, offset: 6108},
						val:        "e",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 229, col: 20, offset: 6112},
						expr: &ruleRefExpr{
							pos:  position{line: 229, col: 20, offset: 6112},
							name: "Sign",
						},
					},
					&oneOrMoreExpr{
						pos: position{line: 229, col: 26, offset: 6118},
						expr: &ruleRefExpr{
							pos:  position{line: 229, col: 26, offset: 6118},
							name: "Digit",
						},
					},
				},
			},
		},
		{
			name: "Digit",
			pos:  position{line: 230, col: 1, offset: 6125},
			expr: &charClassMatcher{
				pos:        position{line: 230, col: 9, offset: 6135},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDigit",
			pos:  position{line: 231, col: 1, offset: 6141},
			expr: &charClassMatcher{
				pos:        position{line: 231, col: 16, offset: 6158},
				val:        "[123456789]",
				chars:      []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "String",
			pos:  position{line: 232, col: 1, offset: 6170},
			expr: &actionExpr{
				pos: position{line: 232, col: 10, offset: 6181},
				run: (*parser).callonString1,
				expr: &seqExpr{
					pos: position{line: 232, col: 10, offset: 6181},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 232, col: 10, offset: 6181},
							val:        "\"",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 232, col: 14, offset: 6185},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 16, offset: 6187},
								name: "string",
							},
						},
						&litMatcher{
							pos:        position{line: 232, col: 23, offset: 6194},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "string",
			pos:  position{line: 235, col: 1, offset: 6226},
			expr: &actionExpr{
				pos: position{line: 235, col: 10, offset: 6237},
				run: (*parser).callonstring1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 235, col: 10, offset: 6237},
					expr: &ruleRefExpr{
						pos:  position{line: 235, col: 10, offset: 6237},
						name: "StringCharacter",
					},
				},
			},
		},
		{
			name: "StringCharacter",
			pos:  position{line: 238, col: 1, offset: 6286},
			expr: &choiceExpr{
				pos: position{line: 238, col: 19, offset: 6306},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 238, col: 19, offset: 6306},
						val:        "[^\\\\\"]",
						chars:      []rune{'\\', '"'},
						ignoreCase: false,
						inverted:   true,
					},
					&seqExpr{
						pos: position{line: 238, col: 28, offset: 6315},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 238, col: 28, offset: 6315},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 238, col: 33, offset: 6320},
								name: "EscapedCharacter",
							},
						},
					},
					&seqExpr{
						pos: position{line: 238, col: 52, offset: 6339},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 238, col: 52, offset: 6339},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 238, col: 57, offset: 6344},
								name: "EscapedUnicode",
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedUnicode",
			pos:  position{line: 239, col: 1, offset: 6359},
			expr: &seqExpr{
				pos: position{line: 239, col: 18, offset: 6378},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 239, col: 18, offset: 6378},
						val:        "u",
						ignoreCase: false,
					},
					&charClassMatcher{
						pos:        position{line: 239, col: 22, offset: 6382},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 239, col: 32, offset: 6392},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 239, col: 42, offset: 6402},
						val:        "[0-9a-f]i",
						ranges:     []rune{'0', '9', 'a', 'f'},
						ignoreCase: true,
						inverted:   false,
					},
					&charClassMatcher{
						pos:        position{line: 239, col: 52, offset: 6412},
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
			pos:  position{line: 240, col: 1, offset: 6422},
			expr: &choiceExpr{
				pos: position{line: 240, col: 20, offset: 6443},
				alternatives: []interface{}{
					&charClassMatcher{
						pos:        position{line: 240, col: 20, offset: 6443},
						val:        "[\"/bfnrt]",
						chars:      []rune{'"', '/', 'b', 'f', 'n', 'r', 't'},
						ignoreCase: false,
						inverted:   false,
					},
					&litMatcher{
						pos:        position{line: 240, col: 32, offset: 6455},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "EnumValue",
			pos:  position{line: 242, col: 1, offset: 6461},
			expr: &actionExpr{
				pos: position{line: 242, col: 13, offset: 6475},
				run: (*parser).callonEnumValue1,
				expr: &seqExpr{
					pos: position{line: 242, col: 13, offset: 6475},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 242, col: 13, offset: 6475},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 242, col: 16, offset: 6478},
								name: "TypeName",
							},
						},
						&litMatcher{
							pos:        position{line: 242, col: 25, offset: 6487},
							val:        ".",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 242, col: 29, offset: 6491},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 242, col: 31, offset: 6493},
								name: "EnumValueName",
							},
						},
					},
				},
			},
		},
		{
			name: "Array",
			pos:  position{line: 248, col: 1, offset: 6597},
			expr: &seqExpr{
				pos: position{line: 248, col: 9, offset: 6607},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 248, col: 9, offset: 6607},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 248, col: 13, offset: 6611},
						expr: &ruleRefExpr{
							pos:  position{line: 248, col: 13, offset: 6611},
							name: "Value",
						},
					},
					&litMatcher{
						pos:        position{line: 248, col: 20, offset: 6618},
						val:        "]",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Object",
			pos:  position{line: 249, col: 1, offset: 6622},
			expr: &seqExpr{
				pos: position{line: 249, col: 10, offset: 6633},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 249, col: 10, offset: 6633},
						val:        "{",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 249, col: 14, offset: 6637},
						expr: &ruleRefExpr{
							pos:  position{line: 249, col: 14, offset: 6637},
							name: "Property",
						},
					},
					&litMatcher{
						pos:        position{line: 249, col: 24, offset: 6647},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 251, col: 1, offset: 6652},
			expr: &choiceExpr{
				pos: position{line: 251, col: 12, offset: 6665},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 251, col: 12, offset: 6665},
						run: (*parser).callonVariable2,
						expr: &seqExpr{
							pos: position{line: 251, col: 12, offset: 6665},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 251, col: 12, offset: 6665},
									label: "vn",
									expr: &ruleRefExpr{
										pos:  position{line: 251, col: 15, offset: 6668},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 251, col: 28, offset: 6681},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 251, col: 32, offset: 6685},
									label: "pn",
									expr: &ruleRefExpr{
										pos:  position{line: 251, col: 35, offset: 6688},
										name: "PropertyName",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 256, col: 5, offset: 6821},
						run: (*parser).callonVariable9,
						expr: &labeledExpr{
							pos:   position{line: 256, col: 5, offset: 6821},
							label: "vn",
							expr: &ruleRefExpr{
								pos:  position{line: 256, col: 8, offset: 6824},
								name: "VariableName",
							},
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 261, col: 1, offset: 6897},
			expr: &actionExpr{
				pos: position{line: 261, col: 16, offset: 6914},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 261, col: 16, offset: 6914},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 261, col: 16, offset: 6914},
							val:        "$",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 261, col: 20, offset: 6918},
							expr: &charClassMatcher{
								pos:        position{line: 261, col: 20, offset: 6918},
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
			pos:  position{line: 269, col: 1, offset: 7241},
			expr: &seqExpr{
				pos: position{line: 269, col: 12, offset: 7254},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 269, col: 12, offset: 7254},
						name: "PropertyName",
					},
					&litMatcher{
						pos:        position{line: 269, col: 25, offset: 7267},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 269, col: 29, offset: 7271},
						name: "Value",
					},
				},
			},
		},
		{
			name: "PropertyName",
			pos:  position{line: 270, col: 1, offset: 7277},
			expr: &ruleRefExpr{
				pos:  position{line: 270, col: 16, offset: 7294},
				name: "Name",
			},
		},
		{
			name: "Directives",
			pos:  position{line: 272, col: 1, offset: 7300},
			expr: &actionExpr{
				pos: position{line: 272, col: 14, offset: 7315},
				run: (*parser).callonDirectives1,
				expr: &labeledExpr{
					pos:   position{line: 272, col: 14, offset: 7315},
					label: "ds",
					expr: &oneOrMoreExpr{
						pos: position{line: 272, col: 17, offset: 7318},
						expr: &ruleRefExpr{
							pos:  position{line: 272, col: 17, offset: 7318},
							name: "Directive",
						},
					},
				},
			},
		},
		{
			name: "Directive",
			pos:  position{line: 279, col: 1, offset: 7467},
			expr: &actionExpr{
				pos: position{line: 279, col: 13, offset: 7481},
				run: (*parser).callonDirective1,
				expr: &seqExpr{
					pos: position{line: 279, col: 13, offset: 7481},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 279, col: 13, offset: 7481},
							val:        "@",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 279, col: 17, offset: 7485},
							label: "d",
							expr: &choiceExpr{
								pos: position{line: 279, col: 20, offset: 7488},
								alternatives: []interface{}{
									&actionExpr{
										pos: position{line: 279, col: 20, offset: 7488},
										run: (*parser).callonDirective6,
										expr: &seqExpr{
											pos: position{line: 279, col: 21, offset: 7489},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 279, col: 21, offset: 7489},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 279, col: 24, offset: 7492},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 279, col: 38, offset: 7506},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 279, col: 42, offset: 7510},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 279, col: 44, offset: 7512},
													label: "v",
													expr: &ruleRefExpr{
														pos:  position{line: 279, col: 46, offset: 7514},
														name: "Value",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 285, col: 5, offset: 7624},
										run: (*parser).callonDirective14,
										expr: &seqExpr{
											pos: position{line: 285, col: 6, offset: 7625},
											exprs: []interface{}{
												&labeledExpr{
													pos:   position{line: 285, col: 6, offset: 7625},
													label: "dn",
													expr: &ruleRefExpr{
														pos:  position{line: 285, col: 9, offset: 7628},
														name: "DirectiveName",
													},
												},
												&litMatcher{
													pos:        position{line: 285, col: 23, offset: 7642},
													val:        ":",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 285, col: 27, offset: 7646},
													name: "_",
												},
												&labeledExpr{
													pos:   position{line: 285, col: 29, offset: 7648},
													label: "t",
													expr: &ruleRefExpr{
														pos:  position{line: 285, col: 31, offset: 7650},
														name: "Type",
													},
												},
											},
										},
									},
									&actionExpr{
										pos: position{line: 291, col: 5, offset: 7759},
										run: (*parser).callonDirective22,
										expr: &labeledExpr{
											pos:   position{line: 291, col: 5, offset: 7759},
											label: "dn",
											expr: &ruleRefExpr{
												pos:  position{line: 291, col: 8, offset: 7762},
												name: "DirectiveName",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 295, col: 5, offset: 7838},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DirectiveName",
			pos:  position{line: 299, col: 1, offset: 7860},
			expr: &ruleRefExpr{
				pos:  position{line: 299, col: 17, offset: 7878},
				name: "Name",
			},
		},
		{
			name: "Type",
			pos:  position{line: 301, col: 1, offset: 7884},
			expr: &actionExpr{
				pos: position{line: 301, col: 8, offset: 7893},
				run: (*parser).callonType1,
				expr: &labeledExpr{
					pos:   position{line: 301, col: 8, offset: 7893},
					label: "t",
					expr: &choiceExpr{
						pos: position{line: 301, col: 11, offset: 7896},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 301, col: 11, offset: 7896},
								name: "OptionalType",
							},
							&ruleRefExpr{
								pos:  position{line: 301, col: 26, offset: 7911},
								name: "GenericType",
							},
						},
					},
				},
			},
		},
		{
			name: "OptionalType",
			pos:  position{line: 302, col: 1, offset: 7942},
			expr: &actionExpr{
				pos: position{line: 302, col: 16, offset: 7959},
				run: (*parser).callonOptionalType1,
				expr: &seqExpr{
					pos: position{line: 302, col: 16, offset: 7959},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 302, col: 16, offset: 7959},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 302, col: 18, offset: 7961},
								name: "GenericType",
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 30, offset: 7973},
							val:        "?",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "GenericType",
			pos:  position{line: 307, col: 1, offset: 8044},
			expr: &actionExpr{
				pos: position{line: 307, col: 15, offset: 8060},
				run: (*parser).callonGenericType1,
				expr: &seqExpr{
					pos: position{line: 307, col: 15, offset: 8060},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 307, col: 15, offset: 8060},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 307, col: 18, offset: 8063},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 307, col: 27, offset: 8072},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 307, col: 29, offset: 8074},
							label: "tps",
							expr: &zeroOrOneExpr{
								pos: position{line: 307, col: 33, offset: 8078},
								expr: &ruleRefExpr{
									pos:  position{line: 307, col: 33, offset: 8078},
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
			pos:  position{line: 312, col: 1, offset: 8145},
			expr: &seqExpr{
				pos: position{line: 312, col: 14, offset: 8160},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 312, col: 14, offset: 8160},
						val:        ":",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 312, col: 18, offset: 8164},
						val:        "<",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 312, col: 22, offset: 8168},
						expr: &ruleRefExpr{
							pos:  position{line: 312, col: 22, offset: 8168},
							name: "Type",
						},
					},
					&litMatcher{
						pos:        position{line: 312, col: 28, offset: 8174},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TypeName",
			pos:  position{line: 313, col: 1, offset: 8178},
			expr: &ruleRefExpr{
				pos:  position{line: 313, col: 12, offset: 8191},
				name: "Name",
			},
		},
		{
			name: "TypeDefinition",
			pos:  position{line: 314, col: 1, offset: 8196},
			expr: &actionExpr{
				pos: position{line: 314, col: 18, offset: 8215},
				run: (*parser).callonTypeDefinition1,
				expr: &seqExpr{
					pos: position{line: 314, col: 18, offset: 8215},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 314, col: 18, offset: 8215},
							val:        "type",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 25, offset: 8222},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 314, col: 27, offset: 8224},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 314, col: 30, offset: 8227},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 39, offset: 8236},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 314, col: 41, offset: 8238},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 314, col: 44, offset: 8241},
								expr: &ruleRefExpr{
									pos:  position{line: 314, col: 44, offset: 8241},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 314, col: 56, offset: 8253},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 314, col: 58, offset: 8255},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 314, col: 62, offset: 8259},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 314, col: 66, offset: 8263},
								expr: &ruleRefExpr{
									pos:  position{line: 314, col: 66, offset: 8263},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 314, col: 83, offset: 8280},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeExtension",
			pos:  position{line: 334, col: 1, offset: 8761},
			expr: &actionExpr{
				pos: position{line: 334, col: 17, offset: 8779},
				run: (*parser).callonTypeExtension1,
				expr: &seqExpr{
					pos: position{line: 334, col: 17, offset: 8779},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 334, col: 17, offset: 8779},
							val:        "extend",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 26, offset: 8788},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 28, offset: 8790},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 31, offset: 8793},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 40, offset: 8802},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 42, offset: 8804},
							label: "is",
							expr: &zeroOrOneExpr{
								pos: position{line: 334, col: 45, offset: 8807},
								expr: &ruleRefExpr{
									pos:  position{line: 334, col: 45, offset: 8807},
									name: "Interfaces",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 57, offset: 8819},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 334, col: 59, offset: 8821},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 334, col: 63, offset: 8825},
							label: "fds",
							expr: &oneOrMoreExpr{
								pos: position{line: 334, col: 67, offset: 8829},
								expr: &ruleRefExpr{
									pos:  position{line: 334, col: 67, offset: 8829},
									name: "FieldDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 334, col: 84, offset: 8846},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Interfaces",
			pos:  position{line: 354, col: 1, offset: 9326},
			expr: &actionExpr{
				pos: position{line: 354, col: 14, offset: 9341},
				run: (*parser).callonInterfaces1,
				expr: &oneOrMoreExpr{
					pos: position{line: 354, col: 14, offset: 9341},
					expr: &ruleRefExpr{
						pos:  position{line: 354, col: 14, offset: 9341},
						name: "GenericType",
					},
				},
			},
		},
		{
			name: "FieldDefinition",
			pos:  position{line: 358, col: 1, offset: 9448},
			expr: &actionExpr{
				pos: position{line: 358, col: 19, offset: 9468},
				run: (*parser).callonFieldDefinition1,
				expr: &seqExpr{
					pos: position{line: 358, col: 19, offset: 9468},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 358, col: 19, offset: 9468},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 358, col: 21, offset: 9470},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 358, col: 24, offset: 9473},
								name: "FieldName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 358, col: 34, offset: 9483},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 358, col: 36, offset: 9485},
							label: "args",
							expr: &zeroOrOneExpr{
								pos: position{line: 358, col: 41, offset: 9490},
								expr: &ruleRefExpr{
									pos:  position{line: 358, col: 41, offset: 9490},
									name: "ArgumentDefinitions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 358, col: 62, offset: 9511},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 358, col: 64, offset: 9513},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 358, col: 68, offset: 9517},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 358, col: 70, offset: 9519},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 358, col: 72, offset: 9521},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 358, col: 77, offset: 9526},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinitions",
			pos:  position{line: 369, col: 1, offset: 9763},
			expr: &actionExpr{
				pos: position{line: 369, col: 23, offset: 9787},
				run: (*parser).callonArgumentDefinitions1,
				expr: &seqExpr{
					pos: position{line: 369, col: 23, offset: 9787},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 369, col: 23, offset: 9787},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 369, col: 27, offset: 9791},
							label: "args",
							expr: &oneOrMoreExpr{
								pos: position{line: 369, col: 32, offset: 9796},
								expr: &ruleRefExpr{
									pos:  position{line: 369, col: 32, offset: 9796},
									name: "ArgumentDefinition",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 369, col: 52, offset: 9816},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgumentDefinition",
			pos:  position{line: 376, col: 1, offset: 9979},
			expr: &actionExpr{
				pos: position{line: 376, col: 22, offset: 10002},
				run: (*parser).callonArgumentDefinition1,
				expr: &seqExpr{
					pos: position{line: 376, col: 22, offset: 10002},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 376, col: 22, offset: 10002},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 376, col: 24, offset: 10004},
							label: "an",
							expr: &ruleRefExpr{
								pos:  position{line: 376, col: 27, offset: 10007},
								name: "ArgumentName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 376, col: 40, offset: 10020},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 376, col: 42, offset: 10022},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 376, col: 46, offset: 10026},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 376, col: 48, offset: 10028},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 376, col: 50, offset: 10030},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 376, col: 55, offset: 10035},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 376, col: 57, offset: 10037},
							label: "dv",
							expr: &zeroOrOneExpr{
								pos: position{line: 376, col: 60, offset: 10040},
								expr: &ruleRefExpr{
									pos:  position{line: 376, col: 60, offset: 10040},
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
			pos:  position{line: 388, col: 1, offset: 10271},
			expr: &actionExpr{
				pos: position{line: 388, col: 18, offset: 10290},
				run: (*parser).callonEnumDefinition1,
				expr: &seqExpr{
					pos: position{line: 388, col: 18, offset: 10290},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 388, col: 18, offset: 10290},
							val:        "enum",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 388, col: 25, offset: 10297},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 388, col: 27, offset: 10299},
							label: "tn",
							expr: &ruleRefExpr{
								pos:  position{line: 388, col: 30, offset: 10302},
								name: "TypeName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 388, col: 39, offset: 10311},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 388, col: 41, offset: 10313},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 388, col: 45, offset: 10317},
							label: "vals",
							expr: &oneOrMoreExpr{
								pos: position{line: 388, col: 50, offset: 10322},
								expr: &ruleRefExpr{
									pos:  position{line: 388, col: 50, offset: 10322},
									name: "EnumValueName",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 388, col: 65, offset: 10337},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EnumValueName",
			pos:  position{line: 398, col: 1, offset: 10518},
			expr: &actionExpr{
				pos: position{line: 398, col: 17, offset: 10536},
				run: (*parser).callonEnumValueName1,
				expr: &seqExpr{
					pos: position{line: 398, col: 17, offset: 10536},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 398, col: 17, offset: 10536},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 398, col: 19, offset: 10538},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 398, col: 21, offset: 10540},
								name: "Name",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 398, col: 26, offset: 10545},
							name: "_",
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"ignored\"",
			pos:         position{line: 400, col: 1, offset: 10566},
			expr: &actionExpr{
				pos: position{line: 400, col: 15, offset: 10582},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 400, col: 15, offset: 10582},
					expr: &choiceExpr{
						pos: position{line: 400, col: 16, offset: 10583},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 400, col: 16, offset: 10583},
								name: "whitespace",
							},
							&ruleRefExpr{
								pos:  position{line: 400, col: 29, offset: 10596},
								name: "Comment",
							},
							&litMatcher{
								pos:        position{line: 400, col: 39, offset: 10606},
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
			pos:  position{line: 401, col: 1, offset: 10632},
			expr: &charClassMatcher{
				pos:        position{line: 401, col: 14, offset: 10647},
				val:        "[ \\n\\t\\r]",
				chars:      []rune{' ', '\n', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 403, col: 1, offset: 10658},
			expr: &notExpr{
				pos: position{line: 403, col: 7, offset: 10666},
				expr: &anyMatcher{
					line: 403, col: 8, offset: 10667,
				},
			},
		},
	},
}

func (c *current) onDocument2(stmts interface{}) (interface{}, error) {
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

func (p *parser) callonDocument2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument2(stack["stmts"])
}

func (c *current) onDocument8() (interface{}, error) {
	return nil, errors.New("no graphql document found")
}

func (p *parser) callonDocument8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument8()
}

func (c *current) onStatement2(s interface{}) (interface{}, error) {
	return s, nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["s"])
}

func (c *current) onStatement14() (interface{}, error) {
	panic(errors.New("expected top-level statement"))
}

func (p *parser) callonStatement14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement14()
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
		Type:         graphql.OperationQuery,
		SelectionSet: sels.(graphql.SelectionSet),
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
		SelectionSet:        sels.(graphql.SelectionSet),
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
