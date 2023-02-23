package repositories

import "github.com/devNica/mochileros/entities"

type BackofficeRepo interface {
	InsertKYCReviewRequest(kycReview entities.KYCReviewRequest) error
}
