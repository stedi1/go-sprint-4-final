package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// получаем из строки "3456,Ходьба,3h00m" данные
	str := strings.Split(data, ",")
	if len(str) != 3 {
		return 0, "", 0, fmt.Errorf("slice length not three")
	}
	// переводим строку в шаги
	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps < 1 {
		return 0, "", 0, fmt.Errorf("incorrect data")
	}
	// переводим строку в время (duration)
	duration, err := time.ParseDuration(str[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration < 1 {
		return 0, "", 0, fmt.Errorf("incorrect data")
	}

	return steps, str[1], duration, nil
}

func distance(steps int, height float64) float64 {
	stepDis := height * stepLengthCoefficient
	return (float64(steps) * stepDis) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration < 1 {
		return 0.0
	}
	d := distance(steps, height)
	return d / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duraration, err := parseTraining(data)
	if err != nil {
		fmt.Println(err)
	}
	t := duraration.Hours()
	d := distance(steps, height)
	s := meanSpeed(steps, height, duraration)

	switch activity {
	case "Ходьба":
		kcal, err := WalkingSpentCalories(steps, weight, height, duraration)
		if err != nil {
			fmt.Println(err)
		}
		str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, t, d, s, kcal)
		return str, nil
	case "Бег":
		kcal, err := RunningSpentCalories(steps, weight, height, duraration)
		if err != nil {
			fmt.Println(err)
		}
		str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, t, d, s, kcal)
		return str, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 1 || weight < 1 || height < 1 || duration < 1 {
		return 0.0, fmt.Errorf("incorrect data")
	}
	speed := meanSpeed(steps, height, duration)
	kcal := (weight * speed * duration.Minutes()) / minInH
	return kcal, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 1 || weight < 1 || height < 1 || duration < 1 {
		return 0.0, fmt.Errorf("incorrect data")
	}
	speed := meanSpeed(steps, height, duration)
	kcal := (weight * speed * duration.Minutes()) / minInH
	return kcal * walkingCaloriesCoefficient, nil
}
