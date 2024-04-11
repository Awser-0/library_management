export { queryBooks };

import { request } from "./request";
import * as consts from "~/consts";
import type { Result } from "./request";

type Book = {
	uuid: number;
	title: string;
	author: string;
	cover: string;
	introduction: string;
	publish_time?: string;
	total: number;
	update_time: string;
	create_time: string;
};

async function queryBooks(query_string: string) {
	return request<Result<{ books: Book[] }>>({
		url: "/v1/book/query",
		method: "POST",
		data: { query_string },
	});
}
