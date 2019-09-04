## K3s Operator
Operate K3s cluster at the edge nodes. This operator should run at the data center Kubernetes nodes, which use the k3sup to automatically provision the k3s cluster on the edge node.

## Cloud vs. Edge
Edge node should connect with the Kubernetes node which running the operator with OpenVPN client. Openvpn server will run separately at a network node, in that case we can realize the cloud-edge corporation

## Goal
The cluster running at edge nodes could be managed and automatically operated by the K3s Operator running on the datacenter

