package Request

type CreateAlumnRequest struct{
	LastName string
	Name string
	Curp string
	Email string
	Phone_cel string
	Suspended *bool
}