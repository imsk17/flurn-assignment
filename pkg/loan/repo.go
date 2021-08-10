package loan

import (
	"assignment/pkg/entities"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Create(request entities.LoanRequest) error
	Read(id uint) (entities.LoanRequest, error)
	ReadByStatus(status []entities.Status) ([]entities.LoanRequest, error)
	ReadAll() ([]entities.LoanRequest, error)
	ReadByLoanAmount(amount int32) ([]entities.LoanRequest, error)
	ReadByLoanAmountAndStatus(amount int32, status []entities.Status) ([]entities.LoanRequest, error)
	UpdateStatus(id uint, status entities.Status) error
	Delete(id uint) error
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) ReadAll() ([]entities.LoanRequest, error) {
	var reqs []entities.LoanRequest
	if err := r.DB.Find(&reqs).Error; err != nil {
		log.Print(err)
		return reqs, err
	}
	return reqs, nil
}

func (r *repo) Create(request entities.LoanRequest) error {
	err := r.DB.Create(&request).Error
	if err != nil {
		log.Print(err)
	}
	return nil
}

func (r *repo) Read(id uint) (entities.LoanRequest, error) {
	var req entities.LoanRequest
	if err := r.DB.First(&req, entities.LoanRequest{Model: gorm.Model{ID: id}}).Error; err != nil {
		log.Print(err)
		return req, err
	}
	return req, nil
}

func (r *repo) ReadByLoanAmountAndStatus(amount int32, status []entities.Status) ([]entities.LoanRequest, error) {
	var reqs []entities.LoanRequest
	if err := r.DB.Find(&reqs, "status IN ? AND loan_amount >= ?", status, amount).Error; err != nil {
		log.Print(err)
		return reqs, err
	}
	return reqs, nil
}

func (r *repo) ReadByStatus(status []entities.Status) ([]entities.LoanRequest, error) {
	var reqs []entities.LoanRequest
	if err := r.DB.Find(&reqs, map[string]interface{}{"status": status}).Error; err != nil {
		log.Print(err)
		return reqs, err
	}
	return reqs, nil
}

func (r *repo) ReadByLoanAmount(amount int32) ([]entities.LoanRequest, error) {
	var reqs []entities.LoanRequest
	if err := r.DB.Where("loan_amount >= ?", amount).Find(&reqs).Error; err != nil {
		log.Print(err)
		return reqs, err
	}
	return reqs, nil
}

func (r *repo) UpdateStatus(id uint, status entities.Status) error {
	var req entities.LoanRequest
	if err := r.DB.First(&req, id).Error; err != nil {
		return err
	}
	req.Status = status
	return r.DB.Save(&req).Error
}

func (r *repo) Delete(id uint) error {
	err := r.DB.Where("id = ?", id).Delete(&entities.LoanRequest{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}
