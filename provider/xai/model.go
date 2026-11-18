package xai

// Model is an xAI model identifier. Any model id can be used by converting a
// string; the constants below name the ones xAI currently documents.
type Model string

// Current models, with the context window xAI documents for each. xAI
// recommends Grok 4.7 for general use, including code.
//
// Reasoning behaviour is only documented for the two grok-4.20-0309 variants,
// where it is part of the model id. For the rest, send a request and inspect
// the response rather than assuming.
//
// Two notes that apply to the models below: grok-4.20 and newer do not support
// logprobs or top_logprobs, and a request whose prompt reaches 200k tokens is
// billed at the higher rate for every token in that request.
const (
	// ModelGrok47 is Grok 4.7: 500k context. xAI's most capable model, and the
	// one it recommends for code and chat. On the Responses API it always
	// returns reasoning.encrypted_content. Knowledge cut-off May 2026.
	ModelGrok47 Model = "grok-4.7"

	// ModelGrok46 is Grok 4.6: 500k context.
	ModelGrok46 Model = "grok-4.6"

	// ModelGrok45 is Grok 4.5: 500k context.
	ModelGrok45 Model = "grok-4.5"

	// ModelGrok43 is Grok 4.3: 1M context.
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

// Older models, kept for existing callers. xAI no longer lists these on its
// models page, and whether they are still served has not been re-checked since
// 2026-09-26. Prefer a model from the block above.
const (
	// ModelGrok3Mini is xAI's Grok 3 mini model.
	ModelGrok3Mini Model = "grok-3-mini"

	// ModelGrok4FastNonReasoning is xAI's fast non-reasoning Grok 4.1 model.
	ModelGrok4FastNonReasoning Model = "grok-4-1-fast-non-reasoning"
)
