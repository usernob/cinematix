import { writable } from 'svelte/store';

import type { Film, Genre, Penayangan } from '../types/modelTypes';

export type FilmData = Film & {
  genre: Genre[];
  penayangan: Penayangan[];
};

export const dataFilm = writable<FilmData | null>(null);

dataFilm.subscribe(value => console.log({value}))
