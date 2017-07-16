package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
func TestGravatarAbatar(t *testing.T) {
	var gravataravatar GravatarAvatar
	client := new(client)
	//client.userData = map[string]interface{}{"email": "MyEmailAddress@example.com"}
	client.userData = map[string]interface{}{"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346"}
	url, err := gravataravatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値あり->gravatarAvatar.GetAvatarURLはエラーをかえすべきでない")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Error("gravatarAvatar.GetAvatarURLが%sという値を返した", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpn")
	ioutil.WriteFile(filename, []byte(), 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値あり->fileSystemAvatar.GetAvatarURLはエラーをかえすべきでない")
	}
	if url != "/avatars/abc.jpg" {
		t.Error("fileSystemAvatar.GetAvatarURLが%sという値を返した", url)
	}

}
