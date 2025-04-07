package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
)

type ApiResponse struct {
	Next    string  `json:"next"`
	Results []Event `json:"results"`
}

type Event struct {
	gorm.Model
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Place          *Place `json:"place" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PlaceID        *uint
	Location       string        `json:"location" `
	AgeRestriction int           `json:"-"`
	Price          int           `json:"-"`
	IsFree         bool          `json:"is_free"`
	Dates          []DateRange   `json:"dates" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Participants   []Participant `json:"participants" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categories     []Category    `json:"-" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Структура для хранения ответа по категориям
type RawCategories struct {
	Categories []string `json:"categories"`
}

// Новая модель для категорий
type Category struct {
	gorm.Model
	EventID uint   `json:"event_id"`             // Внешний ключ на Event
	Name    string `json:"name" gorm:"not null"` // Название категории
}

type DateRange struct {
	gorm.Model
	EventID uint  `json:"-"`
	Start   int64 `json:"start"`
	End     int64 `json:"end"`
}

type Place struct {
	gorm.Model
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	CoordsLat float64 `json:"coords_lat"`
	CoordsLon float64 `json:"coords_lon"`
	Location  string  `json:"location"`
}
type RawPlace struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Coords struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coords"`
	Location string `json:"location"`
}

type Location struct {
	Slug string `json:"slug"`
}

type Participant struct {
	gorm.Model
	EventID uint   `json:"-"`
	Role    string `json:"slug"`
	Title   string `json:"title"`
	Type    string `json:"agent_type"`
}

type RawParticipant struct {
	Role struct {
		Slug string `json:"slug"`
	} `json:"role"`
	Agent struct {
		Title string `json:"title"`
		Type  string `json:"agent_type"`
	} `json:"agent"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	temp := struct {
		Price          string      `json:"price"`
		AgeRestriction interface{} `json:"age_restriction"`
		Location       `json:"location"`
		*RawPlace      `json:"place"`
		RP             []RawParticipant `json:"participants"`
		*Alias
	}{Alias: (*Alias)(e)}

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	//Для location
	e.Location = temp.Location.Slug

	//Для price
	if val, err := extractFirstNumber(temp.Price); err == nil {
		e.Price = val
	}

	//Для place
	if temp.RawPlace != nil {
		e.Place = &Place{
			ID:        temp.RawPlace.ID,
			Title:     temp.RawPlace.Title,
			Location:  temp.RawPlace.Location,
			CoordsLat: temp.RawPlace.Coords.Lat,
			CoordsLon: temp.RawPlace.Coords.Lon,
		}
	}

	//Для age restriction
	if temp.AgeRestriction == nil {
		e.AgeRestriction = 0
	} else {
		switch v := temp.AgeRestriction.(type) {
		case float64:
			e.AgeRestriction = int(v)
		case string:
			digits := strings.TrimRight(v, "+")
			e.AgeRestriction, err = strconv.Atoi(digits)
		}
	}
	//Для participants
	e.Participants = make([]Participant, len(temp.RP))
	for i := range temp.RP {
		e.Participants[i].Role = temp.RP[i].Role.Slug
		e.Participants[i].Title = temp.RP[i].Agent.Title
		e.Participants[i].Type = temp.RP[i].Agent.Type
	}
	return err
}

func extractFirstNumber(text string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(text)
	if match == "" {
		return 0, fmt.Errorf("число не найдено")
	}

	number, err := strconv.Atoi(match)
	if err != nil {
		return 0, fmt.Errorf("ошибка преобразования числа: %v", err)
	}

	return number, nil
}
