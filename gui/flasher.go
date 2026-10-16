package gui

import (
	"fmt"
	"os/exec"
	"strings"
)

type FlashConfig struct {
	HexFile string
	Port    string
	FQBN    string
}

func Flash(cfg FlashConfig, log func(string)) error {
	if cfg.HexFile == "" {
		return fmt.Errorf("hex file is not selected")
	}
	if cfg.Port == "" {
		return fmt.Errorf("serial port is not selected")
	}
	if cfg.FQBN == "" {
		return fmt.Errorf("board (FQBN) is not selected")
	}

	args := []string{
		"upload",
		"-p", cfg.Port,
		"-b", cfg.FQBN,
		"--input-file", cfg.HexFile,
	}

	log(fmt.Sprintf("$ arduino-cli %s", strings.Join(args, " ")))

	cmd := exec.Command("arduino-cli", args...)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		log(strings.TrimRight(string(out), "\n"))
	}
	if err != nil {
		return fmt.Errorf("flash failed: %w", err)
	}
	log("[OK] Flash completed successfully")
	return nil
}

// ListPorts возвращает список портов через arduino-cli board list --format json.
func ListPorts() ([]string, error) {
	cmd := exec.Command("arduino-cli", "board", "list", "--format", "json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Простой парс: ищем "address":"COM3" в JSON-выводе.
	var ports []string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		idx := strings.Index(line, `"address"`)
		if idx < 0 {
			continue
		}
		rest := line[idx:]
		colon := strings.Index(rest, ":")
		if colon < 0 {
			continue
		}
		rest = rest[colon+1:]
		rest = strings.TrimSpace(rest)
		rest = strings.Trim(rest, `",`)
		if rest != "" && rest != "unknown" {
			ports = append(ports, rest)
		}
	}
	return ports, nil
}
