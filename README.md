# podLifecycle
# sample o/p
### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp create --namespace default --pod pod1

Created pod: pod1

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp pods --namespace default

Pods:
Name: pod1, Creation Time: 2023-07-15 19:07:51 +0530 IST

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp create --namespace default --pod pod2

Created pod: pod2

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp pods --namespace default

Pods:
Name: pod2, Creation Time: 2023-07-15 19:08:46 +0530 IST
Name: pod1, Creation Time: 2023-07-15 19:07:51 +0530 IST

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp watch --namespace default

Pod added: pod1
Pod added: pod2
^C

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp delete --namespace default --time 2s

Deleted pod: pod1
Deleted pod: pod2

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp pods --namespace default

Pods:

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp create --namespace default --pod pod1

Created pod: pod1

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp create --namespace default --pod pod2

Created pod: pod2

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp pods --namespace default

Pods:
Name: pod2, Creation Time: 2023-07-15 19:10:55 +0530 IST
Name: pod1, Creation Time: 2023-07-15 19:10:51 +0530 IST

### root@ubuntu:~/Raghav/go_workspace/src/podLifecycle# ./myapp pods --namespace default --ascending

Pods:
Name: pod1, Creation Time: 2023-07-15 19:10:51 +0530 IST
Name: pod2, Creation Time: 2023-07-15 19:10:55 +0530 IST
