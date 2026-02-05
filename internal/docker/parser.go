package docker

import (
	"io"
	"strings"

	"github.com/moby/moby/api/types/container"
)

func ParseStats(statsJSON container.StatsResponse) (cpu float64, mem uint64) {
	cpuDelta := float64(statsJSON.CPUStats.CPUUsage.TotalUsage - statsJSON.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(statsJSON.CPUStats.SystemUsage - statsJSON.PreCPUStats.SystemUsage)
	cpuPercent := 0.0
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(statsJSON.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}

	memUsage := statsJSON.MemoryStats.Usage

	return cpuPercent, memUsage
}

func ParseLogs(rawLogs io.ReadCloser) []string {
	var logs []string
	buf := make([]byte, 4096)

	for {
		n, err := rawLogs.Read(buf)
		if err != nil && err != io.EOF {
			return []string{"Error reading logs"}
		}
		if n == 0 {
			break
		}

		// Handle Docker multiplexed log format
		// Format: [HEADER][PAYLOAD][HEADER][PAYLOAD]...
		pos := 0
		for pos < n {
			if pos+8 > n {
				break // Not enough data for header
			}

			// Header: stream type (1 byte) + size (4 bytes big endian) + padding (3 bytes)
			_ = buf[pos] // stream type (1=stdout, 2=stderr)
			// Parse size from next 4 bytes (big endian)
			size := int(buf[pos+4])<<24 | int(buf[pos+5])<<16 | int(buf[pos+6])<<8 | int(buf[pos+7])

			pos += 8 // Skip header

			if pos+size > n {
				break // Not enough data for payload
			}

			if size > 0 {
				logLine := string(buf[pos : pos+size])
				// Trim newline and any trailing whitespace
				logLine = strings.TrimSpace(logLine)
				if logLine != "" {
					logs = append(logs, logLine)
				}
			}

			pos += size
		}
	}

	if len(logs) == 0 {
		return []string{"No logs available"}
	}

	return logs
}
