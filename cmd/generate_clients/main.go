package main

import (
	"github.com/soustify/data-gateway-model/pkg/generator"
	"github.com/soustify/data-gateway-model/pkg/pb"
)

var (
	address                    *pb.AddressServiceClient
	adminUser                  *pb.AdminUserServiceClient
	applications               *pb.ApplicationServiceClient
	categoriesServices         *pb.CategoriesServicesServiceClient
	contacts                   *pb.ContactServiceClient
	contextDomainChanges       *pb.ContextDomainChangeServiceClient
	contextDomainTablesDetails *pb.ContextDomainTableDetailServiceClient
	contextDomainTransaction   *pb.ContextDomainTransactionServiceClient
	customers                  *pb.CustomerServiceClient
	documentType               *pb.DocumentTypeServiceClient
	partners                   *pb.PartnerServiceClient
	partnersContacts           *pb.PartnerContactServiceClient
	polices                    *pb.PolicyServiceClient
	roles                      *pb.RoleServiceClient
	partnersService            *pb.PartnersServicesServiceClient
	partnersUsers              *pb.PartnerUserServiceClient
	partnerAddresses           *pb.PartnerAddressServiceClient
	rolesApplications          *pb.RoleApplicationServiceClient
	rolesPolices               *pb.RolePolicyServiceClient
	services                   *pb.ServiceServiceClient
	servicesCost               *pb.ServiceCostServiceClient
	categories                 *pb.CategoryServiceClient
	partnersCategories         *pb.PartnerCategoryServiceClient
)

func main() {
	generator.RegisterLambdaClient("Address", "address_service_server", address)
	generator.RegisterLambdaClient("AdminUser", "admin_user_service_server", adminUser)
	generator.RegisterLambdaClient("Applications", "applications_service_server", applications)
	generator.RegisterLambdaClient("Categories", "categories_service_server", categories)
	generator.RegisterLambdaClient("CategoriesServices", "categories_services_service_server", categoriesServices)
	generator.RegisterLambdaClient("Contacts", "contacts_service_server", contacts)
	generator.RegisterLambdaClient("ContextDomainChanges", "context_domain_changes_service_server", contextDomainChanges)
	generator.RegisterLambdaClient("ContextDomainTablesDetails", "context_domain_tables_details_service_server", contextDomainTablesDetails)
	generator.RegisterLambdaClient("ContextDomainTransaction", "context_domain_transaction_service_server", contextDomainTransaction)
	generator.RegisterLambdaClient("Customers", "customers_service_server", customers)
	generator.RegisterLambdaClient("DocumentType", "document_type_service_server", documentType)
	generator.RegisterLambdaClient("Partners", "partners_service_server", partners)
	generator.RegisterLambdaClient("PartnersCategories", "partners_categories_service_server", partnersCategories)
	generator.RegisterLambdaClient("PartnersContacts", "partners_contacts_service_server", partnersContacts)
	generator.RegisterLambdaClient("PartnersService", "partners_service_service_server", partnersService)
	generator.RegisterLambdaClient("PartnersUsers", "partners_users_service_server", partnersUsers)
	generator.RegisterLambdaClient("PartnerAddresses", "partner_addresses_service_server", partnerAddresses)
	generator.RegisterLambdaClient("Polices", "polices_service_server", polices)
	generator.RegisterLambdaClient("Roles", "roles_service_server", roles)
	generator.RegisterLambdaClient("RolesApplications", "roles_applications_service_server", rolesApplications)
	generator.RegisterLambdaClient("RolesPolices", "roles_polices_service_server", rolesPolices)
	generator.RegisterLambdaClient("Services", "services_service_server", services)
	generator.RegisterLambdaClient("ServicesCost", "services_cost_service_server", servicesCost)
}
