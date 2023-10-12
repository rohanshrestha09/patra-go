package service

import (
	"strconv"

	"github.com/rohanshrestha09/patra-go/dto"
	"gorm.io/gorm"
)

type Store[T any] struct {
	DB *gorm.DB
}

func (s Store[T]) GetByID(paramID string, args dto.GetByIDArgs) (T, error) {

	var data T

	id, err := strconv.Atoi(paramID)

	if err != nil {
		return data, err
	}

	dbScopes := []func(*gorm.DB) *gorm.DB{
		Exclude(args.Exclude...),
	}

	for k, v := range args.Include {
		dbScopes = append(dbScopes, Include(k, v...))
	}

	if err := s.DB.Scopes(dbScopes...).Where(id).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil

}

func (s Store[T]) Get(args dto.GetArgs[T]) (T, error) {

	var data T

	dbScopes := []func(*gorm.DB) *gorm.DB{
		Exclude(args.Exclude...),
	}

	for k, v := range args.Include {
		dbScopes = append(dbScopes, Include(k, v...))
	}

	if err := s.DB.Scopes(dbScopes...).Where(&args.Filter).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil

}

func (s Store[T]) GetAll(query dto.Query, args dto.GetAllArgs[T]) ([]T, int64, error) {

	var (
		data  []T
		count int64
	)

	dbScopes := []func(*gorm.DB) *gorm.DB{
		Exclude(args.Exclude...),
		Sort(query.Sort, query.Order),
		Paginate(query.Page, query.Size),
	}

	for k, v := range args.Search {
		dbScopes = append(dbScopes, Search(k, v))
	}

	for k, v := range args.Include {
		dbScopes = append(dbScopes, Include(k.Value(), v...))
	}

	err := s.DB.
		Scopes(dbScopes...).
		Where(&args.Filter).
		Where(args.MapFilter).
		Find(&data).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	if err != nil {
		return data, count, err
	}

	return data, count, err

}

func (s Store[T]) Create(data T) error {

	if err := s.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil

}

func (s Store[T]) Update(model T, data T) error {

	if err := s.DB.Model(&model).Updates(&data).Error; err != nil {
		return err
	}

	return nil

}

func (s Store[T]) Delete(data T) error {

	if err := s.DB.Delete(&data).Error; err != nil {
		return err
	}

	return nil

}

func (s Store[T]) RecordExists(filter T) (bool, error) {

	var data T

	err := s.DB.Where(&filter).First(&data).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return false, nil

	default:
		return true, err
	}

}
