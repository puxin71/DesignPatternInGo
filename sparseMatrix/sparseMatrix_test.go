package sparsematrix_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	sparsematrix "github.com/puxin71/DesignPatternInGo/sparseMatrix"
	"github.com/stretchr/testify/assert"
)

func TestNewSparseMatrix(t *testing.T) {
	matrix, err := sparsematrix.NewSparseMatrix(-1, -1)
	assert.Nil(t, matrix, "negative row and col upperBounds")
	assert.NotNil(t, err, "negative row and col upperBounds")
	matrix, err = sparsematrix.NewSparseMatrix(math.MaxUint32, math.MaxUint32)
	assert.Nil(t, matrix, "large row and col upperBounds")
	assert.NotNil(t, err, "large row and col upperBounds")
	matrix, err = sparsematrix.NewSparseMatrix(0, 0)
	assert.Nil(t, matrix, "zero row and col upperBounds")
	assert.NotNil(t, err, "zero row and col upperBounds")
	matrix, err = sparsematrix.NewSparseMatrix(0, 1)
	assert.Nil(t, matrix, "zero row upperBound")
	assert.NotNil(t, err, "zero row upperBound")
	matrix, err = sparsematrix.NewSparseMatrix(1, 0)
	assert.Nil(t, matrix, "zero col upperBound")
	assert.NotNil(t, err, "zero col upperBound")
	matrix, err = sparsematrix.NewSparseMatrix(1, 1)
	assert.NotNil(t, matrix, "correct row, col upperBounds")
	assert.Nil(t, err, "correct row, col upperBounds")
}

func TestAdd(t *testing.T) {
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
		"first row and col":    {row: 0, col: 0, val: 1, expect_err: false, expect_value: 1},
		"diagonal row and col": {row: 1, col: 1, val: 2, expect_err: false, expect_value: 2},
	}
	for key, test := range tests {
		matrix, _ := sparsematrix.NewSparseMatrix(2, 2)
		err := matrix.Add(sparsematrix.Data{Row: test.row, Col: test.col, Value: test.val})
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

func TestGetRow(t *testing.T) {
	matrix, _ := sparsematrix.NewSparseMatrix(2, 2)
	row, err := matrix.GetRow(0)
	assert.Nil(t, err, "row 0 not present")
	assert.Nil(t, row, "row 0 not present")

	err = matrix.Add(sparsematrix.Data{Row: 0, Col: 1, Value: 10})
	assert.Nil(t, err, "add 10 to (0,1)")
	err = matrix.Add(sparsematrix.Data{Row: 0, Col: 0, Value: 1})
	assert.Nil(t, err, "add 1 to (0,0)")

	row, err = matrix.GetRow(1)
	assert.Nil(t, err, "row 1 not present")
	assert.Nil(t, row, "row 1 not present")
	row, err = matrix.GetRow(100)
	assert.NotNil(t, err, "row out of range error")
	assert.Nil(t, row, "row out of range error")
	row, err = matrix.GetRow(0)
	assert.Nil(t, err, "row 0 returned")
	assert.Equal(t, 2, len(row), "expect 2 data in row 0")
	assert.Equal(t, 1, row[0].Value, "expect first col in row")
	assert.Equal(t, 10, row[1].Value, "expect 2nd col in row")
}

func TestGetData_checkRowAndCol(t *testing.T) {
	matrix, _ := sparsematrix.NewSparseMatrix(2, 2)
	tests := map[string]struct {
		row int
		col int
	}{
		"negative row":        {row: -1},
		"negative col":        {row: 1, col: -1},
		"over row upperBound": {row: 100, col: 1},
		"over col upperBound": {row: 1, col: 100},
	}
	for key, test := range tests {
		data, err := matrix.GetData(test.row, test.col)
		assert.NotNil(t, err, fmt.Sprintf("key: %s", key))
		assert.Nil(t, data, "key: %s", key)
	}
}

func TestGetData(t *testing.T) {
	matrix, _ := sparsematrix.NewSparseMatrix(2, 2)
	matrix.Add(sparsematrix.Data{Row: 0, Col: 0, Value: 1})
	matrix.Add(sparsematrix.Data{Row: 1, Col: 1, Value: 10})

	data, err := matrix.GetData(0, 1)
	assert.Nil(t, err, "data not present at (0,1)")
	assert.Nil(t, data, "data not present at (0,1)")

	data, err = matrix.GetData(1, 1)
	assert.Nil(t, err, "found data at (1,1)")
	assert.Equal(t, 10, data.Value, "found data at (1,1)")

	matrix.Add(sparsematrix.Data{Row: 0, Col: 1, Value: 2})
	data, err = matrix.GetData(0, 1)
	assert.Nil(t, err, "found data at (0,1)")
	assert.Equal(t, 2, data.Value, "found data at (0,1)")
}

func TestPrint(t *testing.T) {
	matrix, _ := sparsematrix.NewSparseMatrix(10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				matrix.Add(sparsematrix.Data{Row: i, Col: j, Value: int(math.Pow10(i))})
			}
		}
	}
	t.Log(matrix.Print())

	matrix, _ = sparsematrix.NewSparseMatrix(5, 9)
	maxLoops := 100
	numberAdded := 0

	for i := 0; i < 30; i++ {
		if generateUniqueData(matrix, maxLoops) {
			numberAdded++
		}
	}
	t.Logf("number added: %d", numberAdded)
	t.Log(matrix.Print())
}

func generateUniqueData(matrix sparsematrix.SparseMatrix, maxRetries int) bool {
	var err error

	if matrix == nil {
		return false
	}
	count := maxRetries
	row := generateRandomIndex(0, matrix.MaxRow())
	col := generateRandomIndex(0, matrix.MaxColumn())
	data, _ := matrix.GetData(row, col)
	for (data != nil || err != nil) && count > 0 {
		row = generateRandomIndex(0, matrix.MaxRow())
		col = generateRandomIndex(0, matrix.MaxColumn())
		data, err = matrix.GetData(row, col)
		count--
	}
	if count == 0 {
		return false
	}

	value := int(math.Mod(float64(rand.Int()), float64(100)))
	count = maxRetries
	for value == 0 && count > 0 {
		value = int(math.Mod(float64(rand.Int()), float64(100)))
		count--
	}
	if count == 0 {
		return false
	}

	newData := sparsematrix.Data{Row: row, Col: col, Value: int(math.Mod(float64(rand.Int()), float64(100)))}
	matrix.Add(newData)
	return true
}

func generateRandomIndex(min, max int) int {
	// generate random row and column index using linear congruential generator.
	// result = rand(int) Mod (max - min + 1)
	if min > max {
		return 0
	}
	idx := int(math.Mod(float64(rand.Int()), float64(max-min+1)))
	return idx
}
