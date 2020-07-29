# Firebase scrypt
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

