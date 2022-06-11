package cashier

import (
	models "cesgo/models/cashier"
	"cesgo/models/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func CreateJournal(c echo.Context) (err error) {

	journal := new(models.Journal)
	c.Bind(journal)
	response := new(response.General)
	if journal.Createjournal() != nil { // method create journal
		response.ErrorCode = 10
		response.Message = "Gagal create data Journal"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data Journal"
		response.Data = *journal
	}
	return c.JSON(http.StatusOK, response)
}

func GetJournal(c echo.Context) (err error) {
	response := new(response.General)
	journal, err := models.Getjournal() // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Failed"
	} else {
		response.Message = "Success"
		response.Data = journal
	}
	return c.JSON(http.StatusOK, response)
}

func SearchJournal(c echo.Context) (err error) {
	response := new(response.General)
	journal, err := models.GetAll(c.QueryParam("year")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Failed"
	} else {
		response.Message = "Success"
		response.Data = journal
	}
	return c.JSON(http.StatusOK, response)
}

func GetJournalReport(c echo.Context) (err error) {
	response := new(response.General)
	f := excelize.NewFile()
	sheet := "Laporan"
	index := f.NewSheet(sheet)
	path := "public/report/Report_" + c.QueryParam("year") + ".xlsx"
	journal, err := models.GetAll(c.QueryParam("year"))
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Failed"
	} else {
		//loop(journal)

		idx := 1
		f.SetCellValue(sheet, "A"+strconv.Itoa(idx), "Journal Date")
		f.SetCellValue(sheet, "B"+strconv.Itoa(idx), "Doc. No")
		f.SetCellValue(sheet, "C"+strconv.Itoa(idx), "Beginnning")
		f.SetCellValue(sheet, "D"+strconv.Itoa(idx), "Debit")
		f.SetCellValue(sheet, "E"+strconv.Itoa(idx), "Credit")
		f.SetCellValue(sheet, "F"+strconv.Itoa(idx), "Ending")

		idx = 2
		for _, jr := range journal {
			f.SetCellValue(sheet, "A"+strconv.Itoa(idx), jr.Journal_date)
			f.SetCellValue(sheet, "B"+strconv.Itoa(idx), jr.Voucher_no)
			f.SetCellValue(sheet, "C"+strconv.Itoa(idx), jr.Amount_beginning)
			f.SetCellValue(sheet, "D"+strconv.Itoa(idx), jr.Amount_debit)
			f.SetCellValue(sheet, "E"+strconv.Itoa(idx), jr.Amount_credit)
			f.SetCellValue(sheet, "F"+strconv.Itoa(idx), jr.Amount_ending)
			idx++
		}
		f.SetActiveSheet(index)
		if err := f.SaveAs(path); err != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.Message = "Success"
			response.Data = path
		}
	}
	return c.JSON(http.StatusOK, response)
}
