import { writable } from 'svelte/store';
import type { User } from '../types/modelTypes';

interface Session {
	user?: User | null;
}

export const session = writable<Session>({ user: null });
