package stack

var pushTestSuite = []struct {
	name          string
	values        []int
	expectedStack []int
}{
	{"no values", []int{}, nil},
	{"single value", []int{1}, []int{1}},
	{"two values", []int{1, 2}, []int{1, 2}},
	{"multiple values", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
}

var popTestSuite = []struct {
	name          string
	values        []int
	expectedValue int
	expectedStack []int
	expectedError error
}{
	{"no values", []int{}, 0, nil, errIsEmptyStack},
	{"single value", []int{1}, 1, []int{}, nil},
	{"two values", []int{1, 2}, 2, []int{1}, nil},
	{"multiple values", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}, nil},
}

var frontTestSuite = []struct {
	name          string
	values        []int
	expectedValue int
	expectedStack []int
	expectedError error
}{
	{"no values", []int{}, 0, nil, errIsEmptyStack},
	{"single value", []int{1}, 1, []int{1}, nil},
	{"two values", []int{1, 2}, 2, []int{1, 2}, nil},
	{"multiple values", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}, nil},
}

var sizeTestSuite = []struct {
	name         string
	values       []int
	expectedSize int
}{
	{"no values", []int{}, 0},
	{"single value", []int{1}, 1},
	{"two values", []int{1, 2}, 2},
	{"multiple values", []int{1, 2, 3, 4, 5}, 5},
}
