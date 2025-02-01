package domain

import (
	"errors"
)

var (
	ErrInternal = errors.New("internal error")

	ErrDataNotFound = errors.New("data not found")

	ErrNoUpdatedData = errors.New("no data to update")

	ErrConfig = errors.New("unable to load configuration file")

	ErrConflictingData = errors.New("data conflicts with existing data in unique column")

	ErrInsufficientStock = errors.New("product stock is not enough")

	ErrInsufficientPayment = errors.New("total paid is less than total price")

	ErrTokenDuration = errors.New("invalid token duration format")

	ErrTokenCreation = errors.New("error creating token")

	ErrExpiredToken = errors.New("access token has expired")

	ErrInvalidToken = errors.New("access token is invalid")

	ErrInvalidCredentials = errors.New("invalid email or password")

	ErrEmptyAuthorizationHeader = errors.New("authorization header is missing")

	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")

	ErrInvalidAuthorizationType = errors.New("authorization type is not supported")

	ErrUnauthorized = errors.New("user is not authorized to access the resources")

	ErrForbidden = errors.New("user is forbidden to access the resources")

	ErrMaintenance = errors.New("this website is currently undergoing maintenance. Please try again later")

	ErrIdNotFound = errors.New("id not found")
)
