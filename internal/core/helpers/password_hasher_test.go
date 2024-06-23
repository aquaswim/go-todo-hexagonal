package helpers

import "testing"

func TestHashPassword(t *testing.T) {
	_, err := HashPassword("password123!!")
	if err != nil {
		t.Errorf("Error hashing password: %s", err)
	}
}

func TestComparePasswordWithHash(t *testing.T) {
	t.Run("compare with correct hash", func(t *testing.T) {
		ok, err := ComparePasswordWithHash("password123!!", "$2a$12$Z0LgLWlyqGTTu0SyTv7cruS3f4R3wyMwh3dd1MTlF53cJoiJrodQW")
		if err != nil {
			t.Errorf("Error compare password: %s", err)
		}
		if !ok {
			t.Errorf("Password should be equal to hash")
		}
	})
	t.Run("compare with incorrect hash", func(t *testing.T) {
		ok, err := ComparePasswordWithHash("password123!!", "$2a$12$5IokLrp/pTZTyfPHxo5d0uGM9oukhc9VzX.xTIlC7SLHPPs0ZlcKO")
		if err != nil {
			t.Errorf("Error compare password: %s", err)
		}
		if ok {
			t.Errorf("Password should not be equal to hash")
		}
	})
	t.Run("compare with empty string", func(t *testing.T) {
		ok, err := ComparePasswordWithHash("password123!!", "")
		if err == nil {
			t.Errorf("compare password must return error object")
		}
		if ok {
			t.Errorf("Password should not be equal to hash")
		}
	})
}
