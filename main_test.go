package scrypt

import (
	"testing"
)

var app *App

func config() {
	var err error
	if app, err = New(
		"YE0dO4bwD4JnJafh6lZZfkp1MtKzuKAXQcDCJNJNyeCHairWHKENOkbh3dzwaCdizzOspwr/FITUVlnOAwPKyw==",
		"Bw==",
		8,
		14,
	); err != nil {
		panic(err)
	}
}

func TestVerify(t *testing.T) {
	config()
	if !app.FirebaseVerify(
		"8x4WjoDbSxJZdR",
		"sPtDhWcd1MfdAw==",
		"xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A==",
	) {
		t.Fail()
	}
}

func BenchmarkVerify(b *testing.B) {
	config()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			app.FirebaseVerify(
				"8x4WjoDbSxJZdR",
				"sPtDhWcd1MfdAw==",
				"xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A==",
			)
		}
	})
}

func TestEncode(t *testing.T) {
	config()
	_, err := app.Encode(
		[]byte("8x4WjoDbSxJZdR"),
		[]byte("sPtDhWcd1MfdAw=="),
	)
	if err != nil {
		t.Fail()
	}
}

func BenchmarkEncode(b *testing.B) {
	config()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := app.Encode(
				[]byte("8x4WjoDbSxJZdR"),
				[]byte("xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A=="),
			)
			if err != nil {
				b.Fail()
			}
		}
	})
}
