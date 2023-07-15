package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func listPods(clientset *kubernetes.Clientset, namespace string, ascending bool) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	sort.SliceStable(pods.Items, func(i, j int) bool {
		if ascending {
			return pods.Items[i].ObjectMeta.CreationTimestamp.Time.Before(pods.Items[j].ObjectMeta.CreationTimestamp.Time)
		}
		return pods.Items[i].ObjectMeta.CreationTimestamp.Time.After(pods.Items[j].ObjectMeta.CreationTimestamp.Time)
	})

	fmt.Println("Pods:")
	for _, pod := range pods.Items {
		fmt.Printf("Name: %s, Creation Time: %s\n", pod.Name, pod.CreationTimestamp.String())
	}
}

func deletePods(clientset *kubernetes.Clientset, namespace string, deleteTime time.Duration) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range pods.Items {
		runningTime := time.Since(pod.CreationTimestamp.Time)
		if runningTime > deleteTime {
			err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(),pod.Name, metav1.DeleteOptions{})
			if err != nil {
				panic(err)
			}
			fmt.Printf("Deleted pod: %s\n", pod.Name)
		}
	}
}

func createPod(clientset *kubernetes.Clientset, namespace, podName string) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range pods.Items {
		if pod.Name == podName {
			fmt.Printf("Pod %s already exists in namespace %s\n", podName, namespace)
			return
		}
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "example-container",
					Image: "nginx",
				},
			},
		},
	}

	createdPod, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(),pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created pod: %s\n", createdPod.Name)
}

func watchPods(clientset *kubernetes.Clientset, namespace string) {
	watch, err := clientset.CoreV1().Pods(namespace).Watch(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	ch := watch.ResultChan()
	for event := range ch {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			fmt.Printf("Unexpected event type: %T\n", event.Object)
			continue
		}

		switch event.Type {
		case "ADDED":
			fmt.Printf("Pod added: %s\n", pod.Name)
		case "DELETED":
			fmt.Printf("Pod deleted: %s\n", pod.Name)
		}
	}
}

func main() {
	cmd := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)
	namespace := cmd.String("namespace", "default", "Namespace")
	ascending := cmd.Bool("ascending", false, "Sort by ascending order of creation time")
	deleteTime := cmd.Duration("time", 0, "Pod running duration for deletion")
	podName := cmd.String("pod", "", "Pod name")
	cmd.Parse(os.Args[2:])

	config, err := clientcmd.BuildConfigFromFlags("","/root/admin.conf")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "pods":
		listPods(clientset, *namespace, *ascending)
	case "delete":
		deletePods(clientset, *namespace, *deleteTime)
	case "create":
		createPod(clientset, *namespace, *podName)
	case "watch":
		watchPods(clientset, *namespace)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
