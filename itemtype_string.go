// Code generated by "stringer -type=itemType"; DO NOT EDIT

package nash

import "fmt"

const _itemType_name = "itemErroritemEOFitemBuiltinitemImportitemCommentitemSetEnvitemShowEnvitemVarNameitemAssignitemAssignCmditemConcatitemVariableitemListOpenitemListCloseitemListElemitemCommanditemPipeitemBindFnitemArgitemLeftBlockitemRightBlockitemLeftParenitemRightParenitemStringitemRedirRightitemRedirRBracketitemRedirLBracketitemRedirFileitemRedirNetAddritemRedirMapEqualitemRedirMapLSideitemRedirMapRSideitemIfitemElseitemComparisonitemRforkitemRforkFlagsitemCditemFnDeclitemFnInv"

var _itemType_index = [...]uint16{0, 9, 16, 27, 37, 48, 58, 69, 80, 90, 103, 113, 125, 137, 150, 162, 173, 181, 191, 198, 211, 225, 238, 252, 262, 276, 293, 310, 323, 339, 356, 373, 390, 396, 404, 418, 427, 441, 447, 457, 466}

func (i itemType) String() string {
	i -= 2
	if i < 0 || i >= itemType(len(_itemType_index)-1) {
		return fmt.Sprintf("itemType(%d)", i+2)
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}
