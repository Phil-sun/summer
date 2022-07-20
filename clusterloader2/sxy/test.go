package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("","/root/.kube/karmada.config")
	fmt.Println(err)
	c, _ := kubernetes.NewForConfig(config)
	//c.RbacV1().ClusterRoles().Create(context.TODO(), &rbacv1.ClusterRole{
	//	ObjectMeta: metav1.ObjectMeta{Name: "apiserver-debug-viewer"},
	//	Rules: []rbacv1.PolicyRule{
	//		{Verbs: []string{"get"}, NonResourceURLs: []string{"/debug/*"}},
	//	},
	//}, metav1.CreateOptions{})
	//
	//c.RbacV1().ClusterRoleBindings().Create(context.TODO(), &rbacv1.ClusterRoleBinding{
	//	ObjectMeta: metav1.ObjectMeta{Name: "anonymous:apiserver-debug-viewer"},
	//	RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "apiserver-debug-viewer"},
	//	Subjects: []rbacv1.Subject{
	//		{Kind: "User", Name: "system:anonymous"},
	//	},
	//}, metav1.CreateOptions{})

	body, err := c.CoreV1().RESTClient().Get().AbsPath("/debug/pprof/profile").DoRaw(context.TODO())
	fmt.Println(string(body))
}

