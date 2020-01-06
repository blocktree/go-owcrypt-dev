package owcrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/blocktree/owcrypt/blake256"
	"github.com/blocktree/owcrypt/blake2s"
	"github.com/blocktree/owcrypt/blake512"
	"github.com/blocktree/owcrypt/sha3"
	"github.com/blocktree/owcrypt/sm3"
	"golang.org/x/crypto/blake2b"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

func Hash(data []byte, digestLen uint16, typeChoose uint32) []byte {

	switch typeChoose {
	case HASH_ALG_MD4:
		h := md4.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_MD5:
		h := md5.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SHA1:
		h := sha1.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_RIPEMD160:
		h := ripemd160.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_HASH160:
		h := sha256.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		tmp := h.Sum(nil)
		h = ripemd160.New()
		_, err = h.Write(tmp)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_KECCAK256_RIPEMD160:
		h := sha3.NewKeccak256()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		tmp := h.Sum(nil)
		h = ripemd160.New()
		_, err = h.Write(tmp)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SHA3_256_RIPEMD160:
		h := sha3.New256()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		tmp := h.Sum(nil)
		h = ripemd160.New()
		_, err = h.Write(tmp)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SHA256:
		h := sha256.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_DOUBLE_SHA256:
		h := sha256.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		tmp := h.Sum(nil)
		h = sha256.New()
		_, err = h.Write(tmp)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SM3:
		h := sm3.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_KECCAK256:
		h := sha3.NewKeccak256()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SHA3_256:
		h := sha3.New256()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_BLAKE256:
		h := blake256.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SHA512:
		h := sha512.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_SHA3_512:
		h := sha3.New512()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_KECCAK512:
		h := sha3.NewKeccak512()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_BLAKE512:
		h := blake512.New()
		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_BLAKE2B:
		h, err := blake2b.New(int(digestLen), nil)
		if err != nil {
			return nil
		}
		_, err = h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	case HASH_ALG_BLAKE2S:
		h := blake2s.NewMAC(uint8(digestLen), nil)

		_, err := h.Write(data)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
		break

	default:
		return nil
	}

	return nil
}

func Hmac(key []byte, data []byte, typeChoose uint32) []byte {
	switch typeChoose {
	case HMAC_SHA256_ALG:
		h := hmac.New(sha256.New, key)
		h.Write(data)
		return h.Sum(nil)
		break
	case HMAC_SM3_ALG:
		h := hmac.New(sm3.New, key)
		h.Write(data)
		return h.Sum(nil)
		break
	case HMAC_SHA512_ALG:
		h := hmac.New(sha512.New, key)
		h.Write(data)
		return h.Sum(nil)
		break
	default:
		break
	}

	return nil
}