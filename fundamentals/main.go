package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	// variables()
	// constants()
	// for_loop()
	// if_else()
	// switch_golang()
	// Arays()
	// Sices_example()
	// Maps_golang()
	// function_call_golang()
	// multiple_return_value_golang()
	// arrays()
	// slices_example()
	// maps_golang()
	// vardic_function_golang()
	// closures_golang()
	// recursion_golang()
	// range_over_builtin_types_golang()
	// pointers_golng()
	// string_and_rums()
	// struct_golang()
	// methods_with_struct()
	// interfaces_golang()
	// enum_golang()
	// struct_embeding_golang()
	// generics_golang()
	// range_over_iteration_golang()
	// Errorsgolang()
	// custom_error_golang()
	// coroutine_golang()
	// channels_golang()
	// channel_buffering_golang()
	// channel_sync_golang()
	// channel_directions_golang()
	// channel_select_golang()
	// channel_timeout_golang()
	// channel_non_blocking_golang()
	// channel_closing_golang()
	// channel_range_golang()
	// channel_timer_golang()
	// channels_tickers_golang()
	// worker_pool_golang()
	// channels_rate_limiting_golang()
	// waitgroup_golang()
	// atomic_counter_golang()
	// mutex_golang()
	// stateful_goroutine_golang()
	// slices_sorting_golang()
	// sorting_by_function_golang()
	// panic_golang()
	// defer_golang()
	// recover_golang()
	// string_function_golang()
	// string_formatting_golang()
	text_template_golang()

	PrintMemUsage()
	elapsed := time.Since(start)
	fmt.Println("total execution time :", elapsed)
}
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
