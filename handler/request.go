package handler

import "fmt"

func errParamsIsRequered(name, typ string) error {
	return fmt.Errorf("parametro: %s (type: %s) é obrigatório", name, typ)
}

// Create Opening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	// Verifica se os campos estão vazios.
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("requisição do body vazia ou malformatada")
	}

	if r.Role == "" {
		return errParamsIsRequered("role", "string")
	}
	if r.Company == "" {
		return errParamsIsRequered("company", "string")
	}
	if r.Location == "" {
		return errParamsIsRequered("location", "string")
	}
	if r.Link == "" {
		return errParamsIsRequered("link", "string")
	}
	if r.Remote == nil {
		return errParamsIsRequered("remote", "string")
	}
	if r.Salary <= 0 {
		return errParamsIsRequered("salary", "string")
	}

	return nil
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("pelo menos algum dos campos precisam ser informados")
}
