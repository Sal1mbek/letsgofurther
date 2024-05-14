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
		// Test case 1: totalRecords = 0
		{
			totalRecords: 0,
			page:         1,
			pageSize:     10,
			expected:     Metadata{},
		},
		// Test case 2: totalRecords > 0
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
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := calculateMetadata(test.totalRecords, test.page, test.pageSize)

		// Compare the actual result with the expected result
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For totalRecords=%d, page=%d, pageSize=%d, expected %v, but got %v",
				test.totalRecords, test.page, test.pageSize, test.expected, result)
		}
	}
}
