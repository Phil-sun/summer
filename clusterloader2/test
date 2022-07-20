package main

import (
	"context"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("","/root/.kube/karmada.config")
	c, _ := kubernetes.NewForConfig(config)
	c.RbacV1().ClusterRoles().Create(context.TODO(), &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{Name: "apiserver-debug-viewer"},
		Rules: []rbacv1.PolicyRule{
			{Verbs: []string{"get"}, NonResourceURLs: []string{"/debug/*"}},
		},
	}, metav1.CreateOptions{})

	c.RbacV1().ClusterRoleBindings().Create(context.TODO(), &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{Name: "anonymous:apiserver-debug-viewer"},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "apiserver-debug-viewer"},
		Subjects: []rbacv1.Subject{
			{Kind: "User", Name: "system:anonymous"},
		},
	}, metav1.CreateOptions{})
}
