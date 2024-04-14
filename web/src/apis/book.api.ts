export { selectBook, queryBooks, addBook, updateBook };

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

type BookInfo = {
	title: string;
	author: string;
	cover: string;
	introduction: string;
	publish_time: Date | null;
	total: number;
};

// 查看书籍
async function selectBook(uuid: number) {
	return request<Result<{ book: Book }>>({
		url: "/v1/book/select",
		method: "POST",
		data: { uuid },
	});
}

// 查询书籍列表，（查询范围：标题和作者）
async function queryBooks(query_string: string) {
	return request<Result<{ books: Book[] }>>({
		url: "/v1/book/query",
		method: "POST",
		data: { query_string },
	});
}

// 更新书籍
async function updateBook(uuid: number, info: BookInfo) {
	return request<Result>({
		url: "/v1/book/update",
		method: "POST",
		data: { uuid, ...info },
	});
}

// 添加书籍
async function addBook(info: BookInfo) {
	return request<Result>({
		url: "/v1/book/add",
		method: "POST",
		data: { ...info },
	});
}
