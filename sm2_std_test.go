package owcrypt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_sm2_enc(t *testing.T) {

		pri, _ := hex.DecodeString("a0dd4b86ab6fd4e9c6b49bb50861493b0324589800a71531988536eaa6c2bf1f ")

		pub, err := genPublicKey(pri, "sm2_std")
		fmt.Println(err)
		fmt.Println(hex.EncodeToString(pub))

		plain, _ := hex.DecodeString("ee4b69642b40d63d6c8459e623dca6d063c85db1f914662d58b596fb5d782f8d2ff6b40e8b454fe55078501e27e7d8be28c73d86f383808c71d92195f2db6bec772ded571db8aae216c64522718860dfd1a9a2d9adaf4e63fc1064f0e70f6cbfd9c6e009d8cb54f353b306ca41be3007cb2c1bd063bba2439df6cb9662e2f3d5391264da140b49dc33007e1b8f46713b44514fca9a92b6b7b69cb2d2bb8f7ab0846e1cf355ee73426ba831354ea945fb45fcb3c867faeda3d29967013b07ea3b26718b8318301419df4a1fb3eb2113bdeb45aef2aa54e1661777850f1cb3e84e53ae46cda58faf71a6bc408c6d5711bd2b00f69d9132d301556a0eae565b5d211807b3e0ec53b58c36b10aaf763257ce82b1d284f0da11b781b16a98d9e4685d4c6fdd2b560dc5bbfbf5cd1fd4253d26f40c2d7912c653f83c9bd4a1cb8a08fbdbb909cae84861e23b59755e02fcbe829be574dff2933e6bc4e16533fc2207ac24dd21dfec1ef6e19a521611d238ea3f6bec850e660ec44284c6fd35c4a539512ffcbd424abe5ff4fad346ac32f414dd0f7ed6e3de767521eb1716beb40a9762c952fff41f07610cfff2642aa2adf59eacfdf8aa1d57a0ef2f9a6fa855acee98a22fff2eb9ed79edfe27d75fc84aaeaf65d91e880f1ef5c1bf60052e183768e60e2bea528abe564a58453c6814def173ab074971447001392e3d4398dc814aa3963a1a224fc1534c22f6cab73032368d7317be2709a61a3961bb99aaee92acddf918cdf81f1d42e314a44637ec7012bea12bf54b1cbd885dc510293d09d466c2782f2e4a355600a44fecab44c66482720202fef9bdc9a6e944a04b7cd3480bbc0c5c88f83596367d201d677900f6fd87a615cc1e527bc5bd5831c61901601501e5a58b72d4fdda9802e51a347a5634af81ada587c165e2af3200a42fc5b39c29dcdd7dc3d7faa394ec56d2fb111ac5ed83b78153372122044173377ee8667a812467f2bcfa3951d438bde63549dabde3b60ec9125a2859f7aa4667d91449d4e4376e97fd78f9afb8992373aec4db60aba6350ab139378a3a483db9445be6f8fcca822fb4b5756c0a56880df9e766e02d06b7c6be9844cd8b730442260dc0b38dcdf7c3b899d76fcefc5fdd64175ecbcdfa630af691237a0a228ab37541a14bd099477e484e31c6584911a1b76cc23873550291ef34500b35a35f1d4e322affe5f344dd5619273a199807de975e34f8affa8e9734ffe3b3a389472e95cc663931a54d7f88e663d3f2cec9944f9c99b0bebcdc8a57e77cbe03b78b95fc437b8f39ab89013ba76d398be6a81b32d1fed3231b99821227bdc16f94ddf33364f5fce6581a13af52ffde3ac0f46712899620c04ed750d28173c7dee5a3a7275c49e57a0a2a094773871e8691d805a525e197bac44c4e04014c7aab87d020eb6b604576")

		cipher, retCode := Encryption(pub, plain, ECC_CURVE_SM2_STANDARD)
		fmt.Println(retCode)
		fmt.Println(hex.EncodeToString(cipher))

		check, retCode := Decryption(pri, cipher, ECC_CURVE_SM2_STANDARD)
		fmt.Println(retCode)
		fmt.Println(hex.EncodeToString(check))

		if hex.EncodeToString(check) != hex.EncodeToString(plain) {
			t.Error("failed----------------")
			return
		}


}