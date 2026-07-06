package resolver

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolve(t *testing.T) {
	// Create a temporary file to test file resolution
	tmpDir := t.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "blink_test_*.html")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpFile.Close()

	// Create a temp subdirectory as well
	tmpSubDir, err := os.MkdirTemp(tmpDir, "blink_subdir_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name: "alias without query",
			args: []string{"yt"},
			want: "https://youtube.com",
		},
		{
			name: "alias with query",
			args: []string{"yt", "golang", "tutorial"},
			want: "https://youtube.com/results?search_query=golang+tutorial",
		},
		{
			name: "alias case insensitive",
			args: []string{"GH", "sapirrior/blink"},
			want: "https://github.com/search?q=sapirrior%2Fblink",
		},
		{
			name: "valid URL with protocol",
			args: []string{"http://example.com"},
			want: "http://example.com",
		},
		{
			name: "valid URL without protocol",
			args: []string{"example.com"},
			want: "https://example.com",
		},
		{
			name: "existing file resolution",
			args: []string{tmpFile.Name()},
			want: "file://" + filepath.ToSlash(tmpFile.Name()),
		},
		{
			name: "existing directory resolution",
			args: []string{tmpSubDir},
			want: "file://" + filepath.ToSlash(tmpSubDir),
		},
		{
			name:    "missing path resolution",
			args:    []string{"./missing_file_abc123.html"},
			wantErr: true,
		},
		{
			name:    "too many arguments without alias",
			args:    []string{"example.com", "extra_arg"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Resolve(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				// Normalize windows drive letter differences (e.g. file:///C:/ vs file:///c:/)
				if strings.ToLower(got) != strings.ToLower(tt.want) {
					t.Errorf("Resolve() = %q, want %q", got, tt.want)
				}
			}
		})
	}
}
