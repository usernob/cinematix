export type BaseModel = {
	id: number;
	created_at: string;
	updated_at: string;
	deleted_at?: string;
};

export type Film = BaseModel & {
	title: string;
	poster_path: string;
	overview: string;
	rating: number;
	sinopsis: string;
	tanggal_rilis: string;
};

export type Genre = BaseModel & {
	nama: string;
};

export type Penayangan = BaseModel & {
	film_id: number;
	audiotorium_id: number;
	harga: number;
	mulai: string;
	selesai: string;
};

export type Tiket = BaseModel & {
	penyangan_id: number;
	total_harga: number;
	user_id: number;
	status_pembayaran: string;
};

export type Kursi = BaseModel & {
	auditorium_id: number;
	nama: string;
};

export type Auditorium = BaseModel & {
	nama: string;
};

export type Seat = BaseModel & {
	kursi_id: number;
	tiket_id: number;
};

export type User = BaseModel & {
	nama: string;
	email: string;
	avatar: string;
	password?: string;
};
