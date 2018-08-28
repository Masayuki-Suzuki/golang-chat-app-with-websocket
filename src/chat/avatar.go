package main

import "errors"

// ErrNoAvatar is an error that is happened when Avatar instance can't return an avatar url.
var ErrNoAvatarURL = errors.New("chat: Error! Cannot get avatar URL")

// Avatar is a type that indicates user profile image.
type Avatar interface {
	// GetAvatarURL returns an avatar url of designated client.
	// If some problems are happened, it will return error.
	// Especially, when it cannot get url, it will return ErrNoAvatarURL.
	GetAvatarURL(c *client) (string, error)
}
