// This file is auto-generated, don't edit it. Thanks.
/**
 * This is a array module
 */
package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

func TestSplit(t *testing.T) {
	array := []string{"a", "b", "c", "d", "e", "f"}
	index := 1
	limit := 3
	result := Split(tea.StringSlice(array), tea.Int(index), tea.Int(limit))
	utils.AssertEqual(t, []string{"b", "c"}, tea.StringSliceValue(result))
}

func TestContains(t *testing.T) {
	array := []string{"a", "b", "c"}
	sep := "j"
	result := Contains(tea.StringSlice(array), tea.String(sep))
	utils.AssertEqual(t, false, tea.BoolValue(result))
	sep = "a"
	result = Contains(tea.StringSlice(array), tea.String(sep))
	utils.AssertEqual(t, true, tea.BoolValue(result))
}

func TestConcat(t *testing.T) {
	array := []string{"a", "b", "c"}
	sep := []string{"p", "o"}
	result := Concat(tea.StringSlice(array), tea.StringSlice(sep))
	utils.AssertEqual(t, []string{"a", "b", "c", "p", "o"}, tea.StringSliceValue(result))
}

func TestIndex(t *testing.T) {
	array := []string{"a", "b", "c", "d"}
	str := "c"
	result := Index(tea.StringSlice(array), tea.String(str))
	utils.AssertEqual(t, 2, tea.IntValue(result))
}

func TestSize(t *testing.T) {
	array := []string{"a", "b", "c"}
	result := Size(tea.StringSlice(array))
	utils.AssertEqual(t, 3, tea.IntValue(result))
}

func TestGet(t *testing.T) {
	array := []string{"a", "b", "c", "d"}
	index := 3
	res := Get(tea.StringSlice(array), tea.Int(index))
	utils.AssertEqual(t, "d", tea.StringValue(res))
}

func TestJoin(t *testing.T) {
	array := []string{"a", "b", "c", "d"}
	sep := "e"
	res := Join(tea.StringSlice(array), tea.String(sep))
	utils.AssertEqual(t, "aebecede", tea.StringValue(res))
}

func TestAscSort(t *testing.T) {
	array := []string{"a", "g", "j", "b", "d", "c", "B"}
	res := AscSort(tea.StringSlice(array))
	utils.AssertEqual(t, []string{"B", "a", "b", "c", "d", "g", "j"}, tea.StringSliceValue(res))
}

func TestDescSort(t *testing.T) {
	array := []string{"a", "g", "j", "b", "d", "c", "B"}
	res := DescSort(tea.StringSlice(array))
	utils.AssertEqual(t, []string{"j", "g", "d", "c", "b", "a", "B"}, tea.StringSliceValue(res))
}
