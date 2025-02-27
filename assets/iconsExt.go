package assets

var Icon_Ext = map[string]*Icon_Info{
	"htm":               Icon_Set["html"],
	"html":              Icon_Set["html"],
	"xhtml":             Icon_Set["html"],
	"html_vm":           Icon_Set["html"],
	"asp":               Icon_Set["html"],
	"jade":              Icon_Set["pug"],
	"pug":               Icon_Set["pug"],
	"md":                Icon_Set["markdown"],
	"markdown":          Icon_Set["markdown"],
	"rst":               Icon_Set["markdown"],
	"blink":             Icon_Set["blink"],
	"css":               Icon_Set["css"],
	"scss":              Icon_Set["sass"],
	"sass":              Icon_Set["sass"],
	"less":              Icon_Set["less"],
	"json":              Icon_Set["json"],
	"tsbuildinfo":       Icon_Set["json"],
	"json5":             Icon_Set["json"],
	"jsonl":             Icon_Set["json"],
	"ndjson":            Icon_Set["json"],
	"jinja":             Icon_Set["jinja"],
	"jinja2":            Icon_Set["jinja"],
	"j2":                Icon_Set["jinja"],
	"jinja-html":        Icon_Set["jinja"],
	"sublime-project":   Icon_Set["sublime"],
	"sublime-workspace": Icon_Set["sublime"],
	"yaml":              Icon_Set["yaml"],
	"yaml-tmlanguage":   Icon_Set["yaml"],
	"yml":               Icon_Set["yaml"],
    "php":               Icon_Set["php"],
	"xml":               Icon_Set["xml"],
	"plist":             Icon_Set["xml"],
	"xsd":               Icon_Set["xml"],
	"dtd":               Icon_Set["xml"],
	"xsl":               Icon_Set["xml"],
	"xslt":              Icon_Set["xml"],
	"resx":              Icon_Set["xml"],
	"iml":               Icon_Set["xml"],
	"xquery":            Icon_Set["xml"],
	"tmLanguage":        Icon_Set["xml"],
	"manifest":          Icon_Set["xml"],
	"project":           Icon_Set["xml"],
	"png":               Icon_Set["image"],
	"jpeg":              Icon_Set["image"],
	"jpg":               Icon_Set["image"],
	"gif":               Icon_Set["image"],
	"ico":               Icon_Set["image"],
	"tif":               Icon_Set["image"],
	"tiff":              Icon_Set["image"],
	"psd":               Icon_Set["image"],
	"psb":               Icon_Set["image"],
	"ami":               Icon_Set["image"],
	"apx":               Icon_Set["image"],
	"bmp":               Icon_Set["image"],
	"bpg":               Icon_Set["image"],
	"brk":               Icon_Set["image"],
	"cur":               Icon_Set["image"],
	"dds":               Icon_Set["image"],
	"dng":               Icon_Set["image"],
	"exr":               Icon_Set["image"],
	"fpx":               Icon_Set["image"],
	"gbr":               Icon_Set["image"],
	"img":               Icon_Set["image"],
	"jbig2":             Icon_Set["image"],
	"jb2":               Icon_Set["image"],
	"jng":               Icon_Set["image"],
	"jxr":               Icon_Set["image"],
	"pbm":               Icon_Set["image"],
	"pgf":               Icon_Set["image"],
	"pic":               Icon_Set["image"],
	"raw":               Icon_Set["image"],
	"webp":              Icon_Set["image"],
	"eps":               Icon_Set["image"],
	"afphoto":           Icon_Set["image"],
	"ase":               Icon_Set["image"],
	"aseprite":          Icon_Set["image"],
	"clip":              Icon_Set["image"],
	"cpt":               Icon_Set["image"],
	"heif":              Icon_Set["image"],
	"heic":              Icon_Set["image"],
	"kra":               Icon_Set["image"],
	"mdp":               Icon_Set["image"],
	"ora":               Icon_Set["image"],
	"pdn":               Icon_Set["image"],
	"reb":               Icon_Set["image"],
	"sai":               Icon_Set["image"],
	"tga":               Icon_Set["image"],
	"xcf":               Icon_Set["image"],
	"js":                Icon_Set["javascript"],
	"esx":               Icon_Set["javascript"],
	"mj":                Icon_Set["javascript"],
	"jsx":               Icon_Set["react"],
	"tsx":               Icon_Set["react_ts"],
	"ini":               Icon_Set["settings"],
	"dlc":               Icon_Set["settings"],
	"dll":               Icon_Set["settings"],
	"config":            Icon_Set["settings"],
	"conf":              Icon_Set["settings"],
	"properties":        Icon_Set["settings"],
	"prop":              Icon_Set["settings"],
	"settings":          Icon_Set["settings"],
	"option":            Icon_Set["settings"],
	"props":             Icon_Set["settings"],
	"toml":              Icon_Set["settings"],
	"prefs":             Icon_Set["settings"],
	"dotsettings":       Icon_Set["settings"],
	"cfg":               Icon_Set["settings"],
	"ts":                Icon_Set["typescript"],
	"marko":             Icon_Set["markojs"],
	"pdf":               Icon_Set["pdf"],
	"xlsx":              Icon_Set["table"],
	"xls":               Icon_Set["table"],
	"csv":               Icon_Set["table"],
	"tsv":               Icon_Set["table"],
	"vscodeignore":      Icon_Set["vscode"],
	"vsixmanifest":      Icon_Set["vscode"],
	"vsix":              Icon_Set["vscode"],
	"code-workplace":    Icon_Set["vscode"],
	"csproj":            Icon_Set["visualstudio"],
	"ruleset":           Icon_Set["visualstudio"],
	"sln":               Icon_Set["visualstudio"],
	"suo":               Icon_Set["visualstudio"],
	"vb":                Icon_Set["visualstudio"],
	"vbs":               Icon_Set["visualstudio"],
	"vcxitems":          Icon_Set["visualstudio"],
	"vcxproj":           Icon_Set["visualstudio"],
	"pdb":               Icon_Set["database"],
	"sql":               Icon_Set["mysql"],
	"pks":               Icon_Set["database"],
	"pkb":               Icon_Set["database"],
	"accdb":             Icon_Set["database"],
	"mdb":               Icon_Set["database"],
	"sqlite":            Icon_Set["sqlite"],
	"sqlite3":           Icon_Set["sqlite"],
	"pgsql":             Icon_Set["postgresql"],
	"postgres":          Icon_Set["postgresql"],
	"psql":              Icon_Set["postgresql"],
	"cs":                Icon_Set["csharp"],
	"csx":               Icon_Set["csharp"],
	"qs":                Icon_Set["qsharp"],
	"zip":               Icon_Set["zip"],
	"tar":               Icon_Set["zip"],
	"gz":                Icon_Set["zip"],
	"xz":                Icon_Set["zip"],
	"br":                Icon_Set["zip"],
	"bzip2":             Icon_Set["zip"],
	"gzip":              Icon_Set["zip"],
	"brotli":            Icon_Set["zip"],
	"7z":                Icon_Set["zip"],
	"rar":               Icon_Set["zip"],
	"tgz":               Icon_Set["zip"],
	"vala":              Icon_Set["vala"],
	"zig":               Icon_Set["zig"],
	"exe":               Icon_Set["exe"],
	"msi":               Icon_Set["exe"],
	"java":              Icon_Set["java"],
	"jar":               Icon_Set["java"],
	"jsp":               Icon_Set["java"],
	"c":                 Icon_Set["c"],
	"m":                 Icon_Set["c"],
	"i":                 Icon_Set["c"],
	"mi":                Icon_Set["c"],
	"h":                 Icon_Set["h"],
	"cc":                Icon_Set["cpp"],
	"cpp":               Icon_Set["cpp"],
	"cxx":               Icon_Set["cpp"],
	"c++":               Icon_Set["cpp"],
	"cp":                Icon_Set["cpp"],
	"mm":                Icon_Set["cpp"],
	"mii":               Icon_Set["cpp"],
	"ii":                Icon_Set["cpp"],
	"hh":                Icon_Set["hpp"],
	"hpp":               Icon_Set["hpp"],
	"hxx":               Icon_Set["hpp"],
	"h++":               Icon_Set["hpp"],
	"hp":                Icon_Set["hpp"],
	"tcc":               Icon_Set["hpp"],
	"inl":               Icon_Set["hpp"],
	"go":                Icon_Set["go"],
	"py":                Icon_Set["python"],
	"pyc":               Icon_Set["python-misc"],
	"whl":               Icon_Set["python-misc"],
	"url":               Icon_Set["url"],
	"sh":                Icon_Set["console"],
	"ksh":               Icon_Set["console"],
	"csh":               Icon_Set["console"],
	"tcsh":              Icon_Set["console"],
	"zsh":               Icon_Set["console"],
	"bash":              Icon_Set["console"],
	"bat":               Icon_Set["console"],
	"cmd":               Icon_Set["console"],
	"awk":               Icon_Set["console"],
	"fish":              Icon_Set["console"],
	"ps1":               Icon_Set["powershell"],
	"psm1":              Icon_Set["powershell"],
	"psd1":              Icon_Set["powershell"],
	"ps1xml":            Icon_Set["powershell"],
	"psc1":              Icon_Set["powershell"],
	"pssc":              Icon_Set["powershell"],
	"gradle":            Icon_Set["gradle"],
	"doc":               Icon_Set["word"],
	"docx":              Icon_Set["word"],
	"rtf":               Icon_Set["word"],
	"cer":               Icon_Set["certificate"],
	"cert":              Icon_Set["certificate"],
	"crt":               Icon_Set["certificate"],
	"pub":               Icon_Set["key"],
	"key":               Icon_Set["key"],
	"pem":               Icon_Set["key"],
	"asc":               Icon_Set["key"],
	"gpg":               Icon_Set["key"],
	"woff":              Icon_Set["font"],
	"woff2":             Icon_Set["font"],
	"ttf":               Icon_Set["font"],
	"eot":               Icon_Set["font"],
	"suit":              Icon_Set["font"],
	"otf":               Icon_Set["font"],
	"bmap":              Icon_Set["font"],
	"fnt":               Icon_Set["font"],
	"odttf":             Icon_Set["font"],
	"ttc":               Icon_Set["font"],
	"font":              Icon_Set["font"],
	"fonts":             Icon_Set["font"],
	"sui":               Icon_Set["font"],
	"ntf":               Icon_Set["font"],
	"mrf":               Icon_Set["font"],
	"lib":               Icon_Set["lib"],
	"bib":               Icon_Set["lib"],
	"rb":                Icon_Set["ruby"],
	"erb":               Icon_Set["ruby"],
	"fs":                Icon_Set["fsharp"],
	"fsx":               Icon_Set["fsharp"],
	"fsi":               Icon_Set["fsharp"],
	"fsproj":            Icon_Set["fsharp"],
	"swift":             Icon_Set["swift"],
	"ino":               Icon_Set["arduino"],
	"dockerignore":      Icon_Set["docker"],
	"dockerfile":        Icon_Set["docker"],
	"tex":               Icon_Set["tex"],
	"sty":               Icon_Set["tex"],
	"dtx":               Icon_Set["tex"],
	"ltx":               Icon_Set["tex"],
	"pptx":              Icon_Set["powerpoint"],
	"ppt":               Icon_Set["powerpoint"],
	"pptm":              Icon_Set["powerpoint"],
	"potx":              Icon_Set["powerpoint"],
	"potm":              Icon_Set["powerpoint"],
	"ppsx":              Icon_Set["powerpoint"],
	"ppsm":              Icon_Set["powerpoint"],
	"pps":               Icon_Set["powerpoint"],
	"ppam":              Icon_Set["powerpoint"],
	"ppa":               Icon_Set["powerpoint"],
	"webm":              Icon_Set["video"],
	"mkv":               Icon_Set["video"],
	"flv":               Icon_Set["video"],
	"vob":               Icon_Set["video"],
	"ogv":               Icon_Set["video"],
	"ogg":               Icon_Set["video"],
	"gifv":              Icon_Set["video"],
	"avi":               Icon_Set["video"],
	"mov":               Icon_Set["video"],
	"qt":                Icon_Set["video"],
	"wmv":               Icon_Set["video"],
	"yuv":               Icon_Set["video"],
	"rm":                Icon_Set["video"],
	"rmvb":              Icon_Set["video"],
	"mp4":               Icon_Set["video"],
	"m4v":               Icon_Set["video"],
	"mpg":               Icon_Set["video"],
	"mp2":               Icon_Set["video"],
	"mpeg":              Icon_Set["video"],
	"mpe":               Icon_Set["video"],
	"mpv":               Icon_Set["video"],
	"m2v":               Icon_Set["video"],
	"vdi":               Icon_Set["virtual"],
	"vbox":              Icon_Set["virtual"],
	"vbox-prev":         Icon_Set["virtual"],
	"ics":               Icon_Set["email"],
	"mp3":               Icon_Set["audio"],
	"flac":              Icon_Set["audio"],
	"m4a":               Icon_Set["audio"],
	"wma":               Icon_Set["audio"],
	"aiff":              Icon_Set["audio"],
	"coffee":            Icon_Set["coffee"],
	"cson":              Icon_Set["coffee"],
	"iced":              Icon_Set["coffee"],
	"txt":               Icon_Set["document"],
	"graphql":           Icon_Set["graphql"],
	"gql":               Icon_Set["graphql"],
	"rs":                Icon_Set["rust"],
	"raml":              Icon_Set["raml"],
	"xaml":              Icon_Set["xaml"],
	"hs":                Icon_Set["haskell"],
	"kt":                Icon_Set["kotlin"],
	"kts":               Icon_Set["kotlin"],
	"patch":             Icon_Set["git"],
	"lua":               Icon_Set["lua"],
	"clj":               Icon_Set["clojure"],
	"cljs":              Icon_Set["clojure"],
	"cljc":              Icon_Set["clojure"],
	"groovy":            Icon_Set["groovy"],
	"r":                 Icon_Set["r"],
	"rmd":               Icon_Set["r"],
	"dart":              Icon_Set["dart"],
	"as":                Icon_Set["actionscript"],
	"mxml":              Icon_Set["mxml"],
	"ahk":               Icon_Set["autohotkey"],
	"swf":               Icon_Set["flash"],
	"swc":               Icon_Set["swc"],
	"cmake":             Icon_Set["cmake"],
	"asm":               Icon_Set["assembly"],
	"a51":               Icon_Set["assembly"],
	"inc":               Icon_Set["assembly"],
	"nasm":              Icon_Set["assembly"],
	"s":                 Icon_Set["assembly"],
	"ms":                Icon_Set["assembly"],
	"agc":               Icon_Set["assembly"],
	"ags":               Icon_Set["assembly"],
	"aea":               Icon_Set["assembly"],
	"argus":             Icon_Set["assembly"],
	"mitigus":           Icon_Set["assembly"],
	"binsource":         Icon_Set["assembly"],
	"vue":               Icon_Set["vue"],
	"ml":                Icon_Set["ocaml"],
	"mli":               Icon_Set["ocaml"],
	"cmx":               Icon_Set["ocaml"],
	"lock":              Icon_Set["lock"],
	"hbs":               Icon_Set["handlebars"],
	"mustache":          Icon_Set["handlebars"],
	"pm":                Icon_Set["perl"],
	"raku":              Icon_Set["perl"],
	"hx":                Icon_Set["haxe"],
	"pp":                Icon_Set["puppet"],
	"ex":                Icon_Set["elixir"],
	"exs":               Icon_Set["elixir"],
	"eex":               Icon_Set["elixir"],
	"leex":              Icon_Set["elixir"],
	"erl":               Icon_Set["erlang"],
	"twig":              Icon_Set["twig"],
	"jl":                Icon_Set["julia"],
	"elm":               Icon_Set["elm"],
	"pure":              Icon_Set["purescript"],
	"purs":              Icon_Set["purescript"],
	"tpl":               Icon_Set["smarty"],
	"styl":              Icon_Set["stylus"],
	"merlin":            Icon_Set["merlin"],
	"v":                 Icon_Set["verilog"],
	"vhd":               Icon_Set["verilog"],
	"sv":                Icon_Set["verilog"],
	"svh":               Icon_Set["verilog"],
	"robot":             Icon_Set["robot"],
	"sol":               Icon_Set["solidity"],
	"yang":              Icon_Set["yang"],
	"mjml":              Icon_Set["mjml"],
	"tf":                Icon_Set["terraform"],
	"tfvars":            Icon_Set["terraform"],
	"tfstate":           Icon_Set["terraform"],
	"applescript":       Icon_Set["applescript"],
	"ipa":               Icon_Set["applescript"],
	"cake":              Icon_Set["cake"],
	"nim":               Icon_Set["nim"],
	"nimble":            Icon_Set["nim"],
	"apib":              Icon_Set["apiblueprint"],
	"apiblueprint":      Icon_Set["apiblueprint"],
	"pcss":              Icon_Set["postcss"],
	"sss":               Icon_Set["postcss"],
	"todo":              Icon_Set["todo"],
	"nix":               Icon_Set["nix"],
	"slim":              Icon_Set["slim"],
	"http":              Icon_Set["http"],
	"rest":              Icon_Set["http"],
	"apk":               Icon_Set["android"],
	"env":               Icon_Set["tune"],
	"jenkinsfile":       Icon_Set["jenkins"],
	"jenkins":           Icon_Set["jenkins"],
	"log":               Icon_Set["log"],
	"ejs":               Icon_Set["ejs"],
	"djt":               Icon_Set["django"],
	"pot":               Icon_Set["i18n"],
	"po":                Icon_Set["i18n"],
	"mo":                Icon_Set["i18n"],
	"d":                 Icon_Set["d"],
	"mdx":               Icon_Set["mdx"],
	"gd":                Icon_Set["godot"],
	"godot":             Icon_Set["godot-assets"],
	"tres":              Icon_Set["godot-assets"],
	"tscn":              Icon_Set["godot-assets"],
	"azcli":             Icon_Set["azure"],
	"vagrantfile":       Icon_Set["vagrant"],
	"cshtml":            Icon_Set["razor"],
	"vbhtml":            Icon_Set["razor"],
	"ad":                Icon_Set["asciidoc"],
	"adoc":              Icon_Set["asciidoc"],
	"asciidoc":          Icon_Set["asciidoc"],
	"edge":              Icon_Set["edge"],
	"ss":                Icon_Set["scheme"],
	"scm":               Icon_Set["scheme"],
	"stl":               Icon_Set["3d"],
	"obj":               Icon_Set["3d"],
	"ac":                Icon_Set["3d"],
	"blend":             Icon_Set["3d"],
	"mesh":              Icon_Set["3d"],
	"mqo":               Icon_Set["3d"],
	"pmd":               Icon_Set["3d"],
	"pmx":               Icon_Set["3d"],
	"skp":               Icon_Set["3d"],
	"vac":               Icon_Set["3d"],
	"vdp":               Icon_Set["3d"],
	"vox":               Icon_Set["3d"],
	"svg":               Icon_Set["svg"],
	"vimrc":             Icon_Set["vim"],
	"gvimrc":            Icon_Set["vim"],
	"exrc":              Icon_Set["vim"],
	"moon":              Icon_Set["moonscript"],
	"iso":               Icon_Set["disc"],
	"f":                 Icon_Set["fortran"],
	"f77":               Icon_Set["fortran"],
	"f90":               Icon_Set["fortran"],
	"f95":               Icon_Set["fortran"],
	"f03":               Icon_Set["fortran"],
	"f08":               Icon_Set["fortran"],
	"tcl":               Icon_Set["tcl"],
	"liquid":            Icon_Set["liquid"],
	"p":                 Icon_Set["prolog"],
	"pro":               Icon_Set["prolog"],
	"coco":              Icon_Set["coconut"],
	"sketch":            Icon_Set["sketch"],
	"opam":              Icon_Set["opam"],
	"dhallb":            Icon_Set["dhall"],
	"pwn":               Icon_Set["pawn"],
	"amx":               Icon_Set["pawn"],
	"dhall":             Icon_Set["dhall"],
	"pas":               Icon_Set["pascal"],
	"unity":             Icon_Set["shaderlab"],
	"nupkg":             Icon_Set["nuget"],
	"command":           Icon_Set["command"],
	"dsc":               Icon_Set["denizenscript"],
	"deb":               Icon_Set["debian"],
	"rpm":               Icon_Set["redhat"],
	"snap":              Icon_Set["ubuntu"],
	"ebuild":            Icon_Set["gentoo"],
	"pkg":               Icon_Set["applescript"],
	"openbsd":           Icon_Set["freebsd"],
	// "ls":                Icon_Set["livescript"],
	// "re":                Icon_Set["reason"],
	// "rei":               Icon_Set["reason"],
	// "cmj":               Icon_Set["bucklescript"],
	// "nb":                Icon_Set["mathematica"],
	// "wl":                Icon_Set["wolframlanguage"],
	// "wls":               Icon_Set["wolframlanguage"],
	// "njk":               Icon_Set["nunjucks"],
	// "nunjucks":          Icon_Set["nunjucks"],
	// "au3":               Icon_Set["autoit"],
	// "haml":              Icon_Set["haml"],
	// "feature":           Icon_Set["cucumber"],
	// "riot":              Icon_Set["riot"],
	// "tag":               Icon_Set["riot"],
	// "vfl":               Icon_Set["vfl"],
	// "kl":                Icon_Set["kl"],
	// "cfml":              Icon_Set["coldfusion"],
	// "cfc":               Icon_Set["coldfusion"],
	// "lucee":             Icon_Set["coldfusion"],
	// "cfm":               Icon_Set["coldfusion"],
	// "cabal":             Icon_Set["cabal"],
	// "rql":               Icon_Set["restql"],
	// "restql":            Icon_Set["restql"],
	// "kv":                Icon_Set["kivy"],
	// "graphcool":         Icon_Set["graphcool"],
	// "sbt":               Icon_Set["sbt"],
	// "cr":                Icon_Set["crystal"],
	// "ecr":               Icon_Set["crystal"],
	// "cu":                Icon_Set["cuda"],
	// "cuh":               Icon_Set["cuda"],
	// "def":               Icon_Set["dotjs"],
	// "dot":               Icon_Set["dotjs"],
	// "jst":               Icon_Set["dotjs"],
	// "pde":               Icon_Set["processing"],
	// "wpy":               Icon_Set["wepy"],
	// "hcl":               Icon_Set["hcl"],
	// "san":               Icon_Set["san"],
	// "red":               Icon_Set["red"],
	// "fxp":               Icon_Set["foxpro"],
	// "prg":               Icon_Set["foxpro"],
	// "wat":               Icon_Set["webassembly"],
	// "wasm":              Icon_Set["webassembly"],
	// "ipynb":             Icon_Set["jupyter"],
	// "bal":               Icon_Set["ballerina"],
	// "balx":              Icon_Set["ballerina"],
	// "rkt":               Icon_Set["racket"],
	// "bzl":               Icon_Set["bazel"],
	// "bazel":             Icon_Set["bazel"],
	// "mint":              Icon_Set["mint"],
	// "vm":                Icon_Set["velocity"],
	// "fhtml":             Icon_Set["velocity"],
	// "vtl":               Icon_Set["velocity"],
	// "prisma":            Icon_Set["prisma"],
	// "abc":               Icon_Set["abc"],
	// "lisp":              Icon_Set["lisp"],
	// "lsp":               Icon_Set["lisp"],
	// "cl":                Icon_Set["lisp"],
	// "fast":              Icon_Set["lisp"],
	// "svelte":            Icon_Set["svelte"],
	// "prw":               Icon_Set["advpl_prw"],
	// "prx":               Icon_Set["advpl_prw"],
	// "ptm":               Icon_Set["advpl_ptm"],
	// "tlpp":              Icon_Set["advpl_tlpp"],
	// "ch":                Icon_Set["advpl_include"],
	// "4th":               Icon_Set["forth"],
	// "fth":               Icon_Set["forth"],
	// "frt":               Icon_Set["forth"],
	// "iuml":              Icon_Set["uml"],
	// "pu":                Icon_Set["uml"],
	// "puml":              Icon_Set["uml"],
	// "plantuml":          Icon_Set["uml"],
	// "wsd":               Icon_Set["uml"],
	// "sml":               Icon_Set["sml"],
	// "mlton":             Icon_Set["sml"],
	// "mlb":               Icon_Set["sml"],
	// "sig":               Icon_Set["sml"],
	// "fun":               Icon_Set["sml"],
	// "cm":                Icon_Set["sml"],
	// "lex":               Icon_Set["sml"],
	// "use":               Icon_Set["sml"],
	// "grm":               Icon_Set["sml"],
	// "imba":              Icon_Set["imba"],
	// "drawio":            Icon_Set["drawio"],
	// "dio":               Icon_Set["drawio"],
	// "sas":               Icon_Set["sas"],
	// "sas7bdat":          Icon_Set["sas"],
	// "sashdat":           Icon_Set["sas"],
	// "astore":            Icon_Set["sas"],
	// "ast":               Icon_Set["sas"],
	// "sast":              Icon_Set["sas"],
}
