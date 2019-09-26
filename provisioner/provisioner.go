package provisioner

import (
	provisionerv1 "github.com/kevin-zhaoshuai/k3s-operator/api/v1"
	"io"
	"os"
	"os/exec"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

func ProvisionEdgeNode(edgeNode provisionerv1.K3s) error {
	IP := edgeNode.Spec.IP
	nodeType := edgeNode.Spec.Type
	_ = edgeNode.Spec.SkipInstall
	sshPort := edgeNode.Spec.SshPort
	user := edgeNode.Spec.User
	serverIP := edgeNode.Spec.ServerIP

	setupLog.Info("============ready to handle: " + IP + " as: " + nodeType + "===========")
	var k3supCmd string
	var args []string{}
	if nodeType == "server" {
		k3supCmd = "install"
		args = []string{k3supCmd, "--user", user, "--ssh-port", sshPort, "--ip", IP}
	} else {
		k3supCmd = "join"
		args = []string{k3supCmd, "--server-ip", serverIP, "--user", user, "--ssh-port", sshPort, "--ip", IP}
	}
	if sshPort == "" {
		sshPort = "22"
	}

	var stdout, stderr []byte
	var errStdout, errStderr error

	cmd := exec.Command("k3sup", args...)
	// 获取输出对象，可以从该对象中读取输出结果
	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		setupLog.Error(err, "Failed to get k3sup stdout")
	}
	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		setupLog.Error(err, "Failed to get k3sup stderr")
	}
	err = cmd.Start()
	if err != nil {
		setupLog.Error(err, "Failed to start command")
	}

	// 保证关闭输出流
	defer stdoutIn.Close()
	defer stderrIn.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		setupLog.Error(err, "Run cmd failed")
	}
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()
	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		setupLog.Error(err, "cmd.Run() failed with %s\n")
	}
	if errStdout != nil {
		setupLog.Error(errStdout, "failed to capture stdout or stderr\n")
	}
	if errStderr != nil {
		setupLog.Error(errStderr, "failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	setupLog.Info("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	return nil
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			os.Stdout.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
	// never reached
}
