package block

type DtoBlockRequest struct {
	Name string `json:"name,omitempty" validate:"required"`
}
