export {
	selectBook,
	queryBooks,
	addBook,
	updateBook,
	queryRecords,
	querySelfRecords,
	applyBorrowBook,
	cancelBorrowBook,
	agreeBorrowBook,
	rejectBorrowBook,
	returnBorrowBook,
};

export type { BorrowRecord };

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

type BorrowRecord = {
	id: number;
	user_id: number;
	book_uuid: number;
	state: string;
	reply_time?: string;
	return_time?: string;
	apply_description: string;
	reply_description: string;
	update_time: string;
	create_time: string;
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

// 查询借阅记录
async function queryRecords(query: { book_uuid?: string; user_id?: string } = {}) {
	return request<Result<{ records: BorrowRecord[] }>>({
		url: "/v1/book/record/query",
		method: "POST",
		data: { ...query },
	});
}

// 查询自己借阅记录
async function querySelfRecords(query: { book_uuid?: string } = {}) {
	return request<Result<{ records: BorrowRecord[] }>>({
		url: "/v1/book/record/query/self",
		method: "POST",
		data: { ...query },
	});
}

// 申请借阅书籍
async function applyBorrowBook(book_uuid: number, desc: string) {
	return request<Result<{ records: BorrowRecord[] }>>({
		url: "/v1/book/borrow/apply",
		method: "POST",
		data: { book_uuid, desc },
	});
}

// 取消借阅书籍
async function cancelBorrowBook(record_id: number) {
	return request<Result<{ records: BorrowRecord[] }>>({
		url: "/v1/book/borrow/cancel",
		method: "POST",
		data: { record_id },
	});
}

// 同意借阅书籍
async function agreeBorrowBook(record_id: number, desc: string) {
	return request<Result>({
		url: "/v1/book/borrow/agree",
		method: "POST",
		data: { record_id, desc },
	});
}

// 拒绝借阅书籍
async function rejectBorrowBook(record_id: number, desc: string) {
	return request<Result>({
		url: "/v1/book/borrow/reject",
		method: "POST",
		data: { record_id, desc },
	});
}

// 归还借阅书籍
async function returnBorrowBook(record_id: number) {
	return request<Result>({
		url: "/v1/book/borrow/return",
		method: "POST",
		data: { record_id },
	});
}
