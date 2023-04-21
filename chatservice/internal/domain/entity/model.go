// o nome do package sempre deve ser o nome da pasta que ele está
package entity

// É como se fosse uma class em POO
// Parece mais uma interface
type Model struct {
	Name      string
	MaxTokens int
}

func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}

// Isso é um metodo de Model
func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
