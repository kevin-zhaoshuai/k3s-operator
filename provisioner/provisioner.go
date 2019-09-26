package provisioner

import (
	"github.com/golang/glog"
	provisionerv1 "github.com/kevin-zhaoshuai/k3s-operator/api/v1"
	"io/ioutil"
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

	setupLog.Info("============ready to handle: " + IP + " as: " + nodeType + "===========")
	var k3supCmd string
	if nodeType == "server" {
		k3supCmd = "install"
	} else {
		k3supCmd = "join"
	}
	if sshPort == "" {
		sshPort = "22"
	}
	cmdStr := k3supCmd + " --user " + user + " --sshPort " + sshPort + " --ip " + IP
	setupLog.Info(cmdStr)
	cmd := exec.Command("/usr/local/bin/k3sup", cmdStr)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		setupLog.Error(err, "Failed to run k3sup")
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		setupLog.Error(err, "Run cmd failed")
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		glog.Fatal(err)
	}
	setupLog.Info(string(opBytes))
	return nil
}
