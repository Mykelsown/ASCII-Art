package methodsandtests

// NewBannerLoader creates and loads the banner file once
func NewBannerLoader() (*BannerLoader, error) {
	charMap, err := LoadBannerMap()
	if err != nil {
		return nil, err
	}

	return &BannerLoader{
		charMap: charMap,
	}, nil
}
