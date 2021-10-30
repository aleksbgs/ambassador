package services

var UserService Service

func Setup() {
	UserService = CreateService("http://host.docker.internal:8001/api/")
}
