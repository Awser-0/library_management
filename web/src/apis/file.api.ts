export { uploadImage };

import { request } from "./request";
import * as consts from "~/consts";
import type { Result } from "./request";

// /v1/file/upload/image

async function uploadImage(file: File) {
	const formData = new FormData();
	formData.append("upload-file", file);
	return request<Result<{ name: string }>>({
		url: "/v1/file/upload/image",
		method: "POST",
		data: formData,
	});
}
