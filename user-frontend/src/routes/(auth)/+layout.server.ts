import { session } from '$lib/stores/session';
import { redirect } from '@sveltejs/kit';
import { get } from 'svelte/store';
import type { LayoutServerLoadEvent } from './$types';

export async function load(event: LayoutServerLoadEvent) {
	const parent_user = (await event.parent())?.user;
	const locals_user = event.locals?.user;
	const session_user = get(session)?.user;

	const user = session_user || locals_user || parent_user;

	if (!user) {
		throw redirect(301, '/login');
	}
	return { user, token: event.cookies.get('token') };
}
