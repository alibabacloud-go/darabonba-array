// This file is auto-generated, don't edit it. Thanks.
/**
 * This is a array module
 */
package client

import (
	"sort"

	"github.com/alibabacloud-go/tea/tea"
)

func Split(raw []*string, index *int, limit *int) (_result []*string) {
	return raw[tea.IntValue(index):tea.IntValue(limit)]
}

func Contains(raw []*string, str *string) (_result *bool) {
	for _, v := range raw {
		if tea.StringValue(v) == tea.StringValue(str) {
			return tea.Bool(true)
		}
	}
	return tea.Bool(false)
}

func Index(raw []*string, str *string) (_result *int) {
	for i, v := range raw {
		if tea.StringValue(v) == tea.StringValue(str) {
			return tea.Int(i)
		}
	}
	return
}

func Size(raw []*string) (_result *int) {
	return tea.Int(len(raw))
}

func Get(raw []*string, index *int) (_result *string) {
	return raw[tea.IntValue(index)]
}

func Join(raw []*string, sep *string) (_result *string) {
	res := ""
	separator := ""
	for _, value := range raw {
		res = res + separator + tea.StringValue(value)
		separator = tea.StringValue(sep)
	}
	return tea.String(res)
}

func Concat(raw []*string, sep []*string) (_result []*string) {
	raw = append(raw, sep...)
	return raw
}

func AscSort(raw []*string) (_result []*string) {
	slice := tea.StringSliceValue(raw)
	sort.Strings(slice)
	return tea.StringSlice(slice)
}

func DescSort(raw []*string) (_result []*string) {
	slice := tea.StringSliceValue(raw)
	sort.Strings(slice)
	for i := 0; i <= (len(slice)-1)/2; i++ {
		temp := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = temp
	}
	return tea.StringSlice(slice)
}
