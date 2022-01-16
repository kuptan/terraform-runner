package internal

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const sshKeyPath string = "/root/.ssh/id_rsa"

func AddSSHKeyIfExist() {
	log.Info("adding ssh key if exist")

	if fileExists(sshKeyPath) {
		log.Infof("ssh key '%s' was detected, adding SSH key", sshKeyPath)

		shell(fmt.Sprintf("echo \"IdentityFile %s\" >> /etc/ssh/ssh_config", sshKeyPath))
		shell(fmt.Sprintf("eval `ssh-agent -s` && ssh-add %s", sshKeyPath))

		log.Info("ssh key was added")
	}
}
