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
	generator.RegisterAdapter("DocumentType", documentType, extraPackages)
	//generator.RegisterAdapter("Address", address, extraPackages)
	//generator.RegisterAdapter("AdminUser", adminUser, extraPackages)
	//generator.RegisterAdapter("Applications", applications, extraPackages)
	//generator.RegisterAdapter("Categories", categories, extraPackages)
	//generator.RegisterAdapter("CategoriesServices", categoriesServices, extraPackages)
	//generator.RegisterAdapter("Contacts", contacts, extraPackages)
	//generator.RegisterAdapter("ContextDomainChanges", contextDomainChanges, extraPackages)
	//generator.RegisterAdapter("ContextDomainTablesDetails", contextDomainTablesDetails, extraPackages)
	//generator.RegisterAdapter("ContextDomainTransaction", contextDomainTransaction, extraPackages)
	//generator.RegisterAdapter("Customers", customers, extraPackages)
	//generator.RegisterAdapter("DocumentType", documentType, extraPackages)
	//generator.RegisterAdapter("Partners", partners, extraPackages)
	//generator.RegisterAdapter("PartnersCategories", partnersCategories, extraPackages)
	//generator.RegisterAdapter("PartnersContacts", partnersContacts, extraPackages)
	//generator.RegisterAdapter("PartnersService", partnersService, extraPackages)
	//generator.RegisterAdapter("PartnersUsers", partnersUsers, extraPackages)
	//generator.RegisterAdapter("PartnerAddresses", partnerAddresses, extraPackages)
	//generator.RegisterAdapter("Polices", polices, extraPackages)
	//generator.RegisterAdapter("Roles", roles, extraPackages)
	//generator.RegisterAdapter("RolesApplications", rolesApplications, extraPackages)
	//generator.RegisterAdapter("RolesPolices", rolesPolices, extraPackages)
	//generator.RegisterAdapter("Services", services, extraPackages)
	//generator.RegisterAdapter("ServicesCost", servicesCost, extraPackages)
}
