package sortedslice

import (
	"slices"
	"testing"
)

func TestSortedSlice(t *testing.T) {
	var slice SortedSlice[string]
	if i := slice.Len(); i != 0 {
		t.Errorf("expected 0; got %d", i)
	}
	if slice.Delete("pie") {
		t.Error("expected false; got true")
	}
	if !slice.Insert("eta") {
		t.Error("expected true; got false")
	}
	if slice.Insert("eta") {
		t.Error("expected false; got true")
	}
	if i := slice.Len(); i != 1 {
		t.Errorf("expected 1; got %d", i)
	}
	greek := []string{"tau", "xi", "pi", "mu", "rho", "nu"}
	for _, text := range greek {
		if !slice.Insert(text) {
			t.Error("expected true; got false")
		}
	}
	if i := slice.Len(); i != 7 {
		t.Errorf("expected 7; got %d", i)
	}
	if i := slice.Find("xi"); i != 6 {
		t.Errorf("expected 6; got %d", i)
	}
	if i := slice.Find("zip"); i != -1 {
		t.Errorf("expected -1; got %d", i)
	}
	//tag::neweg[]
	lines := New[string](250) // length is 0; capacity is 250
	//end::neweg[]
	if slice.Equal(lines) || lines.Equal(slice) {
		t.Error("expected false; got true")
	}
	for _, text := range greek {
		if !lines.Insert(text) {
			t.Error("expected true; got false")
		}
	}
	if slice.Equal(lines) || lines.Equal(slice) {
		t.Error("expected false; got true")
	}
	if !lines.Insert("eta") {
		t.Error("expected true; got false")
	}
	if !(slice.Equal(lines) && lines.Equal(slice)) {
		t.Error("expected true; got false")
	}
	greek = append(greek, "eta")
	slices.Sort(greek)
	for i, element := range slice.All() {
		if element != greek[i] {
			t.Errorf("expected %q; got %q", greek[i], element)
		}
	}
	count := 0
	for i, word := range greek {
		if i%2 == 0 {
			if slice.Delete(word) {
				count++
			}
		}
	}
	if count != 4 {
		t.Errorf("expected 4; got %d", count)
	}
	odds := []string{"mu", "pi", "tau"}
	for i, element := range slice.All() {
		if element != odds[i] {
			t.Errorf("expected %q; got %q", odds[i], element)
		}
	}
}
func TestSortedSliceAPI(t *testing.T) {
	//tag::sortedslice1[]
	var chars SortedSlice[rune] // usable zero value
	//end::sortedslice1[]
	unique := 0
	//tag::sortedslice2[]
	for _, r := range "BLUEBERRYPIE" {
		if chars.Insert(r) { // true if inserted
			//end::sortedslice2[]
			unique++
		}
	}
	if unique != 8 {
		t.Errorf("expected 8; got %d", unique)
	}
	actual := []rune{}
	total := 0
	//tag::sortedslice3[]
	for i, char := range chars.All() { // 0:B 1:E 2:I … 7:Y
		//end::sortedslice3[]
		actual = append(actual, char)
		total += i
	}
	expected := "BEILPRUY"
	if got := string(actual); expected != got {
		t.Errorf("expected %q; got %q", expected, got)
	}
	//tag::sortedslice4[]
	ok := chars.Delete('R') // ok == true
	//end::sortedslice4[]
	if !ok {
		t.Error("expected true; got false")
	}
	//tag::sortedslice5[]
	ok = chars.Delete('Z') // ok == false
	//end::sortedslice5[]
	if ok {
		t.Error("expected false; got true")
	}
	//tag::sortedslice5[]
	i := chars.Find('U') // i == 5
	//end::sortedslice5[]
	if i != 5 {
		t.Errorf("expected 6; got %d", i)
	}
	//tag::sortedslice6[]
	i = chars.Find('A') // i == -1
	//end::sortedslice6[]
	if i != -1 {
		t.Errorf("expected -1; got %d", i)
	}
	//tag::sortedslice7[]
	size := chars.Len() // size == 7
	//end::sortedslice7[]
	if size != 7 {
		t.Errorf("expected 7; got %d", size)
	}
	//tag::sortedslice8[]
	var letters SortedSlice[rune]
	//end::sortedslice8[]
	//tag::sortedslice9[]
	if chars.Equal(letters) {
		//end::sortedslice9[]
		t.Error("expected false; got true")
	}
}
