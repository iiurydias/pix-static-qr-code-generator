package pix_static_qr_code_generator

import (
	"bytes"
	"github.com/liyue201/goqr"
	. "github.com/onsi/gomega"
	"image"
	"io/ioutil"
	"testing"
)

func TestNewPixQRCode(t *testing.T) {
	RegisterTestingT(t)
	pixQRCodeGenerator := New("iury@email.com")
	t.Run("it generates a pix static qr code", func(t *testing.T) {
		pixQRCodeGenerator.SetAmount(15.00)
		pixQRCodeGenerator.SetId("123")
		pixQRCodeGenerator.SetDescription("a new payment")
		pixQRCodeGenerator.SetMerchantName("Iury")
		pixQRCodeGenerator.SetMerchantCity("Salvador")
		path := "./qr.png"
		err := pixQRCodeGenerator.Generate(path)
		Expect(err).ShouldNot(HaveOccurred())
		img := getImage(path)
		qrCode, err := goqr.Recognize(img)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(len(qrCode)).Should(BeEquivalentTo(1))
		Expect(qrCode[0].Payload).Should(BeEquivalentTo("00010256200341br.gov.bcb.pix1104iury@email.com1203a new payment025004500303689045515.005802BR5904Iury0068Salvador02650713063250397318"))
	})
}

func getImage(path string) image.Image {
	imgData, err := ioutil.ReadFile(path)
	Expect(err).ShouldNot(HaveOccurred())
	img, _, err := image.Decode(bytes.NewReader(imgData))
	Expect(err).ShouldNot(HaveOccurred())
	return img
}
