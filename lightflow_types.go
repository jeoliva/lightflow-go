package lightflow

import (
	"encoding/json"
	"time"
)

// PaginatedResponse represents a paginated response.
type IWorkflowStoragePaginatedResponse struct {
	TotalItems int                        `json:"totalItems"`
	PageIndex  int                        `json:"pageIndex"`
	PageSize   int                        `json:"pageSize"`
	Pages      int                        `json:"pages"`
	Items      []IWorkflowStorageResponse `json:"items"`
}

// IWorkflowStorage represents the workflow storage.
type IWorkflowStorage struct {
	UUID         string                     `json:"uuid"`
	Username     string                     `json:"username"`
	AccountName  string                     `json:"accountName"`
	CreationDate time.Time                  `json:"creationDate"`
	StorageID    string                     `json:"storageId"`
	StorageType  string                     `json:"storageType"`
	Name         string                     `json:"name"`
	Label        *string                    `json:"label,omitempty"`
	Args         []IWorkflowStorageArgument `json:"args"`
	RemoveDate   *time.Time                 `json:"removeDate,omitempty"`
}

// IWorkflowStorageResponse represents the JSON-based response.
type IWorkflowStorageResponse struct {
	UUID         string                     `json:"uuid"`
	Username     string                     `json:"username"`
	AccountName  string                     `json:"accountName"`
	CreationDate time.Time                  `json:"creationDate"`
	StorageID    string                     `json:"storageId"`
	StorageType  string                     `json:"storageType"`
	Name         string                     `json:"name"`
	Label        *string                    `json:"label,omitempty"`
	Args         []IWorkflowStorageArgument `json:"args"`
	RemoveDate   *time.Time                 `json:"removeDate,omitempty"`
}

// IWorkflowStorageArgument represents an argument in the workflow storage.
type IWorkflowStorageArgument struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Value       string `json:"value,omitempty"`
}

