package main

import (
	"errors"
)

// ErrNoAvatar is an error that is happened when Avatar instance can't return an avatar url.
var ErrNoAvatarURL = errors.New("chat: Error! Cannot get avatar URL")

// Avatar is a type that indicates user profile image.
type Avatar interface {
	// GetAvatarURL returns an avatar url of designated client.
	// If some problems are happened, it will return error.
	// Especially, when it cannot get url, it will return ErrNoAvatarURL.
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}
type GravatarAvatar struct{}
type FileSystemAvatar struct{}

var UseAuthAvatar AuthAvatar
var UseGravatar GravatarAvatar
var UseFileSystemAvatar FileSystemAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

func (_ FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "/avatars/" + useridStr + ".jpg", nil
		}
	}
	return "", ErrNoAvatarURL
}
