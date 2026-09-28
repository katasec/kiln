package xai

// Model is an xAI model identifier. Any model id can be used by converting a
// string; the constants below name the ones xAI currently documents.
type Model string

// Current models. xAI recommends Grok 4.7 for general use, including code.
const (
	// ModelGrok47 is Grok 4.7: 500k context, non-reasoning.
	ModelGrok47 Model = "grok-4.7"

	// ModelGrok46 is Grok 4.6: 500k context, non-reasoning.
	ModelGrok46 Model = "grok-4.6"

	// ModelGrok45 is Grok 4.5: 500k context, non-reasoning.
	ModelGrok45 Model = "grok-4.5"

	// ModelGrok43 is Grok 4.3: 1M context, non-reasoning.
	ModelGrok43 Model = "grok-4.3"

	// ModelGrok420Reasoning is Grok 4.20: 1M context, reasoning.
	ModelGrok420Reasoning Model = "grok-4.20-0309-reasoning"

	// ModelGrok420NonReasoning is Grok 4.20: 1M context, non-reasoning.
	ModelGrok420NonReasoning Model = "grok-4.20-0309-non-reasoning"

	// ModelGrok420MultiAgent is the Grok 4.20 multi-agent model: 1M context.
	ModelGrok420MultiAgent Model = "grok-4.20-multi-agent-0309"

	// ModelGrokBuild01 is Grok Build 0.1: 256k context.
	ModelGrokBuild01 Model = "grok-build-0.1"
)

// Older models, kept for existing callers. xAI no longer lists these, but they
// were still served as of 2026-09-26.
const (
	// ModelGrok3Mini is xAI's Grok 3 mini model.
	ModelGrok3Mini Model = "grok-3-mini"

	// ModelGrok4FastNonReasoning is xAI's fast non-reasoning Grok 4.1 model.
	ModelGrok4FastNonReasoning Model = "grok-4-1-fast-non-reasoning"
)
