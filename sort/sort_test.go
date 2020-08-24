package sort

import (
	"log"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

type TestCase struct {
	input, output []int
}

var TestCases = []TestCase{
	{[]int{}, []int{}},
	{[]int{2}, []int{2}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{2, 2, 1}, []int{1, 2, 2}},
	{[]int{4, 6, 1, 0, 3, 7, 9, 10, 12}, []int{0, 1, 3, 4, 6, 7, 9, 10, 12}},
}

func TestBubble(t *testing.T) {
	testSortAlgorithm(t, bubble)
	testPropertiesSortAlgorithm(t, bubble)
}

func TestBubble2(t *testing.T) {
	testSortAlgorithm(t, bubble2)
	testPropertiesSortAlgorithm(t, bubble2)
}

func TestSelection(t *testing.T) {
	testSortAlgorithm(t, selection)
	testPropertiesSortAlgorithm(t, selection)
}

func testSortAlgorithm(t *testing.T, sort func([]int)) {
	t.Helper()
	for _, testCase := range TestCases {
		input := make([]int, len(testCase.input))
		copy(input, testCase.input)
		sort(testCase.input)

		if !reflect.DeepEqual(testCase.input, testCase.output) {
			t.Errorf("wrong output on %v", input)
			t.Logf("want: %v", testCase.output)
			t.Logf(" got: %v", testCase.input)
		}
	}
}

func testPropertiesSortAlgorithm(t *testing.T, sort func([]int)) {
	t.Helper()
	assertion := func(length int16) bool {
		// can use uint as type of length here
		if length < 0 {
			log.Printf("length %d out of range in property test", length)
			return true
		}
		sorted, random := generateRandInts(int(length))
		sort(random)

		return reflect.DeepEqual(sorted, random)
	}

	cfg := &quick.Config{
		MaxCount: 5,
	}
	if err := quick.Check(assertion, cfg); err != nil {
		t.Error("failed checks", err)
	}
}

var _, benchmarkSlice = generateRandInts(100)

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble(benchmarkSlice)
	}
}

func BenchmarkBubble2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble2(benchmarkSlice)
	}
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selection(benchmarkSlice)
	}
}

func generateRandInts(length int) (sorted, random []int) {
	sorted = make([]int, length)
	for i := range sorted {
		sorted[i] = i
	}

	random = make([]int, length)
	copy(random, sorted)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(random), func(i, j int) { random[i], random[j] = random[j], random[i] })

	return
}
