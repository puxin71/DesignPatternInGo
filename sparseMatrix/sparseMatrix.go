package sparseMatrix

import (
	"bitbucket.ciena.com/scm/bpp_infrastructure/pmprocessor/calculations"
	"fmt"
	"math"
	"sort"
)

const (
	MaximumRows = math.MaxUint32
	MaximumCols = math.MaxUint32
	InvalidResult = math.Inf(0)
)

type Data struct {
	Row   int
	Col   int
	Value int
}

type RowList []Data

// ByColumn implements sort.Interface for []Data based on
// the col field.
type ByColumn RowList

func (c ByColumn) Len() int           { return len(c) }
func (c ByColumn) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByColumn) Less(i, j int) bool { return c[i].Col < c[j].Col }

type sparseMatrix struct {
	rowUpperBound int             // maximum number of Rows allowed
	colUpperBound int             // maximum of Columns allowed
	dataList      map[int]RowList // use row has the key
	batchedNumOfRows int // number of rows in a batch
}

type SparseMatrix interface {
	Add(number Data) error
	Remove(number Data) error
	GetData(row, col int) (*Data, error)
	GetRowSortByCol(row int) (RowList, error)
	PrintRow(row int) string
	Print() string
	LargeMatrixVectorProduct(colMetrix []int, batchSize int) (int, error)
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

	// note that at this point, the numbers are
	// added to the correct rows, but within a row,
	// the numbers may not be sorted with an increasing column index
	var aList RowList
	if _, ok := m.dataList[number.Row]; ok {
		aList = m.dataList[number.Row]
	}
	aList = append(aList, number)
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

// GetRowSortByCol returns list of integers that are in the same row.
// the numbers are sorted by the column index in the asceding order
func (m *sparseMatrix) GetRowSortByCol(row int) (RowList, error) {
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
	sort.Sort(ByColumn(rowList))
	return rowList, nil
}

// GetData finds the data using the row and column index
func (m *sparseMatrix) GetData(row, col int) (*Data, error) {
	if err := m.isValidColumn(col); err != nil {
		return nil, err
	}
	rowList, err := m.GetRowSortByCol(row)
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

func (m *sparseMatrix) Print() string {
	panic("not impl.")
}

// LargeMatrixVectorProduct performs concurrent multiplication on batch of rows.
// Merge the result in the end. Use this API only if the matrix is large,
// rows > 100,000.
// equation: result[n,1] = matrix[n,m] * columnVector[m,1]
func (m *sparseMatrix) LargeMatrixVectorProduct(colMetrix []int, batchSize int) ([]int, error) {
	numOfRows : = 0

	if err := m.validColMetrix(colMetrix), err != nil {
		return nil, err
	}
	if batchSize <= 0 {
		// assume that the caller wants to use the default batch size, 100
		m.batchedNumOfRows = DefaultPatchSize
	} else {
		m.batchedNumOfRows = batchSize
	}

	// calculate range of rows for each batch operation

	return nil, nil
}

func (m *sparseMatrix) validColMetrix(colMetrix []int) error {
	// num of rows of colMetrix = num of columns of sparseMatrix
	if len(colMetrix) == 0 {
		return errors.New("null column matrix vector")
	}
	if len(colMetrix) != m.colUpperBound {
		return fmt.Errorf("incompatible matrix and vector. vector row: %d, matrix columns: %d", len(colMetrix), m.colUpperBound)
	}
	return nil
}