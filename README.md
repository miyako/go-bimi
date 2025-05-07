# go-bimi
CLI to fetch BIMI icon

```4d
$domain:="amazon.com"
//$domain:="bankofamerica.com"


var $cmd : Text
var $stdin; $stdOut; $stdErr : Blob

$dir:=Folder("/RESOURCES/bin/"+(Is macOS ? "macOS" : "Windows"))

SET ENVIRONMENT VARIABLE("_4D_OPTION_CURRENT_DIRECTORY"; $dir.platformPath)
$cmd:="bimi"+(Is Windows ? ".exe" : "")

$cmd+=" "+$domain

LAUNCH EXTERNAL PROCESS($cmd; $stdin; $stdOut; $stdErr)

If (BLOB size($stdOut)#0)
	
	
	/*
		bimi svg icon is not compatible with 4D SVG engine!
		SVG parser: Wrong base profile (must be equal to 'tiny', 'full' or 'basic'): xpath=/svg/title/@baseProfile
	*/
	
	$dom:=DOM Parse XML variable($stdOut)
	For ($attr; 1; DOM Count XML attributes($dom))
		DOM GET XML ATTRIBUTE BY INDEX($dom; $attr; $attrName; $attrValue)
		If ($attrName="baseProfile")
			DOM REMOVE XML ATTRIBUTE($dom; $attrName)
			DOM EXPORT TO VAR($dom; $stdOut)
			break
		End if 
	End for 
	DOM CLOSE XML($dom)
	
	var $icon : Picture
	BLOB TO PICTURE($stdOut; $icon; ".svg")
	SET PICTURE TO PASTEBOARD($icon)
End if
```
