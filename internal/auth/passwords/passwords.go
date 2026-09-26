package passwords

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"unicode/utf8"

	"golang.org/x/crypto/argon2"
)

const (
	MaxLength = 128

	currentArgon2idVersion     = argon2.Version
	currentArgon2idMemoryKiB   = 19 * 1024
	currentArgon2idIterations  = 2
	currentArgon2idParallelism = 1
	currentArgon2idSaltLength  = 16
	currentArgon2idKeyLength   = 32
)

func Hash(password string) (string, error) {
	if utf8.RuneCountInString(password) > MaxLength {
		return "", fmt.Errorf("password must not exceed %d characters", MaxLength)
	}

	salt := make([]byte, currentArgon2idSaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}

	derivedKey := argon2.IDKey(
		[]byte(password),
		salt,
		currentArgon2idIterations,
		currentArgon2idMemoryKiB,
		currentArgon2idParallelism,
		currentArgon2idKeyLength,
	)

	return encodeArgon2idHash(argon2idHash{
		version:     currentArgon2idVersion,
		memoryKiB:   currentArgon2idMemoryKiB,
		iterations:  currentArgon2idIterations,
		parallelism: currentArgon2idParallelism,
		salt:        salt,
		derivedKey:  derivedKey,
	}), nil
}

func Verify(password, encodedHash string) bool {
	if utf8.RuneCountInString(password) > MaxLength {
		return false
	}

	// Verify legacy SHA-256 hashes.
	if expectedHash, ok := decodeLegacyHash(encodedHash); ok {
		candidateHash := sha256.Sum256([]byte(password))
		return subtle.ConstantTimeCompare(candidateHash[:], expectedHash) == 1
	}

	// Verify Argon2id hashes.
	storedHash, ok := parseArgon2idHash(encodedHash)
	if !ok {
		return false
	}

	if storedHash.version != currentArgon2idVersion {
		return false
	}

	candidateHash := argon2.IDKey(
		[]byte(password),
		storedHash.salt,
		storedHash.iterations,
		storedHash.memoryKiB,
		storedHash.parallelism,
		uint32(len(storedHash.derivedKey)),
	)

	return subtle.ConstantTimeCompare(candidateHash, storedHash.derivedKey) == 1
}

func NeedsRehash(encodedHash string) bool {
	// Every legacy SHA-256 password hash should be upgraded.
	if _, ok := decodeLegacyHash(encodedHash); ok {
		return true
	}

	// Malformed hashes are not candidates for rehashing.
	storedHash, ok := parseArgon2idHash(encodedHash)
	if !ok {
		return false
	}

	// Upgrade valid Argon2id hashes that don't match the current policy.
	return storedHash.version != currentArgon2idVersion ||
		storedHash.memoryKiB != currentArgon2idMemoryKiB ||
		storedHash.iterations != currentArgon2idIterations ||
		storedHash.parallelism != currentArgon2idParallelism ||
		len(storedHash.derivedKey) != currentArgon2idKeyLength
}
