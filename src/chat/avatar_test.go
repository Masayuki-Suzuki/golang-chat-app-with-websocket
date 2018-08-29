package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("If there aren't values, AuthAvatar.GetAvatarURL must return \"ErrorNoAvatarURL\".")
	}
	// Set Values
	testUrl := "http://url-to-avatar"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("If there are values, AuthAvatar.GetAvatarURL must not return error.")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURL must return correct URL.")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346"}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL must NOT return error.")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.GetAvatarURL returned value %s but it's wrong.", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	// Create an avatar file for testing.
	fileName := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(fileName, []byte{}, 0777)
	defer func() { os.Remove(fileName) }()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL must not return error.")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL returned a value: \"%s\"", url)
	}
}
