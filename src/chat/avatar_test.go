package main

import "testing"

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
