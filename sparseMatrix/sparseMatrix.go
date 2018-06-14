package sparseMatrix

import (
	"fmt"
	"math"
	"sort"
)

const (
	MaximumRows = math.MaxUint32
	MaximumCols = math.MaxUint32
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
}

type SparseMatrix interface {
	Add(number Data) error
	Remove(number Data) error
	GetData(row, col int) (*Data, error)
	GetRowSortByCol(row int) (RowList, error)
	PrintRow(row int) string
	Print() string
	isValidRow(row int) error
	isValidColumn(col int) error
}

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
	// the columns of the number are not sorted.
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

func (m *sparseMatrix) isValidRow(row int) error {
	if row < 0 || row >= m.rowUpperBound {
		return fmt.Errorf("invalid row: %d", row)
	}
	return nil
}

func (m *sparseMatrix) isValidColumn(col int) error {
	if col < 0 || col >= m.colUpperBound {
		return fmt.Errorf("invalid col: %d", col)
	}
	return nil
}

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
