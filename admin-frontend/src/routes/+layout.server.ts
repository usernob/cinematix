import { routeApi } from '@/lib/util';
import type { LayoutServerLoadEvent } from './$types';
import type { User } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { redirect } from '@sveltejs/kit';

export async function load(event: LayoutServerLoadEvent) {
	const token = event.cookies.get('admin-token');

	event.depends('reload:user');

	if (!token) {
		return { token: null, user: null };
	}

	if (event.locals?.user) {
		return { token, user: event.locals.user };
	}

	const getData = await event.fetch(routeApi('admin/info'));

	const res: ApiResponse<User> = await getData.json();

	if (res.status === 'ok' || res.data) {
		delete res.data.password;
		event.locals.user = res.data;
	} else {
		throw redirect(302, '/login');
	}

	const user = event.locals?.user;
	if (!user) return { token, user: null };
	return { token, user };
}
