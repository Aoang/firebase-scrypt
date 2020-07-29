# Firebase scrypt

[![GoDoc](https://pkg.go.dev/badge/pkg.go.dev/github.com/Aoang/firebase-scrypt)](https://pkg.go.dev/pkg.go.dev/github.com/Aoang/firebase-scrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/Aoang/firebase-scrypt)](https://goreportcard.com/report/github.com/Aoang/firebase-scrypt)
[![Release](https://img.shields.io/github/v/release/Aoang/firebase-scrypt.svg)](https://github.com/Aoang/firebase-scrypt/releases)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

This is the golang implementation of the modified Scrypt algorithm used by Firebase Auth.

### Sample
```golang
package main

import scrypt "github.com/Aoang/firebase-scrypt"

func main(){
	scrypt.Default = scrypt.New(
		"YE0dO4bwD4JnJafh6lZZfkp1MtKzuKAXQcDCJNJNyeCHairWHKENOkbh3dzwaCdizzOspwr/FITUVlnOAwPKyw==",
		"Bw==",
		8,
		14,
	)
	scrypt.Verify(
		"8x4WjoDbSxJZdR",
		"xbSou7FOl6mChCyzpCPIQ7tku7nsQMTFtyOZSXXd7tjBa4NtimOx7v42Gv2SfzPQu1oxM2/k4SsbOu73wlKe1A==",
		"sPtDhWcd1MfdAw==",
	)
}
```

