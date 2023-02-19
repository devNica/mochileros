package executors

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userAccountExecutor struct {
	*gorm.DB
}

func NewUserAccountExecutor(DB *gorm.DB) repositories.UserAccountRepo {
	return &userAccountExecutor{DB: DB}
}

func (repo *userAccountExecutor) UserInsert(ctx context.Context, userAccount entities.UserAccount) error {
	userAccount.Id = uuid.New()
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userAccount).Error; err != nil {
			tx.Rollback()
			return err
		}

		profile := entities.UserProfiles{
			UserId:    userAccount.Id,
			ProfileId: 2,
		}

		if err := tx.Create(&profile).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo *userAccountExecutor) InsertKYC(ctx context.Context, kyc entities.UserInfo) error {
	kyc.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&kyc).Error
	exceptions.PanicLogging(err)
	return nil
}

func (repo *userAccountExecutor) FetchUserByEmail(ctx context.Context, email string) (response.UserResponseModel, error) {

	var foundUser entities.UserAccount

	result := repo.DB.WithContext(ctx).Where("email = ?", email).First(&foundUser)
	if result.RowsAffected == 0 {
		return response.UserResponseModel{}, errors.New("user Not Found")
	}

	fmt.Println("user account", foundUser)

	account := response.UserResponseModel{
		Id:        foundUser.Id,
		Email:     foundUser.Email,
		IsActive:  foundUser.IsActive,
		CreatedAt: foundUser.CreatedAt,
	}

	return account, nil
}

func (repo *userAccountExecutor) FetchCompleteUserInfo(ctx context.Context, userId string) (response.UserInfoResponseModel, error) {

	type queryModel struct {
		UserId    string
		Email     string
		FirstName string
		LastName  string
		IsActive  bool
		Profile   string
		CreatedAt time.Time
	}

	var model queryModel

	result := repo.DB.Table("user_account").
		Select(`
		    user_account.id as user_id, 
			user_account.email, 
			user_info.first_name, 
			user_info.last_name, 
			user_account.is_active,
			string_agg(distinct profile.profile, ',') as profile,
			user_account.created_at
		`).
		Joins("inner join user_info on user_info.user_id = user_account.id").
		Joins("inner join user_profiles on user_profiles.user_id = user_account.id").
		Joins("inner join profile on profile.id = user_profiles.profile_id").
		Where("user_account.id = ?", userId).Group("user_account.id, user_info.first_name, user_info.last_name").
		Scan(&model)

	if result.RowsAffected == 0 {
		return response.UserInfoResponseModel{}, errors.New("User not found")
	}

	Account := response.UserInfoResponseModel{
		Id:       model.UserId,
		Email:    model.Email,
		IsActive: model.IsActive,
		UserInfo: struct {
			FirstName string
			LastName  string
		}{
			FirstName: model.FirstName,
			LastName:  model.LastName,
		},
		Profile:   strings.Split(model.Profile, ","),
		CreatedAt: model.CreatedAt,
	}

	return Account, nil

}

func (repo *userAccountExecutor) UpdateUserAccountStatus(ctx context.Context, userId string) (response.UserResponseModel, error) {

	var foundUser entities.UserAccount

	q1 := repo.DB.WithContext(ctx).Where("id = ?", userId).First(&foundUser)

	if q1.RowsAffected == 0 {
		return response.UserResponseModel{}, errors.New("User account not found")
	}

	q2 := repo.DB.Model(foundUser).Where("id = ?", userId).Update("is_active", !foundUser.IsActive)

	if q2.RowsAffected == 0 {
		return response.UserResponseModel{}, errors.New("failed to update account")
	}

	User := response.UserResponseModel{
		Id:        foundUser.Id,
		Email:     foundUser.Email,
		IsActive:  foundUser.IsActive,
		CreatedAt: foundUser.CreatedAt,
	}

	return User, nil
}
