package spentcalories

import (
	"fmt"
	"log"
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
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("вкол-во элементов не совпадает с ожиданиями")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования в целое число")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("кол-во шагов неверное")
	}
	durationSteps, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования в длительность")
	}
	if durationSteps <= 0 {
		return 0, "", 0, fmt.Errorf("неверная длительность")
	}
	return steps, parts[1], durationSteps, nil
}

func distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	return (float64(steps) * stepLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceInKm := distance(steps, height)
	return distanceInKm / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeOfTraining, durationOfTraining, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	distanceOftraining := distance(steps, height)
	speed := meanSpeed(steps, height, durationOfTraining)
	switch typeOfTraining {
	case "Ходьба":
		spentCalor, err := WalkingSpentCalories(steps, weight, height, durationOfTraining)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfTraining, durationOfTraining.Hours(), distanceOftraining, speed, spentCalor), nil
	case "Бег":
		spentCalor, err := RunningSpentCalories(steps, weight, height, durationOfTraining)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfTraining, durationOfTraining.Hours(), distanceOftraining, speed, spentCalor), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("кол-во шагов неверное")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("данные о весе неверны")
	}
	if height <= 0 {
		return 0, fmt.Errorf("кол-во шагов неверное")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверная длительность")
	}
	srSpeed := meanSpeed(steps, height, duration)
	durationInMins := duration.Minutes()
	return (weight * srSpeed * durationInMins) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("кол-во шагов неверное")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("данные о весе неверны")
	}
	if height <= 0 {
		return 0, fmt.Errorf("кол-во шагов неверное")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверная длительность")
	}
	srSpeed := meanSpeed(steps, height, duration)
	durationInMins := duration.Minutes()
	return ((weight * srSpeed * durationInMins) / minInH) * walkingCaloriesCoefficient, nil
}
