package handler

import "testing"

func TestCreateOpeningRequest(t *testing.T) {
	t.Run("Validação de CreateOpeningRequest", func(t *testing.T) {
		remote := true
		request := CreateOpeningRequest{
			Role:     "Desenvolvedor Backend",
			Company:  "Laboratório Bridge",
			Location: "Florianópolis",
			Remote:   &remote,
			Link:     "https://portal.bridge.ufsc.br",
			Salary:   2448,
		}

		err := request.Validate()

		if err != nil {
			t.Errorf("Unexpected error ocurred: %v", err.Error())
		}
	})

	t.Run("Validação negativa de CreateOpeningRequest inválido", func(t *testing.T) {
		request := CreateOpeningRequest{}

		err := request.Validate()

		if err == nil {
			t.Errorf("Expected error, but none ocurred")
		}
	})
}

func TestUpdateOpeningRequest(t *testing.T) {
	t.Run("Validação de UpdateOpeningRequest", func(t *testing.T) {
		remote := true
		request := UpdateOpeningRequest{
			Role:     "Desenvolvedor Backend",
			Company:  "Laboratório Bridge",
			Location: "Florianópolis",
			Remote:   &remote,
			Link:     "https://portal.bridge.ufsc.br",
			Salary:   2448,
		}

		err := request.Validate()

		if err != nil {
			t.Errorf("Unexpected error ocurred: %v", err.Error())
		}
	})

	t.Run("Validação negativo de UpdateCreateRequest inválido", func(t *testing.T) {
		request := UpdateOpeningRequest{}

		err := request.Validate()

		if err == nil {
			t.Errorf("Expected error, but none ocurred")
		}
	})
}
