export type Film = {
  id: number;
  title: string;
  poster_path: string;
  overview: string;
  rating: number;
  sinopsis: string;
  tanggal_rilis: string;
};

export type Genre = {
  id: number;
  nama: string;
};

export type Penayangan = {
  id: number,
  film_id: number,
  auditorium_id: number,
  harga: number,
  mulai: string,
  selesai: string,
}

export type Tiket = {
  id: number,
  penyangan_id: number,
  user_id: number,
  status_pembayaran: string,
}

export type Kursi = {
  id: number,
  auditorium_id: number,
  nama: string,
}

export type Auditorium = {
  id: number,
  nama: string,
}

export type Seat = {
  id: number,
  kursi_id: number,
  tiket_id: number,
}

export type User = {
  id: number,
  nama: string,
  email: string,
  password?: string,
}


