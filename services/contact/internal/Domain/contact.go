package Domain

type Contact struct {
	id          int
	fullName    string
	phoneNumber string
}

func newContact(id int, firstName, lastName, fatherName, phoneNumber string) *Contact {
	fullName := lastName + " " + firstName + " " + fatherName
	return &Contact{
		id:          id,
		fullName:    fullName,
		phoneNumber: phoneNumber,
	}
}
