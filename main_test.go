package main

import (
	"reflect"
	"testing"
)

func TestAdd_article(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput []string
	}{
		{[]string{"a", "apple"}, []string{"an", "apple"}},
		{[]string{"a", "banana"}, []string{"a", "banana"}},
		{[]string{"A", "elephant"}, []string{"An", "elephant"}},
		{[]string{"A", "honor"}, []string{"An", "honor"}},
		{[]string{"an", "orange"}, []string{"an", "orange"}},
	}

	for _, testCase := range testCases {
		output := add_article(testCase.input)
		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("For input %v, expected %v, but got %v", testCase.input, testCase.expectedOutput, output)
		}
	}
}

func TestDeal_w_apos(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"'", "awsome", "'"},
			want:  []string{"'awsome'"},
		},
	}

	for _, test := range tests {
		result := deal_w_apos(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("deal_w_apos(%v) == %v, want %v", test.input, result, test.want)
		}
	}
}

func TestPunctuate(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"there", ",", "then", "bamm", "!!"},
			want:  []string{"there,", "then", "bamm!!"},
		},
	}

	for _, test := range tests {
		result := punctuate(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("deal_w_apos(%v) == %v, want %v", test.input, result, test.want)
		}
	}
}

func TestToTitle(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "awesome",
			want:  "Awesome",
		},
	}

	for _, test := range tests {
		result := toTitle(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("deal_w_apos(%v) == %v, want %v", test.input, result, test.want)
		}
	}
}

func TestHex2str(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1E",
			want:  "30",
		},
	}

	for _, test := range tests {
		result := hex2str(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("deal_w_apos(%v) == %v, want %v", test.input, result, test.want)
		}
	}
}

func TestBin2str(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "10",
			want:  "2",
		},
	}

	for _, test := range tests {
		result := bin2str(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("deal_w_apos(%v) == %v, want %v", test.input, result, test.want)
		}
	}
}

func TestRecieveParam(t *testing.T) {
	tests := []struct {
		input       string
		expected    []string
		description string
	}{
		// Add test cases here
		{
			input:       "1E (hex) files were added",
			expected:    []string{"30", "files", "were", "added"},
			description: "Test case for (hex)",
		},
		{
			input:       "10 (bin) years",
			expected:    []string{"2", "years"},
			description: "Test case for (bin)",
		},
		{
			input:       "Ready, set, go (up) !",
			expected:    []string{"Ready,", "set,", "GO!"},
			description: "Test case for (up)",
		},
		{
			input:       "I should stop SHOUTING (low)",
			expected:    []string{"I", "should", "stop", "shouting"},
			description: "Test case for (low)",
		},
		{
			input:       "Welcome to the Brooklyn bridge (cap)",
			expected:    []string{"Welcome", "to", "the", "Brooklyn", "Bridge"},
			description: "Test case for (cap)",
		},
		{
			input:       "This is so exciting (up, 2)",
			expected:    []string{"This", "is", "SO", "EXCITING"},
			description: "Test case for (up, 2)",
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := recieveParam(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test case failed: %s\nExpected: %v\nGot: %v", test.description, test.expected, result)
			}
		})
	}
}
