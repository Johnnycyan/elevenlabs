package types

import "os"

type AddVoiceResponse struct {
	VoiceID string `json:"voice_id"`
}

type AnyOfValidationErrorLocItems struct{}

type HistoryPost struct {
	HistoryItemIds []string `json:"history_item_ids"`
}
type Voice struct {
	Name        string     `json:"name"`                  // The name that identifies this voice. This will be displayed in the dropdown of the website.
	Files       []*os.File `json:"files,omitempty"`       // Audio files to add to the voice
	Description string     `json:"description,omitempty"` // How would you describe the voice?
	Labels      string     `json:"labels,omitempty"`      // Serialized labels dictionary for the voice.
}
type TTS struct {
	Text          string           `json:"text"`                     // The text that will get converted into speech. Currently only English text is supported.
	VoiceSettings SynthesisOptions `json:"voice_settings,omitempty"` // Voice settings are applied only on the given TTS request.
}

type SynthesisOptions struct {
	Stability       float64 `json:"stability"`
	SimilarityBoost float64 `json:"similarity_boost"`
}

type ExtendedSubscriptionResponseModel struct {
	Tier                           string                  `json:"tier"`
	CharacterCount                 int32                   `json:"character_count"`
	CharacterLimit                 int32                   `json:"character_limit"`
	CanExtendCharacterLimit        bool                    `json:"can_extend_character_limit"`
	AllowedToExtendCharacterLimit  bool                    `json:"allowed_to_extend_character_limit"`
	NextCharacterCountResetUnix    int32                   `json:"next_character_count_reset_unix"`
	VoiceLimit                     int32                   `json:"voice_limit"`
	ProfessionalVoiceLimit         int32                   `json:"professional_voice_limit"`
	CanExtendVoiceLimit            bool                    `json:"can_extend_voice_limit"`
	CanUseInstantVoiceCloning      bool                    `json:"can_use_instant_voice_cloning"`
	CanUseProfessionalVoiceCloning bool                    `json:"can_use_professional_voice_cloning"`
	AvailableModels                []TtsModelResponseModel `json:"available_models"`
	CanUseDelayedPaymentMethods    bool                    `json:"can_use_delayed_payment_methods"`
	Currency                       string                  `json:"currency"`
	Status                         string                  `json:"status"`
	NextInvoice                    InvoiceResponseModel    `json:"next_invoice"`
}
type FeedbackResponseModel struct {
	ThumbsUp        bool   `json:"thumbs_up"`
	Feedback        string `json:"feedback"`
	Emotions        bool   `json:"emotions"`
	InaccurateClone bool   `json:"inaccurate_clone"`
	Glitches        bool   `json:"glitches"`
	AudioQuality    bool   `json:"audio_quality"`
	Other           bool   `json:"other"`
	ReviewStatus    string `json:"review_status,omitempty"`
}
type FineTuningResponseModel struct {
	ModelID                   string                             `json:"model_id"`
	IsAllowedToFineTune       bool                               `json:"is_allowed_to_fine_tune"`
	FineTuningRequested       bool                               `json:"fine_tuning_requested"`
	FinetuningState           string                             `json:"finetuning_state"`
	VerificationAttempts      []VerificationAttemptResponseModel `json:"verification_attempts"`
	VerificationFailures      []string                           `json:"verification_failures"`
	VerificationAttemptsCount int32                              `json:"verification_attempts_count"`
	SliceIds                  []string                           `json:"slice_ids"`
}
type GetHistoryResponseModel struct {
	History []HistoryItemResponseModel `json:"history"`
}
type GetVoicesResponseModel struct {
	Voices []VoiceResponseModel `json:"voices"`
}
type HistoryItemResponseModel struct {
	HistoryItemId            string                 `json:"history_item_id"`
	RequestId                string                 `json:"request_id"`
	VoiceId                  string                 `json:"voice_id"`
	VoiceName                string                 `json:"voice_name"`
	Text                     string                 `json:"text"`
	DateUnix                 int32                  `json:"date_unix"`
	CharacterCountChangeFrom int32                  `json:"character_count_change_from"`
	CharacterCountChangeTo   int32                  `json:"character_count_change_to"`
	ContentType              string                 `json:"content_type"`
	State                    string                 `json:"state"`
	Settings                 *interface{}           `json:"settings"`
	Feedback                 *FeedbackResponseModel `json:"feedback"`
}
type HttpValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}
type Invoice struct {
	AmountDueCents         int32 `json:"amount_due_cents"`
	NextPaymentAttemptUnix int32 `json:"next_payment_attempt_unix"`
}
type LanguageResponseModel struct {
	IsoCode     string `json:"iso_code"`
	DisplayName string `json:"display_name"`
}
type RecordingResponseModel struct {
	RecordingId    string `json:"recording_id"`
	MimeType       string `json:"mime_type"`
	SizeBytes      int32  `json:"size_bytes"`
	UploadDateUnix int32  `json:"upload_date_unix"`
	Transcription  string `json:"transcription"`
}
type Sample struct {
	SampleId  string `json:"sample_id"`
	FileName  string `json:"file_name"`
	MimeType  string `json:"mime_type"`
	SizeBytes int32  `json:"size_bytes"`
	Hash      string `json:"hash"`
}

