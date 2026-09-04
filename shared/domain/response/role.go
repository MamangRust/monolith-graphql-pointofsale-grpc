package response

// RoleResponsePayload represents the base payload structure for role validation responses.
// Used to verify role assignments and permissions.
type RoleResponsePayload struct {
	CorrelationID string   `json:"correlation_id"` // Unique ID for request tracing
	Valid         bool     `json:"valid"`          // Indicates if role validation succeeded
	RoleNames     []string `json:"role_names"`     // List of role names associated with the validation
}

type RoleResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RoleResponseDeleteAt struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type ApiResponseRoleAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseRoleDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseRole struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    *RoleResponse `json:"data"`
}

type ApiResponsesRole struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []*RoleResponse `json:"data"`
}

type ApiResponsePaginationRole struct {
	Status     string          `json:"status"`
	Message    string          `json:"message"`
	Data       []*RoleResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}

type ApiResponsePaginationRoleDeleteAt struct {
	Status     string                  `json:"status"`
	Message    string                  `json:"message"`
	Data       []*RoleResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta         `json:"pagination"`
}
