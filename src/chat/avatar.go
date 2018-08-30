package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatar is an error that is happened when Avatar instance can't return an avatar url.
var ErrNoAvatarURL = errors.New("chat: Error! Cannot get avatar URL")

// Avatar is a type that indicates user profile image.
type Avatar interface {
	// GetAvatarURL returns an avatar url of designated client.
	// If some problems are happened, it will return error.
	// Especially, when it cannot get url, it will return ErrNoAvatarURL.
	GetAvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar
type AuthAvatar struct{}
type GravatarAvatar struct{}
type FileSystemAvatar struct{}

var UseAuthAvatar AuthAvatar
var UseGravatar GravatarAvatar
var UseFileSystemAvatar FileSystemAvatar

func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}
