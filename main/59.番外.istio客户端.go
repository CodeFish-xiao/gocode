package main

import (
	"context"
	"log"

	"github.com/owenliang/k8s-client-go/common"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// 需要获取的namespace
	namespace := "bookinfo"
	// 获取kubernetes的 config 文件
	restConfig, err := common.GetRestConf()
	if err != nil {
		return
	}
	// 创建istio的clientset
	istioClient, err := versionedclient.NewForConfig(restConfig)
	if err != nil {
		return
	}
	// 获取namespace下面的所有VS
	vsList, err := istioClient.NetworkingV1alpha3().VirtualServices(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return
	}
	// 遍历输出所有VS 的相关信息
	for i := range vsList.Items {
		vs := vsList.Items[i]
		log.Printf("VirtualService Hosts: %+v -- VirtualService Gateway:  %+v -- VirtualService http:   %+v", vs.Spec.Hosts, vs.Spec.Gateways, vs.Spec.Http)

	}
}
