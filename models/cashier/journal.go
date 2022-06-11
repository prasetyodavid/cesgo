package models

import (
	"cesgo/config"
)

type Journal struct {
	//Journal_id       string `json:"journal_id" form:"journal_id" gorm:"primaryKey"`
	Journal_date     string `json:"journal_date" form:"journal_date"`
	Voucher_no       string `json:"voucher_no" form:"voucher_no"`
	Amount_beginning string `json:"amount_beginning" form:"amount_beginning"`
	Amount_debit     string `json:"amount_debit" form:"amount_debit"`
	Amount_credit    string `json:"amount_credit" form:"amount_credit"`
	Amount_ending    string `json:"amount_ending" form:"amount_ending"`
	Description      string `json:"description" form:"description"`
	Created_by       string `json:"created_by" form:"created_by"`
	Created_at       string `json:"created_at" form:"created_at"`
}

func (journal *Journal) Createjournal() error {
	if err := config.DB.Create(journal).Error; err != nil {
		return err
	}
	return nil
}

func (journal *Journal) Updatejournal(journal_id string) error {
	if err := config.DB.Model(&Journal{}).Where("journal_id = ?", journal_id).Updates(journal).Error; err != nil {
		return err
	}
	return nil
}

func (journal *Journal) Deletejournal() error {
	if err := config.DB.Delete(journal).Error; err != nil {
		return err
	}
	return nil
}

func Getjournal() ([]Journal, error) {
	var Journal []Journal

	result := config.DB.Find(&Journal)

	return Journal, result.Error
}

func GetAll(year string) ([]Journal, error) {
	var Journal []Journal

	result := config.DB.Where(" year( journal_date ) = ?", year).Find(&Journal)

	return Journal, result.Error
}
