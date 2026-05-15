package service

import (
	"context"
	"errors"
	"testing"
	"time"
)

type updateTestCache struct {
	data string
}

func (c *updateTestCache) GetUpdateInfo(context.Context) (string, error) {
	if c.data == "" {
		return "", errors.New("cache miss")
	}
	return c.data, nil
}

func (c *updateTestCache) SetUpdateInfo(_ context.Context, data string, _ time.Duration) error {
	c.data = data
	return nil
}

type updateTestGitHubClient struct {
	release *GitHubRelease
	err     error
}

func (c updateTestGitHubClient) FetchLatestRelease(context.Context, string) (*GitHubRelease, error) {
	if c.err != nil {
		return nil, c.err
	}
	return c.release, nil
}

func (c updateTestGitHubClient) DownloadFile(context.Context, string, string, int64) error {
	return nil
}

func (c updateTestGitHubClient) FetchChecksumFile(context.Context, string) ([]byte, error) {
	return nil, nil
}

func TestUpdateService_CustomBuildDisablesAutomaticUpdate(t *testing.T) {
	svc := NewUpdateService(&updateTestCache{}, updateTestGitHubClient{
		release: &GitHubRelease{
			TagName: "v0.1.127",
			Name:    "v0.1.127",
			HTMLURL: "https://github.com/Wei-Shaw/sub2api/releases/tag/v0.1.127",
		},
	}, "0.1.126-product.1", "release")

	info, err := svc.CheckUpdate(context.Background(), true)
	if err != nil {
		t.Fatalf("CheckUpdate returned error: %v", err)
	}
	if info.HasUpdate {
		t.Fatalf("custom fork build must not expose one-click upstream updates")
	}
	if !info.CustomBuild {
		t.Fatalf("expected custom build")
	}
	if info.UpdateMode != updateModeFork {
		t.Fatalf("expected update mode %q, got %q", updateModeFork, info.UpdateMode)
	}
	if !info.UpstreamUpdateAvailable {
		t.Fatalf("expected upstream update marker")
	}
	if info.LatestVersion != "0.1.127-product.1" {
		t.Fatalf("expected display latest version with custom suffix, got %q", info.LatestVersion)
	}
}

func TestUpdateService_PerformUpdateRejectsCustomBuild(t *testing.T) {
	svc := NewUpdateService(&updateTestCache{}, updateTestGitHubClient{}, "0.1.126-product.1", "release")

	err := svc.PerformUpdate(context.Background())
	if err == nil {
		t.Fatalf("expected custom build update rejection")
	}
}

func TestCompareVersions_IgnoresCustomSuffix(t *testing.T) {
	cases := []struct {
		current string
		latest  string
		want    int
	}{
		{current: "0.1.126-product.1", latest: "0.1.126", want: 0},
		{current: "0.1.126-product.1", latest: "0.1.127", want: -1},
		{current: "0.1.128-product.1", latest: "0.1.127", want: 1},
		{current: "v0.1.126+product.1", latest: "v0.1.126", want: 0},
	}

	for _, tc := range cases {
		if got := compareVersions(tc.current, tc.latest); got != tc.want {
			t.Fatalf("compareVersions(%q, %q) = %d, want %d", tc.current, tc.latest, got, tc.want)
		}
	}
}
