package errorss

var (
	// --- basis msg
	Base5xx    = "This happened due to bad connection on server or database and internal error on server"
	Base4xx    = "This happened due to invalid client request or server can't fulfill client request"
	BaseSvrErr = "Unexpected server error: "

	// --- custom error
	MsgQryErr          = " query failed. " + Base5xx
	MsgSvrErr          = BaseSvrErr + Base5xx
	MsgParseParamIdErr = "Failed to parse param 'id'. " + Base5xx
	Msg404Err          = " not found. " + Base4xx
	Msg400Err          = "Failed to parse incoming input. " + Base4xx
	Msg401Err          = "Unauthorized access, read the error value!" + Base4xx
	Msg409Err          = " already created. " + Base4xx
	MsgNoInput         = "No fields provided for update. " + Base4xx

	Err409Fill = "Caught recreating/re-registering data"

	// --- success msg
	MsgGet200 = " data fetched successfully!"
	MsgPst201 = " data created!"
	MsgPtc200 = " data updated!"
	MsgDel200 = " data deleted!"
)
