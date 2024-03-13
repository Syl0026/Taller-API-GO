package Response

//$ Para enviar datos a los rpoints
type AlumnResponse struct{
	IdAlumn uint
	LastName string
	Name string
	Curp string
	Email string
	Phone_cel string
	Suspended *bool
}