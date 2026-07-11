// Package augmented provides Game of Thrones data augmented with modern demo fields
// such as email addresses, phone numbers, and risk levels.
package augmented

import (
	got "github.com/grokify/gameofthrones"
	"github.com/grokify/goauth/scim"
)

// Character embeds canonical GoT data and adds modern demo fields.
type Character struct {
	got.Character
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"` // E.164 format
}

// NewCharacter creates an augmented Character from a base Character.
func NewCharacter(base got.Character, email, phone string) Character {
	return Character{
		Character: base,
		Email:     email,
		Phone:     phone,
	}
}

// GetEmail returns the character's email, using the embedded SCIM email if
// the augmented email is empty.
func (c *Character) GetEmail() string {
	if c.Email != "" {
		return c.Email
	}
	if len(c.Character.Character.Emails) > 0 {
		return c.Character.Character.Emails[0].Value
	}
	return ""
}

// GetPhone returns the character's phone number in E.164 format, using the
// embedded SCIM phone if the augmented phone is empty.
func (c *Character) GetPhone() string {
	if c.Phone != "" {
		return c.Phone
	}
	if len(c.Character.Character.PhoneNumbers) > 0 {
		return c.Character.Character.PhoneNumbers[0].Value
	}
	return ""
}

// DisplayName returns the character's display name.
func (c *Character) DisplayName() string {
	return c.Character.Character.DisplayName
}

// SCIMUser returns the character as a SCIM user with augmented contact info.
func (c *Character) SCIMUser() scim.User {
	user := c.Character.Character
	if c.Email != "" && len(user.Emails) == 0 {
		user.Emails = append(user.Emails, scim.Item{Value: c.Email})
	}
	if c.Phone != "" && len(user.PhoneNumbers) == 0 {
		user.PhoneNumbers = append(user.PhoneNumbers, scim.Item{Value: c.Phone})
	}
	return user
}
