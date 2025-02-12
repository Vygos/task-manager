package pagination

type Page struct {
	Data          any `json:"data"`
	Page          int `json:"page"`
	Size          int `json:"page_size"`
	TotalElements int `json:"total_elements"`
}

func NewPage(page, size, defaultSize int) Page {
	if size == 0 {
		size = defaultSize
	}

	if page == 0 {
		page = 1
	}

	return Page{
		Page: page,
		Size: size,
	}
}
