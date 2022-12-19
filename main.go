package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/tinode/jsonco"
)

// Theme ...
type Theme struct {
	Name   string
	URL    string
	Target string
	File   string
	Dark   bool
}

func main() {
	themes := []Theme{
		{
			Name:   "aura",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/DaltonMenezes/vsextensions/aura-theme/2.1.2/vspackage",
			Target: "aura",
			File:   "extension/themes/aura-soft-dark-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "ayu-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/teabyii/vsextensions/ayu/1.0.5/vspackage",
			Target: "ayu",
			File:   "extension/ayu-dark.json",
			Dark:   true,
		},
		{
			Name:   "ayu-mirage",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/teabyii/vsextensions/ayu/1.0.5/vspackage",
			Target: "ayu",
			File:   "extension/ayu-mirage.json",
			Dark:   true,
		},
		{
			Name:   "ayu",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/teabyii/vsextensions/ayu/1.0.5/vspackage",
			Target: "ayu",
			File:   "extension/ayu-light.json",
			Dark:   false,
		},
		{
			Name:   "cobalt",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/wesbos/vsextensions/theme-cobalt2/2.3.0/vspackage",
			Target: "cobalt",
			File:   "extension/theme/cobalt2.json",
			Dark:   true,
		},
		{
			Name:   "cyberpunk",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/Endormi/vsextensions/2077-theme/1.5.3/vspackage",
			Target: "2077",
			File:   "extension/theme/2077 theme-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "dracula",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/dracula-theme/vsextensions/theme-dracula/2.24.2/vspackage",
			Target: "dracula",
			File:   "extension/theme/dracula.json",
			Dark:   true,
		},
		{
			Name:   "github-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/GitHub/vsextensions/github-vscode-theme/6.0.0/vspackage",
			Target: "github",
			File:   "extension/themes/dark.json",
			Dark:   true,
		},
		{
			Name:   "github-light",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/GitHub/vsextensions/github-vscode-theme/6.0.0/vspackage",
			Target: "github",
			File:   "extension/themes/light.json",
			Dark:   false,
		},
		{
			Name:   "gruvbox-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/jdinhlife/vsextensions/gruvbox/1.8.0/vspackage",
			Target: "gruvbox",
			File:   "extension/themes/gruvbox-dark-medium.json",
			Dark:   true,
		},
		{
			Name:   "gruvbox-light",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/jdinhlife/vsextensions/gruvbox/1.8.0/vspackage",
			Target: "gruvbox",
			File:   "extension/themes/gruvbox-light-medium.json",
			Dark:   false,
		},
		{
			Name:   "material-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/Equinusocio/vsextensions/vsc-material-theme/33.4.0/vspackage",
			Target: "material",
			File:   "extension/build/themes/Material-Theme-Default.json",
			Dark:   true,
		},
		{
			Name: 	"material-light",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/Equinusocio/vsextensions/vsc-material-theme/33.4.0/vspackage",
			Target: "material",
			File:   "extension/build/themes/Material-Theme-Lighter.json",
			Dark:   false,
		},
		{
			Name: 	"mirage",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/tristanremy/vsextensions/mirage/0.3.10/vspackage",
			Target: "mirage",
			File:   "extension/themes/Mirage-color-theme.json",
			Dark:   true,
		},

		{
			Name: 	"monokai",
			URL:    "https://golinski.org/pub/theme-monokai.vsix",
			Target: "monokai",
			File:   "extension/themes/monokai-color-theme.json",
			Dark:   true,
		},
		{
			Name: 	"noctis-azureus",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/azureus.json",
			Dark:   true,
		},
		{
			Name: 	"noctis-bordo",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/bordo.json",
			Dark:   true,
		},
		{
			Name: 	"noctis-hibernus",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/hibernus.json",
			Dark:   false,
		},
		{
			Name: 	"noctis-lilac",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/lilac.json",
			Dark:   false,
		},
		{
			Name: 	"noctis-lux",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/lux.json",
			Dark:   false,
		},
		{
			Name: 	"noctis-minimus",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/minimus.json",
			Dark:   true,
		},
		{
			Name: 	"noctis-uva",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/uva.json",
			Dark:   true,
		},
		{
			Name: 	"noctis-viola",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/viola.json",
			Dark:   true,
		},
		{
			Name: 	"noctis",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/liviuschera/vsextensions/noctis/10.40.0/vspackage",
			Target: "noctis",
			File:   "extension/themes/noctis.json",
			Dark:   true,
		},
		{
			Name:   "one-dark-pro",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/zhuangtongfa/vsextensions/Material-theme/3.15.6/vspackage",
			Target: "one-dark-pro",
			File:   "extension/themes/OneDark-Pro.json",
			Dark:   true,
		},
		{
			Name:   "panda",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/tinkertrain/vsextensions/theme-panda/1.3.0/vspackage",
			Target: "panda",
			File:   "extension/dist/Panda.json",
			Dark:   true,
		},
		{
			Name:   "pastel",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/cval/vsextensions/pastel-theme/1.4.1/vspackage",
			Target: "pastel",
			File:   "extension/themes/pastel-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "poimandres",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/pmndrs/vsextensions/pmndrs/0.3.6/vspackage",
			Target: "poimandres",
			File:   "extension/themes/poimandres-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "ra-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/rahmanyerli/vsextensions/ra-dark/1.1.0/vspackage",
			Target: "ra-dark",
			File:   "extension/themes/ra-dark.json",
			Dark:   true,
		},
		{
			Name:   "ruby-blue",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/hirofumii/vsextensions/rubyblue-theme/2.1.0/vspackage",
			Target: "ruby-blue",
			File:   "extension/themes/RubyBlue.tmTheme.json",
			Dark:   true,
		},
		{
			Name:   "seti-black",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/bobsparadox/vsextensions/seti-black/1.0.0/vspackage",
			Target: "seti-black",
			File:   "extension/themes/Seti Black-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "shades-of-purple",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ahmadawais/vsextensions/shades-of-purple/7.1.5/vspackage",
			Target: "shades-of-purple",
			File:   "extension/themes/shades-of-purple-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "solarized-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ryanolsonx/vsextensions/solarized/2.0.3/vspackage",
			Target: "solarized",
			File:   "extension/themes/dark-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "solarized-light",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ryanolsonx/vsextensions/solarized/2.0.3/vspackage",
			Target: "solarized",
			File:   "extension/themes/light-color-theme.json",
			Dark:   false,
		},
		{
			Name:   "tokyo-night-storm",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/enkia/vsextensions/tokyo-night/0.9.4/vspackage",
			Target: "tokyoś",
			File:   "extension/themes/tokyo-night-storm-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "tokyo-night",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/enkia/vsextensions/tokyo-night/0.9.4/vspackage",
			Target: "tokyo",
			File:   "extension/themes/tokyo-night-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "zenburn",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ryanolsonx/vsextensions/zenburn/2.1.3/vspackage",
			Target: "zenburn",
			File:   "extension/themes/zenburn.json",
			Dark:   true,
		},
	}

	for _, theme := range themes {
		fmt.Println("Processing theme: ", theme.Name)
		if _, err := os.Stat("./tmp/" + theme.Target + ".zip"); os.IsNotExist(err) {
			fmt.Println("  Download theme")
			downloadTheme(theme)
		}

		fmt.Println("  Extracting theme")
		content, err := extractTheme(theme)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("  Generating template")
		generateTheme(theme, content)
	}
}

