package lightflow

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// APIClient represents the REST API client.
type APIClient struct {
	BaseURL   string
	AuthToken string
}

// NewAPIClient creates a new API client with the given base URL and auth token.
func NewAPIClient(baseURL, authToken string) *APIClient {
	return &APIClient{
		BaseURL:   baseURL,
		AuthToken: authToken,
	}
}

// DoRequest performs an HTTP request with the given method, endpoint, and body.
func (client *APIClient) DoRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := client.BaseURL + endpoint

	var req *http.Request
	var err error
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		println(string(jsonBody[:]))
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.AuthToken)
	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, errors.New("request failed with status code " + resp.Status)
	}
	return resp, nil
}

func (client *APIClient) GetInputOutputs() (*IWorkflowStoragePaginatedResponse, error) {
	resp, err := client.DoRequest("GET", "/inputs-outputs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get input-outputs, status code: " + resp.Status)
	}

	var result IWorkflowStoragePaginatedResponse
	println(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateInputOutput creates a new input-output.
func (client *APIClient) CreateInputOutput(inputOutput IWorkflowStorage) (*IWorkflowStorageResponse, error) {
	resp, err := client.DoRequest("POST", "/inputs-outputs", inputOutput)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to create input-output, status code: " + resp.Status)
	}

	var result IWorkflowStorageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteStorage deletes a storage given a UUID.
func (client *APIClient) DeleteInputOutput(uuid string) error {
	resp, err := client.DoRequest("DELETE", "/inputs-outputs/"+uuid, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to delete storage, status code: " + resp.Status)
	}

	return nil
}

// ListAssets requests the list of assets.
func (client *APIClient) GetAssets() (*IAssetPaginatedResponse, error) {
	resp, err := client.DoRequest("GET", "/assets", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get assets, status code: " + resp.Status)
	}

	var result IAssetPaginatedResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAsset requests an asset by UUID.
func (client *APIClient) GetAsset(uuid string) (*IAsset, error) {
	resp, err := client.DoRequest("GET", "/assets/"+uuid, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get asset, status code: " + resp.Status)
	}

	var result IAsset
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAssetPlayback requests playback information for an asset by UUID.
func (client *APIClient) GetAssetPlayback(uuid string) (*IPlaybackInfo, error) {
	resp, err := client.DoRequest("GET", "/assets/"+uuid+"/playback", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get asset playback information, status code: " + resp.Status)
	}

	var result IPlaybackInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateAsset creates a new asset.
func (client *APIClient) CreateAsset(asset IAsset) (*IAsset, error) {
	resp, err := client.DoRequest("POST", "/assets", asset)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to create asset, status code: " + resp.Status)
	}

	var result IAsset
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *APIClient) CreateBasicAsset(url string, maxBitrate int, targetQuality int) (*IAsset, error) {
	new_asset := IAsset{}
	new_asset.Parameters.Input = IInputParameters{}
	new_asset.Parameters.Input.UrlPath = url
	new_asset.Parameters.PerceptualQuality = &IPerceptualQualityParams{
		Encoder: "none",
		H264: &IPerceptualOptionParams{
			MaxBitrate:               maxBitrate,
			MinBitrate:               500,
			MaxResolution:            1080,
			ComplexityPeaksAwareness: 1,
			MaxFPS:                   "auto",
			TargetQuality:            targetQuality,
		},
	}

	return client.CreateAsset(new_asset)
}

// DeleteAsset deletes an asset given a UUID.
func (client *APIClient) DeleteAsset(uuid string) error {
	resp, err := client.DoRequest("DELETE", "/assets/"+uuid, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to delete asset, status code: " + resp.Status)
	}

	return nil
}
