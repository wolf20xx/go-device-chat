package main

import (
	"crypto/md5"
	"errors"
	"io"
	"strings"
)

var ErrNoAvatarURL = errors.New("char: avatarを取得できません")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
type GravatarAvatar struct{}

type AuthAvatar struct{}

type FileSystemAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(useridStr))
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(useridStr))
			return "/avatars/" + useridStr + ".jpg", nil
		}
	}
	return "", ErrNoAvatarURL
}
