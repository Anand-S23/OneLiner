package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/google/uuid"
)

func NewHashedUUID(hashItem string) string {
	hasher := md5.New()
	hasher.Write([]byte(hashItem))
	hash := hex.EncodeToString(hasher.Sum(nil))

	uuidFromHash := uuid.NewSHA1(uuid.Nil, []byte(hash))
	return uuidFromHash.String()
}

func NewUUID() string {
    return uuid.New().String()
}

