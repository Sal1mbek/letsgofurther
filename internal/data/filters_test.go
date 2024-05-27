package data

import (
	"reflect"
	"testing"
)

func TestCalculateMetadata(t *testing.T) {
	tests := []struct {
		totalRecords int
		page         int
		pageSize     int
		expected     Metadata
	}{
		{
			totalRecords: 0,
			page:         1,
			pageSize:     10,
			expected:     Metadata{},
		},
		{
			totalRecords: 25,
			page:         2,
			pageSize:     10,
			expected: Metadata{
				CurrentPage:  2,
				PageSize:     10,
				FirstPage:    1,
				LastPage:     3,
				TotalRecords: 25,
			},
		},
	}

	for _, test := range tests {
		result := calculateMetadata(test.totalRecords, test.page, test.pageSize)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For totalRecords=%d, page=%d, pageSize=%d, expected %v, but got %v",
				test.totalRecords, test.page, test.pageSize, test.expected, result)
		}
	}
}
