package sysapp

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func runApp(t *testing.T, dir string, args ...string) (int, string, string) {
	t.Helper()

	var stdout, stderr bytes.Buffer
	code := New(Options{
		Dir:    dir,
		Stdout: &stdout,
		Stderr: &stderr,
	}).Run(args)

	return code, stdout.String(), stderr.String()
}

func TestInitScaffoldsProjectAndIsIdempotent(t *testing.T) {
	root := t.TempDir()

	code, out, errOut := runApp(t, root, "init")
	if code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	wantFiles := []string{
		".sys-orchestrator/state.json",
		".sys-orchestrator/freeze.json",
		".sys-orchestrator/allowlists.json",
		"system/architecture/system.md",
		"system/contracts/api.yaml",
		"system/contracts/events.asyncapi.yaml",
		"system/contracts/auth.md",
		"system/modules/frontend.md",
		"system/modules/backend.md",
		"system/data/schema.sql",
		"system/data/schema.md",
		"system/data/db/indexes.md",
		"system/obs/dashboards/grafana.md",
	}

	for _, rel := range wantFiles {
		if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
			t.Fatalf("expected %s to exist: %v", rel, err)
		}
	}

	code, out, errOut = runApp(t, root, "init")
	if code != 0 {
		t.Fatalf("second init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	if !strings.Contains(out, "already initialized") {
		t.Fatalf("second init should report already initialized, got %q", out)
	}
}

