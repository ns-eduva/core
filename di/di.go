package di

import (
	"eduva/core/controller/usercontroller"
	"eduva/core/internal/auth"
	"eduva/core/internal/crm"
)

type AppDependencies struct {
	UserController usercontroller.Interface
	// Add your dependencies here
}

func BuildDependencies() *AppDependencies {
	// TODO remplacer par les ENV variables
	crmBaseURL := "path api crm"
	authBaseURL := "path api auth"

	// services
	crmService := crm.NewCrmService(crmBaseURL)
	authService := auth.NewAuthService(authBaseURL)
	// TODO : ajouter les autres services

	// controllers
	userCtrl := usercontroller.NewUserController(crmService, authService)
	// TODO : ajouter les autres controllers

	return &AppDependencies{
		UserController: userCtrl,
	}
}
