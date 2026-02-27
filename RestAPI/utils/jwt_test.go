package utils

import "testing"

//const secret = "secret"

func TestGenerateToken(t *testing.T) {
	type args struct {
		email  string
		userId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "get token", args: args{email: "test@test.com", userId: "some-user-id-1"}, wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateToken(tt.args.email, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "unauthorized", args: args{tokenString: "some-string"}, wantErr: true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyToken(tt.args.tokenString); (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
