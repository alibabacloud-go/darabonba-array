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
	sep := ","
	res := Join(tea.StringSlice(array), tea.String(sep))
	utils.AssertEqual(t, "a,b,c,d", tea.StringValue(res))
	sep = ""
	res = Join(tea.StringSlice(array), tea.String(sep))
	utils.AssertEqual(t, "abcd", tea.StringValue(res))
	res = Join(tea.StringSlice(array), nil)
	utils.AssertEqual(t, "abcd", tea.StringValue(res))
	array = []string{"a"}
	sep = ","
	res = Join(tea.StringSlice(array), tea.String(sep))
	utils.AssertEqual(t, "a", tea.StringValue(res))
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

type Test struct {
	Tasks []*Task `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

type Task struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func TestAppend(t *testing.T) {
	Append(nil, tea.String("c"))
	array := []*string{tea.String("a"), tea.String("b")}
	Append(array, tea.String("c"))
	// r := reflect.ValueOf(array)
	// utils.AssertEqual(t, 3, r.Len())

	Append(array, tea.Int(1))
	// r = reflect.ValueOf(array)
	// utils.AssertEqual(t, 4, r.Len())

	test := &Test{}
	tmp := &Task{
		DataId: tea.String("test"),
	}
	test.Tasks = []*Task{tmp}
	Append(test.Tasks, tea.ToMap(tmp))
}
