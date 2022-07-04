# Companies crud API


<a name="overview"></a>
## Overview
Provides endpoints to view and manage Company entities.


### Version information
*Version* : 1.0.0


### URI scheme
*Schemes* : HTTPS, HTTP


### Consumes

* `application/json`


### Produces

* `application/json`




<a name="paths"></a>
## Paths

<a name="getcompanies"></a>
### GET /api/v1/companies

#### Description
Returns list of Companies


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**code**  <br>*optional*|Filtring parameter|string|
|**Query**|**country**  <br>*optional*|Filtring parameter|string|
|**Query**|**limit**  <br>*optional*|The numbers of items to return|integer (int64)|
|**Query**|**name**  <br>*optional*|Filtring parameter|string|
|**Query**|**phone**  <br>*optional*|Filtring parameter|string|
|**Query**|**website**  <br>*optional*|Filtring parameter|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|JSON Response|[CompaniesListResponse](#companieslistresponse)|


#### Example HTTP request

##### Request path
```
/api/v1/companies
```


#### Example HTTP response

##### Response 200
```json
{
  "items" : [ {
    "code" : "00101",
    "country" : "Ukraine",
    "id" : 1,
    "name" : "EPAM",
    "phone" : "+38063334422",
    "website" : "https://www.epam.com/"
  } ]
}
```


<a name="createcompanies"></a>
### POST /api/v1/companies/

#### Description
Create a company


#### Body parameter
*Name* : Body  
*Flags* : optional  
*Type* : [CompanyRequest](#companyrequest)


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|JSON Response|[CompanyResponse](#companyresponse)|


#### Security

|Type|Name|Scopes|
|---|---|---|
|**apiKey**|**[bearer](#bearer)**|[]|


#### Example HTTP request

##### Request path
```
/api/v1/companies/
```


##### Request body
```json
{
  "code" : "00101",
  "country" : "Ukraine",
  "name" : "EPAM",
  "phone" : "+38063334422",
  "website" : "https://www.epam.com/"
}
```


#### Example HTTP response

##### Response 200
```json
{
  "code" : "00101",
  "country" : "Ukraine",
  "id" : 1,
  "name" : "EPAM",
  "phone" : "+38063334422",
  "website" : "https://www.epam.com/"
}
```


<a name="updatecompany"></a>
### PUT /api/v1/companies/

#### Description
Update company


#### Body parameter
*Name* : Body  
*Flags* : optional  
*Type* : [CompanyResponse](#companyresponse)


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|JSON Response|[CompanyResponse](#companyresponse)|


#### Security

|Type|Name|Scopes|
|---|---|---|
|**apiKey**|**[bearer](#bearer)**|[]|


#### Example HTTP request

##### Request path
```
/api/v1/companies/
```


##### Request body
```json
{
  "code" : "00101",
  "country" : "Ukraine",
  "id" : 1,
  "name" : "EPAM",
  "phone" : "+38063334422",
  "website" : "https://www.epam.com/"
}
```


#### Example HTTP response

##### Response 200
```json
{
  "code" : "00101",
  "country" : "Ukraine",
  "id" : 1,
  "name" : "EPAM",
  "phone" : "+38063334422",
  "website" : "https://www.epam.com/"
}
```


<a name="getcompany"></a>
### GET /api/v1/companies/{id}

#### Description
Returns company by ID


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**id**  <br>*required*|Product ID|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|JSON Response|[CompanyResponse](#companyresponse)|


#### Example HTTP request

##### Request path
```
/api/v1/companies/0
```


#### Example HTTP response

##### Response 200
```json
{
  "code" : "00101",
  "country" : "Ukraine",
  "id" : 1,
  "name" : "EPAM",
  "phone" : "+38063334422",
  "website" : "https://www.epam.com/"
}
```


<a name="deletecompany"></a>
### DELETE /api/v1/companies/{id}

#### Description
Delete company by ID


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**id**  <br>*required*|Product ID|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|JSON Response|[DeleteCompayResponse](#deletecompayresponse)|


#### Security

|Type|Name|Scopes|
|---|---|---|
|**apiKey**|**[bearer](#bearer)**|[]|


#### Example HTTP request

##### Request path
```
/api/v1/companies/0
```


#### Example HTTP response

##### Response 200
```json
{
  "Success" : true
}
```




<a name="definitions"></a>
## Definitions

<a name="companieslistresponse"></a>
### CompaniesListResponse

|Name|Description|Schema|
|---|---|---|
|**items**  <br>*required*|Company items  <br>**Example** : `[ "[companyresponse](#companyresponse)" ]`|< [CompanyResponse](#companyresponse) > array|


<a name="companyrequest"></a>
### CompanyRequest

|Name|Description|Schema|
|---|---|---|
|**code**  <br>*required*|Code  <br>**Example** : `"00101"`|string|
|**country**  <br>*required*|Country  <br>**Example** : `"Ukraine"`|string|
|**name**  <br>*required*|Name  <br>**Example** : `"EPAM"`|string|
|**phone**  <br>*optional*|Phone  <br>**Example** : `"+38063334422"`|string|
|**website**  <br>*optional*|Website  <br>**Example** : `"https://www.epam.com/"`|string|


<a name="companyresponse"></a>
### CompanyResponse

|Name|Description|Schema|
|---|---|---|
|**code**  <br>*required*|Code  <br>**Example** : `"00101"`|string|
|**country**  <br>*required*|Country  <br>**Example** : `"Ukraine"`|string|
|**id**  <br>*required*|Company ID auto generated for each company  <br>**Example** : `1`|integer (int64)|
|**name**  <br>*required*|Name  <br>**Example** : `"EPAM"`|string|
|**phone**  <br>*optional*|Phone  <br>**Example** : `"+38063334422"`|string|
|**website**  <br>*optional*|Website  <br>**Example** : `"https://www.epam.com/"`|string|


<a name="deletecompayresponse"></a>
### DeleteCompayResponse

|Name|Description|Schema|
|---|---|---|
|**Success**  <br>*optional*|Success  <br>**Example** : `true`|boolean|




<a name="securityscheme"></a>
## Security

<a name="bearer"></a>
### bearer
*Type* : apiKey  
*Name* : Authorization  
*In* : HEADER



