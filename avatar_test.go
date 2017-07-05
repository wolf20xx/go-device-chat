package main

import (
	"log"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値なし->AuthAvatar.GetAvatarURLはErrNoAvatarURLを返すべき")
	}
	testUrl := "htp://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値あり->AuthAvatar.GetAvatarURLをかえすべきでない")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURLは正しいURLをかえすべき")
		}
		log.Println("TestOK")
	}

}
