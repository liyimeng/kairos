package mos_test

import (
	"os"
	"time"

	"github.com/c3os-io/c3os/tests/machine"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("c3os autoinstall test", Label("autoinstall-test"), func() {
	BeforeEach(func() {
		machine.EventuallyConnects()
	})

	AfterEach(func() {
		if CurrentGinkgoTestDescription().Failed {
			gatherLogs()
		}
	})

	Context("live cd", func() {
		It("has default service active", func() {
			if os.Getenv("FLAVOR") == "alpine" {
				out, _ := machine.SSHCommand("sudo rc-status")
				Expect(out).Should(ContainSubstring("c3os"))
				Expect(out).Should(ContainSubstring("c3os-agent"))
			} else {
				// Eventually(func() string {
				// 	out, _ := machine.SSHCommand("sudo systemctl status c3os-agent")
				// 	return out
				// }, 30*time.Second, 10*time.Second).Should(ContainSubstring("no network token"))

				out, _ := machine.SSHCommand("sudo systemctl status c3os")
				Expect(out).Should(ContainSubstring("loaded (/etc/systemd/system/c3os.service; enabled; vendor preset: disabled)"))
			}

			out, _ := machine.SSHCommand("ls -liah /oem")
			Expect(out).To(ContainSubstring("userdata.yaml"))
		})
	})

	Context("auto installs", func() {
		It("to disk with custom config", func() {
			Eventually(func() string {
				out, _ := machine.SSHCommand("sudo ps aux")
				return out
			}, 30*time.Minute, 1*time.Second).Should(
				Or(
					ContainSubstring("elemental install"),
				))
		})
	})

})
