package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// получаем из data количество шагов и время прогулки
	str := strings.Split(data, ",")
	if len(str) != 2 {
		return 0, 0, fmt.Errorf("slice length not two")
	}
	// переводим строку в шаги
	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps less than one")
	}
	// переводим строку в время (duration)
	duration, err := time.ParseDuration(str[1])
	if err != nil {
		return 0, 0, err
	}
	if duration < 1 {
		return 0, 0, fmt.Errorf("incorrect duration")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	kcal, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, kcal)

	return str
}
