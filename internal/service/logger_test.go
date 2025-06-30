package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestMonitorAndLog(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	events := make(chan model.BusinessEntity, 4)

	// Запускаем логгер
	go MonitorAndLog(ctx, events)

	// Создаём тестовые сущности
	v := &vehicle.Vehicle{}
	c := &car.Car{}
	m := &moto.Moto{}
	e := &expense.Expense{}

	output := captureOutput(func() {
		events <- v
		events <- c
		events <- m
		events <- e
		time.Sleep(100 * time.Millisecond)
		cancel()
		time.Sleep(50 * time.Millisecond)
	})

	if !strings.Contains(output, "New vehicle") {
		t.Error("лог для vehicle не найден")
	}
	if !strings.Contains(output, "New car") {
		t.Error("лог для car не найден")
	}
	if !strings.Contains(output, "New moto") {
		t.Error("лог для moto не найден")
	}
	if !strings.Contains(output, "New expense") {
		t.Error("лог для expense не найден")
	}
	if !strings.Contains(output, "Monitor stopped") {
		t.Error("лог о завершении не найден")
	}
}
