package main

import "testing"

func TestRegexp(t *testing.T) {
	var cases = []struct {
		input    []regexp_params
		expected []int
	}{
		{
			expected: []int{
				19,
				2,
				SLRE_NO_MATCH,
				2,
				1,
				2,
				4,
				-1,
				1,
				3,
				5,
				6,
				4,
				8,
				3,
				2,
				3,
				14,
				2,
			},
			input: []regexp_params{
				regexp_params{
					pattern:     "tel:\\+(\\d+[\\d-]+\\d)",
					expr:        "tel:+1-201-555-0123;a=b",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[abc]",
					expr:        "1c2",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[abc]",
					expr:        "1C2",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[abc]",
					expr:        "1C2",
					num_caps:    0,
					ignore_case: true,
				},
				regexp_params{
					pattern:     "[.2]",
					expr:        "1C3",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[\\S]+",
					expr:        "ab cd",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[\\S]+\\s+[tyc]*",
					expr:        "ab cd",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[\\d]",
					expr:        "ab cd",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[^\\d]",
					expr:        "ab cd",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[^\\d]+",
					expr:        "abc123",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[1-5]+",
					expr:        "123456789",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[1-5a-c]+",
					expr:        "123abcdef",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[1-5a-]+",
					expr:        "123abcdef",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[htps]+://",
					expr:        "https://",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[^\\s]+",
					expr:        "abc def",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[^fc]+",
					expr:        "abc def",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "[^d\\sf]+",
					expr:        "abc def",
					num_caps:    0,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "aa ([0-9]*) *([x-z]*)\\s+xy([yz])",
					expr:        "aa 1234 xy\nxyz",
					num_caps:    3,
					ignore_case: false,
				},
				regexp_params{
					pattern:     "^(te)",
					expr:        "tenacity subdues all",
					num_caps:    10,
					ignore_case: false,
				},
			},
		},
	}

	for _, item := range cases {
		for i, input := range item.input {
			result := regexp(input.pattern, input.expr, input.num_caps, input.ignore_case)

			if item.expected[i] != result {
				t.Error("expected", item.expected[i], "have", result)
			}
		}
	}
}

func BenchmarkCGO(b *testing.B) {
	CallCGo(b.N)
}

func BenchmarkGo(b *testing.B) {
	CallGo(b.N)
}