func generateTheme(theme Theme, content []byte) {
	params := makeTemplateParams(theme, content)

	t, err := template.ParseFiles("./template.js")
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("./theme/" + theme.Name + ".ts")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(out, params)
	if err != nil {
		log.Fatal(err)
	}

	out.Close()
}

// Style ...
type Style struct {
	Color     *string
	FontStyle *string
	Prio      *int
}

// TokenColorSettings ...
type TokenColorSettings struct {
	Foreground *string
	FontStyle  *string
}

// TokenColor ...
type TokenColor struct {
	Scope    interface{}
	Settings TokenColorSettings
}

// VsCodeTheme ...
type VsCodeTheme struct {
	Colors      map[string]string
	TokenColors []TokenColor
}

func find(data VsCodeTheme, keys ...string) Style {
	style := Style{}

	for _, key := range keys {
		if value, exist := data.Colors[key]; exist && style.Color == nil {
			return Style{Color: &value}
		}

		for _, tokenColor := range data.TokenColors {
			scopes := []string{}
			rt := reflect.TypeOf(tokenColor.Scope)
			if tokenColor.Scope == nil {
				continue
			}

			switch rt.Kind() {
			case reflect.Slice:
				for _, s := range tokenColor.Scope.([]interface{}) {
					scopes = append(scopes, s.(string))
				}
			case reflect.String:
				splitted := strings.Split(tokenColor.Scope.(string), ",")
				for _, s := range splitted {
					scopes = append(scopes, strings.TrimSpace(s))
				}
			default:
				panic(fmt.Sprintf("Unexpected scope type %s", rt))
			}

			for i, scope := range scopes {
				if scope == key &&
					(style.Color == nil || *style.Prio > i) &&
					tokenColor.Settings.Foreground != nil {
					prio := i
					style.Color = tokenColor.Settings.Foreground
					style.Prio = &prio
				}

				if scope == key && style.FontStyle == nil && tokenColor.Settings.FontStyle != nil {
					style.FontStyle = tokenColor.Settings.FontStyle
				}
			}
		}
	}

	if style.Color == nil {
		panic(fmt.Sprintf("Could not find color by: %s", keys))
	}

	return style
}

