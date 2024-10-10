package main

import (
	"github.com/soustify/data-gateway-model/pkg/generator"
	"github.com/soustify/data-gateway-model/pkg/pb"
)

var (
	address                    *pb.AddressServiceServer
	adminUser                  *pb.AdminUserServiceServer
	applications               *pb.ApplicationServiceServer
	categoriesServices         *pb.CategoriesServicesServiceServer
	contacts                   *pb.ContactServiceServer
	contextDomainChanges       *pb.ContextDomainChangeServiceServer
	contextDomainTablesDetails *pb.ContextDomainTableDetailServiceServer
	contextDomainTransaction   *pb.ContextDomainTransactionServiceServer
	customers                  *pb.CustomerServiceServer
	documentType               *pb.DocumentTypeServiceServer
	partners                   *pb.PartnerServiceServer
	partnersContacts           *pb.PartnerContactServiceServer
	polices                    *pb.PolicyServiceServer
	roles                      *pb.RoleServiceServer
	partnersService            *pb.PartnersServicesServiceServer
	partnersUsers              *pb.PartnerUserServiceServer
	partnerAddresses           *pb.PartnerAddressServiceServer
	rolesApplications          *pb.RoleApplicationServiceServer
	rolesPolices               *pb.RolePolicyServiceServer
	services                   *pb.ServiceServiceServer
	servicesCost               *pb.ServiceCostServiceServer
	categories                 *pb.CategoryServiceServer
	partnersCategories         *pb.PartnerCategoryServiceServer
	extraPackages              = []string{
		"github.com/soustify/data-gateway/pkg/config/database/gorm_dao",
		"gorm.io/gorm",
	}
)

func main() {
	generator.RegisterServer("DocumentType", documentType, extraPackages)
	generator.RegisterServer("Address", address, extraPackages)
	generator.RegisterServer("AdminUser", adminUser, extraPackages)
	generator.RegisterServer("Applications", applications, extraPackages)
	generator.RegisterServer("Categories", categories, extraPackages)
	generator.RegisterServer("CategoriesServices", categoriesServices, extraPackages)
	generator.RegisterServer("Contacts", contacts, extraPackages)
	generator.RegisterServer("ContextDomainChanges", contextDomainChanges, extraPackages)
	generator.RegisterServer("ContextDomainTablesDetails", contextDomainTablesDetails, extraPackages)
	generator.RegisterServer("ContextDomainTransaction", contextDomainTransaction, extraPackages)
	generator.RegisterServer("Customers", customers, extraPackages)
	generator.RegisterServer("DocumentType", documentType, extraPackages)
	generator.RegisterServer("Partners", partners, extraPackages)
	generator.RegisterServer("PartnersCategories", partnersCategories, extraPackages)
	generator.RegisterServer("PartnersContacts", partnersContacts, extraPackages)
	generator.RegisterServer("PartnersService", partnersService, extraPackages)
	generator.RegisterServer("PartnersUsers", partnersUsers, extraPackages)
	generator.RegisterServer("PartnerAddresses", partnerAddresses, extraPackages)
	generator.RegisterServer("Polices", polices, extraPackages)
	generator.RegisterServer("Roles", roles, extraPackages)
	generator.RegisterServer("RolesApplications", rolesApplications, extraPackages)
	generator.RegisterServer("RolesPolices", rolesPolices, extraPackages)
	generator.RegisterServer("Services", services, extraPackages)
	generator.RegisterServer("ServicesCost", servicesCost, extraPackages)
}
