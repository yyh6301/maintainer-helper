package kubernetes

import (
	"github.com/gin-gonic/gin"
)

type kubernetesPlugin struct{}

func (k *kubernetesPlugin) Register(group *gin.RouterGroup) {
	//TODO implement me
	panic("implement me")
}

func (k *kubernetesPlugin) RouterPath() string {
	//TODO implement me
	panic("implement me")
}

func CreateKubernetesPlug() *kubernetesPlugin {

	return &kubernetesPlugin{}
}
