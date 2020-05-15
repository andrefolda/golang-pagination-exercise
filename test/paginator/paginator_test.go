package paginator_test

import (
	"pagination/paginator"
	"strings"
	"testing"
)

func TestRegularPaginationMiddlePage(t *testing.T) {
	controlPaginator := "1 2 3 4 ... 8 9 10 11 12 ... 17 18 19 20"
	p := paginator.NewPaginator(10, 20, 4, 2)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestRegularPaginationLastPage(t *testing.T) {
	controlPaginator := "1 2 ... 9 10 11 12"
	p := paginator.NewPaginator(12, 12, 2, 3)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestRegularPaginationFirstPage(t *testing.T) {
	controlPaginator := "1 2 ... 8"
	p := paginator.NewPaginator(1, 8, 1, 1)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestRegularPaginationCloseToEndingGroup(t *testing.T) {
	controlPaginator := "1 2 3 4 5 6 ... 39 40 41 42 43 44 45 46 47 48 49 50 51"
	p := paginator.NewPaginator(42, 51, 6, 3)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestRegularPaginationCloseToStartingGroup(t *testing.T) {
	controlPaginator := "1 2 3 4 5 6 ... 10 11 12"
	p := paginator.NewPaginator(4, 12, 3, 2)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestMinimalPaginationMiddlePage(t *testing.T) {
	controlPaginator := "... 12 ..."
	p := paginator.NewPaginator(12, 23, 0, 0)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestMinimalPaginationFirstPage(t *testing.T) {
	controlPaginator := "1 ..."
	p := paginator.NewPaginator(1, 7, 0, 0)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestMinimalPaginationLastPage(t *testing.T) {
	controlPaginator := "... 54"
	p := paginator.NewPaginator(54, 54, 0, 0)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestCompletePaginationFirstScenario(t *testing.T) {
	controlPaginator := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22"
	p := paginator.NewPaginator(4, 22, 22, 0)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestCompletePaginationSecondScenario(t *testing.T) {
	controlPaginator := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15"
	p := paginator.NewPaginator(11, 15, 0, 15)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestPaginationWithBrokenParameters(t *testing.T) {
	controlPaginator := "1"
	p := paginator.NewPaginator(-2, -14, -23, -45)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}

func TestEnormousPagination(t *testing.T) {
	controlPaginator := "1 2 3 4 5 ... 254859 254860 254861 254862 254863 254864 254865 254866 254867 ... 965845261 965845262 965845263 965845264 965845265"
	p := paginator.NewPaginator(254863, 965845265, 5, 4)

	if !strings.EqualFold(controlPaginator, p.GetPagination().ToString()) {
		t.Errorf("Expected \"%v\" pagination, got \"%v\"", controlPaginator, p.GetPagination().ToString())
	}
}
