package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/velocity-ci/ssh-keyscan/pkg/sshkeyscan"
	"golang.org/x/crypto/ssh"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("missing args. Usage: ./ssh-keyscan <host>")
		os.Exit(1)
	}

	o := output{
		Host: os.Args[1],
	}

	o.Entry = sshkeyscan.GetSSHKeyForHost(o.Host)

	_, _, pubKey, _, _, _ := ssh.ParseKnownHosts([]byte(o.Entry))

	o.SHA256Fingerprint = ssh.FingerprintSHA256(pubKey)
	o.MD5Fingerprint = ssh.FingerprintLegacyMD5(pubKey)

	outStr, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outStr))
}

type output struct {
	Host              string `json:"host"`
	Entry             string `json:"entry"`
	SHA256Fingerprint string `json:"sha256Fingerprint"`
	MD5Fingerprint    string `json:"md5Fingerprint"`
}
