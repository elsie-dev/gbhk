package tokens

import (
	"github.com/google/uuid"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		userId        string
		secretKey     string
		expiryMinutes int
	}
	tests := []struct {
		name      string
		args      args
		wantEmpty bool
		wantErr   bool
	}{
		{
			name: "Generate token",
			args: args{
				userId:        uuid.New().String(),
				secretKey:     "SUPER_SECRET_KEY",
				expiryMinutes: 60 * 24 * 30 * 12 * 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.userId, tt.args.secretKey, tt.args.expiryMinutes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" != tt.wantEmpty {
				t.Errorf("GenerateToken() got = %v, want %v", got, err)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5OTE4Mjc4NDY4MjQ3NzEzNjgsInV1aWQiOiJlYTM2NDljMy04Y2VjLTQ2NjktYTliNy04NGEzMzUyZTg4ZjQifQ.l1f8c4ZS4ykFLVSMmXO5aEzYgGnoTbhAOPwNZK2S_tE"
	userId := "ea3649c3-8cec-4669-a9b7-84a3352e88f4"

	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAyMzkzOTczNTgzMjA2OTUsInV1aWQiOiI1Y2NmN2ExOC0yMjE4LTQxNTItOWRiYS04ODExZDZhNDU1ZTMifQ.BLIgdkZrZg2h7jnmyl3c--ff0PT24ndwVZMJx5-dikU"

	type args struct {
		jwtToken string
	}
	tests := []struct {
		name       string
		args       args
		wantUserId string
		wantErr    bool
	}{
		{
			name: "VerifyToken",
			args: args{
				jwtToken: jwtToken,
			},
			wantUserId: userId,
			wantErr:    false,
		}, {
			name: "VerifyExpired",
			args: args{
				jwtToken: expiredToken,
			},
			wantUserId: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserId, err := VerifyToken(tt.args.jwtToken, "SUPER_SECRET_KEY")
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserId != tt.wantUserId {
				t.Errorf("VerifyToken() gotUserId = %v, want %v", gotUserId, tt.wantUserId)
			}
		})
	}
}
