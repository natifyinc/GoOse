package goose

// Goose is the main entry point of the program
type Goose struct {
	config configuration
}

// New returns a new instance of the article extractor
func New(args ...string) Goose {
	return Goose{
		config: GetDefaultConfiguration(args...),
	}
}

// ExtractFromURL follows the URL, fetches the HTML page and returns an article object
func (g Goose) ExtractFromURL(url string) *Article {
	cc := NewCrawler(g.config, url, "")
	return cc.Crawl()
}

// ExtractFromRawHTML returns an article object from the raw HTML content
func (g Goose) ExtractFromRawHTML(url string, RawHTML string) *Article {
	cc := NewCrawler(g.config, url, RawHTML)
	return cc.Crawl()
}
