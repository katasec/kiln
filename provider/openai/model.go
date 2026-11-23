package openai

// Model is an OpenAI model identifier. Any model id can be used by converting
// a string; the constants below name the flagship text models OpenAI
// currently documents.
type Model string

// Current models, all with a 1.05M context window and a 128k output cap.
const (
	// ModelGPT62Astra is GPT-6.2 Astra: OpenAI's most capable model, for
	// complex reasoning and coding.
	ModelGPT62Astra Model = "gpt-6.2-astra"

	// ModelGPT62Sol is GPT-6.2 Sol: the go-to model for everyday work.
	ModelGPT62Sol Model = "gpt-6.2-sol"

	// ModelGPT61Terra is GPT-6.1 Terra: fast and efficient.
	ModelGPT61Terra Model = "gpt-6.1-terra"

	// ModelGPT6Luna is GPT-6 Luna: the most efficient model, for high-volume
	// tasks, and the cheapest of the four.
	ModelGPT6Luna Model = "gpt-6-luna"
)

// Deprecated models, kept for existing callers.
const (
	// ModelGPT54Nano is OpenAI's GPT-5.4 nano model.
	//
	// Deprecated: OpenAI gave notice on 2026-10-01 and shuts this model down
	// on 2027-04-01. Use ModelGPT6Luna, the replacement OpenAI names.
	ModelGPT54Nano Model = "gpt-5.4-nano"
)
