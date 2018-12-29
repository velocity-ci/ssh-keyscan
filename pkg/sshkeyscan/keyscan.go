package sshkeyscan

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func GetSSHKeyForHost(hostname string) string {
	var keyScanOutput string

	ssh.Dial("tcp", fmt.Sprintf("%s:22", hostname), &ssh.ClientConfig{
		User: "git",
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			keyScanOutput = fmt.Sprintf("%s %s", hostname[:len(hostname)-3], ssh.MarshalAuthorizedKey(key))
			return nil
		},
	})

	return keyScanOutput
}
