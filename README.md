# Operator Example


## Init GO

	export GOPATH=$HOME/go
	export GO111MODULE=on
	
	
	mkdir -p $GOPATH/src
	cd $GOPATH/src


## Create Project

	operator-sdk new demo-operator-frontend 
	
	cd demo-operator-frontend


---
## CREATE API
---

	operator-sdk add api --api-version=demo.ibm.com/v1beta1 --kind=MyDemoFrontend


### Edit API Parameters 

	nano ./pkg/apis/demo/v1beta1/MyDemoFrontend_types.go

```
type MyDemoFrontendSpec struct {
        // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Label is the value of the 'daemon=' label to set on a node that should run the daemon
	Label string `json:"label"`
	// Image is the Docker image to run for the daemon
	Image string `json:"image"`
	// Image is the Docker backend image to run for the daemon
	Backend string `json:"backend"`
}
```



## Generate API
	operator-sdk generate k8s
	
	operator-sdk generate openapi


## Modify the Custom resource to include the image information 

	nano ./deploy/crds/lab_v1beta1_MyDemoFrontend_cr.yaml


```yaml
apiVersion: demo.ibm.com/v1beta1
kind: MyDemoFrontend
metadata:
  name: example-MyDemoFrontend
spec:
  # Add fields here
  size: 3
  image: busybox
```





---
## CREATE CONTROLLER
---

operator-sdk add controller --api-version=demo.ibm.com/v1beta1 --kind=MyDemoFrontend



nano ./pkg/controller/MyDemoFrontend/MyDemoFrontend_controller.go



	appsv1 "k8s.io/api/apps/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"

---
# FRONTEND
---

### Compile and create Operator and CR 

	kubectl delete -f ./deploy/operator.yaml
	kubectl delete -f ./deploy/crds/demo_v1beta1_mydemofrontend_cr.yaml
	
	operator-sdk build niklaushirt/demo-operator-frontend:v0.0.1
	docker push niklaushirt/demo-operator-frontend:v0.0.1
	
	kubectl create -f ./deploy/operator.yaml
	kubectl create -f ./deploy/crds/demo_v1beta1_mydemofrontend_cr.yaml
	


### Create All

	kubectl create -f ./deploy/crds/demo_v1beta1_mydemofrontend_crd.yaml
	kubectl create -f ./deploy/service_account.yaml
	kubectl create -f ./deploy/role.yaml
	kubectl create -f ./deploy/role_binding.yaml
	kubectl create -f ./deploy/operator.yaml
	
	kubectl create -f ./deploy/crds/demo_v1beta1_mydemofrontend_cr.yaml



### Delete All 

	kubectl delete -f ./deploy/crds/demo_v1beta1_mydemofrontend_crd.yaml
	kubectl delete -f ./deploy/service_account.yaml
	kubectl delete -f ./deploy/role.yaml
	kubectl delete -f ./deploy/role_binding.yaml
	kubectl delete -f ./deploy/operator.yaml
	
	kubectl delete -f ./deploy/crds/demo_v1beta1_mydemofrontend_cr.yaml



---
# BACKEND
---

### Compile and create Operator and CR 

	kubectl delete -f ./deploy/operator.yaml
	kubectl delete -f ./deploy/crds/demo_v1beta1_mydemobackend_cr.yaml
	
	operator-sdk build niklaushirt/demo-operator-backend:v0.0.1
	docker push niklaushirt/demo-operator-backend:v0.0.1
	
	kubectl create -f ./deploy/operator.yaml
	kubectl create -f ./deploy/crds/demo_v1beta1_mydemobackend_cr.yaml
	
	
### Create All 
	
	kubectl create -f ./deploy/crds/demo_v1beta1_mydemobackend_crd.yaml
	kubectl create -f ./deploy/service_account.yaml
	kubectl create -f ./deploy/role.yaml
	kubectl create -f ./deploy/role_binding.yaml
	kubectl create -f ./deploy/operator.yaml
	
	kubectl create -f ./deploy/crds/demo_v1beta1_mydemobackend_crd.yaml
	
	


### Delete All 

	kubectl delete -f ./deploy/crds/demo_v1beta1_mydemobackend_crd.yaml
	kubectl delete -f ./deploy/service_account.yaml
	kubectl delete -f ./deploy/role.yaml
	kubectl delete -f ./deploy/role_binding.yaml
	kubectl delete -f ./deploy/operator.yaml
	
	kubectl delete -f ./deploy/crds/demo_v1beta1_mydemobackend_crd.yaml


# Links

https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md
https://itnext.io/building-an-operator-for-kubernetes-with-operator-sdk-40a029ea056
https://github.com/kubernetes/sample-controller/blob/master/controller_test.go
https://godoc.org/k8s.io/api/core/v1#Service
https://godoc.org/k8s.io/api/apps/v1#Deployment
https://github.com/shubhomoy/loggingoperator/blob/master/pkg/controller/logmanagement/logmanagement_controller.go











