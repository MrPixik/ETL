package static

import "fmt"

const (
	TotalRequests = 1000
	FirstUrl      = "https://kudago.com/public-api/v1.4/events/?actual_since=&actual_until=&categories=&expand=place&fields=id%2Ctitle%2Cplace%2Cis_free%2Clocation%2Cdates%2Cage_restriction%2Cprice%2Cparticipants&ids=&is_free=&lang=&lat=&location=&lon=&order_by=&radius=&text_format="
)

func GetCategoriesUrl(eventID uint) string {
	return fmt.Sprintf("https://kudago.com/public-api/v1.4/events/%d/?lang=&fields=categories&expand=", eventID)
}
