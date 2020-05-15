package paginator

import (
	"fmt"
	"strconv"
	"strings"
)

type Paginator struct {

	// actual page
	currentPage int

	// total available pages
	totalPages int

	// how many pages we want to link in the beginning and in the end
	boundaries int

	// how many pages we want to link before and after the actual page
	around int

	pagination pagination
}

type pagination []string

func (p pagination) ToString() string {
	return strings.Join(p, " ")
}

func NewPaginator(currentPage int, totalPages int, boundaries int, around int) Paginator {

	if currentPage < 1 {
		currentPage = 1
	}
	if totalPages < 1 {
		totalPages = 1
	}
	if boundaries < 0 {
		boundaries = 0
	}
	if around < 0 {
		around = 0
	}

	p := Paginator{
		currentPage: currentPage,
		totalPages: totalPages,
		boundaries: boundaries,
		around: around,
	}

	p.generatePagination()

	return p
}

func (p Paginator) GetPagination() pagination {
	return p.pagination
}

func (p *Paginator) generatePagination() {

	currentPageGroupStart := p.currentPage - p.around
	currentPageGroupEnd := p.currentPage + p.around
	hasReachedPageGroup := false

	endingGroupStart := p.totalPages - p.boundaries
	hasReachedEndingGroup := false

	page := 1
	for page <= p.totalPages {

		shouldPrint := false

		if page >= currentPageGroupStart && page <= currentPageGroupEnd {
			shouldPrint = true
			hasReachedPageGroup = true

		} else if page <= p.boundaries && p.boundaries > 0 {
			shouldPrint = true

		} else if page > endingGroupStart && page <= p.totalPages {
			shouldPrint = true
			hasReachedEndingGroup = true
		}

		if shouldPrint {
			p.pagination = append(p.pagination, strconv.Itoa(page))
			page++

		} else {
			if !hasReachedPageGroup {
				p.pagination = append(p.pagination, "...")
				page = p.currentPage - p.around

			} else if !hasReachedEndingGroup {
				p.pagination = append(p.pagination, "...")
				page = p.totalPages - (p.boundaries - 1)
			}
		}
	}
}

func (p Paginator) PrintPagination() {
	fmt.Println(p.pagination.ToString())
}
