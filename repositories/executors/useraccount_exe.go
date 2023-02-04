package executors

import (
	"context"
	"errors"
	"fmt"
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

		profile := entities.UserHasProfile{
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

	result := repo.DB.Raw(`
		Select 
			ui.user_id, 
			ua.email, 
			ui.first_name, 
			ui.last_name, 
			ua.is_active,
			p.profile,
			ua.created_at

		from user_account ua 
		join user_info ui on ui.user_id = ua.id
		join user_has_profile uhp on uhp.user_id = ua.id
		join profile p on p.id = uhp.profile_id
		`).Scan(&model)

	if result.RowsAffected == 0 {
		return response.UserInfoResponseModel{}, errors.New("Complete User Info Not Found")
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
		Profile:   model.Profile,
		CreatedAt: model.CreatedAt,
	}

	return Account, nil

}

func (repo *userAccountExecutor) UpdateUserAccountStatus(ctx context.Context, userId string) (entities.UserAccount, error) {

	var User entities.UserAccount

	q1 := repo.DB.WithContext(ctx).Where("id = ?", userId).First(&User)

	if q1.RowsAffected == 0 {
		return entities.UserAccount{}, errors.New("User account not found")
	}

	q2 := repo.DB.Model(User).Where("id = ?", userId).Update("is_active", !User.IsActive)

	if q2.RowsAffected == 0 {
		return entities.UserAccount{}, errors.New("failed to update account")
	}

	return User, nil
}
