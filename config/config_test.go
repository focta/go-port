package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	wantPort := 3333
	// テスト時の設定で環境変数を設定したい場合の表現
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}

}
