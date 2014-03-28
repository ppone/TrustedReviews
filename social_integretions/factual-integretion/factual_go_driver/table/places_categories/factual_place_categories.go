package places_categories

const PAGE_SIZE_LIMIT = 50
const ROW_LIMIT = 500

type placesCategoriesData struct {
	category_id int
	parents     int
	en          string
	de          string
	es          string
	fr          string
	it          string
	jp          string
	kr          string
	zh          string
	zh_hant     string
	pt          string
	abstract    bool
}
