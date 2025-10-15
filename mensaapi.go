package mensaapi

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

// Returns the name of the tip or empty string if not found
// Example: S -> Schwein
func GetTipDescription(tip string) string {
	fullTip, ok := TipsMap[tip]
	if !ok {
		return ""
	}
	return fullTip
}

// Returns all menu items for the given date at the specified location
func GetDayByLocation(loc LOCATION, date time.Time) (Day, error) {

	// calculate dates
	dateS := ConvertDateToString(date)
	startThisWeek := date.Add(time.Hour * 24 * time.Duration(1-int(date.Weekday())))
	startNextWeek := startThisWeek.Add(time.Hour * 24 * 7)

	body := "func=make_spl&locId=" + fmt.Sprint(loc) + "&date=" + dateS + "&lang=de&startThisWeek=" + ConvertDateToString(startThisWeek) + "&startNextWeek=" + ConvertDateToString(startNextWeek)

	resp, err := soup.Post("https://sw-ulm-spl51.maxmanager.xyz/inc/ajax-php_konnektor.inc.php", "application/x-www-form-urlencoded; charset=UTF-8", body)
	if err != nil {
		os.Exit(1)
	}

	categories := map[string][]soup.Root{}
	last := ""

	// extract the categories and menu items
	doc := soup.HTMLParse(resp)
	for _, cat := range doc.Find("div", "class", "container-fluid").Children() {
		if strings.Contains(cat.Attrs()["class"], "gruppenkopf") {
			last = cat.Find("div", "class", "gruppenname").Text()
		} else {
			categories[last] = append(categories[last], cat)
		}
	}

	// transform menu items into Item struct
	categoriesM := map[string][]Item{}
	for cat, fItems := range categories {
		for _, fItem := range fItems {
			item, err := extractItem(cat, dateS, fItem)
			if err != nil {
				return Day{}, err
			}
			categoriesM[cat] = append(categoriesM[cat], item)
		}
	}

	return Day{
		Date:       date,
		Categories: categoriesM,
	}, nil
}

func extractItem(cat string, date string, html soup.Root) (Item, error) {

	// extract name
	name := panicR(func() string {
		name := ""
		for _, c := range html.Children()[3].Children()[1].Children()[3].Children() {
			if len(c.Attrs()) == 0 {
				name += strings.TrimRight(c.HTML(), " ")
			}
		}
		return name
	})

	// extract info
	info := panicR(func() string {
		return html.Children()[3].Children()[1].Children()[3].Find("div", "class", "cc").Text()
	})

	// extract prices
	prices := panicR(func() Prices {
		// formating prices and converting to float64
		nah := html.Children()[1].Children()[3].Children()[3].Children()[1].Children()[1].HTML()
		nah = cleanString(nah)
		pricesList := strings.Split(strings.ReplaceAll(nah, ",", "."), "|")

		ps, err := strconv.ParseFloat(pricesList[0], 64)
		if err != nil {
			return Prices{
				PriceStudent: 0.0,
				PriceWorker:  0.0,
				PriceGuest:   0.0,
			}
		}
		pw, err := strconv.ParseFloat(pricesList[1], 64)
		if err != nil {
			return Prices{
				PriceStudent: 0.0,
				PriceWorker:  0.0,
				PriceGuest:   0.0,
			}
		}
		pg, err := strconv.ParseFloat(pricesList[2], 64)
		if err != nil {
			return Prices{
				PriceStudent: 0.0,
				PriceWorker:  0.0,
				PriceGuest:   0.0,
			}
		}

		prices := Prices{
			PriceStudent: ps,
			PriceWorker:  pw,
			PriceGuest:   pg,
		}

		return prices
	})

	// extract nutrition
	nutrition := panicR(func() Nutrition {
		nutrition := Nutrition{}
		parent := html.Children()[3].Children()[1].Children()[3].Find("div", "class", "azn")
		for i, row := range parent.Find("table").Find("tbody").FindAll("tr") {
			main := cleanString(row.Children()[1].Text())

			parseF := func(ftp string) float64 {
				res, err := strconv.ParseFloat(strings.ReplaceAll(ftp, ",", "."), 64)
				if err != nil {
					panic("failed to convert")
				}
				return res
			}

			switch i {
			case 0:
				nutrition.Cal = parseF(strings.Split(main, "/")[1])

			case 1:
				nutrition.Protein = parseF(main)
			case 2:
				nutrition.Fat = parseF(main)
				nutrition.PartSaturated = parseF(cleanString(row.Children()[1].Children()[1].Text()))
			case 3:
				nutrition.Carbs = parseF(main)
				nutrition.PartSuguar = parseF(cleanString(row.Children()[1].Children()[1].Text()))
			case 4:
				nutrition.Salt = parseF(main)
			}
		}

		return nutrition
	})

	item := Item{
		Category:  cat,
		Name:      name,
		Info:      info,
		Prices:    prices,
		Date:      ConvertStringToDate(date),
		Additives: extractAdatives(html.Attrs()["lang"]), // extract addatives
		Nutrition: nutrition,
	}

	return item, nil
}

func extractAdatives(adds string) Additives {
	tips := []string{}
	allergenes := []string{}
	for _, add := range strings.Split(adds, ",") {
		_, ok := TipsMap[add]
		if ok {
			tips = append(tips, add)
		} else {
			allergenes = append(allergenes, add)
		}
	}

	additives := Additives{
		Allergenes: allergenes,
		Tips:       tips,
	}
	return additives
}

func ConvertStringToDate(date string) time.Time {
	r, err := time.Parse("2006-01-02", date)
	if err != nil {
		r = time.Now()
	}
	return r
}

func ConvertDateToString(date time.Time) string {
	r := date.Format("2006-01-02")
	return r
}

func panicR[T any](f func() T) T {
	defer func() {
		if r := recover(); r != nil {
			// recovered
			//fmt.Println("recovered", r) // enable for testing
		}
	}()
	return f()
}

func cleanString(toClean string) string {
	allowed := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',', '|', '/'}
	return strings.Map(func(r rune) rune {
		if slices.Contains(allowed, r) {
			return r
		}
		return -1
	}, toClean)
}