// TemplateParams ...
type TemplateParams struct {
	ExportPrefix string
	Dark         bool

	// Editor
	Background         Style
	Foreground         Style
	Selection          Style
	Cursor             Style
	DropdownBackground Style
	DropdownBorder     Style
	ActiveLine         Style
	MatchingBracket    Style

	// Syntax
	Keyword   Style // if else, etc
	Storage   Style // const, let, etc - Not supported in CM
	Parameter Style // fn(parmater)    - Not supported in CM
	Variable  Style
	Function  Style
	String    Style
	Constant  Style // ???
	Type      Style // x: MyType
	Class     Style // class MyClass
	Number    Style
	Comment   Style
	Heading   Style
	Invalid   Style
	Regexp    Style
}

func makeTemplateParams(theme Theme, content []byte) TemplateParams {
	var data VsCodeTheme
	jr := jsonco.New(bytes.NewBuffer(content))
	json.NewDecoder(jr).Decode(&data)

	params := TemplateParams{
		ExportPrefix: kebabToCamelCase(theme.Name),
		Dark:         theme.Dark,
		// Layout
		Background:         find(data, "editor.background"),
		Foreground:         find(data, "foreground", "editor.foreground", "input.foreground", "variable.other.readwrite"),
		Selection:          find(data, "editor.selectionBackground"),
		Cursor:             find(data, "editorCursor.foreground", "foreground", "editor.foreground", "input.foreground", "variable.other.readwrite"),
		DropdownBackground: find(data, "editor.background"),
		DropdownBorder:     find(data, "dropdown.border", "foreground", "input.foreground", "editor.foreground"),
		ActiveLine:         find(data, "editor.lineHighlightBackground", "editor.selectionBackground"),
		MatchingBracket:    find(data, "editorBracketMatch.background", "editor.lineHighlightBackground", "editor.selectionBackground"),
		// Syntax
		Keyword:   find(data, "keyword", "foreground", "input.foreground", "editor.foreground"),
		Storage:   find(data, "storage", "keyword"),
		Variable:  find(data, "variable", "variable.parameter", "variable.other", "variable.language", "foreground", "input.foreground", "editor.foreground"),
		Parameter: find(data, "variable.parameter", "variable.other", "variable", "foreground", "input.foreground", "editor.foreground"),
		Function:  find(data, "entity.name.function", "support.function", "entity.name", "support", "foreground", "input.foreground", "editor.foreground"),
		String:    find(data, "string", "string.quoted" ),
		Constant:  find(data, "constant", "constant.character", "constant.keyword"),
		Type:      find(data, "entity.name.type", "entity.name.class", "support.type", "support"),
		Class:     find(data, "entity.name.class", "entity.name", "foreground", "input.foreground", "editor.foreground"),
		Number:    find(data, "constant.numeric", "constant"),
		Comment:   find(data, "comment"),
		Heading:   find(data, "markup.heading", "markup.heading.setext", "heading.1.markdown entity.name", "foreground", "input.foreground", "editor.foreground"),
		Invalid:   find(data, "invalid", "editorError.foreground", "errorForeground", "foreground", "input.foreground"),
		Regexp:    find(data, "string.regexp", "string"),
	}

	return params
}

func extractTheme(theme Theme) ([]byte, error) {
	r, err := zip.OpenReader("tmp/" + theme.Target + ".zip")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name != theme.File {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		content, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}

		rc.Close()

		return content, nil
	}

	return nil, fmt.Errorf("Could not find file %s in extension", theme.File)
}

func downloadTheme(theme Theme) {
	resp, err := http.Get(theme.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("Could not download theme: ", theme, "StatusCode: ", resp.StatusCode)
	}

	_ = os.Mkdir("./tmp", 0700)

	// Create the file
	out, err := os.Create("tmp/" + theme.Target + ".zip")
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}

func kebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}
