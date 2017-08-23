package models

type User struct {
	Id       AutoId   `json:"_id" bson:"_id,omitempty"`
	Email    Email    `json:"email" bson:"email"`
	Password Password `json:"password" bson:"password"`
	Name     Name     `json:"name" bson:"name"`
	Bio      Text     `json:"bio" bson:"bio,omitempty"`
}

// Returns validation error or nil if valid
func (u User) Validate() error {
	if err := u.Id.Validate(); err != nil {
		return err
	}
	if err := u.Email.Validate(); err != nil {
		return err
	}
	if err := u.Password.Validate(); err != nil {
		return err
	}
	if err := u.Name.Validate(); err != nil {
		return err
	}
	if err := u.Bio.Validate(); err != nil {
		return err
	}
	return nil
}

type UsersList []User
