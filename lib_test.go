package owcrypt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_ecc_genpubb(t *testing.T) {

	//secp256k1
	pri, _ := hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	pub, err := genPublicKey(pri, "secp256k1")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

	//	secp256r1
	pri, _ = hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	pub, err = genPublicKey(pri, "secp256r1")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

	//sm2
	pri, _ = hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	pub, err = genPublicKey(pri, "sm2_std")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

}

func Test_ecc_sign(t *testing.T) {

	// secp256k1
	pri, _ := hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	pub, err := genPublicKey(pri, "secp256k1")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

	hash, _ := hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")

	sig,v, err := sign(pri, nil, hash, "secp256k1")
	fmt.Println(v)
	fmt.Println(hex.EncodeToString(sig))

	pass := verify(pub, nil, hash, sig, "secp256k1")
	fmt.Println(pass)

	// secp256r1
	pri, _ = hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	pub, err = genPublicKey(pri, "secp256r1")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

	hash, _ = hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")

	sig,v, err = sign(pri, nil, hash, "secp256r1")
	fmt.Println(v)
	fmt.Println(hex.EncodeToString(sig))

	pass = verify(pub, nil, hash, sig, "secp256r1")
	fmt.Println(pass)

	// sm2
	pri, _ = hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	id, _ := hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")
	pub, err = genPublicKey(pri, "sm2_std")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

	hash, _ = hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")

	sig,v, err = sign(pri, id, hash, "sm2_std")
	fmt.Println(v)
	fmt.Println(hex.EncodeToString(sig))

	pass = verify(pub, id, hash, sig, "sm2_std")
	fmt.Println(pass)
}

func Test_ecc_enc(t *testing.T) {
	pri, _ := hex.DecodeString("08bc185e33964486fdf5c862bd571a576483423ef00d2e9f41631a99493fd74b")

	pub, err := genPublicKey(pri, "sm2_std")
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(pub))

	plain, _ := hex.DecodeString("12345678")

	cipher, retCode := Encryption(pub, plain, ECC_CURVE_SM2_STANDARD)
	fmt.Println(retCode)
	fmt.Println(hex.EncodeToString(cipher))

	check, retCode := Decryption(pri, cipher, ECC_CURVE_SM2_STANDARD)
	fmt.Println(retCode)
	fmt.Println(hex.EncodeToString(check))
}
