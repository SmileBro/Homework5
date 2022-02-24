package mycrypt

import (
	"fmt"
	"io/ioutil"
	"lesson6/signature"
	"lesson6/signature/contract"
	"os"
)

type Decryptor struct {
	fileHash   string
	hashString string
	fileSource string
	fileSigned string
	signature  contract.Signature
}

func NewDecryptor(fileHash string, fileSource string, fileSigned string) (dec *Decryptor, err error) {
	hashString, err := ioutil.ReadFile(fileHash)
	if err != nil {
		fmt.Println(err)
		return
	}
	dec = &Decryptor{fileHash: fileHash, hashString: string(hashString), fileSource: fileSource, fileSigned: fileSigned}
	return
}
func (dec *Decryptor) Validate() (err error) {
	file, err := os.Open(dec.fileSource)
	if err != nil {
		return
	}
	defer file.Close()

	//создаем signature из файла source
	sig, err := signature.NewSignatureSha256FromFileSource(file, dec.hashString)
	if err != nil {
		println(err)
		return err
	}
	var sign signature.SignatureSha256
	//создаем signature из файла sign
	data, err := ioutil.ReadFile(dec.fileSigned)
	if err != nil {
		return err
	}
	sign, err = signature.NewSignatureSha256FromString(string(data))
	if err != nil {
		return err
	}
	//проверяем их на равенство
	if sig.Equals(sign) {
		dec.signature = sign
		fmt.Println("декодирование успешно")
	} else {
		dec.signature = nil
		fmt.Println("декодирование не успешно")
	}
	return
}

func (dec Decryptor) SaveToFile(path string) (err error) {
	err = ioutil.WriteFile(path, dec.signature.SignatureByte(), 0644)
	return
}
