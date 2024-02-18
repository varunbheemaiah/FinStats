package main

import (
	"compress/gzip"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

//go:embed images/*
var images embed.FS

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetNASDAQEarningsCalendar(dates []string) ([]byte, error) {

	url := "https://api.nasdaq.com/api/calendar/earnings"

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()

	for _, date := range dates {

		//* Not needed anymore as the dates will be received in the right format
		// Parse the input date string
		// inputFormat := "02-01-2006"
		// t, err := time.Parse(inputFormat, date)
		// if err != nil {
		// 			// 	continue
		// }

		// // Format the date to YYYY-MM-DD format
		// outputFormat := "2006-01-02"
		// requestDate := t.Format(outputFormat)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}

		q := req.URL.Query()
		q.Add("date", date)
		req.URL.RawQuery = q.Encode()

		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Origin", "https://www.nasdaq.com")
		req.Header.Set("Referer", "https://www.nasdaq.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		var body []byte
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err := gzip.NewReader(resp.Body)
			if err != nil {
				continue
			}
			defer reader.Close()
			body, err = io.ReadAll(reader)
			if err != nil {
				continue
			}
		default:
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
		}

		var result map[string]any
		err = json.Unmarshal(body, &result)
		if err != nil {
			continue
		}

		data, ok := result["data"].(map[string]any)
		if !ok {
			continue
		}

		rows, ok := data["rows"].([]any)
		if !ok {
			continue
		}

		headers := result["data"].(map[string]any)["headers"].(map[string]any)

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
		f.SetCellValue(date, "A1", "NASDAQ Earnings Calendar for "+date)

		if err := f.SetColWidth(date, "A", "A", 6); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "B", "B", 9); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "C", "C", 40); err != nil {
			continue
		}
		if err := f.SetColWidth(date, "D", "D", 15); err != nil {
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
			return nil, err
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
		f.SetCellValue(date, "D4", "Market Cap")

		currentRow := 5
		// timeMap := map[any]string{
		// 	"time-not-supplied": "Not Available",
		// 	"time-after-hours":  "After Hours",
		// 	"time-before-hours": "Before Hours",
		// 	"time-pre-market":   "Pre Market",
		// }
		enable := true
		for _, row := range rows {

			f.SetRowHeight(date, currentRow, 24)
			// f.SetCellValue(date, "A"+fmt.Sprintf("%d", currentRow), timeMap[row.(map[string]any)["time"]])

			image, err := images.ReadFile(fmt.Sprintf("images/%s.png", row.(map[string]any)["time"]))
			if err != nil {
				continue
			}
			if err := f.AddPictureFromBytes(date, "A"+fmt.Sprintf("%d", currentRow), &excelize.Picture{
				Extension: ".png",
				File:      image,
				Format: &excelize.GraphicOptions{
					ScaleX:          0.05,
					ScaleY:          0.05,
					LockAspectRatio: true,
					OffsetX:         8,
					OffsetY:         3,
				},
			}); err != nil {
				continue
			}

			// if err := f.AddPicture(date, "A"+fmt.Sprintf("%d", currentRow), imageName, &excelize.GraphicOptions{
			// 	ScaleX:          0.05,
			// 	ScaleY:          0.05,
			// 	LockAspectRatio: true,
			// 	OffsetX:         8,
			// 	OffsetY:         3,
			// }); err != nil {
			// 				// 	continue
			// }
			f.SetCellValue(date, "B"+fmt.Sprintf("%d", currentRow), row.(map[string]any)["symbol"])
			f.SetCellValue(date, "C"+fmt.Sprintf("%d", currentRow), row.(map[string]any)["name"])
			f.SetCellValue(date, "D"+fmt.Sprintf("%d", currentRow), row.(map[string]any)["marketCap"])
			f.SetCellStyle(date, "A"+fmt.Sprintf("%d", currentRow), "D"+fmt.Sprintf("%d", currentRow), bodyStyle)
			currentRow++

		}

		currentColumn := "E"
		for key, header := range headers {

			if key == "time" || key == "symbol" || key == "name" || key == "marketCap" {
				continue
			}

			currentRow := 4
			f.SetCellValue(date, currentColumn+fmt.Sprintf("%d", currentRow), header)
			f.SetCellStyle(date, currentColumn+fmt.Sprintf("%d", currentRow), currentColumn+fmt.Sprintf("%d", currentRow), headerStyle)

			for _, row := range rows {

				currentRow++
				f.SetCellValue(date, currentColumn+fmt.Sprintf("%d", currentRow), row.(map[string]any)[key])
				f.SetCellStyle(date, currentColumn+fmt.Sprintf("%d", currentRow), currentColumn+fmt.Sprintf("%d", currentRow), bodyStyle)

			}

			if err := f.SetColWidth(date, currentColumn, currentColumn, 15); err != nil {
				continue
			}
			currentColumn = string(rune(currentColumn[0]) + 1)

		}

		currentColumn = string(rune(currentColumn[0]) - 1)
		f.SetCellStyle(date, "A4", currentColumn+"4", headerStyle)

		err = f.AddTable(date, &excelize.Table{
			Range:             "A4:" + currentColumn + strconv.Itoa(currentRow-1),
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
	if err := f.SaveAs("earnings.xlsx"); err != nil {
		fmt.Println("Unable to save file:", err)
		return nil, err
	}

	buf, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buf.Bytes(), nil

}
