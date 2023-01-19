package sheets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Service struct {
	*sheets.Service
}

// GetService is the wrapper function for sheets.NewService
// To use OAuth2 Client Authentication, give the client which was returned by oauth2 package.
func GetService(ctx context.Context, client *http.Client) (*Service, error) {
	var srv *sheets.Service
	var err error

	if client != nil {
		srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
	} else {
		srv, err = sheets.NewService(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %w", err)
	}

	return &Service{srv}, nil
}

type option_ struct {
	MajorDimension   string
	ValueInputOption string
}

type Option func(*option_)

func WithMajorDimension(md string) Option {
	return func(o *option_) {
		o.MajorDimension = md
	}
}

func WithValueInputOption(viopt string) Option {
	return func(o *option_) {
		o.ValueInputOption = viopt
	}
}

// Update updates given range (initial position?) in the specified sheet with values.
func (s Service) Update(spreadsheetId string, range_ string, values [][]interface{}, options ...Option) error {
	opt := &option_{
		MajorDimension:   "ROWS",
		ValueInputOption: "RAW",
	}

	for _, option := range options {
		option(opt)
	}

	vr := sheets.ValueRange{
		MajorDimension: string(opt.MajorDimension),
		Values:         values,
	}

	svuCall := s.Spreadsheets.Values.Update(spreadsheetId, range_, &vr)
	svuCall.ValueInputOption(string(opt.ValueInputOption))

	_, err := svuCall.Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet: %w", err)
	}

	return nil
}

// UpdateSheet updates the sheet by src object.
// src must be a Slice of the object that have `column` tag.
// see GetObj for value mapping spec.
func (s Service) UpdateSheet(spreadsheetId string, sheetTitle string, src interface{}, options ...Option) error {
	if t := reflect.TypeOf(src); t.Kind() != reflect.Slice || t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dst is not a Slice of Struct")
	}

	headerVR, err := s.Get(spreadsheetId, sheetTitle + "!1:1")
	if err != nil {
		return fmt.Errorf("unable to retrieve header row from sheet: %w", err)
	}

	slice := reflect.ValueOf(src)
	structType := slice.Type().Elem()

	header := headerVR.Values[0]

	fieldMap := make(map[int]int, len(header))

	for i, col := range header {
		colName := col.(string)

		for j := 0; j < structType.NumField(); j++ {
			tag := structType.Field(j).Tag.Get("column")
			if colName == tag {
				fieldMap[i] = j
				break
			}
		}
	}

	vals := make([][]interface{}, 0, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		val := make([]interface{}, len(header))
		for j := range val {
			if v, ok := fieldMap[j]; ok {
				switch f := slice.Index(i).Field(v); f.Kind() {
				case reflect.String:
					val[j] = f.String()
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					val[j] = f.Int()
				default:
					return fmt.Errorf("unhandled kind %s at index %d, field %d", f.Kind(), i, v)
				}
			}
		}

		vals = append(vals, val)
	}

	return s.Update(spreadsheetId, sheetTitle+"!A2", vals, options...)
}

// Get gets values of the range in specified sheet.
func (s Service) Get(spreadsheetId string, range_ string, options ...Option) (*sheets.ValueRange, error) {
	opt := &option_{
		MajorDimension: "ROWS",
	}

	for _, option := range options {
		option(opt)
	}

	svgCall := s.Spreadsheets.Values.Get(spreadsheetId, range_)
	svgCall.MajorDimension(opt.MajorDimension)

	return svgCall.Do()
}

// GetObj gets values of the range in specified sheet and store them into dst.
// dst must be a pointer to a slice of the Object that have `column` tag.
// The 1st row of the range is used to map columns and fields of dst Object.
//
// mapping example:
//
//   | A          | B   | C    | D       |
//   | ---------- | --- | ---- | ------- |
// 1 | imported   | id  | name | checked |
// 2 | 2021-08-15 | 123 | Foo  | 1       |
// 3 | 2021-08-16 | 456 | Bar  | 0       |
//
// type Row struct {
//     Imported string `column:"imported"`
//     Id       string `column:"id"`
//     Name     string `column:"name"`
//     Checked  string `column:"checked"`
// }
//
// var rows []Row
// GetObj(sheetId, "sheet!A1:D3", &row)
//
// row = [
//     { "2021-08-15", "123", "Foo", "1" },
//     { "2021-08-16", "456", "Bar", "0" },
// ]
func (s Service) GetObj(spreadsheetId string, range_ string, dst interface{}, options ...Option) error {
	if t := reflect.TypeOf(dst); t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Slice || t.Elem().Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dst is not a Ptr to Slice of Struct")
	}

	vr, err := s.Get(spreadsheetId, range_, options...)
	if err != nil {
		return err
	}

	pointer := reflect.ValueOf(dst)
	slice := pointer.Elem()
	structType := slice.Type().Elem()

	fieldMap := make(map[int]int, structType.NumField())

	header := vr.Values[0]

	for i := 0; i < structType.NumField(); i++ {
		tag := structType.Field(i).Tag.Get("column")

		for j, col := range header {
			colName := col.(string)
			if colName == tag {
				fieldMap[i] = j
				break
			}
		}
	}

	slice.Set(reflect.MakeSlice(slice.Type(), 0, len(vr.Values)))

	for _, val := range vr.Values[1:] {
		row := reflect.New(slice.Type().Elem()).Elem()
		for i := 0; i < row.NumField(); i++ {
			f := row.Field(i)
			if j, ok := fieldMap[i]; ok && j < len(val) {
				f.SetString(val[j].(string))
			}
		}

		r := reflect.Append(slice, row)
		slice.Set(r)
	}

	return nil
}

// DeleteRows delets rows from the sheet specified by `spreadsheetId` and `sheetTitle`.
// `startIndex` and `endIndex` are zero-based, and 0 means unbounded.
//
// see also:
// - https://developers.google.com/sheets/api/samples/rowcolumn#delete_rows_or_columns
// - https://developers.google.com/sheets/api/reference/rest/v4/DimensionRange
func (s Service) DeleteRows(spreadsheetId string, sheetTitle string, startIndex int64, endIndex int64) error {
	sheetId, err := s.GetSheetId(spreadsheetId, sheetTitle)
	if err != nil {
		return fmt.Errorf("(Service).GetSheetId: %w", err)
	}

	_, err = s.Spreadsheets.BatchUpdate(spreadsheetId, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteDimension: &sheets.DeleteDimensionRequest{
					Range: &sheets.DimensionRange{
						Dimension: "ROWS",
						SheetId: sheetId,
						StartIndex: startIndex,
						EndIndex: endIndex,
					},
				},
			},
		},
	}).Do()

	return err
}

var ErrGetSheetIdNotFound = fmt.Errorf("sheet not found")

// GetSheetId gets a sheet id by its title from the spreadsheet specified by `spreadsheetId`.
// You can find these IDs in a Google Sheets URL:
//     https://docs.google.com/spreadsheets/d/<spreadsheetId>/edit#gid=<sheetId>
//
// see also: https://developers.google.com/sheets/api/guides/concepts
func (s Service) GetSheetId(spreadsheetId string, title string) (int64, error) {
	spreadsheet, err := s.Spreadsheets.Get(spreadsheetId).Fields("sheets/properties(sheetId,title)").Do()
	if err != nil {
		return -1, err
	}

	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == title {
			return sheet.Properties.SheetId, nil
		}
	}

	return -1, ErrGetSheetIdNotFound
}
