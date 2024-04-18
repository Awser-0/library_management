export { dayjsFormatDate };

import dayjs from "dayjs";

function dayjsFormatDate(time?: string) {
	if (time) {
		return dayjs(time).format("YYYY-MM-DD");
	} else {
		return "";
	}
}
