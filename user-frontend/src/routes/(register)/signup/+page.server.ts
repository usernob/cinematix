import type { ApiResponse } from '@/lib/types/apiResponse';
import type { User } from '@/lib/types/modelTypes';
import { routeApi } from '@/lib/util';
import { fail, type Actions, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const user = event.locals.user;
	if (user) {
		throw redirect(302, '/user');
	}
};

export const actions = {
	default: async ({ request, cookies, locals }) => {
		const formData = await request.formData();
		const email = formData.get('email');
		const password = formData.get('password');
		const username = formData.get('username');

		if (!email || !password || !username) {
			return fail(400, { missing: true, email, password, username });
		}

		const getRegister = await fetch(routeApi('register/user'), {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				email,
				password,
				nama: username,
			}),
		});
		const res: ApiResponse<{ token?: string; user?: User } | null> = await getRegister.json();

		if (res.status !== 'ok' || !res.data?.token || !res.data?.user) {
			return fail(400, { incorrect: true, email, password, username });
		}

		cookies.set('token', res.data?.token, {
			path: '/',
			maxAge: 60 * 60 * 24, // one day
		});

		delete res.data?.user.password;

		locals.user = res.data.user;
		return { success: true, user: res.data?.user };
	},
} satisfies Actions;
