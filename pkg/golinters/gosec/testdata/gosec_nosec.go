//golangcitest:args -Egosec
//golangcitest:config_path testdata/gosec_nosec.yml
package testdata

import (
	"crypto/md5" // want "G501: Blocklisted import crypto/md5: weak cryptographic primitive"
	"log"
)

func Gosec() {
	// #nosec G401
	h := md5.New()
	log.Print(h)
}