// ToJSON converts the IWorkflowStorage struct to a JSON string.
func (r *IWorkflowStorage) ToJSON() (string, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToJSON converts the IWorkflowStorageResponse struct to a JSON string.
func (r *IWorkflowStorageResponse) ToJSON() (string, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type IAssetPaginatedResponse struct {
	TotalItems int      `json:"totalItems"`
	PageIndex  int      `json:"pageIndex"`
	PageSize   int      `json:"pageSize"`
	Pages      int      `json:"pages"`
	Items      []IAsset `json:"items"`
}

// IAsset represents the asset interface.
type IAsset struct {
	UUID              *string             `json:"uuid,omitempty"`
	WorkflowID        *string             `json:"workflowId,omitempty"`
	Username          *string             `json:"username,omitempty"`
	AccountName       *string             `json:"accountName,omitempty"`
	CreationDate      *time.Time          `json:"creationDate,omitempty"`
	StartDate         *time.Time          `json:"startDate,omitempty"`
	EndDate           *time.Time          `json:"endDate,omitempty"`
	Stage             *string             `json:"stage,omitempty"`
	Status            *string             `json:"status,omitempty"`
	Priority          int                 `json:"priority"`
	Parameters        IAssetParameters    `json:"parameters"`
	Callbacks         []IAssetCallback    `json:"callbacks"`
	PlaybackManifests []IPlaybackManifest `json:"playbackManifests,omitempty"`
	AssetInfo         *IAssetInfo         `json:"assetInfo,omitempty"`
}

type IAssetParameters struct {
	Input             IInputParameters          `json:"input"`
	Output            *IOutputParameters        `json:"output,omitempty"`
	PerceptualQuality *IPerceptualQualityParams `json:"perceptual-quality,omitempty"`
}

type IInputParameters struct {
	ID      *string `json:"id,omitempty"`
	UrlPath string  `json:"urlPath"`
}

type IOutputParameters struct {
	ID             *string `json:"id,omitempty"`
	Path           *string `json:"path,omitempty"`
	FileNameFormat *string `json:"fileNameFormat,omitempty"`
	FullPath       *string `json:"fullPath,omitempty"`
}

type IPerceptualQualityParams struct {
	Encoder              string                        `json:"encoder"`
	ContentPreparationID *string                       `json:"contentPreparationId,omitempty"`
	H264                 *IPerceptualOptionParams      `json:"h264,omitempty"`
	H265                 *IPerceptualOptionParams      `json:"h265,omitempty"`
	AAC                  *IPerceptualOptionAudioParams `json:"aac,omitempty"`
}

type IPerceptualOptionParams struct {
	MaxBitrate               int         `json:"maxBitrate"`
	MinBitrate               int         `json:"minBitrate"`
	MaxResolution            int         `json:"maxResolution"`
	ComplexityPeaksAwareness int         `json:"complexityPeaksAwareness"`
	MaxFPS                   interface{} `json:"maxFPS"` // can be number or string
	TargetQuality            int         `json:"targetQuality"`
}

type IPerceptualOptionAudioParams struct {
	Bitrate    *int    `json:"bitrate,omitempty"`
	Channels   *int    `json:"channels,omitempty"`
	SampleRate *int    `json:"sampleRate,omitempty"`
	Profile    *string `json:"profile,omitempty"`
}

type IPerceptualImages struct {
	ConfigFilePath string  `json:"configFilePath"`
	FileNameFormat *string `json:"fileNameFormat,omitempty"`
	FPS            int     `json:"fps"`
	Resolution     string  `json:"resolution"`
	SpriteMap      *string `json:"spriteMap,omitempty"`
}

type IEffects struct {
	FadeIn  *int `json:"fadeIn,omitempty"`
	FadeOut *int `json:"fadeOut,omitempty"`
}

type IAssetCallback struct {
	URL     string                 `json:"url"`
	Method  *AssetCallbackMethod   `json:"method,omitempty"`
	Headers []IAssetCallbackHeader `json:"headers,omitempty"`
	Source  *CallbackSource        `json:"source,omitempty"`
	Type    *CallbackType          `json:"type,omitempty"`
}

type IAssetCallbackHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AssetCallbackMethod string

const (
	GET  AssetCallbackMethod = "get"
	POST AssetCallbackMethod = "post"
	PUT  AssetCallbackMethod = "put"
)

type CallbackSource string

const (
	Lightflow CallbackSource = "lightflow"
	Encoder   CallbackSource = "encoder"
)

type CallbackType string

const (
	LightflowType CallbackType = "lightflow"
	AkamaiType    CallbackType = "akamai"
)

type IPlaybackManifest struct {
	Type  PlaybackManifestType `json:"type"`
	Codec *CodecsType          `json:"codec,omitempty"`
	URL   string               `json:"url"`
}

type PlaybackManifestType string

const (
	DASH PlaybackManifestType = "dash"
	HLS  PlaybackManifestType = "hls"
	MSS  PlaybackManifestType = "mss"
	MP3  PlaybackManifestType = "mp3"
)

type CodecsType string

const (
	X264 CodecsType = "x264"
	X265 CodecsType = "x265"
	VP9  CodecsType = "vp9"
)

type IAssetInfo struct {
	Duration        int               `json:"duration"`
	Resolution      string            `json:"resolution"`
	Bitrate         int               `json:"bitrate"`
	Codec           string            `json:"codec"`
	AudioCodec      string            `json:"audioCodec"`
	AudioBitrate    int               `json:"audioBitrate"`
	FrameRate       float64           `json:"frameRate"`
	AspectRatio     string            `json:"aspectRatio"`
	Container       string            `json:"container"`
	VideoStreams    []IVideoStream    `json:"videoStreams"`
	AudioStreams    []IAudioStream    `json:"audioStreams"`
	SubtitleStreams []ISubtitleStream `json:"subtitleStreams"`
}

type IVideoStream struct {
	Index       int     `json:"index"`
	Codec       string  `json:"codec"`
	Bitrate     int     `json:"bitrate"`
	Resolution  string  `json:"resolution"`
	FrameRate   float64 `json:"frameRate"`
	AspectRatio string  `json:"aspectRatio"`
}

type IAudioStream struct {
	Index      int    `json:"index"`
	Codec      string `json:"codec"`
	Bitrate    int    `json:"bitrate"`
	Channels   int    `json:"channels"`
	SampleRate int    `json:"sampleRate"`
}

type ISubtitleStream struct {
	Index int    `json:"index"`
	Codec string `json:"codec"`
	Lang  string `json:"lang"`
}

type IPlaybackInfo struct {
	Protection        *IProtection          `json:"protection,omitempty"`
	PlaybackManifests []IPlaybackManifest   `json:"playbackManifests,omitempty"`
	RootPath          *string               `json:"rootPath,omitempty"`
	Renditions        []IPlaybackRenditions `json:"renditions,omitempty"`
	FileNameFormat    *string               `json:"fileNameFormat,omitempty"`
}

type IProtection struct {
	Vendor    string       `json:"vendor"`
	ContentID *string      `json:"contentId,omitempty"`
	Widevine  *ILicenseURL `json:"widevine,omitempty"`
	Playready *ILicenseURL `json:"playready,omitempty"`
	Fairplay  *IFairplay   `json:"fairplay,omitempty"`
}

type ILicenseURL struct {
	LicenseURL string `json:"licenseURL"`
}

type IFairplay struct {
	LicenseURL     string  `json:"licenseURL"`
	CertificateURL *string `json:"certificateURL,omitempty"`
}

type IPlaybackImages struct {
	ConfigFilePath string `json:"configFilePath"`
	PathTemplate   string `json:"pathTemplate"`
	NumImages      int    `json:"numImages"`
	FPS            int    `json:"fps"`
	Resolution     string `json:"resolution"`
	SpriteMap      string `json:"spriteMap"`
}

type IPlaybackRenditions struct {
	Codec      string `json:"codec"`
	Resolution *int   `json:"resolution,omitempty"`
	Bitrate    int    `json:"bitrate"`
	Type       string `json:"type"`
	URL        string `json:"url"`
}
