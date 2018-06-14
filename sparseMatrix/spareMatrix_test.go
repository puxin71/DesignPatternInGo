package sparseMatrix_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/DesignPatternInGo/sparseMatrix"
	"github.com/stretchr/testify/assert"
)

func TestNewSparseMatrix(t *testing.T) {
	matrix, err := sparseMatrix.NewSparseMatrix(-1, -1)
	assert.Nil(t, matrix, "negative row and col upperBounds")
	assert.NotNil(t, err, "negative row and col upperBounds")
	matrix, err = sparseMatrix.NewSparseMatrix(math.MaxUint32, math.MaxUint32)
	assert.Nil(t, matrix, "large row and col upperBounds")
	assert.NotNil(t, err, "large row and col upperBounds")
	matrix, err = sparseMatrix.NewSparseMatrix(0, 0)
	assert.Nil(t, matrix, "zero row and col upperBounds")
	assert.NotNil(t, err, "zero row and col upperBounds")
	matrix, err = sparseMatrix.NewSparseMatrix(0, 1)
	assert.Nil(t, matrix, "zero row upperBound")
	assert.NotNil(t, err, "zero row upperBound")
	matrix, err = sparseMatrix.NewSparseMatrix(1, 0)
	assert.Nil(t, matrix, "zero col upperBound")
	assert.NotNil(t, err, "zero col upperBound")
	matrix, err = sparseMatrix.NewSparseMatrix(1, 1)
	assert.NotNil(t, matrix, "correct row, col upperBounds")
	assert.Nil(t, err, "correct row, col upperBounds")
}

func TestAdd(t *testing.T) {
	matrix, _ := sparseMatrix.NewSparseMatrix(2, 2)

	tests := map[string]struct {
		row          int
		col          int
		val          int
		expect_err   bool
		expect_value int
	}{
		"negative row":         {row: -1, expect_err: true},
		"negative col":         {row: 1, col: -1, expect_err: true},
		"over row upperBound":  {row: 100, col: 1, expect_err: true},
		"over col upperBound":  {row: 1, col: 100, expect_err: true},
		"zero value":           {row: 0, col: 0, val: 0, expect_err: false},
		"correct row and col":  {row: 0, col: 0, val: 1, expect_err: false, expect_value: 1},
		"diagonal row and col": {row: 1, col: 1, val: 2, expect_err: false, expect_value: 2},
	}
	for key, test := range tests {
		err := matrix.Add(sparseMatrix.Data{Row: test.row, Col: test.col, Value: test.val})
		if test.expect_err {
			assert.NotNil(t, err, fmt.Sprintf("test case: %s", key))
		} else {
			assert.Nil(t, err, fmt.Sprintf("test case: %s", key))

			//successfully add. check the result
			data, err := matrix.GetData(test.row, test.col)
			assert.Nil(t, err, "find the added data")
			if test.val == 0 {
				assert.Nil(t, data, "zero value is not added")
			} else {
				assert.NotNil(t, data, fmt.Sprintf("test case: %s", key))
				assert.Equal(t, test.expect_value, data.Value, fmt.Sprintf("test case: %s", key))
			}
		}
	}
}
