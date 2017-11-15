package utrl

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"net"
	"bytes"
)

type SshCMD struct {
	rsa ssh.Signer
}


func (k *SshCMD) LoadPEM(file string) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return err
	}
	k.rsa = key
	return nil
}


//e.g. output, err := remoteRun("root", "MY_IP", "password", "ls")
func (k *SshCMD)RemoteRun(user string, addr string, cmd string) (string, error) {
	// privateKey could be read from a file, or retrieved from another storage
	// source, such as the Secret Service / GNOME Keyring
	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(k.rsa),
			//ssh.Password("jack2017"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		//alternatively, you could use a password
		/*
			Auth: []ssh.AuthMethod{
				ssh.Password("PASSWORD"),
			},
		*/
	}
	// Connect
	client, err := ssh.Dial("tcp", addr+":22222", config)
	if err != nil {
		return "", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")
	session.Stderr = os.Stderr
	// Finally, run the command
	err = session.Run(cmd)
	return b.String(), err
}

