package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTCreationAndValidation(t *testing.T) {
	userID := uuid.New()
	secret := "test_secret"

	tests := []struct {
		name           string
		userID         uuid.UUID
		secret         string
		expiresIn      time.Duration
		validateSecret string
		wantErr        bool
	}{
		{
			name:           "valid token",
			userID:         userID,
			secret:         secret,
			expiresIn:      time.Hour,
			validateSecret: secret,
			wantErr:        false,
		},
		{
			name:           "expired token",
			userID:         userID,
			secret:         secret,
			expiresIn:      -time.Hour, // negative duration creates expired token
			validateSecret: secret,
			wantErr:        true,
		},
		{
			name:           "wrong secret",
			userID:         userID,
			secret:         secret,
			expiresIn:      time.Hour,
			validateSecret: "wrong_secret",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create token
			token, err := MakeJWT(tt.userID, tt.secret, tt.expiresIn)
			if err != nil {
				t.Fatalf("MakeJWT() error = %v", err)
			}

			// Validate token
			gotUserID, err := ValidateJWT(token, tt.validateSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && gotUserID != tt.userID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tt.userID)
			}
		})
	}
}
