import { writable } from 'svelte/store';
import type { Film, Genre, Penayangan, User } from '../types/modelTypes';

interface Session {
  user?: User | null;
}

export const session = writable<Session>({ user: null });

type FilmData = Film &{
  genre: Genre[]
  penayangan: Penayangan[]
}

export const dataFilm = writable<FilmData | null>(null)
