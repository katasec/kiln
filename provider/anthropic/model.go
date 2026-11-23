package anthropic

// Model is an Anthropic model identifier. Any model id can be used by
// converting a string; the constants below name the ones Anthropic currently
// documents.
type Model string

// Current models, all with a 1M context window. ModelClaudeOpus55 is the
// usual default; ModelClaudeFable52 is more capable but costs more per token.
const (
	// ModelClaudeFable52 is Claude Fable 5.2: Anthropic's most capable widely
	// released model, for demanding reasoning and long-horizon agentic work.
	// Priced above the Opus tier, and not available to organisations under
	// zero data retention without Anthropic's authorisation.
	ModelClaudeFable52 Model = "claude-fable-5-2"

	// ModelClaudeOpus55 is Claude Opus 5.5: the current Opus and a good
	// default. Thinking is always on and cannot be disabled.
	ModelClaudeOpus55 Model = "claude-opus-5-5"

	// ModelClaudeSonnet55 is Claude Sonnet 5.5: the current Sonnet, balancing
	// speed and capability for everyday coding and agent work.
	ModelClaudeSonnet55 Model = "claude-sonnet-5-5"

	// ModelClaudeHaiku55 is Claude Haiku 5.5: the current Haiku, the cheapest
	// and fastest of the line.
	ModelClaudeHaiku55 Model = "claude-haiku-5-5"
)

// Previous generation, kept for existing callers. Anthropic still serves
// these, but prefer a model from the block above for new code.
const (
	// ModelClaudeOpus5 is Claude Opus 5, superseded by ModelClaudeOpus55.
	ModelClaudeOpus5 Model = "claude-opus-5"

	// ModelClaudeSonnet5 is Claude Sonnet 5, superseded by ModelClaudeSonnet55.
	ModelClaudeSonnet5 Model = "claude-sonnet-5"

	// ModelClaudeHaiku45 is Claude Haiku 4.5, superseded by
	// ModelClaudeHaiku55. Note its 200k context window, where every model
	// above has 1M.
	ModelClaudeHaiku45 Model = "claude-haiku-4-5"
)
