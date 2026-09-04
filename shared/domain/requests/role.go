package requests

import "github.com/go-playground/validator/v10"

// RoleRequestPayload represents the base payload structure for role management requests.
// Contains common fields used across role-related operations.
type RoleRequestPayload struct {
	UserID        int    `json:"user_id"`        // ID of the user performing the role operation
	CorrelationID string `json:"correlation_id"` // Unique identifier for request tracing
	ReplyTopic    string `json:"reply_topic"`    // Topic name for asynchronous response delivery
}

type FindAllRoles struct {
	Search   string `json:"search" validate:"required"`
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRoleRequest struct {
	ID   *int   `json:"id"`
	Name string `json:"name" validate:"required"`
}

func (r *CreateRoleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateRoleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
