package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type NSEBSEEarning struct {
	Data struct {
		List []NSEBSEEarningData `json:"list"`
	} `json:"data"`
}

type NSEBSEEarningData struct {
	Date           string  `json:"date"`
	StockName      string  `json:"stockName"`
	ScID           string  `json:"scId"`
	StockURL       string  `json:"stockUrl"`
	ResultType     string  `json:"resultType"`
	Ltp            string  `json:"ltp"`
	Change         string  `json:"change"`
	Time           string  `json:"time"`
	SeeFinancial   string  `json:"seeFinancial"`
	StockShortName string  `json:"stockShortName"`
	MarketCap      float64 `json:"marketCap"`
	Exchange       string  `json:"exchange"`
}

func GetNSEBSEEarningsDataForDate(date string, page int) ([]NSEBSEEarningData, error) {

	url := "https://api.moneycontrol.com/mcapi/v1/earnings/get-earnings-data"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("indexId", "All")
	q.Add("limit", "100")
	q.Add("startDate", date)
	q.Add("endDate", date)
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return []NSEBSEEarningData{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	earnings := NSEBSEEarning{}
	if err := json.Unmarshal(body, &earnings); err != nil {
		return nil, err
	}

	return earnings.Data.List, nil

}

func (a *App) GetNSEBSEEarningsCalendar(dates []string) (string, error) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()

	enable := true
	for _, date := range dates {

		rows := []NSEBSEEarningData{}
		for i := 1; ; i++ {
			part, err := GetNSEBSEEarningsDataForDate(date, i)
			if err != nil {
				return "", err
			}
			if len(part) == 0 {
				break
			}
			rows = append(rows, part...)
		}

		index, err := f.NewSheet(date)
		if err != nil {
			continue
		}

		f.SetActiveSheet(index)

		err = f.DeleteSheet("Sheet1")
		if err != nil {
			continue
		}

		f.MergeCell(date, "A1", "I2")
		f.SetCellValue(date, "A1", "NSE/BSE Earnings Calendar for "+date)

		if err := f.SetColWidth(date, "A", "A", 14); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "B", "B", 12); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "C", "C", 40); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "D", "D", 15); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "I", "I", 17); err != nil {
			continue
		}

		titleStyle, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Font: &excelize.Font{
				Size: 24,
			},
			Border: []excelize.Border{
				{
					Type:  "left",
					Color: "000000",
					Style: 2,
				},
				{
					Type:  "top",
					Color: "000000",
					Style: 2,
				},
				{
					Type:  "bottom",
					Color: "000000",
					Style: 2,
				},
				{
					Type:  "right",
					Color: "000000",
					Style: 2,
				},
			},
		})
		if err != nil {
			return "", err
		}

		f.SetCellStyle(date, "A1", "I2", titleStyle)
		f.SetRowHeight(date, 4, 40)

		headerStyle, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
			Border: []excelize.Border{
				{
					Type:  "left",
					Color: "000000",
					Style: 1,
				},
				{
					Type:  "top",
					Color: "000000",
					Style: 1,
				},
				{
					Type:  "bottom",
					Color: "000000",
					Style: 1,
				},
				{
					Type:  "right",
					Color: "000000",
					Style: 1,
				},
			},
		})
		if err != nil {
			continue
		}

		bodyStyle, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
		})
		if err != nil {
			continue
		}

		f.SetCellValue(date, "A4", "Time")
		f.SetCellValue(date, "B4", "Symbol")
		f.SetCellValue(date, "C4", "Company Name")
		f.SetCellValue(date, "D4", "Market Cap (Cr)")
		f.SetCellValue(date, "E4", "Exchange")
		f.SetCellValue(date, "F4", "LTP")
		f.SetCellValue(date, "G4", "% Change")
		f.SetCellValue(date, "H4", "Result Type")
		f.SetCellValue(date, "I4", "Stock Info")
		f.SetCellStyle(date, "A4", "I4", headerStyle)

		currentRow := 5
		timeMap := map[any]string{
			"Time Not Available": "NA",
			"After Market":       "After Market",
			"During Market":      "During Market",
		}

		exchangeMap := map[any]string{
			"N": "NSE",
			"B": "BSE",
		}

		redStyle, err := f.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Color: "FF0000",
			},
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
		})
		if err != nil {
			continue
		}

		greenStyle, err := f.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Color: "008000",
			},
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
		})
		if err != nil {
			continue
		}

		hyperlinkStyle, err := f.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Underline: "single",
				Color:     "0000FF",
			},
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
		})
		if err != nil {
			continue
		}

		for _, row := range rows {

			f.SetRowHeight(date, currentRow, 24)
			f.SetCellValue(date, "A"+fmt.Sprintf("%d", currentRow), timeMap[row.Time])
			f.SetCellValue(date, "B"+fmt.Sprintf("%d", currentRow), row.ScID)
			f.SetCellValue(date, "C"+fmt.Sprintf("%d", currentRow), row.StockName)
			f.SetCellValue(date, "D"+fmt.Sprintf("%d", currentRow), row.MarketCap)
			f.SetCellValue(date, "E"+fmt.Sprintf("%d", currentRow), exchangeMap[row.Exchange])
			f.SetCellValue(date, "F"+fmt.Sprintf("%d", currentRow), row.Ltp)
			change, err := strconv.ParseFloat(row.Change, 64)
			if err != nil {
				fmt.Println("Unable to convert change to int:", err.Error())
			}
			f.SetCellValue(date, "G"+fmt.Sprintf("%d", currentRow), fmt.Sprintf("%s%%", row.Change))
			f.SetCellValue(date, "H"+fmt.Sprintf("%d", currentRow), row.ResultType)
			f.SetCellValue(date, "I"+fmt.Sprintf("%d", currentRow), row.StockShortName)
			f.SetCellHyperLink(date, "I"+fmt.Sprintf("%d", currentRow), row.StockURL, "External")
			f.SetCellStyle(date, "A"+fmt.Sprintf("%d", currentRow), "I"+fmt.Sprintf("%d", currentRow), bodyStyle)
			f.SetCellStyle(date, "I"+fmt.Sprintf("%d", currentRow), "I"+fmt.Sprintf("%d", currentRow), hyperlinkStyle)
			if change >= 0 {
				f.SetCellStyle(date, "G"+fmt.Sprintf("%d", currentRow), "G"+fmt.Sprintf("%d", currentRow), greenStyle)
			} else {
				f.SetCellStyle(date, "G"+fmt.Sprintf("%d", currentRow), "G"+fmt.Sprintf("%d", currentRow), redStyle)
			}
			currentRow++
		}

		err = f.AddTable(date, &excelize.Table{
			Range:             "A4:" + "I" + strconv.Itoa(currentRow-1),
			StyleName:         "TableStyleLight6",
			ShowFirstColumn:   true,
			ShowLastColumn:    true,
			ShowRowStripes:    &enable,
			ShowColumnStripes: false,
		})
		if err != nil {
			continue
		}

	}

	f.SetActiveSheet(0)

	buf, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	b64str := base64.StdEncoding.EncodeToString(buf.Bytes())

	return b64str, nil

}
