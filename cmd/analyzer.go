package main

import (
	"github.com/konveyor/tackle2-addon-analyzer/builder"
	"github.com/konveyor/tackle2-addon/command"
	"path"
)

//
// Analyzer application analyzer.
type Analyzer struct {
	*Data
}

//
// Run analyzer.
func (r *Analyzer) Run() (b *builder.Issues, err error) {
	bin := path.Join(
		Dir,
		"opt",
		"konveyor-analyzer")
	output := path.Join(Dir, "report.yaml")
	cmd := command.Command{Path: bin}
	cmd.Options, err = r.options(output)
	if err != nil {
		return
	}
	b = &builder.Issues{Path: output}
	err = cmd.Run()
	if err != nil {
		return
	}
	return
}

//
// options builds Analyzer options.
func (r *Analyzer) options(output string) (options command.Options, err error) {
	settings := &Settings{}
	err = settings.Read(SettingsPath)
	if err != nil {
		return
	}
	options = command.Options{
		"--provider-settings",
		SettingsPath,
		"--output-file",
		output,
	}
	err = r.Tagger.AddOptions(&options)
	if err != nil {
		return
	}
	err = r.Mode.AddOptions(settings)
	if err != nil {
		return
	}
	if r.Rules != nil {
		err = r.Rules.AddOptions(&options)
		if err != nil {
			return
		}
	}
	err = r.Scope.AddOptions(&options)
	if err != nil {
		return
	}
	err = settings.Write(SettingsPath)
	if err != nil {
		return
	}
	return
}

//
// DepAnalyzer application analyzer.
type DepAnalyzer struct {
	*Data
}

//
// Run analyzer.
func (r *DepAnalyzer) Run() (b *builder.Deps, err error) {
	bin := path.Join(
		Dir,
		"opt",
		"konveyor-analyzer-dep")
	output := path.Join(Dir, "deps.yaml")
	cmd := command.Command{Path: bin}
	cmd.Options, err = r.options(output)
	if err != nil {
		return
	}
	b = &builder.Deps{Path: output}
	err = cmd.Run()
	if err != nil {
		return
	}
	return
}

//
// options builds Analyzer options.
func (r *DepAnalyzer) options(output string) (options command.Options, err error) {
	settings := &Settings{}
	err = settings.Read(SettingsPath)
	if err != nil {
		return
	}
	options = command.Options{
		"--provider-settings",
		SettingsPath,
		"--output-file",
		output,
	}
	err = r.Mode.AddOptions(settings)
	if err != nil {
		return
	}
	err = settings.Write(SettingsPath)
	if err != nil {
		return
	}
	return
}
