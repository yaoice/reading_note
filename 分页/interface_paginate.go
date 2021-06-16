/*
接口层分页实现
*/
func CommonPaginate(x interface{}, offset int, limit int) (int, int) {
	xLen := reflect.ValueOf(x).Len()
	if offset+1 > xLen {
		offset = xLen - 1
		if offset < 0 {
			offset = 0
		}
	}
	end := offset + limit
	if end > xLen {
		end = xLen
	}
	return offset, end
}
