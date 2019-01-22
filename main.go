package main

import (
	"log"
	"net/http"

	"./src/server"
)

func main() {

	log.Println("Starting http service")
	http.HandleFunc("/stock", server.Handler)
	// Create a new HTTP multiplexer
	/*mux := http.NewServeMux()


	mux.Handle("/", acl.APIManagementWrapper(handler.GraphQL(
		consumption.MakeExecutableSchema(consumption.NewResolver(sqlLiteSearchDb, companyService, contactService, propertyService, leaseCompService, saleCompService, availabilityService, opportunityService, employeeService, loggerService, aclmappingService)),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Println(err)
			debug.PrintStack()
			return errors.New("Internal Error")
		}),
	)))*/
	log.Fatal(http.ListenAndServe(":8080", nil))

}