func TestRootDiscoveryAndStatusJSON(t *testing.T) {
	root := t.TempDir()
	if code, out, errOut := runApp(t, root, "init"); code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	frontendDir := filepath.Join(root, "frontend", "app")
	if err := os.MkdirAll(frontendDir, 0o755); err != nil {
		t.Fatal(err)
	}

	code, out, errOut := runApp(t, frontendDir, "status", "--json")
	if code != 0 {
		t.Fatalf("status json failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	var status Status
	if err := json.Unmarshal([]byte(out), &status); err != nil {
		t.Fatalf("status output is not json: %v\n%s", err, out)
	}
	if status.Root != root {
		t.Fatalf("root = %q, want %q", status.Root, root)
	}
	if status.Phase != PhaseDesign {
		t.Fatalf("phase = %q, want %q", status.Phase, PhaseDesign)
	}
	if status.Role != RoleFrontend {
		t.Fatalf("role = %q, want %q", status.Role, RoleFrontend)
	}
}

func TestValidateReportsMissingRequiredSystemFile(t *testing.T) {
	root := t.TempDir()
	if code, out, errOut := runApp(t, root, "init"); code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	if err := os.Remove(filepath.Join(root, "system", "contracts", "api.yaml")); err != nil {
		t.Fatal(err)
	}

	code, out, errOut := runApp(t, root, "validate")
	if code == 0 {
		t.Fatalf("validate should fail when required file is missing: stdout=%q stderr=%q", out, errOut)
	}
	if !strings.Contains(out+errOut, "system/contracts/api.yaml") {
		t.Fatalf("missing file warning not found in output: stdout=%q stderr=%q", out, errOut)
	}
}

func TestDesignFreezeRecordsBaselineAndCaptureBlocksInBuild(t *testing.T) {
	root := t.TempDir()
	if code, out, errOut := runApp(t, root, "init"); code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	code, out, errOut := runApp(t, root, "design", "freeze")
	if code != 0 {
		t.Fatalf("design freeze failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	code, out, errOut = runApp(t, root, "capture")
	if code == 0 {
		t.Fatalf("capture should fail during build phase: stdout=%q stderr=%q", out, errOut)
	}
	if !strings.Contains(out+errOut, "design-change") {
		t.Fatalf("capture output should mention design-change: stdout=%q stderr=%q", out, errOut)
	}

	archPath := filepath.Join(root, "system", "architecture", "system.md")
	if err := os.WriteFile(archPath, []byte("changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, errOut = runApp(t, root, "status", "--json")
	if code != 0 {
		t.Fatalf("status json failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	var status Status
	if err := json.Unmarshal([]byte(out), &status); err != nil {
		t.Fatal(err)
	}
	if len(status.Validation.Warnings) == 0 {
		t.Fatalf("expected freeze warning after architecture mutation: %#v", status)
	}
	if !strings.Contains(out, "design-change") {
		t.Fatalf("status should mention design-change after frozen file changes: %s", out)
	}
}

func TestDesignCommandsDoNotRequireOpenSpec(t *testing.T) {
	root := t.TempDir()
	if code, out, errOut := runApp(t, root, "init"); code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	code, out, errOut := runApp(t, root, "explore", "auth")
	if code != 0 {
		t.Fatalf("explore failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	if !strings.Contains(out, "auth") || strings.Contains(out, "openspec new") {
		t.Fatalf("explore output did not look like design guidance: %q", out)
	}

	code, out, errOut = runApp(t, root, "capture")
	if code != 0 {
		t.Fatalf("capture failed in design phase: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	if !strings.Contains(out, "decision record") {
		t.Fatalf("capture output should mention decision records: %q", out)
	}
}

func TestAgentInstallersGenerateExpectedFilesAndPreserveClaudeContent(t *testing.T) {
	root := t.TempDir()
	if code, out, errOut := runApp(t, root, "init"); code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	claudePath := filepath.Join(root, "CLAUDE.md")
	if err := os.WriteFile(claudePath, []byte("# Existing\n\nKeep this.\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	for _, agent := range []string{"codex", "cursor", "claude"} {
		code, out, errOut := runApp(t, root, "agent", "install", agent)
		if code != 0 {
			t.Fatalf("agent install %s failed: code=%d stdout=%q stderr=%q", agent, code, out, errOut)
		}
	}

	wantFiles := []string{
		".codex/skills/sys-explore/SKILL.md",
		".codex/skills/sys-capture/SKILL.md",
		".codex/skills/sys-apply/SKILL.md",
		".codex/skills/sys-design-change/SKILL.md",
		".cursor/rules/sys-orchestrator.mdc",
		"CLAUDE.md",
	}
	for _, rel := range wantFiles {
		if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
			t.Fatalf("expected %s to exist: %v", rel, err)
		}
	}

	claude, err := os.ReadFile(claudePath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(claude), "Keep this.") || !strings.Contains(string(claude), "SYS-ORCHESTRATOR") {
		t.Fatalf("CLAUDE.md did not preserve content and add marked section:\n%s", claude)
	}
}

func TestBuildWorkflowUsesFakeOpenSpecInBuildPhase(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell-script fake executable is POSIX-only")
	}

	root := t.TempDir()
	if code, out, errOut := runApp(t, root, "init"); code != 0 {
		t.Fatalf("init failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	code, out, errOut := runApp(t, root, "change", "propose", "add-login")
	if code == 0 {
		t.Fatalf("change propose should fail before build phase: stdout=%q stderr=%q", out, errOut)
	}

	if code, out, errOut := runApp(t, root, "design", "freeze"); code != 0 {
		t.Fatalf("freeze failed: code=%d stdout=%q stderr=%q", code, out, errOut)
	}

	logPath := filepath.Join(root, "openspec.log")
	fake := filepath.Join(root, "fake-openspec")
	script := "#!/bin/sh\n" +
		"echo \"$@\" >> " + shellQuote(logPath) + "\n" +
		"if [ \"$1\" = \"new\" ]; then mkdir -p " + shellQuote(filepath.Join(root, "openspec", "changes", "add-login")) + "; fi\n" +
		"exit 0\n"
	if err := os.WriteFile(fake, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}

	var stdout, stderr bytes.Buffer
	code = New(Options{Dir: root, Stdout: &stdout, Stderr: &stderr, OpenSpecPath: fake}).Run([]string{"change", "propose", "add-login"})
	if code != 0 {
		t.Fatalf("change propose failed: code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}

	stdout.Reset()
	stderr.Reset()
	code = New(Options{Dir: root, Stdout: &stdout, Stderr: &stderr, OpenSpecPath: fake}).Run([]string{"change", "archive", "add-login"})
	if code != 0 {
		t.Fatalf("change archive failed: code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "new change add-login") || !strings.Contains(log, "archive add-login") {
		t.Fatalf("fake openspec did not receive expected calls:\n%s", log)
	}
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}
