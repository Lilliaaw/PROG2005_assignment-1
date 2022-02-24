# assignment-1

In this assignment I have develpoed a REST web application in Golang that allows the client to retrieve information about universities based on their name, or the country that they pertain in. The assignment got a hold of the information about universities and countries through these already existing REST web services : 
    
- https://restcountries.com/
- http://universities.hipolabs.com/

## Endpoints

**/unisearcher/v1/uniinfo/**

This endpoint provides the client with information about universities with the name that is provided or partialliy provided in the url. 
   
    Required Path: uniinfo/{:partial_or_complete_university_name}/


**/unisearcher/v1/neighbourunis/**

This endpoint provides the client with information about universities in the country that is searched for and neighbouring countries with the name that is provided or partially provided in the url. Optionally one can limit the amount of results for each country when searching as well. 
    
    Required Path: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}


**/unisearcher/v1/diag/**

This endpoint provides the client with information about the availability of the other REST web services this web service depends on, seen as status codes. This also provides the version of the service and the uptime of the service. 
    
    Required Path: diag/
    
