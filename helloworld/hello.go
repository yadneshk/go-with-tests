package helloworld

const (
	greetingEnglish = "Hello, "
	greetingSpanish = "Hola, "
	greetingFrench  = "Bonjour, "
)

func Hello(name, lang string) string {

	if name == "" {
		name = "World"
	}

	return getGreeting(lang) + name
}

func getGreeting(lang string) string {
	switch lang {
	case "Spanish":
		return greetingSpanish
	case "French":
		return greetingFrench
	default:
		return greetingEnglish
	}
}
