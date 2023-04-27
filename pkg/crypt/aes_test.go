/**
 * @Author Nil
 * @Description pkg/crypt/aes_test.go
 * @Date 2023/4/21 13:48
 **/

package crypt

import (
	"fmt"
	"github.com/ha5ky/hu5ky-bot/pkg/util"
	"testing"
)

func TestDecrypt(t *testing.T) {

	t.Run("Decrypt", func(t *testing.T) {
		fmt.Println(Decrypt("321b313c71d510aec548300437a28bf7c1f41fb736c7977d3626143a04ca322fad0ed10841b277617689", AESKey))
	})
}

func TestEncrypt(t *testing.T) {

	t.Run("Encrypt", func(t *testing.T) {
		uuid := util.GetUUID()
		fmt.Println(Encrypt("hu5ky."+uuid, AESKey))
	})
}
