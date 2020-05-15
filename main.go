package main

import (
	"flag"
	"pagination/paginator"
)

func main() {

	currentPage := *flag.Int("current_page", 1, "actual page")
	totalPages := *flag.Int("total_pages", 20, "total available pages")
	boundaries := *flag.Int("boundaries", 3, "how many pages we want to link in the beginning and in the end")
	around := *flag.Int("around", 2, "how many pages we want to link before and after the actual page")

	p := paginator.NewPaginator(currentPage, totalPages, boundaries, around)

	p.PrintPagination()
}
