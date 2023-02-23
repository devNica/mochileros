package executors

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepoExecutor struct {
	*gorm.DB
}

func NewUserRepoExecutor(DB *gorm.DB) repositories.UserRepo {
	return &userRepoExecutor{DB: DB}
}

func (repo *userRepoExecutor) UserInsert(userAccount entities.UserAccount, profileId uint16) error {

	userAccount.Id = uuid.New()
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userAccount).Error; err != nil {
			tx.Rollback()
			return err
		}

		profile := entities.UserProfiles{
			UserId:    userAccount.Id,
			ProfileId: profileId,
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

func (repo *userRepoExecutor) InsertKYC(ctx context.Context, kyc entities.UserInfo) error {
	err := repo.DB.WithContext(ctx).Create(&kyc).Error
	exceptions.PanicLogging(err)
	return nil
}

func (repo *userRepoExecutor) FetchUserByEmail(email string) (response.UserInfoResponseModel, error) {

	type queryModel struct {
		UserId        string
		Email         string
		Password      string
		FirstName     string
		LastName      string
		TwoFactorAuth bool
		Profile       string
		CreatedAt     time.Time
	}

	var model queryModel

	result := repo.DB.Table("user_account").
		Select(`
		    user_account.id as user_id, 
			user_account.email,
			user_account.password, 
			user_info.first_name, 
			user_info.last_name, 
			user_account.two_factor_auth,
			string_agg(distinct profile.profile, ',') as profile,
			user_account.created_at
		`).
		Joins("inner join user_info on user_info.user_id = user_account.id").
		Joins("inner join user_profiles on user_profiles.user_id = user_account.id").
		Joins("inner join profile on profile.id = user_profiles.profile_id").
		Where("user_account.email = ?", email).Group(`
			user_account.id, 
			user_account.email,
			user_account.password,
			user_account.two_factor_auth,
			user_account.created_at, 
			user_info.first_name, 
			user_info.last_name`).
		Scan(&model)

	if result.RowsAffected == 0 {
		return response.UserInfoResponseModel{}, errors.New("User not found")
	}

	Account := response.UserInfoResponseModel{
		Id:            model.UserId,
		Email:         model.Email,
		Password:      model.Password,
		TwoFactorAuth: model.TwoFactorAuth,
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

func (repo *userRepoExecutor) CheckAccountExistByUserId(userId string) (entities.UserAccount, error) {

	var user = &entities.UserAccount{}

	repo.DB.First(&user, "id = ?", userId)

	if reflect.DeepEqual(user, entities.UserAccount{}) {
		return entities.UserAccount{}, errors.New("user account not found")
	} else {
		return *user, nil
	}
}

func (repo *userRepoExecutor) UpdateAccountVerification(userId uuid.UUID, request request.AccVerificationRequestModel, isFull bool) error {

	userAccount := entities.UserAccount{Id: userId}

	if isFull {

		q2 := repo.DB.Model(userAccount).Select("status_id", "two_factor_auth").Updates(entities.UserAccount{
			StatusId:      request.StatusId,
			TwoFactorAuth: request.TwoFactorAuth,
		})

		if err := repo.DB.Error; err != nil {
			fmt.Println("ocurri un error")
			exceptions.PanicLogging(err)
		}

		if q2.RowsAffected == 0 {
			return errors.New("failed to update account")
		}

		return nil

	} else {

		q2 := repo.DB.Model(userAccount).Select("status_id").Updates(entities.UserAccount{
			StatusId: request.StatusId,
		})

		if q2.RowsAffected == 0 {
			return errors.New("failed to update account")
		}

		return nil
	}

}
