package main

import (
	"io"
	"text/template"
)

type Cask struct {
	Desc     string `json:"desc"`
	Homepage string `json:"homepage"`
	Version  string `json:"version"`
	Binary   string `json:"binary"`
	Sha256   string `json:"sha254"`
}

var caskTemplate = `
cask "mondoo-cli" do
  name "Mondoo"
  desc "{{ .Desc }}"
  version "{{ .Version }}"
  homepage "{{ .Homepage }}"

  url "https://releases.mondoo.io/mondoo/#{version}/mondoo_#{version}_darwin_universal.pkg"
  sha256 "{{ .Sha256 }}"

  livecheck do
    url "https://releases.mondoo.io/mondoo/latest/index.html"
    regex(%r{href='\.\./(\d+(?:\.\d+)+)}i)
  end

  pkg "mondoo_#{version}_darwin_universal.pkg"

  uninstall script: {
    executable: "/Library/Mondoo/#{version}/uninstall.sh",
    sudo:       true,
  }

  zap trash: [
    "/Library/Mondoo",
    "/etc/opt/mondoo",
    "/usr/local/bin/mondoo",
  ]
end
`

func (c *Cask) Render(out io.Writer) error {
	t := template.Must(template.New("cask").Parse(caskTemplate))
	return t.Execute(out, c)
}