//%attributes = {}
$domain:="amazon.com"
//$domain:="bankofamerica.com"


var $cmd : Text
var $stdin; $stdOut; $stdErr : Blob

$dir:=Folder:C1567("/RESOURCES/bin/"+(Is macOS:C1572 ? "macOS" : "Windows"))

SET ENVIRONMENT VARIABLE:C812("_4D_OPTION_CURRENT_DIRECTORY"; $dir.platformPath)
$cmd:="bimi"+(Is Windows:C1573 ? ".exe" : "")

$cmd+=" "+$domain

LAUNCH EXTERNAL PROCESS:C811($cmd; $stdin; $stdOut; $stdErr)

If (BLOB size:C605($stdOut)#0)
	
	
/*
bimi svg icon is not compatible with 4D SVG engine!
SVG parser: Wrong base profile (must be equal to 'tiny', 'full' or 'basic'): xpath=/svg/title/@baseProfile
*/
	
	$dom:=DOM Parse XML variable:C720($stdOut)
	For ($attr; 1; DOM Count XML attributes:C727($dom))
		DOM GET XML ATTRIBUTE BY INDEX:C729($dom; $attr; $attrName; $attrValue)
		If ($attrName="baseProfile")
			DOM REMOVE XML ATTRIBUTE:C1084($dom; $attrName)
			DOM EXPORT TO VAR:C863($dom; $stdOut)
			break
		End if 
	End for 
	DOM CLOSE XML:C722($dom)
	
	var $icon : Picture
	BLOB TO PICTURE:C682($stdOut; $icon; ".svg")
	SET PICTURE TO PASTEBOARD:C521($icon)
End if 