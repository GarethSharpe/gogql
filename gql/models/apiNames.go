package models

/********************************************/
/*				Api Names					*/
/********************************************/

func (object Case) ApiName() string {
	return "Case"
}
func (object Asset) ApiName() string {
	return "Asset"
}
func (object Contact) ApiName() string {
	return "Contact"
} 

/********************************************/
/*			External Id Api Names			*/
/********************************************/

func (object Case) ExternalIdApiName() string {
	return "Ext_ID"
}
func (object Asset) ExternalIdApiName() string {
	return "Ext_ID"
}
func (object Contact) ExternalIdApiName() string {
	return "Ext_ID"
} 