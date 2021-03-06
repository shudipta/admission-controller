package apiserver

import (
	clientset "admission-controller/client/clientset/versioned"
	"sync"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"fmt"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"encoding/json"
	"net/http"
	"k8s.io/client-go/rest"
	"log"
)

type SomethingValidationHook struct {
	Client clientset.Interface

	lock        sync.RWMutex
}

func (*SomethingValidationHook) ValidatingResource() (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{
		Group:    "admission.somethingcontroller.kube-ac.com",
		Version:  "v1alpha1",
		Resource: "validations",
	},
		"validation"
}

func (a *SomethingValidationHook) Validate(
	req *admissionv1beta1.AdmissionRequest) *admissionv1beta1.AdmissionResponse {
	log.Println("------------validating...........")

	validatingObjectMeta := &NamedThing{}
	err := json.Unmarshal(req.Object.Raw, validatingObjectMeta)
	if err != nil {
		log.Println("-----------invalid obj...........")

		return &admissionv1beta1.AdmissionResponse{
			Allowed: false,
			Result: &metav1.Status{
				Status: metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
				Message: err.Error(),
			},
		}
	}

	if req.Operation == admissionv1beta1.Create {
		log.Println("----------validating at creating time...........")

		if len(validatingObjectMeta.Name) == 0 {
			log.Println("-----------name field is empty...........")

			return &admissionv1beta1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status: metav1.StatusFailure, Code: http.StatusForbidden, Reason: metav1.StatusReasonForbidden,
					Message: "name is required",
				},
			}
		}

		if validatingObjectMeta.Annotations == nil ||
			validatingObjectMeta.Annotations["sample-label"] != "true" {
			log.Printf("----------%v has no annotations...........\n", validatingObjectMeta.Name)

			return &admissionv1beta1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status: metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: "doesn't contain required annotations ([sample-label: true])",
				},
			}
		}
		//
		//a.lock.RLock()
		//defer a.lock.RUnlock()

		obj, err := a.Client.SomethingcontrollerV1alpha1().
			Somethings(validatingObjectMeta.Namespace).
				Get(validatingObjectMeta.Name, metav1.GetOptions{})
		if err == nil {
			log.Printf("---------%v already exists...........\n", obj.Name)

			return &admissionv1beta1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status: metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: fmt.Sprintf("%q is reserved", obj.Name),
				},
			}
		}
	} else if req.Operation == admissionv1beta1.Delete {
		log.Println("----------validating at deleting time...........")

		obj, err := a.Client.SomethingcontrollerV1alpha1().
			Somethings(validatingObjectMeta.Namespace).
			Get(validatingObjectMeta.Name, metav1.GetOptions{})
		if err != nil {
			log.Printf("----------%v doesn't exist...........\n")

			return &admissionv1beta1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status: metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: fmt.Sprintf("%q doesn't exist", obj.Name),
				},
			}
		}

		if obj.Annotations != nil &&
			obj.Annotations["sample-label"] == "true" {
			log.Printf("----------annotations ([sample-label: true]) must be remmoved from %v...........\n", obj.Name)

			return &admissionv1beta1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status: metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: "annotations ([sample-label: true]) must be remmoved",
				},
			}
		}

	}

	return &admissionv1beta1.AdmissionResponse{Allowed: true}
}

func (*SomethingValidationHook) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	log.Println("SomethingValidator: Initialize")
	return nil
}

type NamedThing struct {
	ObjectMeta `json:"metadata"`
}

type ObjectMeta struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
}
