import { writable } from 'svelte/store';

import type { Film, Genre, Kursi, Penayangan } from '../types/modelTypes';

export type FilmData = Film & {
  genre: Genre[];
  penayangan: Penayangan[];
};

export const dataFilm = writable<FilmData | null>(null);

dataFilm.subscribe((value) => console.log({ value }));

export const seatPesanan = writable<Kursi[]>([]);

seatPesanan.subscribe((value) => console.log({ value }));
