// Code generated by "stringer -type=itemType"; DO NOT EDIT

package nash

import "fmt"

const itemTypeName = "itemErroritemEOFitemCommentitemVarNameitemVarValueitemListOpenitemListCloseitemListElemitemCommanditemArgitemLeftBlockitemRightBlockitemStringitemRedirRightitemRedirRBracketitemRedirLBracketitemRedirFileitemRedirNetAddritemRedirMapEqualitemRedirMapLSideitemRedirMapRSideitemKeyworditemRforkitemRforkFlagsitemCd"

var itemTypeIndex = [...]uint16{0, 9, 16, 27, 38, 50, 62, 75, 87, 98, 105, 118, 132, 142, 156, 173, 190, 203, 219, 236, 253, 270, 281, 290, 304, 310}

func (i itemType) String() string {
	i--
	if i < 0 || i >= itemType(len(itemTypeIndex)-1) {
		return fmt.Sprintf("itemType(%d)", i+1)
	}
	return itemTypeName[itemTypeIndex[i]:itemTypeIndex[i+1]]
}
