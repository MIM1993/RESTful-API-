package sd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//health check show 'ok' as the ping-pong result
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

//DiskCheck check disk usage
func DeskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGb := int(u.Total) / GB
	//percent
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGb)

	//response
	c.String(status, "\n"+message)
}

//CpuCheck check cpu usage
func CpuCheck(c *gin.Context) {
	cores, err := cpu.Counts(false)
	if err != nil {
		log.Println("cpu.Counts err:", err)
	}

	avg, _ := load.Avg()
	l1 := avg.Load1
	l5 := avg.Load5
	l15 := avg.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %2f, %2f ,%2f | cores: %d", text, l1, l5, l15, cores)

	c.String(status, "\n"+message)
}

//RAMCheck check RAM usage
func RAMCheck(c *gin.Context) {
	vm, _ := mem.VirtualMemory()

	usedMB := int(vm.Used) / MB
	usedGB := int(vm.Used) / GB
	totalMB := int(vm.Total) / MB
	totalGb := int(vm.Total) / GB
	//percent
	usedPercent := int(vm.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNNING"
	}

	message := fmt.Sprintf("%s -Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGb, usedPercent)

	c.String(status, "\n"+message)
}