type Subscription struct {
	Tier                           string                  `json:"tier"`
	CharacterCount                 int32                   `json:"character_count"`
	CharacterLimit                 int32                   `json:"character_limit"`
	CanExtendCharacterLimit        bool                    `json:"can_extend_character_limit"`
	AllowedToExtendCharacterLimit  bool                    `json:"allowed_to_extend_character_limit"`
	NextCharacterCountResetUnix    int32                   `json:"next_character_count_reset_unix"`
	VoiceLimit                     int32                   `json:"voice_limit"`
	ProfessionalVoiceLimit         int32                   `json:"professional_voice_limit"`
	CanExtendVoiceLimit            bool                    `json:"can_extend_voice_limit"`
	CanUseInstantVoiceCloning      bool                    `json:"can_use_instant_voice_cloning"`
	CanUseProfessionalVoiceCloning bool                    `json:"can_use_professional_voice_cloning"`
	AvailableModels                []TtsModelResponseModel `json:"available_models"`
	CanUseDelayedPaymentMethods    bool                    `json:"can_use_delayed_payment_methods"`
	Currency                       string                  `json:"currency"`
	Status                         string                  `json:"status"`
}
type TtsModelResponseModel struct {
	ModelId           string                  `json:"model_id"`
	DisplayName       string                  `json:"display_name"`
	SupportedLanguage []LanguageResponseModel `json:"supported_language"`
}
type UserResponseModel struct {
	Subscription Subscription `json:"subscription"`
	IsNewUser    bool         `json:"is_new_user"`
	XiAPIKey     string       `json:"xi_api_key"`
}
type ValidationError struct {
	Loc   []AnyOfValidationErrorLocItems `json:"loc"`
	Msg   string                         `json:"msg"`
	Type_ string                         `json:"type"`
}
type VerificationAttemptResponseModel struct {
	Text                string                  `json:"text"`
	DateUnix            int32                   `json:"date_unix"`
	Accepted            bool                    `json:"accepted"`
	Similarity          float64                 `json:"similarity"`
	LevenshteinDistance float64                 `json:"levenshtein_distance"`
	Recording           *RecordingResponseModel `json:"recording"`
}
type VoiceResponseModel struct {
	VoiceId           string                  `json:"voice_id"`
	Name              string                  `json:"name"`
	Samples           []Sample                `json:"samples"`
	Category          string                  `json:"category"`
	FineTuning        FineTuningResponseModel `json:"fine_tuning"`
	Labels            map[string]string       `json:"labels"`
	Description       string                  `json:"description"`
	PreviewUrl        string                  `json:"preview_url"`
	AvailableForTiers []string                `json:"available_for_tiers"`
	Settings          SynthesisOptions        `json:"settings"`
}