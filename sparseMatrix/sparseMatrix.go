package sparsematrix

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"sort"
)

const (
	// MaximumRows is the most allowed upper boundary for number of rows
	MaximumRows = math.MaxUint32
	// MaximumCols is the most allowed upper boundary for number of columns
	MaximumCols = math.MaxUint32
	// DefaultPatchSize defines number of rows used to calculate Matrix Vector product
	DefaultPatchSize = 30
)

// Data identifies the row and column of a integer value in the spareMatrix
type Data struct {
	Row   int
	Col   int
	Value int
}

// RowList defines the non-zero data in a row
type RowList []Data

// ByColumn implements sort.Interface for []Data based on
// the col field.
type ByColumn RowList

func (c ByColumn) Len() int           { return len(c) }
func (c ByColumn) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByColumn) Less(i, j int) bool { return c[i].Col < c[j].Col }

type sparseMatrix struct {
	rowUpperBound    int             // maximum number of Rows allowed
	colUpperBound    int             // maximum of Columns allowed
	dataList         map[int]RowList // use row has the key
	batchedNumOfRows int             // number of rows in a batch
}

// SparseMatrix defines the public APIs of spareMatrix
type SparseMatrix interface {
	Add(number Data) error
	Remove(number Data) error
	GetData(row, col int) (*Data, error)
	GetRow(row int) (RowList, error)
	PrintRow(row int) string
	Print() string
	LargeMatrixVectorProduct(colMatrix []int, batchSize int) ([]int, error)
	MaxRow() int
	MaxColumn() int
}

// NewSparseMatrix creates a matrix with a row and column upperBound
func NewSparseMatrix(rowUpper, colUpper int) (SparseMatrix, error) {
	if rowUpper <= 0 {
		return nil, fmt.Errorf("invalid row upperBound: %d", rowUpper)
	}
	if rowUpper >= MaximumRows {
		return nil, fmt.Errorf("invalid row upperBound: %d", rowUpper)
	}
	if colUpper <= 0 {
		return nil, fmt.Errorf("invalid col upperBound: %d", colUpper)
	}
	if colUpper >= MaximumCols {
		return nil, fmt.Errorf("invalid col upperBound: %d", colUpper)
	}
	return &sparseMatrix{
		rowUpperBound: rowUpper,
		colUpperBound: colUpper,
		dataList:      make(map[int]RowList),
	}, nil
}

// Add adds non-zero int to the sparse matrix
func (m *sparseMatrix) Add(number Data) error {
	if err := m.isValidRow(number.Row); err != nil {
		return err
	}
	if err := m.isValidColumn(number.Col); err != nil {
		return err
	}
	if number.Value == 0 {
		return nil
	}

	// if data already exists, do thing
	if data, _ := m.GetData(number.Row, number.Col); data != nil {
		return nil
	}

	// add the new value to a row with ascending column order
	// use Golang built-in sort which uses a combination of
	// quick sort and insertion sort
	var aList RowList
	if _, ok := m.dataList[number.Row]; ok {
		aList = m.dataList[number.Row]
	}
	aList = append(aList, number)
	sort.Sort(ByColumn(aList))
	m.dataList[number.Row] = aList

	return nil
}

func (m *sparseMatrix) Remove(number Data) error {
	panic("not impl.")
}

// isValidRow checks the data's row index
func (m *sparseMatrix) isValidRow(row int) error {
	if row < 0 || row >= m.rowUpperBound {
		return fmt.Errorf("invalid row: %d", row)
	}
	return nil
}

// isValidColumn checks the data's column index
func (m *sparseMatrix) isValidColumn(col int) error {
	if col < 0 || col >= m.colUpperBound {
		return fmt.Errorf("invalid col: %d", col)
	}
	return nil
}

// Get all values in a row of the sparse matrix.
func (m *sparseMatrix) GetRow(row int) (RowList, error) {
	if err := m.isValidRow(row); err != nil {
		return nil, err
	}
	if len(m.dataList) == 0 {
		return nil, nil
	}
	rowList, ok := m.dataList[row]
	if !ok {
		return nil, nil
	}
	return rowList, nil
}

// GetData finds the data using the row and column index
func (m *sparseMatrix) GetData(row, col int) (*Data, error) {
	if err := m.isValidColumn(col); err != nil {
		return nil, err
	}
	rowList, err := m.GetRow(row)
	if err != nil {
		return nil, err
	}
	if rowList == nil {
		return nil, nil
	}
	if col > rowList[len(rowList)-1].Col {
		return nil, nil
	}
	// simple approach for now, just search for col linearly
	for _, data := range rowList {
		if data.Col == col {
			return &data, nil
		}
	}
	return nil, nil
}

func (m *sparseMatrix) PrintRow(row int) string {
	panic("not impl.")
}

// Print displays the sparse matrix in the standard matrix form
func (m *sparseMatrix) Print() string {
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	for i := 0; i < m.rowUpperBound; i++ {
		for j := 0; j < m.colUpperBound; j++ {
			data, _ := m.GetData(i, j)
			if data == nil {
				buffer.WriteString(fmt.Sprintf("%-10v", 0))
			} else {
				buffer.WriteString(fmt.Sprintf("%-10v", data.Value))
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// LargeMatrixVectorProduct performs concurrent multiplication on batch of rows.
// Merge the result in the end. Use this API only if the matrix is large.
// equation: result[n,1] = matrix[n,m] * columnVector[m,1]
func (m *sparseMatrix) LargeMatrixVectorProduct(colMatrix []int, batchSize int) ([]int, error) {
	if err := m.validColMatrix(colMatrix); err != nil {
		return nil, err
	}
	if batchSize <= 0 {
		// assume that the caller wants to use the default batch size, 100
		m.batchedNumOfRows = DefaultPatchSize
	} else {
		m.batchedNumOfRows = batchSize
	}

	// calculate range of rows for each batch operation
	// todo...
	return nil, nil
}

func (m *sparseMatrix) validColMatrix(colMatrix []int) error {
	// num of rows of colMatrix = num of columns of sparsematrix
	if len(colMatrix) == 0 {
		return errors.New("null column matrix vector")
	}
	if len(colMatrix) != m.colUpperBound {
		return fmt.Errorf("incompatible matrix and vector. vector row: %d, matrix columns: %d", len(colMatrix), m.colUpperBound)
	}
	return nil
}

func (m *sparseMatrix) MaxRow() int {
	return m.rowUpperBound
}

func (m *sparseMatrix) MaxColumn() int {
	return m.colUpperBound
}
