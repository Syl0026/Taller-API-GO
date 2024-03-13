package Request


type UpdateAlumnRequest struct{
	IdAlumn uint
	LastName string
	Name string
	Curp string
	Email string
	Phone_cel string
	Suspended *bool
}