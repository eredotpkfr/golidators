package golidators

import "testing"

var (
	Md5Cases = [4]StrTestRecord{
		{"fffd51c41b9176ae3ed36a97699018eb", true},
		{"D41D8CD98F00B204E9800998ECF8427E", true},
		{"fffd51c41b9176ae3ed36a97699018eaa", false},
		{"D41D8CD98F00B204E9800998ECF8427EA", false},
	}
	Sha1Cases = [4]StrTestRecord{
		{"7cf42528c35d46dad692c67b985240f63d53f63e", true},
		{"930A0029225AA4C28B8EF095B679285EAAE27078", true},
		{"7cf42528c35d46dad692c67b985240f63d53f63eaasdasd", false},
		{"930A0029225AA4C28B8EF095B679285EAAE27078AKLMCDE", false},
	}
	Sha224Cases = [4]StrTestRecord{
		{"448b9c7aa6f6245ee1d77e7d140c73179f7f7a596dd386f62a0f6912", true},
		{"448B9C7AA6F6245EE1D77E7D140C73179F7F7A596DD386F62A0F6912", true},
		{"z14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", false},
		{"Z14A028C2A3A2BC9476102BB288234C415A2B01F828EA62AC5B3E42F", false},
	}
	Sha256Cases = [4]StrTestRecord{
		{"d8a928b2043db77e340b523547bf16cb4aa483f0645fe0a290ed1f20aab76257", true},
		{"D8A928B2043DB77E340B523547BF16CB4AA483F0645FE0A290ED1F20AAB76257", true},
		{"7815696ecbf1c96e6894b779456d330e", false},
		{"7815696ECBF1C96E6894B779456D330E", false},
	}
	Sha512Cases = [4]StrTestRecord{
		{"7621770eae0880e21dbf3939f045b9a44013f190df54243f0a4b3d28806d2c55c5de651d4fd4160e09ec3af805695275da1933044a70677a3efa361943644577", true},
		{"7621770EAE0880E21DBF3939F045B9A44013F190DF54243F0A4B3D28806D2C55C5DE651D4FD4160E09EC3AF805695275DA1933044A70677A3EFA361943644577", true},
		{"7815696ecbf1c96e6894b779456d330e", false},
		{"7815696ECBF1C96E6894B779456D330E", false},
	}
)

func TestMd5(t *testing.T) {
	for _, record := range Md5Cases {
		if Md5(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestSha1(t *testing.T) {
	for _, record := range Sha1Cases {
		if Sha1(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestSha224(t *testing.T) {
	for _, record := range Sha224Cases {
		if Sha224(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestSha256(t *testing.T) {
	for _, record := range Sha256Cases {
		if Sha256(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}

func TestSha512(t *testing.T) {
	for _, record := range Sha512Cases {
		if Sha512(record.TargetValue) != record.Expected {
			t.Error(record.TargetValue)
		}
	}
}
