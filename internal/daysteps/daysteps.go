package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	cal "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("кол-во элементов не совпадает с ожиданиями")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования в целое значение")
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("кол-во шагов неверное")
	}
	durationSteps, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования в длительность")
	}
	if durationSteps <= 0 {
		return 0, 0, fmt.Errorf("неверная длительность")
	}
	return steps, durationSteps, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	stepsCount, durationOfSteps, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if stepsCount <= 0 {
		return ""
	}
	distance := float64(stepsCount) * stepLength
	distanceInKm := float64(distance) / float64(mInKm)
	caloriesSpent, err := cal.WalkingSpentCalories(stepsCount, weight, height, durationOfSteps)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepsCount, distanceInKm, caloriesSpent)
}
