package scrypt

import (
	"testing"
)

func config() {
	Default = New(
		"YE0dO4bwD4JnJafh6lZZfkp1MtKzuKAXQcDCJNJNyeCHairWHKENOkbh3dzwaCdizzOspwr/FITUVlnOAwPKyw==",
		"Bw==",
		8,
		14,
	)
}

func TestVerify(t *testing.T) {
	config()
	if !Verify(
		"8x4WjoDbSxJZdR",
		"xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A==",
		"sPtDhWcd1MfdAw==",
	) {
		t.Fail()
	}
}

func BenchmarkVerify(b *testing.B) {
	config()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Verify(
				"8x4WjoDbSxJZdR",
				"xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A==",
				"sPtDhWcd1MfdAw==",
			)
		}
	})
}

func TestEncode(t *testing.T) {
	config()
	_, err := Encode(
		"8x4WjoDbSxJZdR",
		"sPtDhWcd1MfdAw==",
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
			_, err := Encode(
				"8x4WjoDbSxJZdR",
				"xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A==",
			)
			if err != nil {
				b.Fail()
			}
		}
	})
}
